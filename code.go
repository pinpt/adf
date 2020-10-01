package adf

func init() {
	registerMark("code", HTML, func(mark Mark, content string) (string, error) {
		return wrapHTMLTagContent("code", content)
	})
}
