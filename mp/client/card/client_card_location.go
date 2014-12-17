package card

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/card"
)

func (c *Client) CardLocationBatchAdd(locations []card.LocationRequest) (locationList []int, err error) {

	var request = struct {
		LocationList []card.LocationRequest `json:"location_list"`
	}{
		LocationList: locations,
	}
	var result struct {
		Error
		LocationIdList []int `json:"location_id_list"`
	}
	hasRetry := false

RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}

	_url := cardLocationBatchAddUrl(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}
	fmt.Println(result)
	switch result.ErrCode {
	case errCodeOK:
		locationList = result.LocationIdList
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

func (c *Client) CardLocationBatchGet(offset int, count int) (location_list []card.LocationResponse, err error) {
	var request = struct {
		Offset int `json:"offset"`
		Count  int `json:"count"`
	}{
		Offset: offset,
		Count:  count,
	}
	var result struct {
		Error
		LocationList []card.LocationResponse `json:"location_list"`
		Count        int                     `json:"count"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}

	_url := cardLocationBatchGetUrl(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}
	fmt.Println(result)

	switch result.ErrCode {
	case errCodeOK:
		location_list = result.LocationList
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
