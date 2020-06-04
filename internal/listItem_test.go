package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListItem(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "listItem",
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
	assert.Equal(`<li><p>Hello world</p></li>`, html)
}
