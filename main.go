package main

import (
	"log"
	"GoProject/routes"
	"GoProject/config"
	"fmt"
)


func main(){
	log.Println("\nServer is up and Running")
     fmt.Println("-------------------Starting Vechicle Parts ordering App-----------------------")

	 config.ConnectDatabase()
	 fmt.Println("DATABASE CONNECTED")
	r := routes.SetupRouter();
	r.Run()

}
