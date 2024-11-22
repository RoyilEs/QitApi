package action

import (
	"awesomeProject/qitApi"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CourseStruct struct {
	//TODO 课表关键信息
}

func QitCoursePull(cache *qitApi.UserCache) {
	path, err := cache.CourseInfo()
	if err != nil {
		fmt.Printf("获取课程path失败：%v", err)
	}
	url := "http://58.56.101.14:881" + path
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", " 58.56.101.14:881")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Add("Accept", " application/json, text/javascript, */*; q=0.01")
	req.Header.Add("X-Requested-With", " XMLHttpRequest")
	req.Header.Add("Origin", " http://58.56.101.14:881")
	req.Header.Add("Referer", " http://58.56.101.14:881/student/courseSelect/thisSemesterCurriculum/index")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Cookie", cache.Cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
