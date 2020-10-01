package adf

func init() {
	registerNode("text", HTML, func(node Node, next func() (string, error)) (string, error) {
		return node.Text, nil
	})
}
