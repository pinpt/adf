package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParagraph(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "paragraph",
		"content": [
		  {
			"type": "text",
			"text": "Hello world"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<p>Hello world</p>`, html)
}
