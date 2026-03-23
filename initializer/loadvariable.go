package initializer;

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Load_variables(){
	err:= godotenv.Load()
	if err!=nil{
		fmt.Println("Error in loading environment variables")
		log.Fatal("Environment variables loading failed")
	}
}