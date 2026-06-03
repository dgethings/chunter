package features

import "fmt"

type Router struct {
	features map[string]Feature
}

func NewRouter() *Router {
	return &Router{
		features: make(map[string]Feature),
	}
}

func (r *Router) Register(f Feature) {
	r.features[f.LanguageID()] = f
}

func (r *Router) Route(languageID string) (Feature, error) {
	f, ok := r.features[languageID]
	if !ok {
		return nil, fmt.Errorf("no feature registered for language: %s", languageID)
	}
	return f, nil
}
