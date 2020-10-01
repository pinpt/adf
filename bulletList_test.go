package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulletList(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "bulletList",
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
	assert.Equal(`<ul><li><p>Hello world</p></li></ul>`, html)
}
