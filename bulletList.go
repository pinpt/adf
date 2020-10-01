package adf

func init() {
	registerNode("bulletList", HTML, func(node Node, next func() (string, error)) (string, error) {
		return wrapHTMLTag("ul", next)
	})
}
