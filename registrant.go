package namecheap

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// Registrant is a struct that contains all the data necesary to register a domain.
// That is to say, every field in this struct is REQUIRED by the namecheap api to
// crate a new domain.
// In order for `addValues` method to work, all fields must remain strings.
type Registrant struct {
	RegistrantFirstName string `xml:"Registrant>FirstName"`
	RegistrantLastName  string `xml:"Registrant>LastName"`

	RegistrantAddress1 string `xml:"Registrant>Address1"`
	RegistrantAddress2 string `xml:"Registrant>Address2"`
	RegistrantCity     string `xml:"Registrant>City"`

	RegistrantStateProvince string `xml:"Registrant>StateProvince"`
	RegistrantPostalCode    string `xml:"Registrant>PostalCode"`
	RegistrantCountry       string `xml:"Registrant>Country"`

	RegistrantPhone        string `xml:"Registrant>Phone"`
	RegistrantEmailAddress string `xml:"Registrant>EmailAddress"`

	TechFirstName string `xml:"Tech>FirstName"`
	TechLastName  string `xml:"Tech>LastName"`

	TechAddress1 string `xml:"Tech>Address1"`
	TechAddress2 string `xml:"Tech>Address2"`
	TechCity     string `xml:"Tech>City"`

	TechStateProvince string `xml:"Tech>StateProvince"`
	TechPostalCode    string `xml:"Tech>PostalCode"`
	TechCountry       string `xml:"Tech>Country"`

	TechPhone        string `xml:"Tech>Phone"`
	TechEmailAddress string `xml:"Tech>EmailAddress"`

	AdminFirstName string `xml:"Admin>FirstName"`
	AdminLastName  string `xml:"Admin>LastName"`

	AdminAddress1 string `xml:"Admin>Address1"`
	AdminAddress2 string `xml:"Admin>Address2"`
	AdminCity     string `xml:"Admin>City"`

	AdminStateProvince string `xml:"Admin>StateProvince"`
	AdminPostalCode    string `xml:"Admin>PostalCode"`
	AdminCountry       string `xml:"Admin>Country"`

	AdminPhone        string `xml:"Admin>Phone"`
	AdminEmailAddress string `xml:"Admin>EmailAddress"`

	AuxBillingFirstName string `xml:"AuxBilling>FirstName"`
	AuxBillingLastName  string `xml:"AuxBilling>LastName"`

	AuxBillingAddress1 string `xml:"AuxBilling>Address1"`
	AuxBillingAddress2 string `xml:"AuxBilling>Address2"`
	AuxBillingCity     string `xml:"AuxBilling>City"`

	AuxBillingStateProvince string `xml:"AuxBilling>StateProvince"`
	AuxBillingPostalCode    string `xml:"AuxBilling>PostalCode"`
	AuxBillingCountry       string `xml:"AuxBilling>Country"`

	AuxBillingPhone        string `xml:"AuxBilling>Phone"`
	AuxBillingEmailAddress string `xml:"AuxBilling>EmailAddress"`
}

// newRegistrant return a new registrant where all the required fields are the same.
// Feel free to change them as needed
func newRegistrant(
	firstName, lastName,
	addr1, addr2,
	city, state, postalCode, country,
	phone, email string,
) *Registrant {
	return &Registrant{
		RegistrantFirstName:     firstName,
		RegistrantLastName:      lastName,
		RegistrantAddress1:      addr1,
		RegistrantAddress2:      addr2,
		RegistrantCity:          city,
		RegistrantStateProvince: state,
		RegistrantPostalCode:    postalCode,
		RegistrantCountry:       country,
		RegistrantPhone:         phone,
		RegistrantEmailAddress:  email,
		TechFirstName:           firstName,
		TechLastName:            lastName,
		TechAddress1:            addr1,
		TechAddress2:            addr2,
		TechCity:                city,
		TechStateProvince:       state,
		TechPostalCode:          postalCode,
		TechCountry:             country,
		TechPhone:               phone,
		TechEmailAddress:        email,
		AdminFirstName:          firstName,
		AdminLastName:           lastName,
		AdminAddress1:           addr1,
		AdminAddress2:           addr2,
		AdminCity:               city,
		AdminStateProvince:      state,
		AdminPostalCode:         postalCode,
		AdminCountry:            country,
		AdminPhone:              phone,
		AdminEmailAddress:       email,
		AuxBillingFirstName:     firstName,
		AuxBillingLastName:      lastName,
		AuxBillingAddress1:      addr1,
		AuxBillingAddress2:      addr2,
		AuxBillingCity:          city,
		AuxBillingStateProvince: state,
		AuxBillingPostalCode:    postalCode,
		AuxBillingCountry:       country,
		AuxBillingPhone:         phone,
		AuxBillingEmailAddress:  email,
	}
}

// addValues adds the fields of this struct to the passed in url.Values.
// It is important that all the fields of Registrant remain string type.
func (reg *Registrant) addValues(u url.Values) error {
	if u == nil {
		return errors.New("nil value passed as url.Values")
	}

	val := reflect.ValueOf(*reg)
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldName := t.Field(i).Name
		field := val.Field(i).String()
		if ty := val.Field(i).Kind(); ty != reflect.String {
			return fmt.Errorf(
				"Registrant cannot have types that aren't string; %s is type %s",
				fieldName, ty,
			)
		}
		if field == "" {
			if strings.Contains(fieldName, "ddress2") {
				continue
			}

			return fmt.Errorf("Field %s cannot be empty", fieldName)
		}

		u.Set(fieldName, fmt.Sprintf("%v", field))
	}

	return nil
}
