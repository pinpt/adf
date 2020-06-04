package internal

func init() {
	registerMark("strong", HTML, func(mark Mark, content string) (string, error) {
		return wrapHTMLTagContent("strong", content)
	})
}
