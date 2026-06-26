package cisco_ios_test

import (
	"context"
	"testing"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/features/cisco_ios"
	"github.com/dgethings/chunter/internal/protocol"
)

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
			f := cisco_ios.New()
			defer f.Close()
			doc := document.New("file:///test.cfg", "cisco_ios", 1, []byte(tc.src))
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
