package qitApi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func (cache *UserCache) StudentRollApi() ([]string, error) {
	url := "http://58.56.101.14:881/student/rollManagement/rollInfo/index"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Host", " 58.56.101.14:881")
	req.Header.Add("Cache-Control", " max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", " 1")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Referer", " http://58.56.101.14:881/login")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Cookie", cache.Cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 创建正则表达式对象，用于匹配电子邮件地址
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	// 创建正则表达式对象，用于匹配 dh 和 sfzh 字段的值
	dhRegex := regexp.MustCompile(`"dh",\s*"(\d{11})"`)
	sfzhRegex := regexp.MustCompile(`"sfzh",\s*"(\d{18})"`)
	nameRegex := regexp.MustCompile(`<div class="profile-info-value">\s*([^<]*)\s*</div>`)

	// 在字符串中查找所有匹配项
	emailMatch := emailRegex.FindStringSubmatch(string(body))
	phoneMatch := dhRegex.FindStringSubmatch(string(body))
	cardMatch := sfzhRegex.FindStringSubmatch(string(body))
	nameMatch := nameRegex.FindAllStringSubmatch(string(body), -1)
	var rolls []string
	adds := func(str string) {
		rolls = append(rolls, strings.TrimSpace(str))
	}
	adds(nameMatch[14][1])
	adds(nameMatch[49][1])
	adds(cache.Account)
	adds(nameMatch[23][1])
	adds(nameMatch[19][1])
	adds(nameMatch[22][1])
	adds(nameMatch[55][1])
	adds(emailMatch[0])
	adds(phoneMatch[1])
	adds(cardMatch[1])

	return rolls, nil
}
