package pkmn

import "math/rand"

const NationalDexCount = 1025

func RandomNationalDexID() int {
	return rand.Intn(NationalDexCount)
}
