package adf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubSupSubscript(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "text",
		"text": "Hello world",
		"marks": [
		  {
			"type": "subsup",
			"attrs": {
			  "type": "sub"
			}
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<sub>Hello world</sub>`, html)
}

func TestSubSupSuperscript(t *testing.T) {
	assert := assert.New(t)
	buf := []byte(`{
		"type": "text",
		"text": "Hello world",
		"marks": [
		  {
			"type": "subsup",
			"attrs": {
			  "type": "sup"
			}
		  }
		]
	  }`)
	var node Node
	assert.NoError(node.Parse(buf))
	html, err := node.Generate(HTML)
	assert.NoError(err)
	assert.Equal(`<sup>Hello world</sup>`, html)
}
