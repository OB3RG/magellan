package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"magellan/pkg/config"
	"magellan/pkg/db"

	"magellan/pkg/services/author"
	"magellan/pkg/services/content"
	"magellan/pkg/services/site"
	"magellan/pkg/services/tag"
)

func main() {
	log.Print("Setting up config")
	cfg := config.DefaultCfg()
	arg.MustParse(cfg)

	log.Print(cfg)
	db, err := db.Connect(cfg.DBConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Print("Setting up new router")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	authorServer := author.NewServer(cfg, db)
	contentServer := content.NewServer(cfg)
	siteServer := site.NewServer(cfg)
	tagServer := tag.NewServer(cfg)

	authorRoutes := authorServer.SetupRoutes()
	contentRoutes := contentServer.SetupRoutes()
	siteRoutes := siteServer.SetupRoutes()
	tagRoutes := tagServer.SetupRoutes()

	r.Mount("/authors", authorRoutes)
	r.Mount("/contents", contentRoutes)
	r.Mount("/sites", siteRoutes)
	r.Mount("/tags", tagRoutes)

	log.Print("Succesfully loaded routes")

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerConfig.Port), r)
}
