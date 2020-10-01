package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrike(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "text",
		"text": "Hello world",
		"marks": [
		  {
			"type": "strike"
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<s>Hello world</s>`, html)
}
