package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()

	setupAPI(ctx)

	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func setupAPI(ctx context.Context) {

	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/login", manager.loginHandler)
	http.HandleFunc("/ws", manager.serveWS)

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, len(manager.clients))
	})
}
