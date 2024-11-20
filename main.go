package main

import (
	"awesomeProject/action"
	"awesomeProject/qitApi"
	"awesomeProject/utils"
	"fmt"
)

func main() {
	utils.YatoriCoreInit()
	cache := qitApi.UserCache{Account: "XXX", Password: "XXX"}
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
