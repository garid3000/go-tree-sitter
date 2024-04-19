package org_nvim_test

import (
	"context"
	"testing"

	sitter "github.com/garid3000/go-tree-sitter"
	"github.com/garid3000/go-tree-sitter/org_nvim"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("* asdf"), org_nvim.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document subsection: (section headline: (headline stars: (stars) item: (item (expr)))))",
		n.String(),
	)
}
