package validate

import (
	"fmt"
	"strings"
	"unicode"
)

func ObjectName(key string) error {
	if strings.TrimSpace(key) == "" {
		return fmt.Errorf("object key cannot be empty")
	}

	if len(key) > 1024 {
		return fmt.Errorf("object key is too long")
	}

	if strings.Contains(key, "\\") {
		return fmt.Errorf("object key cannot contain backslash")
	}

	if strings.Contains(key, "..") {
		return fmt.Errorf("object key cannot contain '..'")
	}

	for _, r := range key {
		if unicode.IsControl(r) {
			return fmt.Errorf("object key must not contain control characters")
		}
	}

	return nil
}
