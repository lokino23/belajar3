package main

import (
	"fmt"
)

func main() {
	dungeonName := "Mist of Tirna "
	dungeonTime := 30
	finishTime := 31
	originalKey := 15
	plusSatu := 1
	plusDua := 2
	plusTiga := 3
	var finalkey int
	finalkey = originalKey
	//originalScore := 10

	if finishTime < dungeonTime {
		fmt.Print("Congrats u Time : ", dungeonName)
	} else {
		fmt.Println("Sorry u depleted ur Key")
	}
	if finishTime < dungeonTime {
		fmt.Print("You + ")
		if dungeonTime%finishTime >= 13 {
			fmt.Println(plusTiga)
		} else if dungeonTime%finishTime >= 7 {
			fmt.Println(plusDua)
		} else if dungeonTime%finishTime >= 1 {
			fmt.Println(plusSatu)
		}
	}
	fmt.Println("Your Original Key  :", originalKey)
	fmt.Print("Your End Key Become : ")
	if dungeonTime%finishTime >= 13 {
		fmt.Println(finalkey + plusTiga)
		//untuk variable +1
	} else if dungeonTime%finishTime >= 7 {
		fmt.Println(finalkey + plusDua)
		//untuk variable +2
	} else if dungeonTime%finishTime >= 1 {
		fmt.Println(finalkey + plusSatu)
		//untuk variable +3
	}
	// max point 15 dungeon = 150 ( 10 every kenaikan key)
	point := 10
	if finishTime < dungeonTime {

		fmt.Print("You Get = ")
		if dungeonTime%finishTime >= 13 {
			fmt.Println(point*originalKey+30, " score ")
			//untuk +3
		} else if dungeonTime%finishTime >= 7 {
			fmt.Println(point*originalKey+20, " score ")
			//untuk +2
		} else if dungeonTime%finishTime >= 1 {
			fmt.Println(point*originalKey+10, " score ")
			//untuk +1
		}
	}
	/*if dungeonTime%finishTime >= 1 {
		fmt.Println(finalkey + plusSatu)
		//finalkey = originalKey + plusSatu
	}*/
	// Next Pengerjaan point every time doing keys
	// max point 15 dungeon = 150 ( 10 every kenaikan key)

	// ------------------------------------------------
	// Belajar Loop
	// 3 loops
	sum := 0
	for i := 2; i < 6; i++ {
		sum += i
	}
	fmt.Println(sum) // 1+2+3+4

	f := 1
	for f < 5 {
		f *= 2
	}
	fmt.Println(f) //

	/*rum := 0
	for {
		rum++ // forever??
	}
	fmt.Println(rum)*/

	// Conditional
	x := 5
	if x < 0 {
		fmt.Println("A")
	} else if x < 5 {
		fmt.Println("B")
	} else if x < 10 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

}

// Belajar Loop
// 3 loops
