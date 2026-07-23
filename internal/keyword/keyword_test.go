package keyword_test

import (
	"testing"

	"github.com/dgethings/chunter/internal/keyword"
)

func TestSetLookup(t *testing.T) {
	kws := []keyword.Keyword{
		{Keyword: "clock", Section: "config-if"},
		{Keyword: "hostname", Section: "config"},
	}
	s := keyword.NewSet(kws)

	kw, ok := s.Lookup("clock")
	if !ok {
		t.Fatal("expected to find clock")
	}
	if kw.Keyword != "clock" {
		t.Errorf("expected clock, got %q", kw.Keyword)
	}
	if kw.Section != "config-if" {
		t.Errorf("expected section config-if, got %q", kw.Section)
	}

	if _, ok := s.Lookup("nope"); ok {
		t.Error("expected Lookup to return false for a missing keyword")
	}
}

func TestSetInSection_ExactMatch(t *testing.T) {
	kws := []keyword.Keyword{
		{Keyword: "clock", Section: "config-if"},
		{Keyword: "hostname", Section: "config"},
	}
	s := keyword.NewSet(kws)
	got := s.InSection("config-if")
	if len(got) != 1 {
		t.Fatalf("expected 1 keyword for config-if, got %d", len(got))
	}
	if got[0].Keyword != "clock" {
		t.Errorf("expected clock, got %q", got[0].Keyword)
	}
}

func TestSetInSection_EmptySectionIsUniversal(t *testing.T) {
	kws := []keyword.Keyword{
		{Keyword: "do", Section: ""},
	}
	s := keyword.NewSet(kws)
	for _, section := range []string{"config", "config-if", "anything"} {
		got := s.InSection(section)
		if len(got) != 1 {
			t.Errorf("InSection(%q): expected universal keyword, got %d results", section, len(got))
			continue
		}
		if got[0].Keyword != "do" {
			t.Errorf("InSection(%q): expected do, got %q", section, got[0].Keyword)
		}
	}
}

func TestSetInSection_ProductionDataHasUniversalKeywords(t *testing.T) {
	kws := []keyword.Keyword{
		{Keyword: "hostname", Section: "config"},
		{Keyword: "do", Section: ""},
		{Keyword: "clock", Section: "config-if"},
	}
	s := keyword.NewSet(kws)
	got := s.InSection("config")
	labels := map[string]bool{}
	for _, kw := range got {
		labels[kw.Keyword] = true
	}
	if !labels["hostname"] {
		t.Errorf("expected config keyword hostname in config section")
	}
	if !labels["do"] {
		t.Errorf("expected empty-Section keyword do to appear in config section")
	}
	if labels["clock"] {
		t.Errorf("config-if keyword clock must not appear in config section")
	}
}

func TestSetLookupOverwrite(t *testing.T) {
	// When two entries share the same Keyword field, the last one wins in the
	// by-name index built by NewSet.
	kws := []keyword.Keyword{
		{Keyword: "dup", Section: "config"},
		{Keyword: "dup", Section: "config-if"},
	}
	s := keyword.NewSet(kws)

	kw, ok := s.Lookup("dup")
	if !ok {
		t.Fatal("expected to find dup")
	}
	if kw.Section != "config-if" {
		t.Errorf("expected last entry to win (config-if), got %q", kw.Section)
	}
}

func TestIsValidInSection(t *testing.T) {
	kws := []keyword.Keyword{
		{Keyword: "foo", Section: "config-if"},
		{Keyword: "bar", Section: ""}, // global
	}
	s := keyword.NewSet(kws)

	cases := []struct {
		name    string
		section string
		want    bool
	}{
		{"foo", "config-if", true},      // exact section match
		{"foo", "config-router", false}, // known but wrong section
		{"bar", "anything", true},       // global valid everywhere
		{"baz", "config", false},        // unknown keyword
	}
	for _, tc := range cases {
		got := s.IsValidInSection(tc.name, tc.section)
		if got != tc.want {
			t.Errorf("IsValidInSection(%q, %q) = %v, want %v", tc.name, tc.section, got, tc.want)
		}
	}
}

func TestLookupSection(t *testing.T) {
	kws := []keyword.Keyword{
		{Keyword: "foo", Section: "config-if"},
		{Keyword: "bar", Section: ""}, // global only
		{Keyword: "baz", Section: "config-router"},
	}
	s := keyword.NewSet(kws)

	if got := s.LookupSection("foo"); got != "config-if" {
		t.Errorf("LookupSection(foo) = %q, want config-if", got)
	}
	if got := s.LookupSection("baz"); got != "config-router" {
		t.Errorf("LookupSection(baz) = %q, want config-router", got)
	}
	// bar is global only -> no non-empty section to report.
	if got := s.LookupSection("bar"); got != "" {
		t.Errorf("LookupSection(bar) = %q, want \"\"", got)
	}
	// Unknown keyword.
	if got := s.LookupSection("missing"); got != "" {
		t.Errorf("LookupSection(missing) = %q, want \"\"", got)
	}
}
