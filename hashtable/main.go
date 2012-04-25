package main

import (
	"fmt"
)

func main() {
	fmt.Println("hash(1)=", hash(1, P))
	msgs := []string{"hola que tal", "hola", "Hola", "Hola "}
	for _, s := range msgs {
		fmt.Println("hash(hashcode('"+s+"'),P)=", hash(hashcode(s), P))
	}
}
