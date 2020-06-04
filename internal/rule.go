package internal

func init() {
	registerNode("rule", HTML, func(node Node, next func() (string, error)) (string, error) {
		return "<hr/>", nil
	})
}
