package internal

func init() {
	registerNode("doc", HTML, func(node Node, next func() (string, error)) (string, error) {
		return next()
	})
}
