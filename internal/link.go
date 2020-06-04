package internal

import "strings"

func init() {
	registerMark("link", HTML, func(mark Mark, content string) (string, error) {
		href, ok := mark.Attrs["href"].(string)
		if !ok || href == "" {
			return "", nil
		}
		attrs := []string{`href="` + href + `"`}
		if title, ok := mark.Attrs["title"].(string); ok {
			attrs = append(attrs, `title="`+title+`"`)
		}
		return wrapHTMLTagContent("a", content, strings.Join(attrs, " "))
	})
}
