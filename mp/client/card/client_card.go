package card

import (
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/card"
)

// pass
func (c *Client) CardCreate(_card card.Card) (card_id string, err error) {
	var request = struct {
		Card card.Card `json:"card"`
	}{
		Card: _card,
	}
	var result struct {
		CardId string `json:"card_id"`
		Error
	}
	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := cardCreateUrl(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}
	fmt.Println(request)
	fmt.Sprintf("%v", request)
	switch result.ErrCode {
	case errCodeOK:
		card_id = result.CardId
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return

	}

}

// pass
func (c *Client) CardGet(card_id string) (_card card.Card, err error) {
	var request = struct {
		CardId string `json:"card_id"`
	}{
		CardId: card_id,
	}
	var result struct {
		Error
		Card card.Card `json:"card"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := cardGetUrl(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}
	switch result.ErrCode {
	case errCodeOK:
		_card = result.Card
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return

	}
}

//
func (c *Client) CardBatchGet(offset, count int) (cards_id_list []string, total_sum int, err error) {
	var request = struct {
		Offset int `json:"offset"`
		Count  int `json:"count"`
	}{
		Offset: offset,
		Count:  count,
	}
	var result struct {
		Error
		CardIdList []string `json:"card"`
		TotalSum   int      `json:"total_sum"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := cardBatchGetUrl(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}
	switch result.ErrCode {
	case errCodeOK:
		cards_id_list = result.CardIdList
		total_sum = result.TotalSum
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return

	}
}

//
func (c *Client) CardDelete(card_id string) (err error) {
	var request = struct {
		CardId string `json:"card_id"`
	}{
		CardId: card_id,
	}
	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := cardDeleteUrl(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}
	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return

	}
}

/*
import (
	"github.com/magicshui/wechat/mp/card"
)

func (c *Client) CardCreate() (card_id string, err error) {
	return
}
func (c *Client) CardQrcodeCreate() (ticket string, err error) {
	return
}
func (c *Client) CardCodeConsume() error {
	return
}
func (c *Client) CardCodeDecrypt() error {
	return
}
func (c *Client) cardDelete() error {

}
func (c *Client) CardEncyptUrl() string {
	return ""

}
*/
/*
	为了满足商户基于卡券本身的扩展诉求，允许卡券内页添加url跳转外链。以下字段可支持卡券页
	面内跳转外链：
	custom_url，添加该字段后，卡券页面内允许用户点击自定义cell后跳转到该URL。
	activate_url，会员卡页面内允许用户点击“会员卡激活”后跳转URL进行激活操作。
	bind_old_card_url，会员卡页面内允许用户点击“会员卡绑定”跳转URL进行绑定操作。
	check_in_url，飞机票页面允许用户点击“在线办理登机牌”后跳转URL进行在线值机。
	guide_url，门票页面内允许用户点击cell跳转至导览图URL。
	*特殊卡票中添加custom_url字段，仅在完成数据更新后显示自定义cell，即会员卡完成激活/绑定、
	飞机票完成在线值机、电影票完成选座信息更新后在卡券页面内显示自定义cell。
	为检测跳转外链请求来自微信，会在URL参数里加上签名。参与签名的字段有appsecret、
	encrypt_code、card_id。
	签名字段
	字段 说明 是否必填
	appsecret 第三方用户唯一凭证密钥 是
	encrypt_code 指定的卡券 code 码，只能被领一次。 是
	card_id 创建卡券时获得的卡券 ID 是
	签名方式：
	1. 将appsecret、encrypt_code、card_id的value值进行字符串的字典序排序。
	2. 将所有参数字符串拼接成一个字符串进行sha1加密，得到signature。
	示例：encrypt_code= encrypt(456) = xcfsda12312, card_id = 123, app_secrect=789
	signature = sha1(123456789) = f7c3bc1d808e04732adf679965ccc34ca7ae3441
	*签名中对 code 进行加密处理，需调用解码接口（3.3code 解码接口）获取真实 code。
	假如指定的 url 为 http://www.qq.com，用户点击时，跳转的 url 则为：
	http://www.qq.com?encrypt_code=CODE&card_id=CARDID&sinature=SIGNATRUE
*/
