package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLink(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "text",
		"text": "Hello world",
		"marks": [
		  {
			"type": "link",
			"attrs": {
			  "href": "http://atlassian.com",
			  "title": "Atlassian"
			}
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<a href="http://atlassian.com" title="Atlassian">Hello world</a>`, html)
}
