package tree_sitter_dockerfile_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_dockerfile "github.com/sureshjoshi/tree-sitter-dockerfile/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_dockerfile.Language())
	if language == nil {
		t.Errorf("Error loading Dockerfile grammar")
	}
}
