package greeting

import (
	"github.com/ahmetkca/golang-hello-world-modules/hello"
	"github.com/ahmetkca/golang-hello-world-modules/world"
)

func SayHi() string {
	return "Hi!"
}

func SayHelloWorld() string {
	return hello.Greet() + world.Greet()
}
