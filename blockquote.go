package adf

func init() {
	registerNode("blockquote", HTML, func(node Node, next func() (string, error)) (string, error) {
		return wrapHTMLTag("blockquote", next)
	})
}
