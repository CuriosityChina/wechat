// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     ZachBergh(berghzach@gmail.com)

package tester

type Scanticket struct {
	Keystandard string `json:"keystandard"`
	Keystr      string `json:"keystr"`
	Openid      string `json:"openid"`
}

type Tester struct {
	Openid   []string `json:"openid"`
	Username []string `json:"username"`
}
