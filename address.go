package namecheap

import "net/url"

const (
	addressGetList = "namecheap.users.address.getList"
)

type AddressGetListResult struct {
	ID   int    `xml:"AddressId,attr"`
	Name string `xml:"AddressName,attr"`
}

func (client *Client) AddressGetList() ([]AddressGetListResult, error) {
	requestInfo := &ApiRequest{
		command: addressGetList,
		method:  "POST",
		params:  url.Values{},
	}

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}

	return resp.AddressGetList, nil
}
