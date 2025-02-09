package puppy

import "github.com/GoesToEleven/dog"

func Bark() string {
	return "woof"
}

func Barks() string {
	return "Woof, woof, woof"
}

func BigBark() string {
	return dog.WhenGrownUp("I am too old for this")
}