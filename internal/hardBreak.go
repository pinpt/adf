package internal

func init() {
	registerNode("hardBreak", HTML, func(node Node, next func() (string, error)) (string, error) {
		return "<br/>", nil
	})
}
