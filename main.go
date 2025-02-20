package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/senforsce/mathfighter/web/pages"
	"github.com/senforsce/tndr0cean/router"
)

func main() {
	app := router.New()
	Plugs(app)
	log.Fatal(app.Start())

}

func init() {
	if envErr := godotenv.Load(".env"); envErr != nil {
		fmt.Println(".env file missing")
	}
}

func Plugs(app *router.Tndr0cean) {
	//WithSparQlServer(app)
	WithHTMXServer(app)
	// WithHTMXPreviews(app)
}

func WithHTMXServer(app *router.Tndr0cean) func(h router.Handler) {
	app.Get("/", pages.Index)
	// Other routes will be injected with tree-shaking on build inside ./injected_routes.go
	return nil
}
