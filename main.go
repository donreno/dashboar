package main

import (
	"log"

	"github.com/donreno/dashboar/internal/config"
	"github.com/donreno/dashboar/internal/service"
	"github.com/donreno/dashboar/internal/webapp"
)

func main() {
	log.Println("Starting DashBOAR 🐗💨")

	envCfg, dashboarCfg, err := config.LoadConfiguration()
	if err != nil {
		log.Panicf("Error laoding initial configuration!! cause: %v", err.Error())
	}

	dashboarRetriever := service.MakeDashboarRetriever(dashboarCfg)

	log.Fatal(webapp.InitWebApp(dashboarRetriever, envCfg))
}
