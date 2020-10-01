package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRule(t *testing.T) {
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
				"text": "Hello world"
			  }
			]
		  },
		  {
			"type": "rule"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<p>Hello world</p><hr/>`, html)
}
