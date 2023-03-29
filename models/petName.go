package models

import (
	"flag"
	"math/rand"
	"strings"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
)

var (
	words     = flag.Int("words", 1, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func PetName() string {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	a := petname.Generate(*words, *separator)
	name := strings.Title(a)
	return name
}
