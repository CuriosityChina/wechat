// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     ZachBergh(berghzach@gmail.com)

package category

type Category struct {
	Cate_list []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"cate_list"`
}
