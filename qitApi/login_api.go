package qitApi

import (
	"awesomeProject/utils"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type UserCache struct {
	Account    string //账号
	Password   string //密码
	Cookie     string //Cookie,session,主要用于验证码的绑定
	TokenValue string
	Ver        string //验证码
}

// TokenValue TokenValue截取
func TokenValue(cache *UserCache) (string, error) {

	url := "http://58.56.101.14:881/login"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	compileRegex := regexp.MustCompile(`"tokenValue" value="([^"]*)">`) // 正则表达式的分组，以括号()表示，每一对括号就是我们匹配到的一个文本，可以把他们提取出来。
	matchArr := compileRegex.FindStringSubmatch(string(body))
	if len(matchArr) > 0 {
		//设置TokenValue
		cache.TokenValue = matchArr[len(matchArr)-1]
		//设置Cookie
		//cache.Cookie = res.Header.Get("Set-Cookie")
		cache.Cookie = "" //先清空Cookie
		for i := range res.Cookies() {
			cache.Cookie += strings.ReplaceAll(res.Cookies()[i].String(), "Path=/", "")
		}
		return matchArr[len(matchArr)-1], nil
	}

	return "", errors.New("未找到tokenValue")
}

// CaptchaPull 获取验证码，这里返回的应该是一个IO文件流，所以你要保存文件形式到本地，也就是保存验证码图片到本地，方便进行识别
func CaptchaPull(cache *UserCache) (string, int, error) {
	var filepath string
	url := "http://58.56.101.14:881/img/captcha.jpg"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	req.Header.Add("Cookie", cache.Cookie)

	res, err := client.Do(req)
	if res.StatusCode == http.StatusInternalServerError {
		return "", res.StatusCode, nil
	}
	if err != nil {
		fmt.Println(err)
		return "", res.StatusCode, err
	}
	defer res.Body.Close()

	codeFileName := "code" + strconv.Itoa(rand.Intn(99999)) + ".png" //生成验证码文件名称
	utils.PathExistForCreate("./")                                   //检测是否存在路径，如果不存在则创建
	filepath = fmt.Sprintf("./%s", codeFileName)
	file, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		return "", res.StatusCode, err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Println(err)
		return "", res.StatusCode, err
	}
	return filepath, res.StatusCode, nil
}

// LoginApi 登录接口
func LoginApi(cache *UserCache) (string, error) {

	url := "http://58.56.101.14:881/j_spring_security_check"
	method := "POST"
	// 密码属于MD5哈希化
	hash := md5.New()
	hash.Write([]byte(cache.Password))
	md5Pass := hex.EncodeToString(hash.Sum(nil))
	//md5Pass := "dc7d605e1d10f1cb2508388347106338"
	payload := strings.NewReader("tokenValue=" + cache.TokenValue + "&j_username=" + cache.Account + "&j_password=" + md5Pass + "&j_captcha=" + cache.Ver)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		time.Sleep(2 * time.Second)
	}
	req.Header.Add("Host", " 58.56.101.14:881")
	req.Header.Add("Cache-Control", " max-age=0")
	req.Header.Add("Origin", " http://58.56.101.14:881")
	req.Header.Add("Upgrade-Insecure-Requests", " 1")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Referer", " http://58.56.101.14:881/login")
	req.Header.Add("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Cookie", cache.Cookie)
	req.Header.Add("Content-Type", " application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Received 500 Internal Server Error\n")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
