package qitApi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// CourseInfo 课表Get渲染请求 正则获取ajax中的实际json数据请求
func (cache *UserCache) CourseInfo() (string, error) {
	url := "http://58.56.101.14:881/student/courseSelect/thisSemesterCurriculum/index"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Host", " 58.56.101.14:881")
	req.Header.Add("Upgrade-Insecure-Requests", " 1")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Referer", " http://58.56.101.14:881/student/integratedQuery/scoreQuery/allTermScores/index")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Cookie", cache.Cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	urlRegex := regexp.MustCompile(`url:\s*["']([^"']+)["']`)
	match := urlRegex.FindStringSubmatch(string(body))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if len(match) > 1 {
		return match[1], nil
	} else {
		return "未找到url", nil
	}
}
