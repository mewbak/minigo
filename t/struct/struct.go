package main

import "fmt"

func incr() {
	var u User = User{
		id:  0,
		age: 5,
	}

	//u.age = u.age + 1
	u.age++
	fmt.Printf("%d\n", u.age)
}

func decr() {
	var u User = User{
		id:  0,
		age: 8,
	}

	u.age--
	fmt.Printf("%d\n", u.age)
}

func main() {
	var i int = 1
	var u User = User{
		id:  3,
		age: 2,
	}
	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", u.age)
	fmt.Printf("%d\n", u.id)

	u.id = 4
	fmt.Printf("%d\n", u.id)

	u = User{id: 3, age: 5}
	fmt.Printf("%d\n", u.age)

	incr()
	decr()
}

type User struct {
	id  int
	age int
}
