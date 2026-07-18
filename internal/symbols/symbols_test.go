package symbols_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/dgethings/chunter/internal/symbols"
	ts_ci "github.com/dgethings/tree-sitter-cisco-ios-jinja2/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func parseRoot(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	p := sitter.NewParser()
	p.SetLanguage(sitter.NewLanguage(ts_ci.Language()))
	tree := p.Parse([]byte(src), nil)
	t.Cleanup(func() {
		tree.Close()
		p.Close()
	})
	return tree.RootNode(), []byte(src)
}

// namesByKind groups symbols by Kind, returning the sorted names within each
// kind for stable comparison in tests.
func namesByKind(syms []symbols.Symbol) map[symbols.Kind][]string {
	out := make(map[symbols.Kind][]string)
	for _, s := range syms {
		out[s.Kind] = append(out[s.Kind], s.Name)
	}
	for k := range out {
		sort.Strings(out[k])
	}
	return out
}

func TestExtract_Sections(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want map[symbols.Kind][]string
	}{
		{
			name: "interface",
			src:  "interface GigabitEthernet0/0\n!\n",
			want: map[symbols.Kind][]string{symbols.KindInterface: {"GigabitEthernet0/0"}},
		},
		{
			name: "router bgp",
			src:  "router bgp 100\n!\n",
			want: map[symbols.Kind][]string{symbols.KindRouter: {"100"}},
		},
		{
			name: "route-map with action and sequence",
			src:  "route-map FOO permit 10\n!\n",
			want: map[symbols.Kind][]string{symbols.KindRouteMap: {"FOO"}},
		},
		{
			name: "route-map bare",
			src:  "route-map BAR\n!\n",
			want: map[symbols.Kind][]string{symbols.KindRouteMap: {"BAR"}},
		},
		{
			name: "class-map with match-type",
			src:  "class-map match-any VOICE\n!\n",
			want: map[symbols.Kind][]string{symbols.KindClassMap: {"VOICE"}},
		},
		{
			name: "class-map bare name",
			src:  "class-map DATA\n!\n",
			want: map[symbols.Kind][]string{symbols.KindClassMap: {"DATA"}},
		},
		{
			name: "policy-map",
			src:  "policy-map QOS\n!\n",
			want: map[symbols.Kind][]string{symbols.KindPolicyMap: {"QOS"}},
		},
		{
			name: "vlan",
			src:  "vlan 10\n name MANAGEMENT\n!\n",
			want: map[symbols.Kind][]string{symbols.KindVlan: {"10"}},
		},
		{
			name: "line vty with range",
			src:  "line vty 0 4\n!\n",
			want: map[symbols.Kind][]string{symbols.KindLine: {"vty-0-4"}},
		},
		{
			name: "line console single number",
			src:  "line console 0\n!\n",
			want: map[symbols.Kind][]string{symbols.KindLine: {"console-0"}},
		},
		{
			name: "redundancy singleton",
			src:  "redundancy\n auto-sync running-config\n!\n",
			want: map[symbols.Kind][]string{symbols.KindRedundancy: {"redundancy"}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root, content := parseRoot(t, tc.src)
			got := namesByKind(symbols.Extract("file:///test.ios.j2", root, content))
			for k, wantNames := range tc.want {
				gotNames, ok := got[k]
				if !ok {
					t.Fatalf("kind %s: no symbols extracted; want %v", k, wantNames)
				}
				if len(gotNames) != len(wantNames) {
					t.Fatalf("kind %s: got %v, want %v", k, gotNames, wantNames)
				}
				for i := range wantNames {
					if gotNames[i] != wantNames[i] {
						t.Errorf("kind %s[%d]: got %q, want %q", k, i, gotNames[i], wantNames[i])
					}
				}
			}
		})
	}
}

func TestExtract_ACLs(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want []string // expected ACL names
	}{
		{
			name: "named standard ACL",
			src:  "ip access-list standard FOO\n!\n",
			want: []string{"FOO"},
		},
		{
			name: "named extended ACL",
			src:  "ip access-list extended BAR\n!\n",
			want: []string{"BAR"},
		},
		{
			name: "numbered ACL",
			src:  "access-list 101 permit ip any any\n!\n",
			want: []string{"access-list 101"},
		},
		{
			name: "mixed ACL forms",
			src: "ip access-list standard NAMED\n" +
				"access-list 100 permit ip any any\n" +
				"ip access-list extended NAMED2\n" +
				"access-list 101 deny tcp any any\n",
			want: []string{"NAMED", "access-list 100", "NAMED2", "access-list 101"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root, content := parseRoot(t, tc.src)
			syms := symbols.Extract("file:///test.ios.j2", root, content)
			var got []string
			for _, s := range syms {
				if s.Kind == symbols.KindACL {
					got = append(got, s.Name)
				}
			}
			sort.Strings(got)
			if len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			// Compare as sets since numbered ACL ordering within document may
			// not match the input order across mixed forms.
			wantSorted := append([]string(nil), tc.want...)
			sort.Strings(wantSorted)
			for i := range wantSorted {
				if got[i] != wantSorted[i] {
					t.Errorf("[%d]: got %q, want %q", i, got[i], wantSorted[i])
				}
			}
		})
	}
}

func TestExtract_Integration(t *testing.T) {
	src := `!
hostname r1
!
interface GigabitEthernet0/0
 ip address 10.0.0.1 255.255.255.0
!
router bgp 100
 neighbor 10.0.0.2 remote-as 200
!
route-map RM-OUT permit 10
 match ip address NAMED-ACL
!
class-map match-any VOICE
 match ip address NAMED-ACL
!
policy-map QOS
 class VOICE
  priority
!
ip access-list standard NAMED-ACL
 permit 10.0.0.0 0.0.0.255
!
vlan 10
 name MGMT
!
line vty 0 4
 transport input ssh
!
redundancy
!
`
	root, content := parseRoot(t, src)
	syms := symbols.Extract("file:///test.ios.j2", root, content)

	want := map[symbols.Kind]int{
		symbols.KindInterface:  1,
		symbols.KindRouter:     1,
		symbols.KindRouteMap:   1,
		symbols.KindClassMap:   1,
		symbols.KindPolicyMap:  1,
		symbols.KindACL:        1,
		symbols.KindVlan:       1,
		symbols.KindLine:       1,
		symbols.KindRedundancy: 1,
	}
	got := make(map[symbols.Kind]int)
	for _, s := range syms {
		got[s.Kind]++
	}
	for k, wantCount := range want {
		if got[k] != wantCount {
			t.Errorf("kind %s: got %d symbols, want %d", k, got[k], wantCount)
		}
	}
}

func TestExtract_NameRangeNonZero(t *testing.T) {
	// Every symbol's NameRange must be non-empty (Start < End on at least
	// one dimension). This catches bugs where NameRange was left at zero.
	src := "route-map FOO permit 10\n!\ninterface Gi0/0\n!\nvlan 10\n!\n"
	root, content := parseRoot(t, src)
	syms := symbols.Extract("file:///test.ios.j2", root, content)
	if len(syms) != 3 {
		t.Fatalf("got %d symbols, want 3", len(syms))
	}
	for _, s := range syms {
		if s.NameRange.Start.Line == 0 && s.NameRange.Start.Character == 0 &&
			s.NameRange.End.Line == 0 && s.NameRange.End.Character == 0 {
			t.Errorf("symbol %s %q has zero NameRange", s.Kind, s.Name)
		}
		if s.Name == "" {
			t.Errorf("symbol %v has empty Name", s.Kind)
		}
	}
}

func TestTable_Lookup(t *testing.T) {
	src := "route-map FOO permit 10\n!\nroute-map BAR deny 5\n!\ninterface Gi0/0\n!\n"
	root, content := parseRoot(t, src)
	tbl := symbols.NewTable()
	tbl.Index("file:///test", root, content)

	if got := tbl.Lookup("file:///test", symbols.KindRouteMap, "FOO"); len(got) != 1 {
		t.Errorf("Lookup(FOO): got %d, want 1", len(got))
	}
	if got := tbl.Lookup("file:///test", symbols.KindRouteMap, "MISSING"); len(got) != 0 {
		t.Errorf("Lookup(MISSING): got %d, want 0", len(got))
	}
	if got := tbl.Lookup("file:///test", symbols.KindInterface, "FOO"); len(got) != 0 {
		t.Errorf("Lookup(FOO/interface): got %d, want 0", len(got))
	}
	if got := tbl.LookupAny("file:///test", "BAR"); len(got) != 1 {
		t.Errorf("LookupAny(BAR): got %d, want 1", len(got))
	}
	if got := tbl.All("file:///test"); len(got) != 3 {
		t.Errorf("All: got %d, want 3", len(got))
	}
	if got := tbl.All("file:///other"); len(got) != 0 {
		t.Errorf("All(other URI): got %d, want 0", len(got))
	}

	tbl.Clear("file:///test")
	if got := tbl.All("file:///test"); len(got) != 0 {
		t.Errorf("after Clear: got %d, want 0", len(got))
	}
}

func TestExtract_NilRoot(t *testing.T) {
	if got := symbols.Extract("file:///test", nil, nil); got != nil {
		t.Errorf("Extract(nil): got %v, want nil", got)
	}
}

func TestExtract_IgnoresNonDefinitionLines(t *testing.T) {
	// Body commands (match, set, class, permit, deny, neighbor, etc.) must
	// NOT be indexed as symbols. Only headers / definitions count.
	src := `route-map FOO permit 10
 match ip address ACL1
 set metric 100
!
ip access-list extended ACL1
 permit tcp any any
!
router bgp 100
 neighbor 10.0.0.2 remote-as 200
!
`
	root, content := parseRoot(t, src)
	syms := symbols.Extract("file:///test", root, content)

	// Expected: route-map FOO, ACL ACL1, router 100.
	wantKinds := map[symbols.Kind]int{
		symbols.KindRouteMap: 1,
		symbols.KindACL:      1,
		symbols.KindRouter:   1,
	}
	got := make(map[symbols.Kind]int)
	for _, s := range syms {
		got[s.Kind]++
		// No body command keyword should ever appear as a name.
		if strings.HasPrefix(s.Name, "match ") ||
			strings.HasPrefix(s.Name, "set ") ||
			strings.HasPrefix(s.Name, "permit ") ||
			strings.HasPrefix(s.Name, "neighbor ") {
			t.Errorf("body command leaked into symbols: %s %q", s.Kind, s.Name)
		}
	}
	for k, c := range wantKinds {
		if got[k] != c {
			t.Errorf("kind %s: got %d, want %d", k, got[k], c)
		}
	}
}
