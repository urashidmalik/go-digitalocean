package main

import (
	"github.com/urashidmalik/go-digitalocean/digitalocean"
	"github.com/urashidmalik/go-digitalocean/digitalocean/v2"
	"log"
)

var client digitalocean.DoProvisioner
var dropletID int64
var myDomainName string

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	client = digitalocean.NewDoClient(digitalocean.APIV2)
	dropletID = 2170797
	myDomainName = "testmeum.com"
	// dropletCreateTest()
	// dropletListTest()
	// dropletDeleteTest()
	// dropletEventTest()
	// dropletRebootTest()
	// dropletPowerCycleTest()
	// dropletShutdownTest()
	// dropletPowerOffTest()
	// dropletPowerOnTest()
	// dropletPasswordResetTest()
	// dropletResizeTest()
	// dropletRestoreTest()
	// dropletRebuildTest()
	// dropletRenameTest()
	// dropletEnableIpv6Test()
	// dropletDisableBackupsTest()
	// dropletEnablePrivateNetworkingTest()
	// dropletSnapshotTest()
	// domainListTest()
	// showDomainTest()
	// createDomainTest()
	// deleteDomainTest()
	// createAAAADomainRecordTest()
	// createADomainRecordTest()
	// createMXDomainRecordTest()
	// createTXTDomainRecordTest()
	// createSRVDomainRecordTest()
	// createNSDomainRecordTest()
	// listDomainRecordsTest()
	// showDomainRecordTest()
	// deleteDomainRecordTest()
	ListImagesTest()

}

func dropletListTest() {

	if listDropletRes, errL := client.ListDroplets(); errL != nil {
		log.Printf("Houston We have Problem :%v", errL)
	} else {
		log.Printf("Droplet List %v", listDropletRes)
	}

}

func dropletCreateTest() {

	if createDropletRes, err := client.CreateDroplet("Auto2", v2.REGION_NEWYORK_1, "512mb", "3100616"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Created %v", createDropletRes.(*v2.CreateDropletResponse).Droplet)
	}
}
func dropletDeleteTest() {
	var err error
	if err = client.DeleteDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Deleted ")
	}
}
func dropletEventTest() {

	if dropletActionsResponse, err := client.ActionsDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Events %v", dropletActionsResponse.(*v2.DropletActionsResponse))
	}
}
func dropletRebootTest() {
	if dropletActionResponse, err := client.RebootDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Rebooted %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletPowerCycleTest() {

	if dropletActionResponse, err := client.PowerCycleDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Power Cycled %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletShutdownTest() {

	if dropletActionResponse, err := client.ShutdownDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Shutdown %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletPowerOffTest() {

	if dropletActionResponse, err := client.PowerOffDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Power Off %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletPowerOnTest() {

	if dropletActionResponse, err := client.PowerOnDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Power On %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletPasswordResetTest() {

	if dropletActionResponse, err := client.PasswordResetDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Password Reset %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletResizeTest() {
	if dropletActionResponse, err := client.ResizeDroplet(dropletID, "1gb"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Resize %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletRestoreTest() {

	if dropletActionResponse, err := client.RestoreDroplet(dropletID, "5101820"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Restore %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletRebuildTest() {

	if dropletActionResponse, err := client.RebuildDroplet(dropletID, "3100616"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Rebuild %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletRenameTest() {

	if dropletActionResponse, err := client.RenameDroplet(dropletID, "Auto1UmairApitest"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Rename %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletEnableIpv6Test() {

	if dropletActionResponse, err := client.EnableIpv6Droplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Enable Ipv6 %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletDisableBackupsTest() {

	if dropletActionResponse, err := client.DisableBackupsDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Disable Backups %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletEnablePrivateNetworkingTest() {

	if dropletActionResponse, err := client.EnablePrivateNetworkingDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Enable Private Networking %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func dropletSnapshotTest() {

	if dropletActionResponse, err := client.SnapshotDroplet(dropletID); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Droplet Snapshot %v", dropletActionResponse.(*v2.DropletActionResponse))
	}
}
func domainListTest() {

	if listDomainResponse, err := client.ListDomains(); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("All Domains %v", listDomainResponse.(*v2.ListDomainResponse))
	}
}

func showDomainTest() {
	var domainResponse interface{}
	var err error
	if domainResponse, err = client.ShowDomain(myDomainName); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain %v", domainResponse.(*v2.DomainResponse).Domain.Name)
	}
}
func createDomainTest() {
	var domainResponse interface{}
	var err error
	if domainResponse, err = client.CreateDomain(myDomainName, "20.195.2.1"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Create Domain %v", domainResponse.(*v2.DomainResponse))
	}
}
func deleteDomainTest() {
	var err error
	if err = client.DeleteDomain(myDomainName); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Deleted Domain ")
	}
}
func createAAAADomainRecordTest() {
	var domainRecordResponse interface{}
	var err error
	if domainRecordResponse, err = client.CreateAAAADomainRecord(myDomainName, "ipv4","2001:db8::ff00:42:8329"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain Record AAAA : %v", domainRecordResponse.(*v2.DomainRecordResponse))
	}
}
func createADomainRecordTest() {
	var domainRecordResponse interface{}
	var err error
	if domainRecordResponse, err = client.CreateADomainRecord(myDomainName, "subdomain", "127.0.0.1"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain Record A: %v", domainRecordResponse.(*v2.DomainRecordResponse))
	}
}
func createMXDomainRecordTest() {
	var domainRecordResponse interface{}
	var err error
	if domainRecordResponse, err = client.CreateMXDomainRecord(myDomainName, "subdomain", 65535); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain Record MX: %v", domainRecordResponse.(*v2.DomainRecordResponse))
	}
}
func createTXTDomainRecordTest() {
	var domainRecordResponse interface{}
	var err error
	if domainRecordResponse, err = client.CreateTXTDomainRecord(myDomainName, "recordname", "arbitrary data here"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain Record TXT: %v", domainRecordResponse.(*v2.DomainRecordResponse))
	}
}
func createSRVDomainRecordTest() {
	var domainRecordResponse interface{}
	var err error
	if domainRecordResponse, err = client.CreateSRVDomainRecord(myDomainName, "servicename", "targethost", 0, 1, 0); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain Record SRV: %v", domainRecordResponse.(*v2.DomainRecordResponse))
	}
}
func createNSDomainRecordTest() {
	var domainRecordResponse interface{}
	var err error
	if domainRecordResponse, err = client.CreateNSDomainRecord(myDomainName, "ns1.urashidmalik.com."); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain Record NS: %v", domainRecordResponse.(*v2.DomainRecordResponse))
	}
}
func listDomainRecordsTest() {
	var listDomainRecordResponse interface{}
	var err error
	if listDomainRecordResponse, err = client.ListDomainRecords(myDomainName); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("List Domain Records: %v", listDomainRecordResponse.(*v2.ListDomainRecordResponse))
	}
}
func showDomainRecordTest() {
	var domainRecordResponse interface{}
	var err error
	ldomainRecordResponse, _ := client.ListDomainRecords(myDomainName)
	myDomainRecId := ldomainRecordResponse.(*v2.ListDomainRecordResponse).Domain_records[0].Id

	if domainRecordResponse, err = client.ShowDomainRecord(myDomainName, myDomainRecId); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Show Domain Record: %v", domainRecordResponse.(*v2.DomainRecordResponse))
	}
}
func deleteDomainRecordTest() {
	var err error
	ldomainRecordResponse, _ := client.ListDomainRecords(myDomainName)
	myDomainRecId := ldomainRecordResponse.(*v2.ListDomainRecordResponse).Domain_records[0].Id
	if err = client.DeleteDomainRecord(myDomainName, myDomainRecId); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain Record Deleted: ")
	}
}
func ListImagesTest() {
	if lListImageResponse, err := client.ListImages(); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		ImageArray := lListImageResponse.(*v2.ListImageRepose).Images
		for _, img := range ImageArray {
			log.Printf("Images Info: %v", img)

		}
	}

}
