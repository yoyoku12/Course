package utils

import "fmt"

func PromptDataWithText(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}

func PromptDataString() string {
	var result string
	fmt.Scanln(&result)
	return result
}

func PromptDataInt() int {
	var result int
	fmt.Scanln(&result)
	return result
}

func Menu() {
	fmt.Println("=====")
	fmt.Println("Меню")
	fmt.Println("=====")
	fmt.Println("1 - Создать аккаунт")
	fmt.Println("2 - Найти аккаунт")
	fmt.Println("3 - Удалить аккаунт")
	fmt.Println("4 - Выйти")
}

func SearchByChoice() {
	fmt.Println("Каким образом вы хотите найти аккаунт?")
	fmt.Println("1 - По имени")
	fmt.Println("2 - По паролю")
	fmt.Println("3 - По URL")
}
