package internal

import "fmt"

func init() {
	registerMark("subsup", HTML, func(mark Mark, content string) (string, error) {
		switch mark.Attrs["type"].(string) {
		case "sup":
			return wrapHTMLTagContent("sup", content)
		case "sub":
			return wrapHTMLTagContent("sub", content)
		}
		return "", fmt.Errorf("error unsupported subsup type value")
	})
}
