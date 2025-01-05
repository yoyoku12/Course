package main

import (
	"fmt"
	"math/rand"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.password)
}

func (acc *account) outputAllInfo() {
	fmt.Println(acc)
}

func (acc *account) generatePassword(n int) {
	var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!")
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, url string) *account {
	return &account{
		url:      url,
		login:    login,
		password: password,
	}
}

func main() {
	var myAccount account
	insertData(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}

func insertData(acc *account) {
	acc.login = promptData("Введите логин: ")
	acc.url = promptData("Введите URL: ")
	fmt.Println("Ваш пароль будет сгенерирован автоматически.")
	acc.generatePassword(12)
	fmt.Println("Запишите ваш новый пароль:", acc.password)
}
