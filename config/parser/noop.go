package parser

type NoopParser struct{}

func (n *NoopParser) Parse(rawData []byte) error {
	return nil
}

func (n *NoopParser) ConfigMap() map[string]any {
	return nil
}

func (p *NoopParser) Support(contentType string) bool {
	return false
}
