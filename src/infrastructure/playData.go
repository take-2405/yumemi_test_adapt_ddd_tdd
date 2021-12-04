package infrastructure

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/domain/repository"
)

type playDataPersistence struct {
	PlayData domain.PlayDatas
}

func NewPlayDataPersistence() repository.PlayDataRepository {
	return &playDataPersistence{}
}

func (p playDataPersistence) ReadPlayData(filePath string) (*csv.Reader, error) {
	//渡されたファイルパスのファイルを開く
	file, err := os.Open(filePath)
	//defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//csv fileの中身を全部取得
	csv := csv.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	header, err := csv.Read()
	if err != nil {
		log.Fatal(err)
	}
	if checkLogHeader(header) {
		errors.New("header is wrong.")
	}

	return csv, nil
}

func checkLogHeader(header []string) bool {
	//ヘッダーを確認
	if header[0] != "create_timestamp" || header[1] != "player_id" || header[2] != "score" {
		return false
	}
	return true
}

func (p playDataPersistence) ParsePlayData(csv *csv.Reader) (domain.PlayDatas, error) {
	//問題文では10000人以下と書いているため、10000がいいと思うが今回は500
	p.PlayData.Data = make(map[string]domain.PlayData, 500)

	for {

		record, err := csv.Read()
		if len(record) == 0 {
			break
		}

		score, err := strconv.Atoi(record[2])
		if err != nil {
			return p.PlayData, err
		}

		if p.PlayData.Data[record[1]].Count == 0 {

			p.PlayData.Data[record[1]] = domain.PlayData{Count: 1, Score: score}
		} else {
			p.PlayData.Data[record[1]] = domain.PlayData{Count: p.PlayData.Data[record[1]].Count + 1, Score: p.PlayData.Data[record[1]].Score + score}
		}
	}
	return p.PlayData, nil
}
