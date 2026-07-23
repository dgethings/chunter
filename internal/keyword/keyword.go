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
	byName        map[string]Keyword
	bySection     map[string][]Keyword
	byNameSection map[string]map[string]bool
	tree          *SectionTree
	all           []Keyword
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

	// Build the by-name-section index: for each keyword name, the set of
	// sections it is valid in (including "" for global keywords). This lets
	// IsValidInSection answer in O(1) without scanning the keyword list.
	s.byNameSection = make(map[string]map[string]bool)
	for _, kw := range kws {
		if s.byNameSection[kw.Keyword] == nil {
			s.byNameSection[kw.Keyword] = make(map[string]bool)
		}
		s.byNameSection[kw.Keyword][kw.Section] = true
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

// IsValidInSection returns true if the keyword name is known AND valid in the
// given section (either the keyword is global — Section == "" — or it has an
// entry whose Section matches). Returns false for unknown keywords.
func (s *Set) IsValidInSection(name, section string) bool {
	sections, ok := s.byNameSection[name]
	if !ok {
		return false
	}
	return sections[""] || sections[section]
}

// LookupSection returns the first non-empty Section for the given keyword
// name, or "" if the keyword is unknown or only global. Used to build
// diagnostic messages that point the user at where the keyword does belong.
func (s *Set) LookupSection(name string) string {
	for _, kw := range s.all {
		if kw.Keyword == name && kw.Section != "" {
			return kw.Section
		}
	}
	return ""
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
