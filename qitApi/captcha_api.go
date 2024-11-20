package qitApi

import (
	"awesomeProject/utils"
	"fmt"
	"net/http"
)

func (cache *UserCache) GetCaptcha() error {
	for {
		//第二部：拉取对应Session所绑定的验证码
		verPath, httpCode, err := CaptchaPull(cache)
		if err != nil {
			panic(err)
		}
		if httpCode == http.StatusInternalServerError {
			continue
		}
		//第三步：AI识别验证码
		img, _ := utils.ReadImg(verPath)
		if img == nil {
			continue
		}
		verification := utils.AutoVerification(img)
		cache.Ver = verification
		fmt.Printf("验证码：%s", verification)
		return nil
	}
	return nil
}
