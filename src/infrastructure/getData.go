package infrastructure

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadLogData(filePath string)(*csv.Reader ,error){
	//渡されたファイルパスのファイルを開く
	file, err := os.Open(filePath)
	//defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//csv fileの中身を全部取得
	logData := csv.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	return logData,nil
}

func CheckLogHeader(header []string)bool{
	//ヘッダーを確認
	if header[0]!="create_timestamp"||header[1]!="player_id"||header[2]!="score"{
		return false
	}
	return true
}
