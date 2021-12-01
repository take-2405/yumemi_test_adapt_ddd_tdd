package domain

import (
	"encoding/csv"
	"strconv"
)

type playData struct{
	Count int `csv:"player_id"`
	Score int `csv:"score"`
}

type PlayDatas struct{
	Datas map[string] playData
}

type RankingDatas struct{
	Datas map[int] string
}

func (p RankingDatas) Len() int {
	return len(p.Datas)
}

func (p RankingDatas) Swap(i, j int) {
	p.Datas[i], p.Datas[j] = p.Datas[j], p.Datas[i]
}

func (p RankingDatas) Less(i, j int) bool {
	return i< j
}

func ParseLogData(csv *csv.Reader)*PlayDatas{
	 //問題文では10000人以下と書いているため、10000がいいと思うが今回は500
	var playDatas PlayDatas
	playDatas.Datas=make(map[string] playData,500)
	 for {
		 record,_:= csv.Read()
		 if len(record)==0{
			break
		 }
		 score,err:=strconv.Atoi(record[2])
		 if err!=nil{

		 }
		 if playDatas.Datas[record[1]].Count==0{
			 playDatas.Datas[record[1]]=playData{Count: 1,Score:score }
		 }else{
		 	playDatas.Datas[record[1]]=playData{Count: playDatas.Datas[record[1]].Count+1,Score:playDatas.Datas[record[1]].Score+score }
		 }
	 }
	 return &playDatas
}

func (r *RankingDatas)CalcAgerageScore(data *PlayDatas){
	r.Datas=make(map[int] string,500)
	for playerId,playData:=range data.Datas{
		//fmt.Println(playData.Count)
		index :=playData.Score/playData.Count
		r.Datas[index]=playerId
	}
}