package card

// 注意：如果填写的openid和username不合法，系统也不会自动检查
func (c *Client) CardTestWhiteListSet(openid []string, username []string) (err error) {
	var request = struct {
		Openid   []string `json:"openid"`
		Username []string `json:"username"`
	}{
		Openid:   openid,
		Username: username,
	}
	var result struct {
		Error
	}
	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url := cardTestWhiteListSetUrl(token)
	if err = c.postJSON(url, &request, &result); err != nil {
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
		err = &result.Error
		return
	}

}
