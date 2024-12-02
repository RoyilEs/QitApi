package action

import (
	"awesomeProject/qitApi"
	"fmt"
)

type StudentRollStruct struct {
	Name      string `json:"name"`
	Brith     string `json:"brith"`
	StudentID string `json:"student_id"`
	Colleges  string `json:"college"` //院校
	Campus    string `json:"campus"`  //校区
	Class     string `json:"class"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CardID    string `json:"card_id"` //证件
}

func RollAction(cache *qitApi.UserCache) StudentRollStruct {
	rolls, err := cache.StudentRollApi()
	if err != nil {
		fmt.Printf("获取学生信息失败：%v", err)
		return StudentRollStruct{}
	}
	studentRoll := StudentRollStruct{
		rolls[0],
		rolls[1],
		rolls[2],
		rolls[3],
		rolls[4],
		rolls[5],
		rolls[6],
		rolls[7],
		rolls[8],
		rolls[9],
	}
	return studentRoll
}
