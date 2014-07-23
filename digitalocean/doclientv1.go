package digitalocean

import (
	"log"
)

// Digital Ocean API v1 Implementation

type DoClientV1 struct {
}

// Create a new Droplet
func (this *DoClientV1) createDroplet() {
	log.Println("Droplet V2 Created!!")
}

// List all droplets
func (this *DoClientV1) listDroplets(authToken string) {
	log.Println("List of all Droplets!!")

}
