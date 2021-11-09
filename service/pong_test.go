package service

import (
	"github.com/Elder6Driver/grayReleaseProject/model"
	"testing"
)

func TestExtractRespUrl(t *testing.T) {
	//准备 （准备数据）
	type inputArgs struct {
		args []string
	}

	testCases := []struct {
		input  inputArgs
		rules  *[]model.Rule
		expect string
	}{
		{
			// 正常case
			input: inputArgs{args: []string{"15", "15"}},
			rules: &[]model.Rule{
				{
					MinVersion: 10,
					MaxVersion: 20,
					MinUserDID: 10,
					MaxUserDID: 20,
					GrayLink:   "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
				},
			},
			expect: "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
		},
		{
			// 边界值
			input: inputArgs{args: []string{"10", "20"}},
			rules: &[]model.Rule{
				{
					MinVersion: 10,
					MaxVersion: 20,
					MinUserDID: 10,
					MaxUserDID: 20,
					GrayLink:   "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
				},
			},
			expect: "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
		},
		{
			// 异常值
			input: inputArgs{args: []string{"-10", "20"}},
			rules: &[]model.Rule{
				{
					MinVersion: 10,
					MaxVersion: 20,
					MinUserDID: 10,
					MaxUserDID: 20,
					GrayLink:   "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
				},
			},
			expect: "",
		},
		{
			// 异常值
			input: inputArgs{args: []string{"a", "b"}},
			rules: &[]model.Rule{
				{
					MinVersion: 10,
					MaxVersion: 20,
					MinUserDID: 10,
					MaxUserDID: 20,
					GrayLink:   "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
				},
			},
			expect: "",
		},
	}

	// 调用 + 断言
	for _, v := range testCases {
		if v.expect != ExtractRespUrl(v.rules, v.input.args[0], v.input.args[1]) {
			t.Errorf("exprct is %+v , respUrl is %+v, input args are %+v", v.expect, ExtractRespUrl(v.rules, v.input.args[0], v.input.args[1]), v.input.args)
		}
	}
}
