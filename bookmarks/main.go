package main

import (
	"fmt"
)

type stringMap = map[string]string // создаем алиас мап стринг стринг для удобства.

func main() {
	m := stringMap{
		"PurpleSchool": "https://purpleschool.ru",
		"Yandex":       "https://yandex.ru",
		"Google":       "https://google.com",
		"Brave":        "https://brave.com",
		"Dota2":        "https://Dota2.com",
		"CS":           "https://CS.com",
	}

	fmt.Println("___Приложение для создания закладок____")
Menu:
	for {
		getMenu()
		choice := readInputInt()
		if choice == 1 {
			fmt.Println("====================================")
			fmt.Println("Список добавленных вкладок")
			fmt.Println("====================================")
			showAllBookmarks(m)
		} else if choice == 2 {
			createBookmark(m)
		} else if choice == 3 {
			deleteBookmark(m)
		} else if choice == 4 {
			fmt.Println("====================================")
			fmt.Println("Выход из приложения")
			fmt.Println("====================================")
			break Menu
		} else {
			fmt.Println("Вы ввели число не из списка")
			fmt.Println("Попробуйте еще раз")
			continue
		}
	}
}

func readInputInt() int {
	var input int
	fmt.Scan(&input)
	return input
}

func readInputString() string {
	var input string
	fmt.Scan(&input)
	return input
}

func createBookmark(m stringMap) stringMap {
	var name string
	var site string
	fmt.Println("====================================")
	fmt.Println("Введите наименование закладки: ")
	fmt.Println("====================================")
	name = readInputString()
	fmt.Println("====================================")
	fmt.Println("Введите сайт связанный с именем закладки: ")
	fmt.Println("====================================")
	site = readInputString()
	m[name] = site
	fmt.Println("====================================")
	fmt.Println("Закладка успешно создана!")
	fmt.Println("====================================")
	return m
}

func deleteBookmark(m stringMap) stringMap {
	fmt.Println("====================================")
	fmt.Println("Вот полный список имеющихся закладок:")
	showAllBookmarks(m)
	fmt.Println("====================================")
	fmt.Println("====================================")
	fmt.Println("Введите имя закладки для удаления:")
	fmt.Println("====================================")
	name := readInputString()
	delete(m, name)
	fmt.Println("====================================")
	fmt.Println("Закладка успешно удалена!")
	fmt.Println("====================================")
	return m
}

func showAllBookmarks(m stringMap) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func getMenu() {
	fmt.Println("====================================")
	fmt.Println("Введите 1 - для просмотра закладок")
	fmt.Println("Введите 2 - для добавления закладок")
	fmt.Println("Введите 3 - для удаления закладок")
	fmt.Println("Введите 4 - для выхода из приложения")
	fmt.Println("====================================")
}
