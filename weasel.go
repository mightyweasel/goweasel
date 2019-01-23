package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)


//const fav_icon_url = "http://www.canada.ca/etc/designs/canada/wet-boew/assets/favicon.ico"

func main() {
	env := os.Getenv("ENVIRONMENT")
	if "" == env {
	  env = "development"
	}
	godotenv.Load(".env." + env + ".local")

	if os.Getenv("ENVIRONMENT") == "production" {
		fmt.Println("Init "+os.Getenv("ENVIRONMENT")+" Webserver")
	}
	if os.Getenv("ENVIRONMENT") == "development" {
		fmt.Println("Init "+os.Getenv("ENVIRONMENT")+" Webserver")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = "8080"

	initalizeGuiTemplates()
	weaselCreateRouter(port) 	
}

