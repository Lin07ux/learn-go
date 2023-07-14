package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HelloWorld struct {
	content string
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, h.content)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &HelloWorld{content: "Hello World Struct"})

	server := http.Server{
		Addr:    ":8800",
		Handler: mux,
	}

	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server failed:", err)
		}
	}()

	log.Println("Starting Http server on: http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal(err)
		}
	}
}
