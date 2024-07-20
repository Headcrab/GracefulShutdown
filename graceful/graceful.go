package graceful

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"main/task"
)

// GracefulShutdown предоставляет механизм для graceful shutdown
type GracefulShutdown struct {
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
	timeout time.Duration
}

// NewGracefulShutdown создает новый экземпляр GracefulShutdown
func NewGracefulShutdown(timeout time.Duration) *GracefulShutdown {
	ctx, cancel := context.WithCancel(context.Background())
	return &GracefulShutdown{
		ctx:     ctx,
		cancel:  cancel,
		timeout: timeout,
	}
}

// AddTask добавляет задачу в GracefulShutdown
func (gs *GracefulShutdown) AddTask(task task.Task) {
	gs.wg.Add(1)
	go func() {
		defer gs.wg.Done()
		task.Run(gs.ctx)
	}()
}

// Wait ожидает сигнала завершения и затем ожидает завершения всех задач
func (gs *GracefulShutdown) Wait() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\nПолучен сигнал завершения")
	gs.cancel()
	// Создаем канал для отслеживания завершения задач
	done := make(chan struct{})
	go func() {
		gs.wg.Wait()
		close(done)
	}()
	// Ожидаем завершения задач или истечения времени ожидания
	select {
	case <-done:
		fmt.Println("Все задачи завершены")
	case <-time.After(gs.timeout):
		fmt.Println("Время ожидания истекло, принудительное завершение")
	}
}
