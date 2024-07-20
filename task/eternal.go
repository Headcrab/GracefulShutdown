package task

import (
	"context"
	"fmt"
	"time"
)

// EternalTask представляет нашу "вечную" задачу
type EternalTask struct {
	interval time.Duration
}

// NewEternalTask создает новый экземпляр EternalTask
func NewEternalTask(interval time.Duration) *EternalTask {
	return &EternalTask{interval: interval}
}

// Run выполняет "вечную" задачу
func (e *EternalTask) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("EternalTask получила сигнал о завершении")
			return
		default:
			fmt.Println("Это почти вечный цикл!")
			time.Sleep(e.interval)
		}
	}
}
