package main

import (
	"Lab5Go/task1"
	"Lab5Go/task2"
	"Lab5Go/task3"
	"Lab5Go/task4"
	"fmt"
)

func main() {
	fmt.Println("Выберите задание:")
	fmt.Println("1 задание")
	fmt.Println("2 задание")
	fmt.Println("3 задание")
	fmt.Println("4 задание")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		task1.RunTask1()
	case 2:
		task2.RunTask2()
	case 3:
		task3.RunTask3()
	case 4:
		task4.RunTask4()
	default:
		fmt.Println("Нужно выбрать задание")
	}
}
