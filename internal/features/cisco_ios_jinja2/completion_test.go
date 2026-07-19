package cisco_ios_jinja2_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/features/cisco_ios_jinja2"
	"github.com/dgethings/chunter/internal/protocol"
)

// placeholderDefaultRe matches LSP snippet placeholders of the form ${N:default}
// where N is a tabstop index and default is non-empty. Mirrors the regex in
// completion.go so this test fails if the production regex lets one through.
var placeholderDefaultRe = regexp.MustCompile(`\$\{(\d+):[^}]*\}`)

func TestCompletionWhileTypingValue(t *testing.T) {
	cases := []struct {
		name string
		src  string
		line uint
		col  uint
	}{
		{"hostname_space_col9", "!\nhostname ", 1, 9},
		{"hostname_space_col10", "!\nhostname ", 1, 10},
		{"hostname_name_on_value", "!\nhostname name", 1, 9},
		{"hostname_name_on_value", "!\nhostname name", 1, 12},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := cisco_ios_jinja2.New()
			defer f.Close()
			doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, []byte(tc.src))
			if _, err := f.DidOpen(context.Background(), doc); err != nil {
				t.Fatalf("DidOpen: %v", err)
			}
			items, err := f.Completion(context.Background(), doc, protocol.Position{Line: tc.line, Character: tc.col})
			if err != nil {
				t.Fatalf("Completion: %v", err)
			}
			t.Logf("%q @ {%d,%d} -> %d items", tc.src, tc.line, tc.col, len(items))
			if len(items) != 0 {
				t.Errorf("expected no keyword suggestions while typing a value, got %d", len(items))
			}
		})
	}
}

// TestCompletionSnippetsStripPlaceholderDefaults verifies that completion items
// never emit ${N:default} placeholders in their insertText. vim.snippet's
// selection of non-empty placeholders can leave the editor in INSERT mode
// instead of SELECT mode on some clients (notably blink.cmp on Neovim 0.12),
// which causes typed characters to insert before the default text rather than
// replacing it. See completion.go for the full rationale.
func TestCompletionSnippetsStripPlaceholderDefaults(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	src := []byte("!\n")
	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, src)
	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen: %v", err)
	}

	items, err := f.Completion(context.Background(), doc, protocol.Position{Line: 1, Character: 0})
	if err != nil {
		t.Fatalf("Completion: %v", err)
	}
	if len(items) == 0 {
		t.Fatalf("expected at least one completion item")
	}

	for _, item := range items {
		if item.InsertText == nil {
			t.Errorf("item %q: missing insertText", item.Label)
			continue
		}
		if loc := placeholderDefaultRe.FindStringIndex(*item.InsertText); loc != nil {
			t.Errorf("item %q: insertText %q still contains a ${N:default} placeholder at %v",
				item.Label, *item.InsertText, loc)
		}
		if item.FilterText == nil {
			t.Errorf("item %q: missing filterText", item.Label)
			continue
		}
		if *item.FilterText != item.Label {
			t.Errorf("item %q: filterText %q != label %q",
				item.Label, *item.FilterText, item.Label)
		}
	}
}

// TestCompletionHostnameSnippetHasEmptyTabstop pins the exact insertText we
// emit for the hostname keyword after the fix: the original data file ships
// `hostname ${1:name}` but the LSP response must drop the `name` default so
// editors land the cursor in INSERT mode at the tabstop, fixing the bug where
// typing `r1` produced `hostname r1name` instead of `hostname r1`.
func TestCompletionHostnameSnippetHasEmptyTabstop(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, []byte("!\n"))
	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen: %v", err)
	}

	items, err := f.Completion(context.Background(), doc, protocol.Position{Line: 1, Character: 0})
	if err != nil {
		t.Fatalf("Completion: %v", err)
	}

	var hostnameItem *protocol.CompletionItem
	for i := range items {
		if items[i].Label == "hostname" {
			hostnameItem = &items[i]
			break
		}
	}
	if hostnameItem == nil {
		t.Fatalf("hostname completion item not found")
	}
	if hostnameItem.InsertText == nil {
		t.Fatalf("hostname item has no insertText")
	}
	want := "hostname ${1}"
	if got := *hostnameItem.InsertText; got != want {
		t.Errorf("hostname insertText = %q, want %q", got, want)
	}
}

// TestCompletionNoDuplicateLabels pins the fix for the duplicate-suggestions
// bug. The keyword database both (a) carries multiple snippets per keyword —
// the positive form and the `no <kw>` form — and (b) contains ~900 exact
// duplicate keyword records. The completion output must contain at most one
// item per distinct label regardless of how the underlying data repeats.
func TestCompletionNoDuplicateLabels(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, []byte("!\n"))
	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen: %v", err)
	}
	items, err := f.Completion(context.Background(), doc, protocol.Position{Line: 1, Character: 0})
	if err != nil {
		t.Fatalf("Completion: %v", err)
	}
	if len(items) == 0 {
		t.Fatalf("expected at least one completion item")
	}
	seen := map[string]bool{}
	for _, item := range items {
		if seen[item.Label] {
			t.Errorf("duplicate completion label %q", item.Label)
			continue
		}
		seen[item.Label] = true
		if item.FilterText == nil {
			t.Errorf("item %q: missing filterText", item.Label)
			continue
		}
		if *item.FilterText != item.Label {
			t.Errorf("item %q: filterText %q != label %q",
				item.Label, *item.FilterText, item.Label)
		}
	}
}

// TestCompletionNoFormHasDistinctLabel verifies the `no <kw>` snippet gets its
// own distinguishable label instead of collapsing onto the positive form. The
// downward-compatible-config keyword ships both
// `downward-compatible-config ${1:version}` and `no downward-compatible-config`;
// before the fix both surfaced as a single indistinguishable
// `downward-compatible-config` entry (twice).
func TestCompletionNoFormHasDistinctLabel(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, []byte("!\n"))
	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen: %v", err)
	}
	items, err := f.Completion(context.Background(), doc, protocol.Position{Line: 1, Character: 0})
	if err != nil {
		t.Fatalf("Completion: %v", err)
	}
	labels := map[string]bool{}
	for _, item := range items {
		labels[item.Label] = true
	}
	if !labels["downward-compatible-config"] {
		t.Errorf("expected label %q (positive form)", "downward-compatible-config")
	}
	if !labels["no downward-compatible-config"] {
		t.Errorf("expected label %q (negation form)", "no downward-compatible-config")
	}
}
