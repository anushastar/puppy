package puppy

import "github.com/anushastar/dog"
import fmt

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof woof woof!!!"
}

func Bigbark() string {
	return dog.Grownup(Bark())
}

func t1() {
	fmt.Printf("abc")
}
