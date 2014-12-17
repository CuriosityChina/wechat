// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package card

const (
	CARD_TYPE_GENERAL_COUPON = "GENERAL_COUPON" // 通用券
	CARD_TYPE_GROUPON        = "GROUPON"        //
	CARD_TYPE_DISCOUNT       = "DISCOUNT"
	CARD_TYPE_GIFT           = "GIFT"
	CARD_TYPE_CASH           = "CASH"
	CARD_TYPE_MEMBER_CARD    = "MEMBER_CARD"
	CARD_TYPE_SCENIC_TICKET  = "SCENIC_TICKET"
	CARD_TYPE_MOVIE_TICKET   = "MOVIE_TICKET"
	CARD_TYPE_BOARDING_PASS  = "BOARDING_PASS"
	CARD_TYPE_LUCKY_MONEY    = "LUCKY_MONEY"
)

const (
	CODE_TYPE_TEXT    = "CODE_TYPE_TEXT"    // 文本
	CODE_TYPE_BARCODE = "CODE_TYPE_BARCODE" // 一维码（条形码）
	CODE_TYPE_QRCODE  = "CODE_TYPE_QRCODE"  // 二位啊
)

const (
	Date_Info_Type_Begin_End  = 1
	Date_Info_Type_Fixed_Term = 2
)

const (
	URL_NAME_TYPE_TAKE_AWAY           = "URL_NAME_TYPE_TAKE_AWAY"
	URL_NAME_TYPE_RESERVATION         = "URL_NAME_TYPE_RESERVATION"
	URL_NAME_TYPE_USE_IMMEDIATELY     = "URL_NAME_TYPE_USE_IMMEDIATELY"
	URL_NAME_TYPE_APPOINTMENT         = "URL_NAME_TYPE_APPOINTMENT"
	URL_NAME_TYPE_EXCHANGE            = "URL_NAME_TYPE_EXCHANGE"
	URL_NAME_TYPE_MALL                = "URL_NAME_TYPE_MALL"
	URL_NAME_TYPE_VEHICLE_INFORMATION = "URL_NAME_TYPE_VEHICLE_INFORMATION"
)
