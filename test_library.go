package main

import (
	"./digitalocean"
	"./digitalocean/v2"
	"log"
)

var client digitalocean.DoProvisioner
var dropletID string

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	client = digitalocean.NewDoClient(digitalocean.APIV2)
	dropletID = "2151455"
	//dropletCreateTest()
	//dropletListTest()
	//dropletDeleteTest()
	//dropletEventTest()
	//dropletRebootTest()
	//dropletPowerCycleTest()
	//dropletShutdownTest()
	//dropletPowerOffTest()
	//dropletPowerOnTest()
	//dropletPasswordResetTest()
	//dropletResizeTest()
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
	 deleteDomainTest()

	//http://attilaolah.eu/2013/11/29/json-decoding-in-go/

}

func dropletListTest() {
	
	if listDropletRes, errL := client.ListDroplets(); errL != nil {
		log.Printf("Houston We have Problem :%v", errL)
	} else {
		log.Printf("Droplet List %v", listDropletRes.(*v2.ListDropletResponse).Droplets[0].Id)
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
	if domainResponse, err = client.ShowDomain("mydn.com"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Domain %v", domainResponse.(*v2.DomainResponse).Domain.Name)
	}
}
func createDomainTest() {
	var domainResponse interface{}
	var err error
	if domainResponse, err = client.CreateDomain("mydnauto.com", "20.195.2.1"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Create Domain %v", domainResponse.(*v2.DomainResponse))
	}
}
func deleteDomainTest() {
	var err error
	if err = client.DeleteDomain("mydnauto.com"); err != nil {
		log.Printf("Houston We have Problem :%v", err)
	} else {
		log.Printf("Deleted Domain ")
	}
}
