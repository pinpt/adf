package internal

func init() {
	registerMark("underline", HTML, func(mark Mark, content string) (string, error) {
		return wrapHTMLTagContent("u", content)
	})
}
