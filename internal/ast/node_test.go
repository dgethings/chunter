package ast_test

import (
	"testing"

	"github.com/dgethings/chunter/internal/ast"
	ts_ci "github.com/dgethings/tree-sitter-cisco_ios/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

const fullConfig = "!\n! version 26.1.0\n!\nversion 26.2.0\n!\narchive\n" +
	"   path disk0:someconfig\n   end\n!\n" +
	"async-bootp bootfile :172.30.1.1 \"pcboot\"\n!\n" +
	"line console\n   activation-character 127\n!\n" +
	"hostname test\n!\n" +
	"alias exec siib show ip interface brief\n" +
	"alias exec sibnb show ip bgp neighbor brief\n!\n"

func parseRoot(t *testing.T, src string) *sitter.Node {
	t.Helper()
	p := sitter.NewParser()
	p.SetLanguage(sitter.NewLanguage(ts_ci.Language()))
	tree := p.Parse([]byte(src), nil)
	t.Cleanup(func() {
		tree.Close()
		p.Close()
	})
	return tree.RootNode()
}

func TestFindNodeAtPosition_NilRoot(t *testing.T) {
	if got := ast.FindNodeAtPosition(nil, 0, 0); got != nil {
		t.Fatalf("want nil for nil root, got %q", got.GrammarName())
	}
}

func TestFindNodeAtPosition(t *testing.T) {
	cases := []struct {
		name string
		src  string
		line uint
		col  uint
		want string // expected GrammarName of the deepest node
	}{
		// The motivating example: an incomplete hostname line with no trailing "!".
		// The hostname node spans only "hostname bob" (cols 0-12), so positions at
		// or past its end fall back to the config root.
		{"user/bang", "!\nhostname bob\n", 0, 0, "!"},
		{"user/hostname_keyword", "!\nhostname bob\n", 1, 0, "hostname"},
		{"user/hostname_mid", "!\nhostname bob\n", 1, 4, "hostname"},
		{"user/value", "!\nhostname bob\n", 1, 9, "value"},
		{"user/complation_accepted", "!\nhostname ", 1, 10, "hostname"},
		{"user/eol_falls_back_to_root", "!\nhostname bob\n", 1, 12, "config"},
		{"user/trailing_line_falls_back_to_root", "!\nhostname bob\n", 2, 0, "config"},

		// hostname section (system.txt)
		{"hostname/running_version", "!\n! version 26.1.0\n!\nhostname test\n!\n", 1, 10, "running_version"},
		{"hostname/keyword", "!\n! version 26.1.0\n!\nhostname test\n!\n", 3, 0, "hostname"},
		{"hostname/value", "!\n! version 26.1.0\n!\nhostname test\n!\n", 3, 9, "value"},
		{"hostname/trailing_bang", "!\n! version 26.1.0\n!\nhostname test\n!\n", 4, 0, "!"},

		// service section (service.txt)
		{"service/version_keyword", "!\n! version 26.1.0\n!\nversion 26.2.0\n!\n", 3, 0, "version"},
		{"service/configured_version", "!\n! version 26.1.0\n!\nversion 26.2.0\n!\n", 3, 8, "configured_version"},

		// line section (line.txt)
		{"line/keyword", "!\n! version 26.1.0\n!\nline console\n   activation-character 127\n!\n", 3, 0, "line"},
		{"line/console", "!\n! version 26.1.0\n!\nline console\n   activation-character 127\n!\n", 3, 5, "console"},
		{"line/activation_character", "!\n! version 26.1.0\n!\nline console\n   activation-character 127\n!\n", 4, 3, "activation-character"},
		{"line/ascii_value", "!\n! version 26.1.0\n!\nline console\n   activation-character 127\n!\n", 4, 24, "ascii_value"},

		// archive section (archive.txt)
		{"archive/keyword", "!\n! version 26.1.0\n!\narchive\n   path disk0:someconfig\n   end\n!\n", 3, 0, "archive"},
		{"archive/file_path", "!\n! version 26.1.0\n!\narchive\n   path disk0:someconfig\n   end\n!\n", 4, 11, "file_path"},
		{"archive/end", "!\n! version 26.1.0\n!\narchive\n   path disk0:someconfig\n   end\n!\n", 5, 3, "end"},

		// async-bootp section (async-bootp.txt)
		{"async/keyword", "!\n! version 26.1.0\n!\nasync-bootp bootfile :172.30.1.1 \"pcboot\"\n!\n", 3, 0, "async-bootp"},
		{"async/hostname", "!\n! version 26.1.0\n!\nasync-bootp bootfile :172.30.1.1 \"pcboot\"\n!\n", 3, 24, "async_bootp_hostname"},
		{"async/data", "!\n! version 26.1.0\n!\nasync-bootp bootfile :172.30.1.1 \"pcboot\"\n!\n", 3, 33, "async_bootp_data"},

		// full config (full-config.txt) - cross-section checks
		{"full/file_path", fullConfig, 6, 8, "file_path"},
		{"full/ascii_value", fullConfig, 12, 24, "ascii_value"},
		{"full/hostname", fullConfig, 14, 0, "hostname"},
		{"full/alias_mode", fullConfig, 16, 6, "exec"},
		{"full/alias_command", fullConfig, 17, 16, "command"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := parseRoot(t, tc.src)
			got := ast.FindNodeAtPosition(root, tc.line, tc.col)
			if got == nil {
				t.Fatalf("FindNodeAtPosition(%d,%d) returned nil; want %q", tc.line, tc.col, tc.want)
			}
			if name := got.GrammarName(); name != tc.want {
				t.Errorf("FindNodeAtPosition(%d,%d) = %q, want %q", tc.line, tc.col, name, tc.want)
			}
		})
	}
}
