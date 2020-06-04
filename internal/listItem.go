package internal

func init() {
	registerNode("listItem", HTML, func(node Node, next func() (string, error)) (string, error) {
		return wrapHTMLTag("li", next)
	})
}
