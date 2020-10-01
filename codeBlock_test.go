package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeBlock(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "codeBlock",
		"attrs": {
		  "language": "javascript"
		},
		"content": [
		  {
			"type": "text",
			"text": "var foo = {};var bar = [];"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<code language=javascript>var foo = {};var bar = [];</code>`, html)
}
