package puppy

import "strings"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
