package main

import (
	"fmt"
	"password/account"
	"password/utils"
)

func main() {
	for {
		utils.Menu()
		choice := utils.PromptDataInt()
		if choice == 1 {
			account.CreateAccount()
		} else if choice == 2 {
			utils.SearchByChoice()
			choice := utils.PromptDataInt()
			if choice == 1 {
				fmt.Println(account.FindAccountByName())
			} else if choice == 2 {
				fmt.Println(account.FindAccountByPassword())
			} else if choice == 3 {
				fmt.Println(account.FindAccountByUrl())
			}
		} else if choice == 3 {
			account.DeleteAccount()
		} else if choice == 4 {
			fmt.Println("Работа программы завершена")
			break
		}

	}
}
