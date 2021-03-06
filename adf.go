package adf

// GenerateHTMLFromADF will generate HTML from the provided ADF formatted buffer
func GenerateHTMLFromADF(buf []byte) (string, error) {
	var node Node
	if err := node.Parse(buf); err != nil {
		return "", err
	}
	return node.Generate(HTML)
}
