package adf

func init() {
	registerMark("strike", HTML, func(mark Mark, content string) (string, error) {
		return wrapHTMLTagContent("s", content)
	})
}
