package parser

import "strings"

// HumanizeLabel humanizes label names to better visualization on email template.
func HumanizeLabel(key string) string {
	switch key {
	case "acompanantes":
		return "Acompañantes"
	default:
		return strings.Title(strings.Replace(key, "_", " ", -1))
	}
}
