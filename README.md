go-digitalocean
===============

GO digital ocean API for version 2 and version 1. As of now only version 2 is implemented

# Usage
use go get github/urashidmalik/go-digitalocean to get the library. For Quick reference use test_library.go file to understand how to use the library.

#### Configuration
You need to open config.json file and add your digitalocean API KEY
	
	{
		"API_TOKEN":"API_KEY_GOES_HERE"
	}

#### Importing and Using go-digitalocean Library in you program
###### Basic Usage

	package main

	import (
		"./github.com/urashidmalik/go-digitalocean/digitalocean"
		"./github.com/urashidmalik/go-digitalocean/digitalocean/v2"
		"log"
	)

	func main() {
		// Basic Usage
		listDropletResponseBasic, _ := client.ListDroplets()
		// Please remember you will get Interface Object which you can cast to *v2.listDropletResponseBasic
		log.Printf("Droplet List %v", listDropletRes.(*v2.listDropletResponseBasic))

	}

###### Usage with Error Handling

	package main

	import (
		"github.com/urashidmalik/go-digitalocean/digitalocean"
		"github.com/urashidmalik/go-digitalocean/digitalocean/v2"
		"log"
	)

	func main() {
		// Usage with error Handling
		if listDropletRes, errL := client.ListDroplets(); errL != nil {
			log.Printf("Houston We have Problem :%v", errL)
		} else {
			log.Printf("Droplet List %v", listDropletRes.(*v2.ListDropletResponse).Droplets[0].Id)
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