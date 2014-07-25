package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type DomainRecord struct {
	Id       int64  //	The unique id for the individual record.
	Type     string // The DNS record type (A, MX, CNAME, etc).
	Name     string // The host name, alias, or service being defined by the record. See the [domain record] object to find out more.
	Data     string // Variable data depending on record type. See the [domain record] object for more detail on each record type.
	Priority int64  //	The priority of the host (for SRV and MX records. null otherwise).
	Port     int64  // The port that the service is accessible on (for SRV records only. null otherwise).
	Weight   int64  //The weight of records with the same priority (for SRV records only. null otherwise).
}
type DomainRecordResponse struct {
	Domain_Record DomainRecord
}
type ListDomainRecordResponse struct {
	Domain_records []DomainRecord
}
// Create AAAA Domain Record
func (this *DomainRecord) CreateAAAARecord(domainName, alias, data string) (interface{}, error) {
	log.Printf("Create Domain Record AAA: %v for Domain: %v ", alias, domainName)
	var domainRecordResponse *DomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/domains/%v/records", DIGITALOCEAN_API_URL_PREFIX, domainName),
		fmt.Sprintf(`{
		  "name": "%v",
		  "data": "%v",
		  "type": "%v"
		}`, alias, data, "AAAA"))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainRecordResponse, nil

}

// Create A Domain Record
func (this *DomainRecord) CreateARecord(domainName, subdomain, data string) (interface{}, error) {
	log.Printf("Create Domain Record A: %v for Domain: %v ", subdomain, domainName)
	var domainRecordResponse *DomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/domains/%v/records", DIGITALOCEAN_API_URL_PREFIX, domainName),
		fmt.Sprintf(`{
		  "name": "%v",
		  "data": "%v",
		  "type": "%v"
		}`, subdomain, data, "A"))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainRecordResponse, nil

}

// Create MX Domain Record
func (this *DomainRecord) CreateMXRecord(domainName string, data string, prority int64) (interface{}, error) {
	log.Printf("Create Domain Record MX: %v for Domain: %v ", data, domainName)
	if !this.inRange(prority, 0, 65535){
		return nil, errors.New("Value of Prority should be in range 0 to 65535")
	}
	var domainRecordResponse *DomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/domains/%v/records", DIGITALOCEAN_API_URL_PREFIX, domainName),
		fmt.Sprintf(`{
		  "data": "%v",
		  "priority": %v,
		  "type": "%v"
		}`, data, prority, "MX"))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainRecordResponse, nil

}
// Create TXT Domain Record
func (this *DomainRecord) CreateTXTRecord(domainName, name ,data  string) (interface{}, error) {
	log.Printf("Create Domain Record TXT: %v for Domain: %v ", name, domainName)
	
	var domainRecordResponse *DomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/domains/%v/records", DIGITALOCEAN_API_URL_PREFIX, domainName),
		fmt.Sprintf(`{
		  "name": "%v",
		  "data": "%v",
		  "type": "%v"
		}`, name, data, "TXT"))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainRecordResponse, nil

}
// Create NS Domain Record
func (this *DomainRecord) CreateNSRecord(domainName, data  string) (interface{}, error) {
	log.Printf("Create Domain Record NS: %v for Domain: %v ", data, domainName)
	
	var domainRecordResponse *DomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/domains/%v/records", DIGITALOCEAN_API_URL_PREFIX, domainName),
		fmt.Sprintf(`{
		  "data": "%v",
		  "type": "%v"
		}`, data, "NS"))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainRecordResponse, nil

}
// Create SRV Record Domain Record
func (this *DomainRecord) CreateSRVRecord(domainName string, name string,data  string,  prority int64, port int64, weight int64) (interface{}, error) {
	log.Printf("Create Domain Record SRV: %v for Domain: %v ", name, domainName)
	if !this.inRange(prority, 0, 65535){
		return nil, errors.New("Value of Prority should be in range 0 to 65535")
	}
	if !this.inRange(port, 1, 65535){
		return nil, errors.New("Value of Port should be in range 1 to 65535")
	}
	if !this.inRange(weight, 0, 65535){
		return nil, errors.New("Value of Weight should be in range 0 to 65535")
	}
	var domainRecordResponse *DomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/domains/%v/records", DIGITALOCEAN_API_URL_PREFIX, domainName),
		fmt.Sprintf(`{
			"type": "%v", 
			"name": "%v", 
			"data": "%v", 
			"priority":  %v, 
			"port":  %v, 
			"weight":  %v
		}`, "SRV", name,data,prority,port,weight))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainRecordResponse, nil

}

// Show Domain Record
func (this *DomainRecord) Show(domainName string, domainRecordId int64) (interface{}, error) {
	log.Printf("Domain Record NS: ")
	
	var domainRecordResponse *DomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"GET",
		fmt.Sprintf("%v/domains/%v/records/%v", DIGITALOCEAN_API_URL_PREFIX, domainName, domainRecordId),
		"")
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &domainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return domainRecordResponse, nil

}
// List all Domain Records
func (this *DomainRecord) List(domainName string) (interface{}, error) {
	log.Printf("List Domain Records")
	
	var listDomainRecordResponse *ListDomainRecordResponse
	jsonBody, err := GetJsonResponse(
		"GET",
		fmt.Sprintf("%v/domains/%v/records", DIGITALOCEAN_API_URL_PREFIX, domainName),
		"")
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &listDomainRecordResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return listDomainRecordResponse, nil

}
// Delete Domain Record
func (this *DomainRecord) Delete(domainName string, domainRecordId int64) error {
	log.Printf("Domain Record NS: ")
	
	_, err := GetJsonResponse(
		"DELETE",
		fmt.Sprintf("%v/domains/%v/records/%v", DIGITALOCEAN_API_URL_PREFIX, domainName, domainRecordId),
		"")
	if err != nil {
		return err

	}
	
	return nil

}
// check if value is in specified range
func (this *DomainRecord) inRange(val, start, end int64)bool {
	if !((val >= start) && (val <= end)) {
		return false
	}else {
		return true
	}
	
	

}
 
 