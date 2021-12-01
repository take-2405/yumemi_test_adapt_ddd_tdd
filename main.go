package main

import(
	"log"
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/infrastructure"
)
func main(){
	csv,err:=infrastructure.ReadLogData("./game_score_log.csv")
	if err!=nil{
		log.Fatal(err)
	}
	header,err:=csv.Read()
	if err!=nil{
		log.Fatal(err)
	}
	if !infrastructure.CheckLogHeader(header){
		if err!=nil{
			log.Fatal(err)
		}
	}
	domain.ParseLogData(csv)
}
