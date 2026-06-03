package ast

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func QueryLanguage(lang *sitter.Language, query string) (*sitter.Query, error) {
	q, qerr := sitter.NewQuery(lang, query)
	if qerr != nil {
		return nil, qerr
	}
	return q, nil
}
