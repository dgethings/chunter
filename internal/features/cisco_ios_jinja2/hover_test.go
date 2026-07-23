package cisco_ios_jinja2_test

import (
	"context"
	"testing"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/features/cisco_ios_jinja2"
	"github.com/dgethings/chunter/internal/protocol"
)

func TestHoverHostname(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	//       line 0: !
	//       line 1: ! version 26.1.0
	//       line 2: !
	//       line 3: hostname test
	//       line 4: !
	content := []byte("!\n! version 26.1.0\n!\nhostname test\n!\n")
	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, content)

	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen failed: %v", err)
	}

	// Hover on "hostname" — line 3, character 0 (zero-based)
	result, err := f.Hover(context.Background(), doc, protocol.Position{Line: 3, Character: 0})
	if err != nil {
		t.Fatalf("Hover failed: %v", err)
	}
	if result == nil {
		t.Fatal("expected hover result, got nil")
	}

	want := "To specify or modify the hostname for the network server, use the hostname command in global configuration mode."
	if result.Contents.Value != want {
		t.Errorf("hover value = %q\nwant %q", result.Contents.Value, want)
	}
	if result.Contents.Kind != "plaintext" {
		t.Errorf("hover kind = %q, want %q", result.Contents.Kind, "plaintext")
	}
}

func TestHoverNoKeyword(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	content := []byte("!\n! version 26.1.0\n!\nhostname test\n!\n")
	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, content)

	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen failed: %v", err)
	}

	// Hover on the version line — not a documented keyword
	result, err := f.Hover(context.Background(), doc, protocol.Position{Line: 1, Character: 1})
	if err != nil {
		t.Fatalf("Hover failed: %v", err)
	}
	// Currently returns a zero-value HoverResult for unknown nodes; adjust
	// if you change nodeByPosition to return nil for misses.
	if result != nil && result.Contents.Value != "" {
		t.Errorf("expected empty/nil hover for undocumented node, got %q", result.Contents.Value)
	}
}

func TestHoverHostnameKeywordEnd(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	//       line 0: !
	//       line 1: ! version 26.1.0
	//       line 2: !
	//       line 3: hostname test
	//       line 4: !
	content := []byte("!\n! version 26.1.0\n!\nhostname test\n!\n")
	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, content)

	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen failed: %v", err)
	}

	// Hover at the END of "hostname" — line 3, character 8 (zero-based).
	// This boundary resolves to the hostname_statement node, not the leaf.
	result, err := f.Hover(context.Background(), doc, protocol.Position{Line: 3, Character: 8})
	if err != nil {
		t.Fatalf("Hover failed: %v", err)
	}
	if result == nil {
		t.Fatal("expected hover result, got nil")
	}

	want := "To specify or modify the hostname for the network server, use the hostname command in global configuration mode."
	if result.Contents.Value != want {
		t.Errorf("hover value = %q\nwant %q", result.Contents.Value, want)
	}
	if result.Contents.Kind != "plaintext" {
		t.Errorf("hover kind = %q, want %q", result.Contents.Kind, "plaintext")
	}
}

func TestHoverInterface(t *testing.T) {
	f := cisco_ios_jinja2.New()
	defer f.Close()

	//       line 0: !
	//       line 1: interface GigabitEthernet0/0
	//       line 2: !
	content := []byte("!\ninterface GigabitEthernet0/0\n!\n")
	doc := document.New("file:///test.cfg", "cisco_ios_jinja2", 1, content)

	if _, err := f.DidOpen(context.Background(), doc); err != nil {
		t.Fatalf("DidOpen failed: %v", err)
	}

	// Hover on "interface" — line 1, character 0 (zero-based)
	result, err := f.Hover(context.Background(), doc, protocol.Position{Line: 1, Character: 0})
	if err != nil {
		t.Fatalf("Hover failed: %v", err)
	}
	if result == nil {
		t.Fatal("expected hover result, got nil")
	}

	want := "To configure an interface type and to enter interface configuration mode, use the interface command in the appropriate configuration mode."
	if result.Contents.Value != want {
		t.Errorf("hover value = %q\nwant %q", result.Contents.Value, want)
	}
	if result.Contents.Kind != "plaintext" {
		t.Errorf("hover kind = %q, want %q", result.Contents.Kind, "plaintext")
	}
}
