package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedList(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "orderedList",
		"attrs": {
		  "order": 3
		},
		"content": [
		  {
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
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<ol><li><p>Hello world</p></li></ol>`, html)
}
