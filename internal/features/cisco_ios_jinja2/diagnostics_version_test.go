package cisco_ios_jinja2_test

import (
	"testing"
)

// TestVersionMismatchDiagnostics covers the version-mismatch pass: matching
// versions produce no diagnostic; a mismatch emits an Error diagnostic with
// Code "version-mismatch" anchored on the version_statement; missing comment
// or missing version_statement produce no diagnostic.
func TestVersionMismatchDiagnostics(t *testing.T) {
	cases := []struct {
		name    string
		src     string
		wantHit bool
	}{
		{
			name:    "matching versions -> no diag",
			src:     "! version 17.3\n!\nversion 17.3\n!\n",
			wantHit: false,
		},
		{
			name:    "mismatch -> Error with code",
			src:     "! version 17.3\n!\nversion 17.4\n!\n",
			wantHit: true,
		},
		{
			name:    "missing running-version comment -> no diag",
			src:     "! this is some other comment\n!\nversion 17.4\n!\n",
			wantHit: false,
		},
		{
			name:    "missing version_statement -> no diag",
			src:     "! version 17.3\n!\nhostname r1\n!\n",
			wantHit: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			diags := openDiags(t, tc.src)
			d, found := findDiagByCode(diags, "version-mismatch")
			if tc.wantHit && !found {
				t.Fatalf("expected version-mismatch diagnostic; got %d diags: %+v", len(diags), diags)
			}
			if !tc.wantHit && found {
				t.Fatalf("did not expect version-mismatch diagnostic; got: %+v", d)
			}
			if found {
				if d.Severity != 1 { // SeverityError
					t.Errorf("severity: got %d, want Error (1)", d.Severity)
				}
				if d.Source != "chunter" {
					t.Errorf("source: got %q, want \"chunter\"", d.Source)
				}
				if d.Message == "" {
					t.Errorf("message is empty")
				}
			}
		})
	}
}

// TestVersionMismatchDiagnostics_FirstCommentWins verifies that when multiple
// top-level comments exist, the FIRST one matching runningVersionRe supplies
// the running version (matching show-run emit order).
func TestVersionMismatchDiagnostics_FirstCommentWins(t *testing.T) {
	src := "! version 17.3\n! version 17.4\n!\nversion 17.3\n!\n"
	diags := openDiags(t, src)
	if _, found := findDiagByCode(diags, "version-mismatch"); found {
		t.Errorf("first comment said 17.3 == config 17.3; expected no mismatch, got: %+v", diags)
	}
}
