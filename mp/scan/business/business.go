// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     ZachBergh(berghzach@gmail.com)

package business

type Info struct {
	Source              string `json:"source"`
	Service_title       string `json:"service_title"`
	Service_icon_buffer string `json:"service_icon_buffer"`
}

type GetInfo struct {
	Source                string `json:"source"`
	Service_title         string `json:"service_title"`
	Service_icon_url      string `json:"service_icon_url"`
	Verified_cate_id_list []int  `json:"verified_cate_id_list"`
}
