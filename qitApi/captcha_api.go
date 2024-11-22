package qitApi

import (
	"awesomeProject/utils"
	"fmt"
	"net/http"
)

func (cache *UserCache) GetCaptcha() error {
	for {
		//拉取对应Session所绑定的验证码
		verPath, httpCode, err := CaptchaPull(cache)
		if err != nil {
			panic(err)
		}
		if httpCode == http.StatusInternalServerError {
			continue
		}
		//AI识别验证码
		img, _ := utils.ReadImg(verPath)
		if img == nil {
			continue
		}
		verification := utils.AutoVerification(img)
		cache.Ver = verification
		fmt.Printf("验证码：%s\n", verification)
		return nil
	}
	return nil
}
