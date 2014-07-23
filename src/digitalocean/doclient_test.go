package digitalocean

import (
	"log"
	"testing"
	"v2"
)

func TestList(t *testing.T) {
	client := digitalocean.NewDoClient()

	var listDropletRes *v2.ListDropletResponse
	var err error
	if listDropletRes, err = client.ListDroplets(); err != nil {
		log.Printf("Oppsi We have Issue :%v", err)
	} else {
		log.Printf("Data %v", len(listDropletRes.Droplets))
	}

}
