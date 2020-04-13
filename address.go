package namecheap

import (
	"fmt"
	"net/url"
)

const (
	addressGetList = "namecheap.users.address.getList"
	addressGetInfo = "namecheap.users.address.getInfo"
)

type AddressGetListResult struct {
	ID   int    `xml:"AddressId,attr"`
	Name string `xml:"AddressName,attr"`
}

type AddressGetInfoResult struct {
	ID       int    `xml:"AddressId"`
	UserName string `xml:"UserName"`
	Name     string `xml:"AddressName"`
	Default  bool   `xml:"Default_YN"`

	FirstName string `xml:"FirstName"`
	LastName  string `xml:"LastName"`

	JobTitle     string `xml:"JobTitle"`
	Organization string `xml:"Organization"`

	Address1 string `xml:"Address1"`
	Address2 string `xml:"Address2"`
	City     string `xml:"City"`

	StateProvince       string `xml:"StateProvince"`
	StateProvinceChoice string `xml:"StateProvinceChoice"`
	PostalCode          string `xml:"Zip"`
	Country             string `xml:"Country"`

	Phone    string `xml:"Phone"`
	Fax      string `xml:"Fax"`
	PhoneExt string `xml:"PhoneExt"`

	EmailAddress string `xml:"EmailAddress"`
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

func (client *Client) AddressGetInfo(addressID int) (*AddressGetInfoResult, error) {
	requestInfo := &ApiRequest{
		command: addressGetInfo,
		method:  "POST",
		params:  url.Values{},
	}

	requestInfo.params.Set("AddressId", fmt.Sprintf("%d", addressID))

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}

	return resp.AddressGetInfo, nil
}
