package card

// https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
func cardCreateUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/getcolors?access_token=TOKEN
func cardGetColorsUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/getcolors?access_token=" +
		accesstoken

}

// https://api.weixin.qq.com/card/code/consume?access_token=TOKEN
func cardCodeConsumeUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/consume?access_token=" +
		accesstoken
}
func cardCodeDecryptUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/decrypt?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/delete?access_token=TOKEN
func cardDeleteUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/delete?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/code/get?access_token=TOKEN
func cardCodeGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/batchget?access_token=TOKEN
func cardBatchGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/batchget?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/get?access_token=TOKEN
func cardGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/code/update?access_token=TOKEN
func cardCodeUpdateUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/code/unavailable?access_token=TOKEN
func cardCodeUnavailableUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/code/unavailable?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/update?access_token=TOKEN
func cardUpdateUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/code/consume?access_token=TOKEN
func cardLocationBatchAddUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/location/batchadd?access_token=" +
		accesstoken
}

//
func cardLocationBatchGetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/location/batchget?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/card/location/batchget?access_token=TOKEN
func cardTestWhiteListSetUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/card/location/batchget?access_token=" +
		accesstoken
}
