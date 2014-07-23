package v2

import (
	"encoding/json"
	"fmt"
	"log"
)

// Digital Ocean domain type A, AAA, CNAME, MX

type Domain struct {
	Name string // The name of the domain itself. This should follow the standard domain format of domain.TLD. For instance, example.com is a valid domain name.
	Ttl  int64  // This value is the time to live for the records on this domain, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested.
	// Reason for assigning a Pointer is that it returns a null value from JSON
	Zone_file string // This attribute contains the complete contents of the zone file for the selected domain. Individual domain record resources should be used to get more granular control over records. However, this attribute can also be used to get information about the SOA record, which is created automatically and is not accessible as an individual record resource.
}
type ListDomainResponse struct {
	Domains []Domain
}

type DomainResponse struct {
	Domain Domain
}

// List of all domains
func (this *Domain) List() (interface{}, error) {
	log.Println("List all Domain: !!: ")
	var listDomainResponse *ListDomainResponse
	jsonBody, err := GetJsonResponse(
		"GET",
		fmt.Sprintf("%v/domains", DIGITALOCEAN_API_URL_PREFIX),
		"")
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &listDomainResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return listDomainResponse, nil
}

// Show Domain Information
func (this *Domain) Show(domainName string) (interface{}, error) {
	log.Println("List all Domain: !!: ")
	var domainResponse *DomainResponse
	jsonBody, err := GetJsonResponse(
		"GET",
		fmt.Sprintf("%v/domains/%v", DIGITALOCEAN_API_URL_PREFIX, domainName),
		"")
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainResponse, nil
}

// Create Domain Information
func (this *Domain) Create(domainName, ipAddress string) (interface{}, error) {
	log.Printf("Create Domain: %v with IP Address: %v "+domainName, ipAddress)
	var domainResponse *DomainResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/domains", DIGITALOCEAN_API_URL_PREFIX),
		fmt.Sprintf(`{"name":"%v","ip_address":"%v"}`, domainName, ipAddress))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainResponse, nil
}

// Delete Domain Information
func (this *Domain) Delete(domainName string) error {
	log.Println("List all Domain: !!: ")
	_, err := GetJsonResponse(
		"DELETE",
		fmt.Sprintf("%v/domains/%v", DIGITALOCEAN_API_URL_PREFIX, domainName),
		"")
	if err != nil {
		return err

	}

	return nil
}
