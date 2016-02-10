package main

import "fmt"

func main() {
	/*
		var a int
		var b int

		fmt.Print("Enter a number:")
		fmt.Scan(&a)
		fmt.Print("Enter a second number:")
		fmt.Scan(&b)
		fmt.Println(a * b)
	*/

	/*
		a := 100
		var b *int = &a
		fmt.Println(*b)
		*b = 101
		fmt.Println(a)
		a = 102
		fmt.Println(*b)
	*/

	//a := 100
	//fmt.Println(a % 3)

	//Basic loop
	//for i := 0; i <= 100; i++ {
	//	fmt.Println(i)
	//}

	//for i := 0; i <= 100; i++ {
	//	fmt.Println(i)
	//}
	/*
		var b bool = true

		for b {

			for i := 0; i <= 100; i++ {
				fmt.Println(i)
				b = false
			}

		}
	*/

	/*
		char := 'S'
		fmt.Println(reflect.TypeOf(char))
		var char2 rune = 'S'
		fmt.Println(reflect.TypeOf(char2))
	*/

	/*
		myPhrase := "Pillar Hill Willow Kathy"
		for key, value := range myPhrase {
			fmt.Println(key, " - ", value, " - ", string(value))
		}
	*/

	/*
		myPhrase := "Pillar Hill Willow Kathy"
		for i := 0; i < len(myPhrase); i++ {
			fmt.Printf("%d - %d - %q - %b\n", i, myPhrase[i], myPhrase[i], myPhrase[i])
			//fmt.Printf("%d\n", i)
		}
	*/

	//reverse a
	//for i,j:= 0, len(a)-1; i < j; i,j = i+1,j-1 {
	//	a[i],a[j] = a[j],a[i]
	//}
	/*
		switch "X" {
				case "B":
					fmt.Println("A")
				case "C","X"
					fmt.Println("A")
					fallthrough
				case "D"
					fmt.Println("D")
		}
	*/

	//sample := "pillar hill willow kathy"
	//fmt.Println([]byte(sample))

	//slic := []string{"Derek", "Emily", "Alycia", "Lisa", "Heidi"}
	//fmt.Println(slic)
	//fmt.Println(slic[2:4])
	//printer(slic...)
	/*
		slicA := make([]int, 3, 5)
		slicA[0] = 5
		slicA[1] = 10
		slicA[2] = 15

		slicA = append(slicA, 20)
		for k, v := range slicA {
			fmt.Println(k, v)
		}

		fmt.Println(len(slicA))
		fmt.Println(cap(slicA))

		fmt.Println("******")

		mp := make(map[string]string, 10)
		mp["Derek"] = "WolfA"
		mp["Wolf"] = "DerekB"

		fmt.Println(mp["Derek"])

		for k, v := range mp {
			fmt.Println(k, v)
		}
		fmt.Println("******")

		//delete(mp, "Derek")

		//for k, v := range mp {
		//		fmt.Println(k, v)
		//}

		testMap := map[string]string{
			"derek":  "wolf",
			"alycia": "wallin",
			"lisa":   "mckenzie",
		}

		fmt.Println(testMap)

		//common ok idiom
		if _, exists := testMap["derek1"]; exists {
			delete(testMap, "derek1")

		} else {
			fmt.Println("val does not exists")
		}

		age := new([]int)
		fmt.Println(age)
		fmt.Println(*age)
		age1 := new(int)
		fmt.Println(age1)
		fmt.Println(*age1)

		var p *[]int = new([]int)      //allocates slice structure
		var v []int = make([]int, 100) //the slice v now refers to a new arrat of 100 ints
		fmt.Println(*p)
		fmt.Println(p)
		fmt.Println(v)

		//unnecessaryily complex
		var pp *[]int = new([]int)
		*pp = make([]int, 100, 100)
		fmt.Println(*pp)
		fmt.Println(pp)

		//idomatic
		vv := make([]int, 100)
		fmt.Println(vv)
	*/
	/*
		fmt.Println(half(1))
		fmt.Println(half(2))
		h, v := half(57)
		fmt.Println(h, v)

		numbers := []int{99, 44, 33, 24, 90, 100, 1000}
		fmt.Println(max(numbers...))

		fmt.Printf("%T\n", "string")
		fmt.Printf("%T\n", 99.99)
		fmt.Printf("%T\n", 10000)
	*/

	//Example of closure
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}

func printer(str ...string) {

	for k, v := range str {

		fmt.Println(k, v)

	}

}

func half(input int) (int, bool) {
	var result int = input / 2
	return result, (input%2 == 0)
}

func max(input ...int) int {
	var largest int
	for _, v := range input {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func makeEvenGenerator() func() int {
	var i int
	return func() int {
		i += 2
		return i
	}

}
