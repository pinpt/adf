package internal

import (
	"encoding/json"
	"fmt"
	"strings"
)

// schema https://unpkg.com/@atlaskit/adf-schema@10.0.0/dist/json-schema/v1/full.json

// ContentType definition
type ContentType string

// HTML content type
var HTML ContentType = "text/html"

// NodeVisitor is a visitor interface for handling node transformation
type NodeVisitor interface {
	Accept(contentType ContentType, node Node) bool
	Visit(contentType ContentType, node Node, next func() (string, error)) (string, error)
}

// MarkVisitor is a visitor interface  for handling mark transformation
type MarkVisitor interface {
	Accept(contentType ContentType, mark Mark) bool
	Visit(contentType ContentType, mark Mark, content string) (string, error)
}

// Mark is a mark detail
type Mark struct {
	Type  string                 `json:"type"`
	Attrs map[string]interface{} `json:"attrs"`
}

// Node is the base primitive for all node types
type Node struct {
	Type    string                 `json:"type"`
	Content []Node                 `json:"content"`
	Text    string                 `json:"text"`
	Marks   []Mark                 `json:"marks"`
	Attrs   map[string]interface{} `json:"attrs"`
}

// Parse will parse ADF formatted content into Node
func (n *Node) Parse(buf []byte) error {
	return json.Unmarshal(buf, &n)
}

// Generate will generate the content by contentType
func (n *Node) Generate(contentType ContentType) (string, error) {
	var builder strings.Builder
	for _, visitor := range nodeVisitors {
		if visitor.Accept(contentType, *n) {
			next := func() (string, error) {
				var builder strings.Builder
				for _, child := range n.Content {
					val, err := child.Generate(contentType)
					if err != nil {
						return "", err
					}
					builder.WriteString(val)
				}
				return builder.String(), nil
			}
			val, err := visitor.Visit(contentType, *n, next)
			if err != nil {
				return "", err
			}
			builder.WriteString(val)
		}
	}
	if len(n.Marks) > 0 {
		wrappedContent := builder.String()
		for _, mark := range n.Marks {
			for _, mvisitor := range markVisitors {
				if mvisitor.Accept(contentType, mark) {
					val, err := mvisitor.Visit(contentType, mark, wrappedContent)
					if err != nil {
						return "", err
					}
					wrappedContent = val
				}
			}
		}
		return wrappedContent, nil
	}
	return builder.String(), nil
}

type nodeVisitor struct {
	contentType ContentType
	typename    string
	renderer    func(node Node, next func() (string, error)) (string, error)
}

var _ NodeVisitor = (*nodeVisitor)(nil)

func (n *nodeVisitor) Accept(contentType ContentType, node Node) bool {
	return n.typename == node.Type
}

func (n *nodeVisitor) Visit(contentType ContentType, node Node, next func() (string, error)) (string, error) {
	return n.renderer(node, next)
}

type markVisitor struct {
	contentType ContentType
	markname    string
	renderer    func(mark Mark, content string) (string, error)
}

func (m *markVisitor) Accept(contentType ContentType, mark Mark) bool {
	return contentType == m.contentType && mark.Type == m.markname
}

func (m *markVisitor) Visit(contentType ContentType, mark Mark, content string) (string, error) {
	return m.renderer(mark, content)
}

var nodeVisitors = make([]NodeVisitor, 0)
var markVisitors = make([]MarkVisitor, 0)

func registerNode(typename string, contentType ContentType, renderer func(node Node, next func() (string, error)) (string, error)) {
	nodeVisitors = append(nodeVisitors, &nodeVisitor{contentType, typename, renderer})
}

func registerMark(markname string, contentType ContentType, renderer func(mark Mark, content string) (string, error)) {
	markVisitors = append(markVisitors, &markVisitor{contentType, markname, renderer})
}

func wrapHTMLTag(tag string, next func() (string, error), attrs ...string) (string, error) {
	val, err := next()
	if err != nil {
		return "", err
	}
	var attr string
	if len(attrs) > 0 {
		attr = " " + strings.Join(attrs, " ")
	}
	return fmt.Sprintf("<%[1]s%[3]s>%[2]s</%[1]s>", tag, val, attr), nil
}

func wrapHTMLTagContent(tag string, val string, attrs ...string) (string, error) {
	var attr string
	if len(attrs) > 0 {
		attr = " " + strings.Join(attrs, " ")
	}
	return fmt.Sprintf("<%[1]s%[3]s>%[2]s</%[1]s>", tag, val, attr), nil
}
