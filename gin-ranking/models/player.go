package models

import (
	"fmt"
	"gin-ranking/dao"
	"time"
	// "github.com/jinzhu/gorm"
)

// type Player struct {
// 	Id          int    `json:"id"`
// 	Aid         int    `json:"aid"`
// 	Ref         string `json:"ref"`
// 	Nickname    string `json:"nickname"`
// 	Declaration string `json:"declaration"`
// 	Avatar      string `json:"avatar"`
// 	Score       int    `json:"score"`
// 	//AddTime     int64  `json:"addTime"`
// 	//UpdateTime  int64  `json:"updateTime"`
// }

type Player struct {
	ID        int64      `json:"ID"`  // 根据实际情况，可能需要使用int64来存储NUMBER(15)
	Aid       string     `json:"aid"` //  实际对应数据库表中的MSISDN字段，但数据库中的类型是string，需要考虑是否有影响？？
	FLAG      string     `json:"FLAG"`
	MSG       string     `json:"MSG"`
	STS       string     `json:"STS"`
	GET_DATE  time.Time  `json:"GET_DATE"`
	SEND_DATE time.Time  `json:"SEND_DATE"`
	RECV      string     `json:"RECV"`
	DONE_CODE int64      `json:"DONE_CODE"`
	END_DATE  *time.Time `json:"END_DATE"`
}

func (Player) TableName() string {
	return "player"
}

// 此处是查询数据库的入口
func GetPlayers(aid string, sort string) ([]Player, error) {
	var players []Player
	querySQL := fmt.Sprintf("SELECT ID,MSISDN,FLAG,MSG,STS,GET_DATE,SEND_DATE,RECV,DONE_CODE,END_DATE FROM SMS_INFO_TEST WHERE MSISDN = %s", aid)
	fmt.Println("querySQL:", querySQL)

	rows, err := dao.DB.Query(querySQL)
	// err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&players).Error

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("第一处查询错误 : %s", aid)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.ID, &player.Aid, &player.FLAG, &player.MSG, &player.STS, &player.GET_DATE, &player.SEND_DATE, &player.RECV, &player.DONE_CODE, &player.END_DATE); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("第二处查询错误 : %s", aid)
		}
		players = append(players, player)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("第三处查询错误 : %s", aid)
	}
	return players, err
}

// func GetPlayerInfo(id int) (Player, error) {
// 	var player Player
// 	err := dao.Db.Where("id = ?", id).First(&player).Error
// 	return player, err
// }

// func UpdatePlayerScore(id int) {
// 	var player Player
// 	dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
// }
