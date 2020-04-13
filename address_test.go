package namecheap

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestAddressGetList(t *testing.T) {
	setup()
	defer teardown()

	respXML := `
    <?xml version="1.0" encoding="UTF-8"?>
    <ApiResponse xmlns="http://api.namecheap.com/xml.response" Status="OK">
      <Errors />
      <RequestedCommand>namecheap.users.address.getList</RequestedCommand>
      <CommandResponse Type="namecheap.users.address.getList">
        <AddressGetListResult>
          <List AddressId="49" AddressName="newaddress_test" />
          <List AddressId="21" AddressName="newaddress_test2" />
        </AddressGetListResult>
      </CommandResponse>
      <Server>SERVER-NAME</Server>
      <GMTTimeDifference>+5</GMTTimeDifference>
      <ExecutionTime>0.047</ExecutionTime>
    </ApiResponse>`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		correctParams := fillDefaultParams(url.Values{})
		correctParams.Set("Command", "namecheap.users.address.getList")
		testBody(t, r, correctParams)
		testMethod(t, r, "POST")
		fmt.Fprint(w, respXML)
	})
	addresses, err := client.AddressGetList()

	if err != nil {
		t.Errorf("AddressGetList returned error: %v", err)
	}

	// AddressGetListResult we expect, given the respXML above
	wantD := []AddressGetListResult{{
		ID:   49,
		Name: "newaddress_test",
	}, {
		ID:   21,
		Name: "newaddress_test2",
	}}

	if !reflect.DeepEqual(addresses, wantD) {
		t.Errorf("AddressGetListResult returned:\n%+v, want:\n%+v", addresses, wantD)
	}
}

func TestAddressGetInfo(t *testing.T) {
	setup()
	defer teardown()

	respXML := `
    <?xml version="1.0" encoding="UTF-8"?>
    <ApiResponse xmlns="http://api.namecheap.com/xml.response" Status="OK">
      <Errors />
      <RequestedCommand>namecheap.users.address.getInfo</RequestedCommand>
      <CommandResponse Type="namecheap.users.address.getInfo">
        <GetAddressInfoResult>
          <AddressId>49</AddressId>
          <UserName>apisample</UserName>
          <AddressName>newaddress_test</AddressName>
          <Default_YN>false</Default_YN>
          <FirstName>api</FirstName>
          <LastName>sample</LastName>
          <JobTitle>jtitle</JobTitle>
          <Organization>org_Test</Organization>
          <Address1>add1test</Address1>
          <Address2>add2test</Address2>
          <City>city_test</City>
          <StateProvince>state_test</StateProvince>
          <StateProvinceChoice>province_test</StateProvinceChoice>
          <Zip>641004</Zip>
          <Country>IN</Country>
          <Phone>91.1111111111</Phone>
          <Fax>91.11111111</Fax>
          <PhoneExt>23423</PhoneExt>
          <EmailAddress>contact@apisample.com</EmailAddress>
        </GetAddressInfoResult>
      </CommandResponse>
      <Server>SERVER-NAME</Server>
      <GMTTimeDifference>+5</GMTTimeDifference>
      <ExecutionTime>0.047</ExecutionTime>
    </ApiResponse>`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		correctParams := fillDefaultParams(url.Values{})
		correctParams.Set("Command", "namecheap.users.address.getInfo")
		correctParams.Set("AddressId", "49")
		testBody(t, r, correctParams)
		testMethod(t, r, "POST")
		fmt.Fprint(w, respXML)
	})

	result, err := client.AddressGetInfo(49)
	if err != nil {
		t.Errorf("AddressGetInfo returned error: %v", err)
	}

	// AddressGetInfoResult we expect, given the respXML above
	wantD := &AddressGetInfoResult{
		ID:                  49,
		UserName:            "apisample",
		Name:                "newaddress_test",
		Default:             false,
		JobTitle:            "jtitle",
		Organization:        "org_Test",
		FirstName:           "api",
		LastName:            "sample",
		Address1:            "add1test",
		Address2:            "add2test",
		City:                "city_test",
		StateProvince:       "state_test",
		StateProvinceChoice: "province_test",
		PostalCode:          "641004",
		Country:             "IN",
		Phone:               "91.1111111111",
		Fax:                 "91.11111111",
		PhoneExt:            "23423",
		EmailAddress:        "contact@apisample.com",
	}

	if !reflect.DeepEqual(result, wantD) {
		t.Errorf("AddressGetListResult returned:\n%+v, want:\n%+v", result, wantD)
	}
}
