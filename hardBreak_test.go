package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHardBreak(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"version": 1,
		"type": "doc",
		"content": [
		  {
			"type": "paragraph",
			"content": [
			  {
				"type": "text",
				"text": "Hello"
			  },
			  {
				"type": "hardBreak"
			  },
			  {
				"type": "text",
				"text": "world"
			  }
			]
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<p>Hello<br/>world</p>`, html)
}
