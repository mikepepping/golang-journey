package main
import "github.com/mikepepping/golang-journey/package-handling/greet"

func main() {
	greet.Hello()
	greet.hello() // Should not compile, hello is not exported
}
