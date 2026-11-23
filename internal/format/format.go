package format

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Renderable interface {
	Render() string
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}

	first, size := utf8.DecodeRuneInString(s)

	return string(unicode.ToUpper(first)) + s[size:]
}

func Number(n int) string {
	digits := [...]string{"0️⃣", "1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣"}

	var result strings.Builder
	for _, char := range strconv.Itoa(n) {
		if char == '-' {
			result.WriteRune(char)
			continue
		}

		result.WriteString(digits[char-'0'])
	}

	return result.String()
}

func List[T Renderable](l []T) string {
	var message strings.Builder
	for idx, i := range l {
		if idx > 0 {
			message.WriteString("\n")
		}

		fmt.Fprintf(&message, "%s %s", Number(idx+1), i.Render())
	}

	return message.String()
}
