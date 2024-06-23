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
	PARA1     string     `json:"PARA1"`
	PARA2     string     `json:"PARA2"`
	PARA3     string     `json:"PARA3"`
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
	fmt.Println("代码跑到这里了！！！")
	fmt.Println("aid:", aid)
	fmt.Println("sort:", sort)
	rows, err := dao.DB.Query("SELECT *  FROM SMS_INFO_TEST WHERE msisdn = ?", aid)
	// err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&players).Error
	fmt.Println("代码跑到第二个这里了！！！")
	if err != nil {
		return nil, fmt.Errorf("第一处查询错误 : %s", aid)
	}
	fmt.Println("代码跑到第三个这里了！！！")
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.ID, &player.Aid, &player.FLAG, &player.MSG, &player.STS, &player.GET_DATE, &player.SEND_DATE, &player.PARA1, &player.PARA2, &player.PARA3, &player.RECV, &player.DONE_CODE, &player.END_DATE); err != nil {
			return nil, fmt.Errorf("第二处查询错误 : %s", aid)
		}

		fmt.Println("代码跑到第四个这里了！！！")
		players = append(players, player)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("第三处查询错误 : %s", aid)
	}
	fmt.Println("代码跑到第五个这里了！！！")
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
