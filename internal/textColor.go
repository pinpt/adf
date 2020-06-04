package internal

import (
	"fmt"
	"strings"
)

func init() {
	registerMark("textColor", HTML, func(mark Mark, content string) (string, error) {
		color, ok := mark.Attrs["color"].(string)
		if !ok || color == "" {
			return "", fmt.Errorf("missing color attribute")
		}
		attrs := []string{`style="color:` + color + `"`}
		return wrapHTMLTagContent("span", content, strings.Join(attrs, " "))
	})
}
