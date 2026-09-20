package placeholder

import (
	"fmt"
	"regexp"
	"strings"
)

// pattern recognizes ${NAME} and ${NAME:=default} placeholders.
var pattern = regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)(?::=([^}]*))?\}`)

// Expand replaces placeholders with values supplied by the caller.
// A placeholder without a value or default returns an error.
func Expand(command string, values map[string]string) (string, error) {
	var missing string
	result := pattern.ReplaceAllStringFunc(command, func(match string) string {
		parts := pattern.FindStringSubmatch(match)
		if value, ok := values[parts[1]]; ok {
			return value
		}
		if strings.Contains(match, ":=") {
			return parts[2]
		}
		missing = parts[1]
		return match
	})
	if missing != "" {
		return "", fmt.Errorf("missing value for placeholder %q", missing)
	}
	return result, nil
}
