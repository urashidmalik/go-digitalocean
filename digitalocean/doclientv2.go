package digitalocean

import (
	"github.com/urashidmalik/go-digitalocean/digitalocean/v2"
	"log"
)

// Digital Ocean API v2 Implementation

type DoClientV2 struct {
	droplet      v2.Droplet
	domain       v2.Domain
	domainRecord v2.DomainRecord
}

func (this *DoClientV2) CreateDroplet(name string, region string, size string, image string) (interface{}, error) {
	log.Println("Initiation Droplet Creation!!")
	return this.droplet.Create(name, region, size, image)
}
func (this *DoClientV2) ListDroplets() (interface{}, error) {
	return this.droplet.List()

}
func (this *DoClientV2) DeleteDroplet(dropletId int64) error {
	return this.droplet.Delete(dropletId)

}
func (this *DoClientV2) ActionsDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.Actions(dropletId)

}
func (this *DoClientV2) RebootDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.Reboot(dropletId)

}
func (this *DoClientV2) PowerCycleDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.PowerCycle(dropletId)

}
func (this *DoClientV2) ShutdownDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.Shutdown(dropletId)

}
func (this *DoClientV2) PowerOffDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.PowerOff(dropletId)

}
func (this *DoClientV2) PowerOnDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.PowerOn(dropletId)

}

func (this *DoClientV2) PasswordResetDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.PasswordReset(dropletId)
}

func (this *DoClientV2) ResizeDroplet(dropletId int64, size string) (interface{}, error) {
	return this.droplet.Resize(dropletId, size)
}

func (this *DoClientV2) RestoreDroplet(dropletId int64, backupImage string) (interface{}, error) {
	return this.droplet.Restore(dropletId, backupImage)
}

func (this *DoClientV2) RebuildDroplet(dropletId int64, image string) (interface{}, error) {
	return this.droplet.Rebuild(dropletId, image)
}

func (this *DoClientV2) RenameDroplet(dropletId int64, dropletName string) (interface{}, error) {
	return this.droplet.Rename(dropletId, dropletName)
}

func (this *DoClientV2) EnableIpv6Droplet(dropletId int64) (interface{}, error) {
	return this.droplet.EnableIpv6(dropletId)
}

func (this *DoClientV2) DisableBackupsDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.DisableBackups(dropletId)
}

func (this *DoClientV2) EnablePrivateNetworkingDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.EnablePrivateNetworking(dropletId)
}
func (this *DoClientV2) SnapshotDroplet(dropletId int64) (interface{}, error) {
	return this.droplet.Snapshot(dropletId)
}
func (this *DoClientV2) ListDomains() (interface{}, error) {
	return this.domain.List()
}
func (this *DoClientV2) ShowDomain(domainName string) (interface{}, error) {
	return this.domain.Show(domainName)
}
func (this *DoClientV2) CreateDomain(domainName, ipAddress string) (interface{}, error) {
	return this.domain.Create(domainName, ipAddress)
}

func (this *DoClientV2) DeleteDomain(domainName string) error {
	return this.domain.Delete(domainName)
}
func (this *DoClientV2) CreateAAAADomainRecord(domainName, alias, data string) (interface{}, error) {
	return this.domainRecord.CreateAAAARecord(domainName, alias, data)
}
func (this *DoClientV2) CreateADomainRecord(domainName, subdomain, data string) (interface{}, error) {
	return this.domainRecord.CreateARecord(domainName, subdomain, data)
}
func (this *DoClientV2) CreateMXDomainRecord(domainName string, data string, prority int64) (interface{}, error) {
	return this.domainRecord.CreateMXRecord(domainName, data, prority)
}
func (this *DoClientV2) CreateTXTDomainRecord(domainName, name, data string) (interface{}, error) {
	return this.domainRecord.CreateTXTRecord(domainName, name, data)
}
func (this *DoClientV2) CreateNSDomainRecord(domainName, data string) (interface{}, error) {
	return this.domainRecord.CreateNSRecord(domainName, data)
}
func (this *DoClientV2) CreateSRVDomainRecord(domainName string, name string, data string, prority int64, port int64, weight int64) (interface{}, error) {
	return this.domainRecord.CreateSRVRecord(domainName, name, data, prority, port, weight)
}
func (this *DoClientV2) ShowDomainRecord(domainName string, domainRecordId int64) (interface{}, error) {
	return this.domainRecord.Show(domainName, domainRecordId)
}
func (this *DoClientV2) ListDomainRecords(domainName string) (interface{}, error) {
	return this.domainRecord.List(domainName)
}
func (this *DoClientV2) DeleteDomainRecord(domainName string, domainRecordId int64) error {
	return this.domainRecord.Delete(domainName, domainRecordId)
}
