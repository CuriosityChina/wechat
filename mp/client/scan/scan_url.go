// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     ZachBergh(berghzach@gmail.com)

package scan

// https://api.weixin.qq.com/scan/merchantinfo/set?access_token=ACCESS_TOKEN
func setBusinessInfoURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/merchantinfo/set?access_token=" +
		accesstoken
}

func getBusinessInfoURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/merchantinfo/get?access_token=" +
		accesstoken
}

func getCategoryURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/category/getsub?access_token=" +
		accesstoken
}

func createProductURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/product/create?access_token=" +
		accesstoken
}

func getProductURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/product/get?access_token=" +
		accesstoken
}

func clearProductURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/product/clear?access_token=" +
		accesstoken
}

func getProductListURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/product/getlist?access_token=" +
		accesstoken
}

func updateProductURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/product/update?access_token=" +
		accesstoken
}

func ModProductStatusURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/product/modstatus?access_token=" +
		accesstoken
}

func GetQRCodeURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/product/getqrcode?access_token=" +
		accesstoken
}

func CheckScanTicketURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/scanticket/check?access_token=" +
		accesstoken
}

func SetTestWhiteListURL(accesstoken string) string {
	return "https://api.weixin.qq.com/scan/testwhitelist/set?access_token=" +
		accesstoken
}
