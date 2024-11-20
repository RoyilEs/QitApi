package action

import (
	"awesomeProject/qitApi"
	"fmt"
	"regexp"
	"strings"
)

func QitLoginAction(cache *qitApi.UserCache) (api string, err error) {
	//第一步：拉取TokenValue以及设置Cookie
	value, err := qitApi.TokenValue(cache)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TokenValue：%s", value)
	err = cache.GetCaptcha()
	api, err = qitApi.LoginApi(cache)

	re := regexp.MustCompile(`<strong>([^<]*)</strong>`)
	// 查找所有匹配项
	matches := re.FindStringSubmatch(api)
	if len(matches) == 0 { //等于0说明没有问题
		return "登录成功", nil
	}
	if strings.Contains(api, "验证码错误") || strings.Contains(api, "token校验失败") {
		return QitLoginAction(cache)
	}

	return api, nil
}
