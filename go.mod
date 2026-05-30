module github.com/dgethings/chunter

go 1.26.3

require github.com/tree-sitter/go-tree-sitter v0.25.0

require (
	github.com/dgethings/chunter/tree-sitter-cisco_ios v0.0.0
	github.com/mattn/go-pointer v0.0.1 // indirect
)

replace github.com/dgethings/chunter/tree-sitter-cisco_ios => ./tree-sitter-cisco_ios
