package keyword

import "github.com/dgethings/chunter/internal/protocol"

type Keyword struct {
	Description
	Section string
}

type Description struct {
	Format protocol.MarkupKind
	Value  string
}

type Set struct {
	entries map[string]Keyword
}

func NewSet(entries map[string]Keyword) *Set {
	return &Set{entries: entries}
}

func (s *Set) Lookup(nodeKind string) (Keyword, bool) {
	kw, ok := s.entries[nodeKind]
	return kw, ok
}

func (s *Set) InSection(section string) map[string]Keyword {
	in := map[string]Keyword{}
	for kw, item := range s.entries {
		if item.Section == section {
			in[kw] = item
		}
	}
	return in
}
