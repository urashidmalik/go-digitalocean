package digitalocean

import (
	"github.com/urashidmalik/go-digitalocean/digitalocean/v2"
	"log"
)

// Digital Ocean API v2 Implementation

type DoClientV2 struct {
	droplet v2.Droplet
	domain  v2.Domain
}

func (this *DoClientV2) CreateDroplet(name string, region string, size string, image string) (interface{}, error) {
	log.Println("Initiation Droplet Creation!!")
	return this.droplet.Create(name, region, size, image)
}
func (this *DoClientV2) ListDroplets() (interface{}, error) {
	return this.droplet.List()

}
func (this *DoClientV2) DeleteDroplet(dropletId string) error {
	return this.droplet.Delete(dropletId)

}
func (this *DoClientV2) ActionsDroplet(dropletId string) (interface{}, error) {
	return this.droplet.Actions(dropletId)

}
func (this *DoClientV2) RebootDroplet(dropletId string) (interface{}, error) {
	return this.droplet.Reboot(dropletId)

}
func (this *DoClientV2) PowerCycleDroplet(dropletId string) (interface{}, error) {
	return this.droplet.PowerCycle(dropletId)

}
func (this *DoClientV2) ShutdownDroplet(dropletId string) (interface{}, error) {
	return this.droplet.Shutdown(dropletId)

}
func (this *DoClientV2) PowerOffDroplet(dropletId string) (interface{}, error) {
	return this.droplet.PowerOff(dropletId)

}
func (this *DoClientV2) PowerOnDroplet(dropletId string) (interface{}, error) {
	return this.droplet.PowerOn(dropletId)

}

func (this *DoClientV2) PasswordResetDroplet(dropletId string) (interface{}, error) {
	return this.droplet.PasswordReset(dropletId)
}

func (this *DoClientV2) ResizeDroplet(dropletId, size string) (interface{}, error) {
	return this.droplet.Resize(dropletId, size)
}

func (this *DoClientV2) RestoreDroplet(dropletId, backupImage string) (interface{}, error) {
	return this.droplet.Restore(dropletId, backupImage)
}

func (this *DoClientV2) RebuildDroplet(dropletId, image string) (interface{}, error) {
	return this.droplet.Rebuild(dropletId, image)
}

func (this *DoClientV2) RenameDroplet(dropletId, dropletName string) (interface{}, error) {
	return this.droplet.Rename(dropletId, dropletName)
}

func (this *DoClientV2) EnableIpv6Droplet(dropletId string) (interface{}, error) {
	return this.droplet.EnableIpv6(dropletId)
}

func (this *DoClientV2) DisableBackupsDroplet(dropletId string) (interface{}, error) {
	return this.droplet.DisableBackups(dropletId)
}

func (this *DoClientV2) EnablePrivateNetworkingDroplet(dropletId string) (interface{}, error) {
	return this.droplet.EnablePrivateNetworking(dropletId)
}
func (this *DoClientV2) SnapshotDroplet(dropletId string) (interface{}, error) {
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
