package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Указываем директорию для файлов
	fs := http.FileServer(http.Dir("./"))

	// Регистрируем обработчик для корневого пути
	http.Handle("/", fs)

	// Задаём порт для сервера
	port := 8000
	fmt.Printf("Сервер запущен на http://localhost:%d\n", port)

	// Запускаем сервер на порту 8000
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
