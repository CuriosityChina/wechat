// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     ZachBergh(berghzach@gmail.com)

package scan

import (
	"github.com/chanxuehong/wechat/mp/scan/product"
)

// 创建产品信息
func (c *Client) CreateProduct(product_ product.Product) (pid_ string, err error) {
	var result struct {
		Pid string `json:"pid"`
		Error
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := createProductURL(token)

	if err = c.postJSON(url_, product_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		pid_ = result.Pid
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

// 获取产品信息
func (c *Client) GetProductInfo() (product_ product.Brand_info, err error) {
	var result struct {
		Brand_info product.Brand_info `json:"brand_info"`
		Error
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := getProductURL(token)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		product_ = result.Brand_info
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

// 清除产品信息
func (c *Client) ClearProductInfo(clear_ product.Clear) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := clearProductURL(token)

	if err = c.postJSON(url_, clear_, &result); err != nil {
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

// 批量获取产品信息
func (c *Client) GetProductList(listinfo_ product.ListInfo) (keylist_ product.Keylist, err error) {
	var result struct {
		Keylist product.Keylist
		Error
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := getProductListURL(token)

	if err = c.postJSON(url_, listinfo_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		keylist_ = result.Keylist
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

// 更新产品信息
func (c *Client) UpdateProductInfo(product_ product.Product) (pid_ string, err error) {
	var result struct {
		Pid string
		Error
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := updateProductURL(token)

	if err = c.postJSON(url_, product_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		pid_ = result.Pid
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

// 更变商品上下架状态
func (c *Client) ModProductStatus(mod_ product.Mod) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := ModProductStatusURL(token)

	if err = c.postJSON(url_, mod_, &result); err != nil {
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

// 获取二维码
func (c *Client) GetQRCode(qrinfo_ product.Qrinfo) (qrcode_ product.Qrcode, err error) {
	var result struct {
		Qrcode product.Qrcode
		Error
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := GetQRCodeURL(token)

	if err = c.postJSON(url_, qrinfo_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		qrcode_ = result.Qrcode
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
