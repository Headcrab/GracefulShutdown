package main

import (
	"fmt"
	"time"

	"main/graceful"
	"main/server"
	"main/task"
)

func main() {
	// Устанавливаем время ожидания в 10 секунд
	gs := graceful.NewGracefulShutdown(10 * time.Second)

	// Создаем и добавляем нашу "вечную" задачу
	eternalTask := task.NewEternalTask(time.Second)
	gs.AddTask(eternalTask)

	// Создаем и добавляем HTTP-сервер
	httpServer := server.NewHTTPServer(":8080")
	gs.AddTask(httpServer)

	fmt.Println("Сервер запущен на http://localhost:8080")

	// Ожидаем завершения
	gs.Wait()

	fmt.Println("Программа успешно завершена")
}
