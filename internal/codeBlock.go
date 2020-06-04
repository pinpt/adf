package internal

func init() {
	registerNode("codeBlock", HTML, func(node Node, next func() (string, error)) (string, error) {
		var attrs string
		if lang, ok := node.Attrs["language"]; ok {
			attrs = "language=" + lang.(string)
		}
		return wrapHTMLTag("code", next, attrs)
	})
}
