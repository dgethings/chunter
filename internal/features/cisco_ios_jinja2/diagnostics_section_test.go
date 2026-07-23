package cisco_ios_jinja2_test

import (
	"strings"
	"testing"

	"github.com/dgethings/chunter/internal/protocol"
)

// findHintByKeyword returns the first Hint diagnostic whose message contains
// the given keyword token, reporting whether one was found.
func findHintByKeyword(diags []protocol.Diagnostic, keyword string) (protocol.Diagnostic, bool) {
	for _, d := range diags {
		if d.Severity == protocol.SeverityHint && strings.Contains(d.Message, keyword) {
			return d, true
		}
	}
	return protocol.Diagnostic{}, false
}

// assertNoDiagAbout fails the test if any diagnostic message contains needle.
func assertNoDiagAbout(t *testing.T, diags []protocol.Diagnostic, needle string) {
	t.Helper()
	for _, d := range diags {
		if strings.Contains(d.Message, needle) {
			t.Errorf("did not expect a diagnostic about %q: %s", needle, d.Message)
		}
	}
}

// TestWrongSectionDiagnostics_InterfaceCmdInRouter covers the headline case: an
// interface command (speed) copied into a router section is flagged as a Hint.
func TestWrongSectionDiagnostics_InterfaceCmdInRouter(t *testing.T) {
	src := "!\nrouter bgp 100\n speed 1000\n!\n"
	diags := openDiags(t, src)

	d, found := findHintByKeyword(diags, "speed")
	if !found {
		t.Fatalf("expected Hint diagnostic for 'speed' in router section; got %d diags: %+v", len(diags), diags)
	}
	if !strings.Contains(d.Message, "config-router") {
		t.Errorf("expected message to mention config-router; got %q", d.Message)
	}
	if !strings.Contains(d.Message, "config-if") {
		t.Errorf("expected message to mention config-if (where speed is valid); got %q", d.Message)
	}
	if d.Source != "chunter" {
		t.Errorf("source: got %q, want \"chunter\"", d.Source)
	}
}

// TestWrongSectionDiagnostics_CorrectSection asserts known commands in their
// own section produce no wrong-section Hint.
func TestWrongSectionDiagnostics_CorrectSection(t *testing.T) {
	cases := []struct {
		name string
		src  string
		kw   string // keyword that must NOT be flagged
	}{
		{
			name: "speed in interface (config-if)",
			src:  "!\ninterface Gi0/0\n speed 1000\n!\n",
			kw:   "speed",
		},
		{
			name: "aggregate-address in router (config-router)",
			src:  "!\nrouter bgp 100\n aggregate-address 10.0.0.0 255.0.0.0\n!\n",
			kw:   "aggregate-address",
		},
		{
			name: "hostname at top level (config)",
			src:  "hostname r1\n",
			kw:   "hostname",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			diags := openDiags(t, tc.src)
			if d, found := findHintByKeyword(diags, tc.kw); found {
				t.Errorf("did not expect a Hint for %q in its own section: %s", tc.kw, d.Message)
			}
		})
	}
}

// TestWrongSectionDiagnostics_RouterCmdInInterface flags a router command that
// was pasted into an interface section.
func TestWrongSectionDiagnostics_RouterCmdInInterface(t *testing.T) {
	src := "!\ninterface Gi0/0\n aggregate-address 10.0.0.0 255.0.0.0\n!\n"
	diags := openDiags(t, src)

	d, found := findHintByKeyword(diags, "aggregate-address")
	if !found {
		t.Fatalf("expected Hint diagnostic for 'aggregate-address' in interface section; got %d diags: %+v", len(diags), diags)
	}
	if !strings.Contains(d.Message, "config-if") {
		t.Errorf("expected message to mention config-if (enclosing section); got %q", d.Message)
	}
	if !strings.Contains(d.Message, "config-router") {
		t.Errorf("expected message to mention config-router (valid section); got %q", d.Message)
	}
}

// TestWrongSectionDiagnostics_GlobalCmdInSubmode flags a global/config command
// placed inside a sub-mode (e.g. hostname inside a router section).
func TestWrongSectionDiagnostics_GlobalCmdInSubmode(t *testing.T) {
	src := "!\nrouter bgp 100\n hostname r1\n!\n"
	diags := openDiags(t, src)

	d, found := findHintByKeyword(diags, "hostname")
	if !found {
		t.Fatalf("expected Hint diagnostic for 'hostname' in router section; got %d diags: %+v", len(diags), diags)
	}
	if !strings.Contains(d.Message, "config-router") {
		t.Errorf("expected message to mention config-router (enclosing section); got %q", d.Message)
	}
}

// TestWrongSectionDiagnostics_UnknownKeywordNotFlagged ensures commands that
// are not in the keyword database (and therefore cannot be validated against a
// section) produce no Hint.
func TestWrongSectionDiagnostics_UnknownKeywordNotFlagged(t *testing.T) {
	src := "!\nrouter bgp 100\n frobnicate 42\n!\n"
	diags := openDiags(t, src)
	assertNoDiagAbout(t, diags, "frobnicate")
}

// TestWrongSectionDiagnostics_MultiSectionKeyword verifies that a keyword
// valid in more than one section (mtu is valid in config-if and
// config-if-xconn) is not flagged when used in any of its valid sections.
func TestWrongSectionDiagnostics_MultiSectionKeyword(t *testing.T) {
	src := "!\ninterface Gi0/0\n mtu 1500\n!\n"
	diags := openDiags(t, src)
	assertNoDiagAbout(t, diags, "mtu")
}

// TestWrongSectionDiagnostics_OnlyHintsForKeywords confirms the pass does not
// accidentally flag the well-known commands used elsewhere in the same config
// (regression guard against over-eager matching).
func TestWrongSectionDiagnostics_NoFalsePositives(t *testing.T) {
	// speed is correct here (interface); aggregate-address is correct here
	// (router). Neither should be flagged.
	src := "!\ninterface Gi0/0\n speed 1000\n!\nrouter bgp 100\n aggregate-address 10.0.0.0 255.0.0.0\n!\n"
	diags := openDiags(t, src)
	assertNoDiagAbout(t, diags, "speed")
	assertNoDiagAbout(t, diags, "aggregate-address")
}
