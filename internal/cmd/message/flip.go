package message

import "fmt"

func Flip(resource, content string) string {
	return fmt.Sprintf("🪙 A new **%s** has been rolled!\n\n%s\n",
		resource,
		content,
	)
}
