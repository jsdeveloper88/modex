package modex

//https://habr.com/ru/post/421411/

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s !", name)
}
