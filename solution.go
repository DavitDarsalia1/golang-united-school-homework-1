package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	return emoji.Sprintf("Hello %q!", ":world_map:")
}
