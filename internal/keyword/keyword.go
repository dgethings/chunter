package keyword

import "github.com/dgethings/chunter/internal/protocol"

type Keyword struct {
	Keyword string
	Description
	Section  string
	Snippets []string
}

type Description struct {
	Format protocol.MarkupKind
	Value  string
}

type Keywords []Keyword

// Set provides O(1) keyword lookups by name and by section. Build it once at
// startup via NewSet from the full keyword slice.
type Set struct {
	byName    map[string]Keyword
	bySection map[string][]Keyword
	tree      *SectionTree
	all       []Keyword
}

// NewSet builds a Set from a keyword slice, indexing by name and section in a
// single pass. Keywords whose Section is "" are global and are returned by
// every InSection query. When two entries share the same Keyword field, the
// last one wins in the by-name index.
func NewSet(kws []Keyword) *Set {
	s := &Set{
		byName:    make(map[string]Keyword, len(kws)),
		bySection: make(map[string][]Keyword),
		all:       kws,
	}
	for _, kw := range kws {
		s.byName[kw.Keyword] = kw
		s.bySection[kw.Section] = append(s.bySection[kw.Section], kw)
	}

	// Build the section hierarchy tree from all known sections.
	var sections []string
	for sec := range s.bySection {
		if sec != "" {
			sections = append(sections, sec)
		}
	}
	s.tree = BuildSectionTree(sections)

	return s
}

func (s *Set) Lookup(name string) (Keyword, bool) {
	kw, ok := s.byName[name]
	return kw, ok
}

func (s *Set) InSection(section string) []Keyword {
	result := append([]Keyword{}, s.bySection[""]...)
	result = append(result, s.bySection[section]...)
	return result
}

// SectionTree returns the section hierarchy tree built from the keyword data.
func (s *Set) SectionTree() *SectionTree {
	return s.tree
}

// SectionsWithKeywords returns the set of section IDs that have at least one
// keyword (excluding the global "" section).
func (s *Set) SectionsWithKeywords() map[string]bool {
	result := make(map[string]bool, len(s.bySection))
	for sec, kws := range s.bySection {
		if sec != "" && len(kws) > 0 {
			result[sec] = true
		}
	}
	return result
}
