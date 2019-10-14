package adjective

import (
	"math/rand"
	"time"
)

// Get returns a singular adjective
func Get() string {
	rand.Seed(time.Now().Unix())
	words := words()
	n := rand.Int() % len(words)
	return words[n]
}

func words() []string {
	return []string{
		"astonishing",
		"astounding",
		"blindsiding",
		"dumbfounding",
		"eye-opening",
		"flabbergasting",
		"jarring",
		"jaw-dropping",
		"jolting",
		"shocking",
		"startling",
		"stunning",
		"stupefying",
		"surprising",
		"amazing",
		"awesome",
		"fabulous",
		"marvelous",
		"miraculous",
		"portentous",
		"prodigious",
		"staggering",
		"stupendous",
		"sublime",
		"wondrous",
		"incomprehensible",
		"inconceivable",
		"incredible",
		"unbelievable",
		"unimaginable",
		"unthinkable",
		"extraordinary",
		"phenomenal",
		"rare",
		"sensational",
		"spectacular",
		"uncommon",
		"unique",
		"unwonted",
		"conspicuous",
		"notable",
		"noticeable",
		"outstanding",
		"remarkable",
		"impressive",
		"smashing",
		"striking",
	}
}
