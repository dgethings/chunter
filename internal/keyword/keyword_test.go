package keyword_test

import (
	"testing"

	"github.com/dgethings/chunter/internal/keyword"
)

func TestInSection_ExactMatch(t *testing.T) {
	kws := keyword.Keywords{
		{Keyword: "clock", Section: "config-if"},
		{Keyword: "hostname", Section: "config"},
	}
	got := kws.InSection("config-if")
	if len(got) != 1 {
		t.Fatalf("expected 1 keyword for config-if, got %d", len(got))
	}
	if got[0].Keyword != "clock" {
		t.Errorf("expected clock, got %q", got[0].Keyword)
	}
}

func TestInSection_EmptySectionIsUniversal(t *testing.T) {
	kws := keyword.Keywords{
		{Keyword: "do", Section: ""},
	}
	for _, section := range []string{"config", "config-if", "anything"} {
		got := kws.InSection(section)
		if len(got) != 1 {
			t.Errorf("InSection(%q): expected universal keyword, got %d results", section, len(got))
			continue
		}
		if got[0].Keyword != "do" {
			t.Errorf("InSection(%q): expected do, got %q", section, got[0].Keyword)
		}
	}
}

func TestInSection_ProductionDataHasUniversalKeywords(t *testing.T) {
	kws := keyword.Keywords{
		{Keyword: "hostname", Section: "config"},
		{Keyword: "do", Section: ""},
		{Keyword: "clock", Section: "config-if"},
	}
	got := kws.InSection("config")
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
