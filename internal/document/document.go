package document

type Document struct {
	URI        string
	LanguageID string
	Version    int
	Content    []byte
	Lines      []string
}

func New(uri, languageID string, version int, content []byte) *Document {
	return &Document{
		URI:        uri,
		LanguageID: languageID,
		Version:    version,
		Content:    content,
	}
}
