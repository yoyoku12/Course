package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"password/utils"
	"time"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!")

const path = "D:\\go\\Course\\accounts+password\\main\\accounts.json"

type Account struct {
	Login     string    `json:"Login"`
	Password  string    `json:"Password"`
	Url       string    `json:"Url"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (acc *Account) OutputPassword() {
	fmt.Println(acc.Password)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) GeneratePassword(n int) {

	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(Login, Password, UrlString string) (*Account, error) {
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       UrlString,
		Login:     Login,
		Password:  Password,
	}

	if Password == "" {
		newAcc.GeneratePassword(12)
	}
	return newAcc, nil
}

func CreateAccount() {
	login := utils.PromptDataWithText("Введите логин: ")
	password := utils.PromptDataWithText("Введите пароль: ")
	url := utils.PromptDataWithText("Введите URL: ")
	myAccount, err := NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Ошибка при чтении файла:", err)
			return
		}

		data = []byte("[]")
	}

	var accounts []Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		fmt.Println("Ошибка при десериализации данных:", err)
		return
	}

	accounts = append(accounts, *myAccount)

	updatedData, err := json.Marshal(accounts)
	if err != nil {
		fmt.Println("Ошибка при сериализации данных:", err)
		return
	}

	err = os.WriteFile("accounts.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл:", err)
		return
	}

	fmt.Println("Аккаунт успешно создан")
}

func DeleteAccount() (string, error) {
	login := utils.PromptDataWithText("Введите имя аккаунта: ")

	fmt.Println("Читаем файл:", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("не удалось прочитать файл: %v", err)
	}

	var accounts []Account
	fmt.Println("Десериализуем содержимое файла...")
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return "", fmt.Errorf("ошибка при десериализации JSON: %v", err)
	}

	var updatedAccounts []Account
	found := false
	for _, account := range accounts {
		if account.Login != login {
			updatedAccounts = append(updatedAccounts, account)
		} else {
			found = true
		}
	}

	if !found {
		return "", errors.New("аккаунт с таким логином не найден")
	}

	fmt.Println("Сериализуем обновлённые данные...")
	updatedData, err := json.Marshal(updatedAccounts)
	if err != nil {
		return "", fmt.Errorf("ошибка при сериализации обновленных данных: %v", err)
	}

	fmt.Println("Записываем данные обратно в файл:", path)
	err = os.WriteFile(path, updatedData, 0644)
	if err != nil {
		return "", fmt.Errorf("ошибка при записи обновленных данных в файл: %v", err)
	}

	return "Аккаунт успешно удален", nil
}

func FindAccountByName() string {
	name := utils.PromptDataWithText("Введите имя аккаунта: ")

	data, err := os.ReadFile(path)
	if err != nil {
		return "не удалось прочитать файл"
	}

	var accounts []Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return "ошибка при десериализации JSON"
	}

	for _, account := range accounts {
		if account.Login == name {
			return fmt.Sprintf("Логин: %s\nПароль: %s\nURL: %s\nДата создания: %s\nДата обновления: %s\n", account.Login, account.Password, account.Url, account.CreatedAt, account.UpdatedAt)
		}
	}

	return "аккаунт с таким логином не найден"
}

func FindAccountByPassword() string {
	password := utils.PromptDataWithText("Введите пароль аккаунта: ")

	data, err := os.ReadFile(path)
	if err != nil {
		return "не удалось прочитать файл"
	}

	var accounts []Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return "ошибка при десериализации JSON"
	}

	for _, account := range accounts {
		if account.Password == password {
			return fmt.Sprintf("Логин: %s\nПароль: %s\nURL: %s\nДата создания: %s\nДата обновления: %s\n", account.Login, account.Password, account.Url, account.CreatedAt, account.UpdatedAt)
		}
	}

	return "аккаунт с таким паролем не найден"
}

func FindAccountByUrl() string {
	url := utils.PromptDataWithText("Введите URL аккаунта: ")

	data, err := os.ReadFile(path)
	if err != nil {
		return "не удалось прочитать файл"
	}

	var accounts []Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return "ошибка при десериализации JSON:"
	}

	for _, account := range accounts {
		if account.Url == url {
			return fmt.Sprintf("Логин: %s\nПароль: %s\nURL: %s\nДата создания: %s\nДата обновления: %s\n", account.Login, account.Password, account.Url, account.CreatedAt, account.UpdatedAt)
		}
	}

	return "аккаунты с таким URL s не найдены"

}
