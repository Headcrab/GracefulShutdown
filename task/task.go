package task

import "context"

// Task представляет интерфейс для задач, которые могут быть запущены и остановлены
type Task interface {
  Run(context.Context)
}
