package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

//go:embed templates/*
var templatesFS embed.FS

type Response struct {
	Token *oauth2.Token
	// https://pkg.go.dev/html/template#URL
	Next            template.URL
	ConfigurationID string
}

func configurationHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.ExecuteTemplate(w, "configuration.html.tmpl", Response{
			ConfigurationID: r.URL.Query().Get("configurationId"),
		})
	}
}

func callbackHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "code url parameter required", http.StatusBadRequest)
			return
		}

		token, err := getAccessToken(code)
		if err != nil {
			http.Error(w, "error exchanging token:\n"+err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: save this token for the logged in user, we need it to update vercel env vars!

		next := r.URL.Query().Get("next")
		configurationId := r.URL.Query().Get("configurationId")
		// TODO: we likely need teamId as well
		// only applicable when vercel user is part of a team

		response := Response{
			Token:           token,
			Next:            template.URL(next),
			ConfigurationID: configurationId,
		}
		_ = tmpl.ExecuteTemplate(w, "callback.html.tmpl", response)
	}
}

func main() {

	tmpl := template.Must(template.New("app").ParseFS(templatesFS, "templates/*.html.tmpl"))

	mux := http.NewServeMux()
	mux.HandleFunc("/vercel/callback", callbackHandler(tmpl))
	mux.HandleFunc("/vercel/configuration", configurationHandler(tmpl))

	log.Println("serving on http://localhost:8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
