package infrastructure

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadLogData(filePath string)([][]string ,error){
	//渡されたファイルパスのファイルを開く
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	logData := csv.NewReader(file)
	record, err := logData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//reader := csv.NewReader(file)
	//// ヘッダー行の取得
	//header, _ := reader.Read()
	//if !checkLogHeader(header){
	//	log.Fatal("file is incorrect.")
	//}

	return record,nil
}

func checkLogHeader(header []string)bool{
	if header[0]!="create_timestamp"||header[1]!="player_id"||header[2]!="score"{
		return false
	}
	return true
}