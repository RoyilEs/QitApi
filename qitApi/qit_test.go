package qitApi

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	// 给定的 JavaScript 代码
	jsCode := `
	var index;
	$.ajax({
		url: "/student/courseSelect/thisSemesterCurriculum/d0202355k7/ajaxStudentSchedule/curr/callback",
		cache: false,
		type: "post",
		data: null,
		dataType: "json",
		beforeSend: function () {
			index = layer.load(0, {
				shade: [0.2, "#000"] //0.1透明度的白色背景
			});
		}
	});
	`

	// 使用正则表达式匹配 url 属性的值
	urlRegex := regexp.MustCompile(`url:\s*["']([^"']+)["']`)
	match := urlRegex.FindStringSubmatch(jsCode)

	if len(match) > 1 {
		// 提取的 URL
		url := match[1]
		fmt.Println("Extracted URL:", url)
	} else {
		fmt.Println("URL not found")
	}
}
