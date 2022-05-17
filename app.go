package main

import (
	"fmt"
)

func main() {
	dungeonName := "Mist of Tirna "
	dungeonTime := 30
	finishTime := 20
	originalKey := 15
	plusSatu := 1
	plusDua := 2
	plusTiga := 3
	var finalkey int
	finalkey = originalKey
	sum := 0
	//originalScore := 10

	if finishTime < dungeonTime {
		fmt.Print("Congrats u Time : ", dungeonName)
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
			fmt.Print(point*originalKey+30, " score")
			//untuk +3
		} else if dungeonTime%finishTime >= 7 {
			fmt.Print(point*originalKey+20, " score")
			//untuk +2
		} else if dungeonTime%finishTime >= 1 {
			fmt.Print(point*originalKey+10, " score")
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

	for i := 1; i < 6; i++ {
		sum += i
	}
	fmt.Println(sum)

}

// Belajar Loop
// 3 loops
