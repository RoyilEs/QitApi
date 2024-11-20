package main

import (
	"awesomeProject/action"
	"awesomeProject/qitApi"
	"awesomeProject/utils"
	"fmt"
)

func main() {
	utils.YatoriCoreInit()
	cache := qitApi.UserCache{Account: "221171010606", Password: "zxd2839182980."}
	loginAction, err := action.QitLoginAction(&cache)
	if err != nil {
		panic(err)
	}
	utils.DeleteFileS()
	fmt.Println(loginAction)
	courseUrl, _ := cache.CourseInfo()
	fmt.Println(courseUrl)
	action.QitCoursePull(&cache)
}
