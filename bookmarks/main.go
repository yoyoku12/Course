package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"PurpleSchool": "https://purpleschool.ru",
		"Yandex":       "https://yandex.ru",
		"Google":       "https://google.com",
		"Brave":        "https://brave.com",
		"Dota2":        "https://Dota2.com",
		"CS":           "https://CS.com",
	}

	fmt.Println("___Приложение для создания закладок____")
	for {
		fmt.Println("Введите 1 - для просмотра закладок")
		fmt.Println("Введите 2 - для добавления закладок")
		fmt.Println("Введите 3 - для удаления закладок")
		fmt.Println("Введите 4 - для выхода из приложения")
		choice := readInputInt()
		if choice == 1 {
			fmt.Println("===========================")
			fmt.Println("Список добавленных вкладок")
			fmt.Println("===========================")
			showAllBookmarks(m)
			fmt.Println("===========================")
		} else if choice == 2 {
			createBookmark(m)
		} else if choice == 3 {
			deleteBookmark(m)
		} else if choice == 4 {
			fmt.Println("Выход из приложения")
			break
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

func createBookmark(m map[string]string) map[string]string {
	var name string
	var site string
	fmt.Println("Введите наименование закладки")
	name = readInputString()
	fmt.Println("Введите сайт связанный с именем закладки")
	site = readInputString()
	m[name] = site
	fmt.Println("Закладка успешно создана!")
	return m
}

func deleteBookmark(m map[string]string) map[string]string {
	fmt.Println("Введите имя закладки для удаления")
	fmt.Println("Вот полный список имеющихся закладок:")
	for key, value := range m {
		fmt.Println(key, value)
	}
	name := readInputString()
	delete(m, name)
	fmt.Println("Закладка успешно удалена!")
	return m
}

func showAllBookmarks(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
