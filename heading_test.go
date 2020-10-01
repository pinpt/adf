package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeading(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "heading",
		"attrs": {
		  "level": 1
		},
		"content": [
		  {
			"type": "text",
			"text": "Heading 1"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<h1>Heading 1</h1>`, html)
}
