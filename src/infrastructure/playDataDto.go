package infrastructure

import (
	"encoding/csv"
	"strconv"
	"yumemi-coding-test-practice/src/domain"
)

type playData struct {
	Count int `csv:"player_id"`
	Score int `csv:"score"`
}

type PlayDatas struct {
	Datas map[string]playData
}

func ParseLogData(csv *csv.Reader) (*PlayDatas, error) {
	//問題文では10000人以下と書いているため、10000がいいと思うが今回は500
	var playDatas PlayDatas
	playDatas.Datas = make(map[string]playData, 500)
	for {
		record, _ := csv.Read()

		if len(record) == 0 {
			break
		}
		score, err := strconv.Atoi(record[2])
		if err != nil {
			return &playDatas, err
		}
		if playDatas.Datas[record[1]].Count == 0 {
			playDatas.Datas[record[1]] = playData{Count: 1, Score: score}
		} else {
			playDatas.Datas[record[1]] = playData{Count: playDatas.Datas[record[1]].Count + 1, Score: playDatas.Datas[record[1]].Score + score}
		}
	}
	return &playDatas, nil
}

func (p *PlayDatas) CalcAgerageScore() *domain.PlayerScores {
	var playScore domain.PlayerScores
	playScore.PlayerScore = make(map[string]int, 500)

	for playerId, playData := range p.Datas {
		score := playData.Score / playData.Count
		playScore.PlayerScore[playerId] = score
	}

	return &playScore
}
