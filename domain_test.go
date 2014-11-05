package namecheap

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDomain_DomainsGetList(t *testing.T) {
	setup()
	defer teardown()

	response_xml := `
    <?xml version="1.0" encoding="utf-8"?>
    <ApiResponse Status="OK" xmlns="http://api.namecheap.com/xml.response">
      <Errors />
      <Warnings />
      <RequestedCommand>namecheap.domains.getList</RequestedCommand>
      <CommandResponse Type="namecheap.domains.getList">
        <DomainGetListResult>
          <Domain ID="57579" Name="example.com" User="anUser" Created="11/04/2014" Expires="11/04/2015" IsExpired="false" IsLocked="false" AutoRenew="false" WhoisGuard="ENABLED" />
        </DomainGetListResult>
        <Paging>
          <TotalItems>12</TotalItems>
          <CurrentPage>1</CurrentPage>
          <PageSize>20</PageSize>
        </Paging>
      </CommandResponse>
      <Server>WEB1-SANDBOX1</Server>
      <GMTTimeDifference>--5:00</GMTTimeDifference>
      <ExecutionTime>0.009</ExecutionTime>
    </ApiResponse>`

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// verify that the URL exactly matches...brittle, I know.
		correct_url := "/?ApiKey=anToken&ApiUser=anApiUser&ClientIp=127.0.0.1&Command=namecheap.domains.getList&UserName=anUser"
		if r.URL.String() != correct_url {
			t.Errorf("URL = %v, want %v", r.URL, correct_url)
		}
		testMethod(t, r, "GET")
		fmt.Fprint(w, response_xml)
	})

	domains, err := client.DomainsGetList()

	if err != nil {
		t.Errorf("DomainsGetList returned error: %v", err)
	}

	// DomainGetListResult we expect, given the response_xml above
	want := []DomainGetListResult{{
		ID:         57579,
		Name:       "example.com",
		User:       "anUser",
		Created:    "11/04/2014",
		Expires:    "11/04/2015",
		IsExpired:  false,
		IsLocked:   false,
		AutoRenew:  false,
		WhoisGuard: "ENABLED",
	}}

	if !reflect.DeepEqual(domains, want) {
		t.Errorf("Domains returned %+v, want %+v", domains, want)
	}
}
