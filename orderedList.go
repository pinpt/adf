package adf

func init() {
	registerNode("orderedList", HTML, func(node Node, next func() (string, error)) (string, error) {
		return wrapHTMLTag("ol", next)
	})
}
