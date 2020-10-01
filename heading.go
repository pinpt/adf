package adf

import "fmt"

func init() {
	registerNode("heading", HTML, func(node Node, next func() (string, error)) (string, error) {
		level := 1
		if l, ok := node.Attrs["level"]; ok {
			level = int(l.(float64))
		}
		return wrapHTMLTag(fmt.Sprintf("h%d", level), next)
	})
}
