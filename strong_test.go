package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrong(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "text",
		"text": "Hello world",
		"marks": [
		  {
			"type": "strong"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<strong>Hello world</strong>`, html)
}
