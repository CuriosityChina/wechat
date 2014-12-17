package card

import (
	"github.com/chanxuehong/wechat/mp/card"
)

func (c *Client) CardGetColors() (colors []card.Color, err error) {
	var result struct {
		Colors []card.Color `json:"colors"`
		Error
	}
	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := cardGetColorsUrl(token)
	if err = c.postJSON(_url, nil, &result); err != nil {
		return
	}
	switch result.ErrCode {
	case errCodeOK:
		colors = result.Colors
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
