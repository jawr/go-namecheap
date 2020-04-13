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
		t.Errorf("AddressGetListResult returned %+v, want %+v", addresses, wantD)
	}
}
