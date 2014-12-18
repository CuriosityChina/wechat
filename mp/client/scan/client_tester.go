// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     ZachBergh(berghzach@gmail.com)

package scan

import (
	"github.com/chanxuehong/wechat/mp/scan/tester"
)

// 检查wxticket参数
func (c *Client) CheckScanTicket(ticket_ string) (scanticket_ tester.Scanticket, err error) {
	var result struct {
		Wxticket tester.Scanticket
		Error
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := CheckScanTicketURL(token)

	if err = c.postJSON(url_, ticket_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		scanticket_ = result.Wxticket
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

// 添加开发者白名单
func (c *Client) SetTestWhiteList(tester_ tester.Tester) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := SetTestWhiteListURL(token)

	if err = c.postJSON(url_, tester_, &result); err != nil {
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
