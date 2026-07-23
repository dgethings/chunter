package keyword

import "strings"

// SectionNode represents one node in the IOS configuration-mode hierarchy.
type SectionNode struct {
	ID       string
	Parent   string
	Children []string
}

// SectionTree models the parent-child relationships between IOS config modes.
// The root is always "config" (global configuration mode).
type SectionTree struct {
	nodes map[string]*SectionNode
}

// BuildSectionTree constructs the hierarchy from a set of known section IDs.
// For each section, the parent is found by progressively stripping the last
// "-segment" until another known section is found. "config" is always the
// root (parent = "").
//
// Example derivations:
//
//	config-if            → parent: config
//	config-router        → parent: config
//	config-router-af     → parent: config-router
//	config-pmap-c        → parent: config-pmap
//	config-archive-log-config → strip -config → config-archive-log (missing)
//	                             → strip -log → config-archive (found!) → parent: config-archive
func BuildSectionTree(sections []string) *SectionTree {
	known := make(map[string]bool, len(sections)+1)
	known["config"] = true
	for _, s := range sections {
		if s != "" {
			known[s] = true
		}
	}

	t := &SectionTree{
		nodes: make(map[string]*SectionNode, len(known)),
	}

	t.nodes["config"] = &SectionNode{ID: "config", Parent: ""}

	for s := range known {
		if s == "config" {
			continue
		}
		parent := findParent(s, known)
		t.nodes[s] = &SectionNode{ID: s, Parent: parent}
	}

	for s, node := range t.nodes {
		if node.Parent != "" {
			if p, ok := t.nodes[node.Parent]; ok {
				p.Children = append(p.Children, s)
			}
		}
	}

	return t
}

// findParent determines the parent section by progressively stripping the last
// "-segment" until a known section is found. Falls back to "config".
func findParent(section string, known map[string]bool) string {
	parts := strings.Split(section, "-")
	for i := len(parts) - 1; i > 1; i-- {
		candidate := strings.Join(parts[:i], "-")
		if known[candidate] {
			return candidate
		}
	}
	return "config"
}

// Parent returns the immediate parent section, or "" for the root.
func (t *SectionTree) Parent(section string) string {
	if node, ok := t.nodes[section]; ok {
		return node.Parent
	}
	return ""
}

// Ancestors returns the chain of parent sections from the given section up to
// and including the root. The section itself is not included. The root
// "config" has no ancestors and returns nil.
//
//	Ancestors("config-router-af-topology") → ["config-router-af", "config-router", "config"]
//	Ancestors("config")                    → nil
func (t *SectionTree) Ancestors(section string) []string {
	var result []string
	cur := t.Parent(section)
	for cur != "" {
		result = append(result, cur)
		cur = t.Parent(cur)
	}
	return result
}

// IsAncestor returns true if "ancestor" appears in the ancestor chain of
// "section" (or if ancestor == section).
func (t *SectionTree) IsAncestor(ancestor, section string) bool {
	if ancestor == section {
		return true
	}
	for _, a := range t.Ancestors(section) {
		if a == ancestor {
			return true
		}
	}
	return false
}

// NearestKnown walks up from "section" until it finds a section in the provided
// known set, returning that section. Returns "config" as fallback if no known
// ancestor is found. Used by completion when the grammar doesn't model the
// current section precisely.
func (t *SectionTree) NearestKnown(section string, known map[string]bool) string {
	if known[section] {
		return section
	}
	for _, ancestor := range t.Ancestors(section) {
		if known[ancestor] {
			return ancestor
		}
	}
	return "config"
}

// Children returns the direct children of the given section.
func (t *SectionTree) Children(section string) []string {
	if node, ok := t.nodes[section]; ok {
		return node.Children
	}
	return nil
}
