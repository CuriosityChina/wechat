// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package card

type LocationRequest struct {
	BusinessName string `json:"business_name"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	Category     string `json:"category"`
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
}
type LocationResponse struct {
	LocationId int     `json:"location_id"`
	Name       string  `json:"name"`
	Phone      string  `json:"phone"`
	Address    string  `json:"address"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}
