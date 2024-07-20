package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HTTPServer представляет нашу задачу HTTP-сервера
type HTTPServer struct {
	server *http.Server
}

// NewHTTPServer создает новый экземпляр HTTPServer
func NewHTTPServer(addr string) *HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет, мир!")
	})
	return &HTTPServer{
		server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

// Run запускает HTTP-сервер
func (h *HTTPServer) Run(ctx context.Context) {
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Ошибка запуска HTTP-сервера: %v\n", err)
		}
	}()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Ошибка при graceful shutdown HTTP-сервера: %v\n", err)
	} else {
		fmt.Println("HTTP-сервер успешно остановлен")
	}
}
