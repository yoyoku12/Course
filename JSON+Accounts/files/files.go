package files

import (
	"fmt"
	"os"
)

func ReadFiles(file string) string {
	text, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return "Файл не считался"
	}
	return string(text)
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Аккаунт успешно создан!")
	defer file.Close()
}
