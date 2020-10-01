package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEm(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "text",
		"text": "Hello world",
		"marks": [
		  {
			"type": "em"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<em>Hello world</em>`, html)
}
