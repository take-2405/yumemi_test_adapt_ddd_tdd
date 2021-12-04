package domain

type PlayerScores struct {
	PlayerScore map[string]int
}

//func (p *PlayerScores) CalcAgerageScore(playDatas *infrastructure.PlayDatas) {
//	p.PlayerScore = make(map[string]int, 500)
//	for playerId, playData := range playDatas.Datas {
//		score := playData.Score / playData.Count
//		p.PlayerScore[playerId] = score
//	}
//}
