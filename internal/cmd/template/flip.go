package template

import "fmt"

func Flip(resource, content string) string {
	const pattern = `🪙 A new **%s** has been rolled!

%s
`

	return fmt.Sprintf(
		pattern,
		resource,
		content,
	)
}
