package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnderline(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "text",
		"text": "Hello world",
		"marks": [
		  {
			"type": "underline"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<u>Hello world</u>`, html)
}
