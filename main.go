package main

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"
)

//go:embed views/*.html
var views embed.FS

var t *template.Template

func init() {
	var err error
	t, err = template.ParseFS(views, "views/*.html")
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar templates: %v", err))
	}
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Renderiza o template views/index.html
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, "Erro ao renderizar template: "+err.Error(), http.StatusInternalServerError)
		}
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Printf("Servidor rodando na porta %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Erro ao inicializar o servidor: %v\n", err)
	}
}
