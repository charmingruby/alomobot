package set

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set/ucraine"
)

var Sets = map[booster.SetName]booster.Set{
	ucraine.Name: ucraine.New().Set,
}

func FindSet(name string) (booster.Set, bool) {
	selected, ok := Sets[booster.SetName(name)]

	return selected, ok
}
