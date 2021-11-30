package infrastructure

import (
	"log"
	"os"
)

func ReadLogData(filePath string)(error){
	_, err := os.Open("file.csv")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
