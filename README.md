go-digitalocean
===============

GO digital ocean API for version 2 and version 1. As of now only version 2 is implemented

# Usage
use go get github/urashidmalik/go-digitalocean to get the library. 

#### Configuration
* Copy github.com/urashidmalik/go-digitalocean/config.json file 
* Update your digitalocean API KEY
* Place file in your $GOPATH
* E.G 


>**myProj**  
>>**├──config.json**  
>>**├──bin**  
>>**├──pkg**  
>>**└──src**  
>>>> **└─────github.com/**  

      

#### Sample config.json file
	
	{
		"API_TOKEN":"API_KEY_GOES_HERE"
	}

#### Importing and Using go-digitalocean Library in you program
###### Basic Usage
  
	package main

	import (
		"github.com/urashidmalik/go-digitalocean/digitalocean"
		"github.com/urashidmalik/go-digitalocean/digitalocean/v2"
		"log"
	)

	func main() {
		var client digitalocean.DoProvisioner
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		client = digitalocean.NewDoClient(digitalocean.APIV2)
		
		// Basic Usage
		listDropletResponseBasic, _ := client.ListDroplets()
		// Please remember you will get Interface Object which you can cast to *v2.ListDropletResponse
		log.Printf("Droplet List %v", listDropletRes.(*v2.ListDropletResponse))
	}

###### Usage with Error Handling

  
	package main

	import (
	    "github.com/urashidmalik/go-digitalocean/digitalocean"
	    "github.com/urashidmalik/go-digitalocean/digitalocean/v2"
	    "log"
	)

	func main() {
	    var client digitalocean.DoProvisioner
	    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	    client = digitalocean.NewDoClient(digitalocean.APIV2)
	        // Usage with error Handling
	    if listDropletRes, errL := client.ListDroplets(); errL != nil {
	        log.Printf("Houston We have Problem :%v", errL)
	    } else {
	        log.Printf("Droplet List %v", listDropletRes.(*v2.ListDropletResponse))
	    }
	}

#### Implemented Functionality (API version 2)

###### Droplet
* Create Droplet
* List all Droplets
* Delete a Droplet
* All Droplet Events
* Reboot a Droplet
* Power Cycle a Droplet
* Shutdown a Droplet
* Power Off a Droplet
* Power On a Droplet
* Password Reset for Droplet
* Resize a Droplet
* Restore a Droplet
* Rebuild a Droplet
* Rename a Droplet
* Enable IPv6 for Droplet
* Disable Backups for Droplet
* Enable Private Networking on Droplet
* Take a Droplet Snapshot

###### Domain
* List all Domains
* Show Domain Information
* Create a Domain
* Delete a Domain

###### Domain Records
* Create AAAA Domain Record
* Create A Domain Record
* Create MX Domain Record
* Create TXT Domain Record
* Create SRV Domain Record
* Create NS Domain Record
* List Domain Records Test
* Show Domain Record Test
* Delete DomainRecordTest