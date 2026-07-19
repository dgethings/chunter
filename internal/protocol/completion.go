package protocol

type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionItem struct {
	Label               string              `json:"label"`
	Kind                *CompletionItemKind `json:"kind,omitempty"`
	Tags                []CompletionItemTag `json:"tags,omitempty"`
	Detail              string              `json:"detail,omitempty"`
	Documentation       any                 `json:"documentation,omitempty"`
	Deprecated          *bool               `json:"deprecated,omitempty"`
	Preselect           *bool               `json:"preselect,omitempty"`
	SortText            *string             `json:"sortText,omitempty"`
	FilterText          *string             `json:"filterText,omitempty"`
	InsertText          *string             `json:"insertText,omitempty"`
	InsertTextFormat    *InsertTextFormat   `json:"insertTextFormat,omitempty"`
	InsertTextMode      *InsertTextMode     `json:"insertTextMode,omitempty"`
	TextEdit            any                 `json:"textEdit,omitempty"` // nil | TextEdit | InsertReplaceEdit
	AdditionalTextEdits []TextEdit          `json:"additionalTextEdits,omitempty"`
	CommitCharacters    []string            `json:"commitCharacters,omitempty"`
	Command             *Command            `json:"command,omitempty"`
	Data                any                 `json:"data,omitempty"`
}

type CompletionItemKind int
type CompletionItemTag int
type InsertTextFormat int
type InsertTextMode int

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type Command struct {
	/**
	 * Title of the command, like `save`.
	 */
	Title string `json:"title"`

	/**
	 * The identifier of the actual command handler.
	 */
	Command string `json:"command"`

	/**
	 * Arguments that the command handler should be
	 * invoked with.
	 */
	Arguments []any `json:"arguments,omitempty"`
}

type CompletionList struct {
	IsIncomplete bool             `json:"isIncomplete"`
	Items        []CompletionItem `json:"items"`
}

const (
	CompletionItemKindText          = CompletionItemKind(1)
	CompletionItemKindMethod        = CompletionItemKind(2)
	CompletionItemKindFunction      = CompletionItemKind(3)
	CompletionItemKindConstructor   = CompletionItemKind(4)
	CompletionItemKindField         = CompletionItemKind(5)
	CompletionItemKindVariable      = CompletionItemKind(6)
	CompletionItemKindClass         = CompletionItemKind(7)
	CompletionItemKindInterface     = CompletionItemKind(8)
	CompletionItemKindModule        = CompletionItemKind(9)
	CompletionItemKindProperty      = CompletionItemKind(10)
	CompletionItemKindUnit          = CompletionItemKind(11)
	CompletionItemKindValue         = CompletionItemKind(12)
	CompletionItemKindEnum          = CompletionItemKind(13)
	CompletionItemKindKeyword       = CompletionItemKind(14)
	CompletionItemKindSnippet       = CompletionItemKind(15)
	CompletionItemKindColor         = CompletionItemKind(16)
	CompletionItemKindFile          = CompletionItemKind(17)
	CompletionItemKindReference     = CompletionItemKind(18)
	CompletionItemKindFolder        = CompletionItemKind(19)
	CompletionItemKindEnumMember    = CompletionItemKind(20)
	CompletionItemKindConstant      = CompletionItemKind(21)
	CompletionItemKindStruct        = CompletionItemKind(22)
	CompletionItemKindEvent         = CompletionItemKind(23)
	CompletionItemKindOperator      = CompletionItemKind(24)
	CompletionItemKindTypeParameter = CompletionItemKind(25)
	InsertTextFormatPlainText       = InsertTextFormat(1)
	InsertTextFormatSnippet         = InsertTextFormat(2)
	InsertTextModePlainText         = InsertTextMode(1)
	InsertTextModeSnippet           = InsertTextMode(2)
)
