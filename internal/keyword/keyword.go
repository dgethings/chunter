package keyword

import "github.com/dgethings/chunter/internal/protocol"

type Keyword struct {
	Keyword string
	Description
	Section string
}

type Description struct {
	Format protocol.MarkupKind
	Value  string
}

type Keywords []Keyword

func (k Keywords) Lookup(name string) (Keyword, bool) {
	for _, kw := range k {
		if kw.Keyword == name {
			return kw, true
		}
	}
	return Keyword{}, false
}

func (k Keywords) InSection(section string) []Keyword {
	result := []Keyword{}
	for _, kw := range k {
		if kw.Section == section {
			result = append(result, kw)
		}
	}
	return result
}
