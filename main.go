package main

import (
	"log"

	"github.com/donreno/dashboar/internal/service"
	"github.com/donreno/dashboar/internal/webapp"
)

func main() {
	log.Println("Starting DashBOAR 🐗💨")

	log.Fatal(webapp.InitWebApp(service.GetDashboard))
}
