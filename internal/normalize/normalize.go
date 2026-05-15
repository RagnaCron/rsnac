// Package normalize
package normalize

import (
	"path/filepath"
	"strings"
)

func ToSnakeCase(filename string) string {
	lower := strings.ToLower(filename)
	ext := filepath.Ext(lower)
	s := strings.TrimSuffix(lower, ext)
	return normalize(s) + ext
}

func normalize(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	lastUnderscore := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch {
		case c >= 'a' && c <= 'z':
			b.WriteByte(c)
			lastUnderscore = false

		case c >= '0' && c <= '9':
			b.WriteByte(c)
			lastUnderscore = false

		default:
			if !lastUnderscore {
				b.WriteByte('_')
				lastUnderscore = true
			}
		}
	}

	out := b.String()
	return strings.Trim(out, "_")
}
