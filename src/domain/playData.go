package domain

type PlayData struct {
	Count int `csv:"player_id"`
	Score int `csv:"score"`
}

type PlayDatas struct {
	Data map[string]PlayData
}
