package main

import "fmt"

func main() {
	var bulan = [...]string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	bulan[5] = "mecel"
	fmt.Print(slice1)

	var slice2 = bulan[10:]
	fmt.Print(slice2)

	var slice3 = append(slice2, "Mecel 1234")
	fmt.Println(slice3)

	// Append jika sanggup menampung maka akan di replace, jika kapasitas array sudah full maka akan membuat array baru

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Mecel"
	newSlice[1] = "Araona"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
