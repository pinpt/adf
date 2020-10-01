package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockQuote(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "blockquote",
		"content": [
		  {
			"type": "paragraph",
			"content": [
			  {
				"type": "text",
				"text": "Hello world"
			  }
			]
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<blockquote><p>Hello world</p></blockquote>`, html)
}
