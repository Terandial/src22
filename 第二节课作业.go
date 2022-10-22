package main

import "fmt"

func main() {

	father := struct {
		name string
		age  int
	}{
		"Lothric",
		18,
	}
	fmt.Println(father.name)
	p := father.age
	q := &p

	fmt.Println(&q)

}
