package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("______Программа для добавления транзакций_______")

	transactions := []int{}

	for {
		fmt.Println("Выберите действие")
		fmt.Println("1 - Добавить транзакцию")
		fmt.Println("2 - Посмотреть сумму транзакций")
		fmt.Println("3 - Выход")
		a := readInput()
		if a == 1 {
			fmt.Println("Введите транзакцию")
			transactions = addTransaction(transactions, readTransactions())
			fmt.Println(transactions)
		} else if a == 2 {
			fmt.Println("Сумма транзакций:", sumOfTransactions(transactions))
		} else if a == 3 {
			fmt.Println("Работа программы завершена")
			break
		} else {
			panic("Вы ввели неверное число")
		}

	}

}

func readInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choiceStr := scanner.Text()
	b, _ := strconv.Atoi(choiceStr)
	return b
}

func readTransactions() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choiceStr := scanner.Text()
	strValues := strings.Split(choiceStr, " ")
	var transactions []int
	for _, value := range strValues {
		if value == "" {
			continue
		}
		transaction, _ := strconv.Atoi(value)
		transactions = append(transactions, transaction)
	}
	return transactions
}

func addTransaction(a []int, b []int) []int {
	a = append(a, b...)
	return a
}

func sumOfTransactions(a []int) int {
	var sum int
	for _, x := range a {
		sum = sum + x
	}
	return sum
}
