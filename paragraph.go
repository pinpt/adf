package adf

func init() {
	registerNode("paragraph", HTML, func(node Node, next func() (string, error)) (string, error) {
		return wrapHTMLTag("p", next)
	})
}
