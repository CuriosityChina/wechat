// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     ZachBergh(berghzach@gmail.com)

package product

type Product struct {
	Keystandard string     `json:"keystandard"`
	Keystr      string     `json:"keystr"`
	BrandInfo   Brand_info `json:"brand_info"`
}

type Brand_info struct {
	Base_info struct {
		Title       string `json:"title,omitempty"`
		Thumb_url   string `json:"thumb_url,omitempty"`
		Category_id string `json:"category_id,omitempty"`
	} `json:"base_info,omitempty"`
	Detail_info struct {
		Banner_list []struct {
			Link string `json:"link,omitempty"`
		} `json:"banner_list,omitempty"`
		Detail_list []struct {
			Title string `json:"title,omitempty"`
			Desc  string `json:"desc,omitempty"`
		} `json:"detail_list,omitempty"`
	} `json:"detail_info,omitempty"`
	Action_info struct {
		Action_list []struct {
			Type      string `json:"type,omitempty"`
			Name      string `json:"name,omitempty"`
			Appid     string `json:"appid,omitempty"`
			Link      string `json:"link,omitempty"`
			Cardid    string `json:"cardid,omitempty"`
			Digest    string `json:"digest,omitempty"`
			Image     string `json:"image,omitempty"`
			Productid string `json:"productid,omitempty"`
		} `json:"action_list,omitempty"`
	} `json:"action_info,omitempty"`
}

type Clear struct {
	Keystandard string `json:"keystandard"`
	Keystr      string `json:"keystr"`
}

type ListInfo struct {
	Offset string `json:"offset"`
	Limit  string `json:"limit"`
}

type Keylist struct {
	Totle    int   `json:"totle"`
	Key_list []Key `json:"key_list"`
}

type Key struct {
	Keystandard   string `json:"keystandard"`
	Keystr        string `json:"keystr"`
	Category_id   int    `json:"category_id"`
	Update_time   int    `json:"update_time"`
	Status        int    `json:"status"`
	Category_name string `json:"category_name"`
}

type Mod struct {
	Keystandard string `json:"keystandard"`
	Keystr      string `json:"keystr"`
	Status      string `json:"status"`
}

type Qrinfo struct {
	Keystandard string `json:"keystandard"`
	Keystr      string `json:"keystr"`
	Extinfo     string `json:"extinfo"`
	Qrcode_size string `json:"qrcode_size"`
}

type Qrcode struct {
	Pic_url    string `json:"pic_url"`
	Qrcode_url string `json:"qrcode_url"`
}
