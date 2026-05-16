package main

import (
	"fmt"
	"log"
	"net/http"

	"servidorHTTP/app/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}

	// Serve arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Rotas da API
	http.HandleFunc("/api/create", handlers.FormHandler)
	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.HandleFunc("/api/update", handlers.UpdateAccountHandler)
	http.HandleFunc("/api/delete", handlers.DeleteAccountHandler)

	addr := "127.0.0.1:3000"
	fmt.Printf("\n╔══════════════════════════════════════╗\n")
	fmt.Printf("║   🏥  Sistema de Saúde - Servidor    ║\n")
	fmt.Printf("║   Acesse: http://%s    ║\n", addr)
	fmt.Printf("╚══════════════════════════════════════╝\n\n")

	log.Fatal(http.ListenAndServe(addr, nil))
}
