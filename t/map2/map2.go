package main

import "fmt"

func f1() {
	var lmap map[int]int = map[int]int{
		1: 2,
		3: 4,
	}

	for i, v := range lmap {
		fmt.Printf("%d\n", i)
		fmt.Printf("%d\n", v)
	}

	fmt.Printf("%d\n", lmap[1]+3) // 5
	fmt.Printf("%d\n", lmap[3]+2) // 6

	lmap[7] = 8
	fmt.Printf("%d\n", lmap[4]+7) // 7
	fmt.Printf("%d\n", lmap[7])   // 8
}

func f2() {

	var lmap map[string]string = map[string]string{
		"15": "10",
		"17": "11",
	}

	fmt.Printf("9%s\n", lmap["noexists"])
	fmt.Printf("%s\n", lmap["15"]) // 10
	fmt.Printf("%s\n", lmap["17"]) // 11

	fmt.Printf("%d\n", len(lmap)+10) // 12

	var lenmap int

	lmap["19"] = "13"
	v19, ok := lmap["19"]
	if ok {
		fmt.Printf("%s\n", v19) // 13
	} else {
		fmt.Printf("ERROR\n")
	}

	lenmap = len(lmap)            // 3
	fmt.Printf("%d\n", lenmap+11) // 14

	lmap["15"] = "16"
	lmap["17"] = "18"
	lmap["19"] = "20"
	for k, v := range lmap {
		fmt.Printf("%s\n%s\n", k, v) // 15,16,17,18,19,20
	}
}

func f3() {
	var lmap map[int]int = map[int]int{
		1: 2,
		3: 21,
	}
	var ok bool
	var val int
	val, ok = lmap[3]
	fmt.Printf("%d\n", val) // 21
	if ok {
		fmt.Printf("%d\n", 22)
	}

	val, ok = lmap[2]
	if !ok {
		fmt.Printf("%d\n", 23)
	}
	fmt.Printf("%d\n", val+24) //24
}

var gkeyFoo2 string = "keyfoo"

func f4() {
	gkeyFoo := "keyfoo"

	var lmap map[string]string = map[string]string{
		gkeyFoo:  "26",
		"keybar": "valuebar",
	}

	var ok bool
	var v string
	v, ok = lmap[string(gkeyFoo2)]
	if ok {
		fmt.Printf("%d\n", 25)
		fmt.Printf("%s\n", v) // 26
	} else {
		fmt.Printf("ERROR\n")
	}

	v, ok = lmap["noexits"]
	if !ok {
		fmt.Printf("%d\n", 27)
	}
	fmt.Printf("28%s\n", v) // 28
}

func main() {
	f1()
	f2()
	f3()
	f4()
}
