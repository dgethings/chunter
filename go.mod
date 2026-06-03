module github.com/dgethings/chunter

go 1.26.3

require (
	github.com/creachadair/jrpc2 v1.3.5
	github.com/dgethings/chunter/grammars/tree-sitter-cisco_ios v0.0.0
	github.com/spf13/cobra v1.10.2
	github.com/tree-sitter/go-tree-sitter v0.25.0
)

require (
	github.com/creachadair/mds v0.26.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-pointer v0.0.1 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/sync v0.19.0 // indirect
)

replace github.com/dgethings/chunter/grammars/tree-sitter-cisco_ios => ./grammars/tree-sitter-cisco_ios
