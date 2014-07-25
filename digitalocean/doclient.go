// Digital Ocean SDK
package digitalocean

import (
	"errors"
)

/*
Constants to be used accrodd digital ocean sdk
*/
const (
	APIV1 = "v1"
	APIV2 = "v2"
)

func init() {

}

type DoClient struct {
	apiVersion string
	client     DoProvisioner
}

//Create Client for DigitalOcean version 1 or 2
// :apiVersion can be "v2" or "v1"
func NewDoClient(apiVersion string) *DoClient {
	tmpClient, _ := getClient(apiVersion)
	return &DoClient{
		apiVersion: apiVersion,
		client:     tmpClient,
	}

}

/*
Public funciton for DoClient for user to Create Droplet
*/
func (this *DoClient) CreateDroplet(name string, region string, size string, image string) (interface{}, error) {
	return this.client.CreateDroplet(name, region, size, image)

}

/*
Public funciton for DoClient for user to Get List of Droplet
*/
func (this *DoClient) ListDroplets() (interface{}, error) {
	return this.client.ListDroplets()

}

/*
Public funciton for DoClient for user to Delete Droplet
*/
func (this *DoClient) DeleteDroplet(dropletId int64) error {
	return this.client.DeleteDroplet(dropletId)

}

/*
Public funciton To get all the events for Droplet
*/
func (this *DoClient) ActionsDroplet(dropletId int64) (interface{}, error) {
	return this.client.ActionsDroplet(dropletId)

}

/*
Public funciton To reboot droplet
*/
func (this *DoClient) RebootDroplet(dropletId int64) (interface{}, error) {
	return this.client.RebootDroplet(dropletId)

}

/*
Public funciton To PowerCycle droplet
*/
func (this *DoClient) PowerCycleDroplet(dropletId int64) (interface{}, error) {
	return this.client.PowerCycleDroplet(dropletId)

}

/*
Public funciton To Shutdown droplet in graceful way
*/
func (this *DoClient) ShutdownDroplet(dropletId int64) (interface{}, error) {
	return this.client.ShutdownDroplet(dropletId)

}

/*
Public funciton To Power off, hard shutdown
*/
func (this *DoClient) PowerOffDroplet(dropletId int64) (interface{}, error) {
	return this.client.PowerOffDroplet(dropletId)

} /*
Public funciton To Power On
*/
func (this *DoClient) PowerOnDroplet(dropletId int64) (interface{}, error) {
	return this.client.PowerOnDroplet(dropletId)

}

/*
Public funciton To Password Reset
*/
func (this *DoClient) PasswordResetDroplet(dropletId int64) (interface{}, error) {
	return this.client.PasswordResetDroplet(dropletId)
}

/*
Public funciton To Resize Droplet
*/
func (this *DoClient) ResizeDroplet(dropletId int64, size string) (interface{}, error) {
	return this.client.ResizeDroplet(dropletId, size)
}

/*
Public funciton To Restore Droplet from backupImage
*/
func (this *DoClient) RestoreDroplet(dropletId int64, backupImage string) (interface{}, error) {
	return this.client.RestoreDroplet(dropletId, backupImage)
}

/*
Public funciton To Rebuild, this will be done from new image
*/
func (this *DoClient) RebuildDroplet(dropletId int64, image string) (interface{}, error) {
	return this.client.RebuildDroplet(dropletId, image)
}

/*
Public funciton To Rename Droplet name
*/
func (this *DoClient) RenameDroplet(dropletId int64, dropletName string) (interface{}, error) {
	return this.client.RenameDroplet(dropletId, dropletName)
}

/*
Public funciton To Enable IPv6
*/
func (this *DoClient) EnableIpv6Droplet(dropletId int64) (interface{}, error) {
	return this.client.EnableIpv6Droplet(dropletId)
}

/*
Public funciton To Disable Backups
*/
func (this *DoClient) DisableBackupsDroplet(dropletId int64) (interface{}, error) {
	return this.client.DisableBackupsDroplet(dropletId)
}

/*
Public funciton To Enable Private Networking
*/
func (this *DoClient) EnablePrivateNetworkingDroplet(dropletId int64) (interface{}, error) {
	return this.client.EnablePrivateNetworkingDroplet(dropletId)
}

/*
Public funciton To Take Snapshot
*/
func (this *DoClient) SnapshotDroplet(dropletId int64) (interface{}, error) {
	return this.client.SnapshotDroplet(dropletId)
}

/*
Public funciton To List All Domains
*/
func (this *DoClient) ListDomains() (interface{}, error) {
	return this.client.ListDomains()

}

/*
Public funciton To Get Domain Information
*/
func (this *DoClient) ShowDomain(domainName string) (interface{}, error) {
	return this.client.ShowDomain(domainName)
}

/*
Public funciton To Create Domain
*/
func (this *DoClient) CreateDomain(domainName, ipAddress string) (interface{}, error) {
	return this.client.CreateDomain(domainName, ipAddress)
}

/*
Public funciton To Delete Domain
*/
func (this *DoClient) DeleteDomain(domainName string) error {
	return this.client.DeleteDomain(domainName)
}
/*
Public funciton To Create AAAA Domain Record for domain
*/
func (this *DoClient) CreateAAAADomainRecord(domainName, alias, data string) (interface{}, error) {
	return this.client.CreateAAAADomainRecord(domainName, alias, data)
}
/*
Public funciton To Create A Domain Record for domain
*/
func (this *DoClient) CreateADomainRecord(domainName, alias, data string) (interface{}, error) {
	return this.client.CreateADomainRecord(domainName, alias, data)
}
/*
Public funciton To Create MX Domain Record for domain
*/
func (this *DoClient) CreateMXDomainRecord(domainName string,  data string, prority int64) (interface{}, error){
	return this.client.CreateMXDomainRecord(domainName, data, prority)
}
/*
Public funciton To Create TXT Domain Record for domain
*/
func (this *DoClient)  CreateTXTDomainRecord(domainName, name ,data  string) (interface{}, error){
	return this.client.CreateTXTDomainRecord(domainName, name, data)
}
/*
Public funciton To Create NS Domain Record for domain
*/
func (this *DoClient) CreateNSDomainRecord(domainName, data  string) (interface{}, error) {
	return this.client.CreateNSDomainRecord(domainName,  data)
}
/*
Public funciton To Create SRV Domain Record for domain
*/
func (this *DoClient) CreateSRVDomainRecord(domainName string, name string,data  string,  prority int64, port int64, weight int64) (interface{}, error) {
	return this.client.CreateSRVDomainRecord(domainName, name,data ,  prority, port, weight)
}
/*
Public funciton To Show Domain Record for domain
*/
func (this *DoClient) ShowDomainRecord(domainName string, domainRecordId int64)  (interface{}, error){
	return this.client.ShowDomainRecord(domainName, domainRecordId )
}
/*
Public funciton To List all Domain Records for domain
*/
func (this *DoClient) ListDomainRecords(domainName string) (interface{}, error){
	return this.client.ListDomainRecords(domainName)
}
/*
Public funciton To Delete Domain Record for domain
*/
func (this *DoClient) DeleteDomainRecord(domainName string, domainRecordId int64) error{
	return this.client.DeleteDomainRecord(domainName, domainRecordId )
}

/*
Interface to be implemented by Internal Clients based on Differnt versions of
*/
type DoProvisioner interface {
	CreateDroplet(name string, region string, size string, image string) (interface{}, error)
	ListDroplets() (interface{}, error)
	DeleteDroplet(dropletId int64) error
	ActionsDroplet(dropletId int64) (interface{}, error)
	RebootDroplet(dropletId int64) (interface{}, error)
	PowerCycleDroplet(dropletId int64) (interface{}, error)
	ShutdownDroplet(dropletId int64) (interface{}, error)
	PowerOffDroplet(dropletId int64) (interface{}, error)
	PowerOnDroplet(dropletId int64) (interface{}, error)
	PasswordResetDroplet(dropletId int64) (interface{}, error)
	ResizeDroplet(dropletId int64, size string) (interface{}, error)
	RestoreDroplet(dropletId int64, image string) (interface{}, error)
	RebuildDroplet(dropletId int64, image string) (interface{}, error)
	RenameDroplet(dropletId int64 , dropletName string) (interface{}, error)
	EnableIpv6Droplet(dropletId int64) (interface{}, error)
	DisableBackupsDroplet(dropletId int64) (interface{}, error)
	EnablePrivateNetworkingDroplet(dropletId int64) (interface{}, error)
	SnapshotDroplet(dropletId int64) (interface{}, error)
	ListDomains() (interface{}, error)
	ShowDomain(domainName string) (interface{}, error)
	CreateDomain(domainName, ipAddress string) (interface{}, error)
	DeleteDomain(domainName string) error
	CreateAAAADomainRecord(domainName, alias, data string) (interface{}, error)
	CreateADomainRecord(domainName, subdomain, data string) (interface{}, error)
	CreateMXDomainRecord(domainName string,  data string, prority int64) (interface{}, error)
	CreateTXTDomainRecord(domainName, name ,data  string) (interface{}, error)
	CreateNSDomainRecord(domainName, data  string) (interface{}, error)
	CreateSRVDomainRecord(domainName string, name string,data  string,  prority int64, port int64, weight int64) (interface{}, error) 
	ShowDomainRecord(domainName string, domainRecordId int64)  (interface{}, error)
	ListDomainRecords(domainName string) (interface{}, error)
	DeleteDomainRecord(domainName string, domainRecordId int64)  error
}

// this function makes DoClientV... implment DoProvisioner
func getClient(apiVersion string) (DoProvisioner, error) {
	if apiVersion == "v2" {
		return &DoClientV2{}, nil

	} else if apiVersion == "v1" {
		//return &DoClientV1{}
		return nil, errors.New("Api Version v1 is not yet Implemented")

	}
	return nil, errors.New("Valid values are v1 and v2 (Only v2 is implemented)")
}
