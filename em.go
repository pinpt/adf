package adf

func init() {
	registerMark("em", HTML, func(mark Mark, content string) (string, error) {
		return wrapHTMLTagContent("em", content)
	})
}
