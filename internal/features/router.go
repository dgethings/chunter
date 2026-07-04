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

func (r *Router) Close() error {
	var firstErr error
	for _, f := range r.features {
		if err := f.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	r.features = make(map[string]Feature)
	return firstErr
}
