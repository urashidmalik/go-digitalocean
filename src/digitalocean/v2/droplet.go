package v2

import (
	"encoding/json"
	"fmt"
	"log"
)

// List Droplet Response Structure

type ListDropletResponse struct {
	Droplets []Droplet
}

// Create Droplet Response Structure

type CreateDropletResponse struct {
	Droplet Droplet
}

// Droplet structure

type Droplet struct {
	Id           int64   // A unique identifier for each Droplet instance. This is automatically generated upon Droplet creation.
	Name         string  // The human-readable name set for the Droplet instance.
	Region       Region  // The region that the Droplet instance is deployed in. When setting a region, the value should be the slug identifier for the region. When you query a Droplet, the entire region object will be returned.
	Image        Image   // The base image used to create the Droplet instance. When setting an image, the value is set to the image id or slug. When querying the Droplet, the entire image object will be returned.
	Size         Size    // The size of the Droplet instance. When setting a size, the value should be the slug identifier for a particular size. When querying the Droplet, the entire size object will be returned.
	Locked       bool    // A boolean value indicating whether the Droplet has been locked, preventing actions by users.
	Created_at   string  // A time value given in ISO8601 combined date and time format that represents when the Droplet was created.
	Status       string  // A status string indicating the state of the Droplet instance. This may be "new", "active", or "archive"..
	Networks     Network // The details of the network that are configured for the Droplet instance. This is an object that contains keys for IPv4 and IPv6. The value of each of these is an array that contains objects describing an individual IP resource allocated to the Droplet. These will define attributes like the IP address, netmask, and gateway of the specific network depending on the type of network it is.
	Kernel       Kernel  // The current kernel. This will initially be set to the kernel of the base image when the Droplet is created.
	Features     []string
	Backup_ids   []int64 // An array of backup IDs of any backups that have been taken of the Droplet instance. Droplet backups are enabled at the time of the instance creation.
	Snapshot_ids []int64 // An array of snapshot IDs of any snapshots created from the Droplet instance.
	Action_ids   []int64 // An array of action IDs of any actions that have been taken on this Droplet.
}

// Create a new Droplet
func (this *Droplet) Create(name string, region string, size string, image string) (interface{}, error) {
	var createDropletResponse *CreateDropletResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		DIGITALOCEAN_API_URL_PREFIX+"/droplets",
		fmt.Sprintf(`{"name":"%v","region":"%v","size":"%v","image":"%v"}`, name, region, size, image))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &createDropletResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return createDropletResponse, nil

}

// List all droplets
func (this *Droplet) List() (*ListDropletResponse, error) {
	log.Println("Listing all Droplets!!")
	var listDropletResponse *ListDropletResponse
	jsonBody, err := GetJsonResponse(
		"GET",
		DIGITALOCEAN_API_URL_PREFIX+"/droplets",
		"")
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &listDropletResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return listDropletResponse, nil
}

// Delete a droplet
func (this *Droplet) Delete(dropletId string) error {
	log.Println("Delete Droplet!!: " + dropletId)
	_, err := GetJsonResponse(
		"DELETE",
		DIGITALOCEAN_API_URL_PREFIX+"/droplets/"+dropletId,
		"")
	if err != nil {
		return err

	}

	return nil
}

// Get droplet Events
func (this *Droplet) Actions(dropletId string) (interface{}, error) {
	log.Println("Droplet Actions for: !!: " + dropletId)
	var dropletActionsResponse *DropletActionsResponse
	jsonBody, err := GetJsonResponse(
		"GET",
		fmt.Sprintf("%v/droplets/%v/%v", DIGITALOCEAN_API_URL_PREFIX, dropletId, "actions"),
		"")
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &dropletActionsResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return dropletActionsResponse, nil
}

// Reboot droplet
func (this *Droplet) Reboot(dropletId string) (interface{}, error) {
	log.Println("Rebooting Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_REBOOT)

}

// PowerCycle droplet
func (this *Droplet) PowerCycle(dropletId string) (interface{}, error) {
	log.Println("Power Cycle Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_POWER_CYCLE)

}

// Shutdown droplet
func (this *Droplet) Shutdown(dropletId string) (interface{}, error) {
	log.Println("Graceful Shutdown Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_SHUTDOWN)

}

// Power off droplet
func (this *Droplet) PowerOff(dropletId string) (interface{}, error) {
	log.Println("Hard Power Off Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_POWER_OFF)

}

// Power On droplet
func (this *Droplet) PowerOn(dropletId string) (interface{}, error) {
	log.Println("Power On Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_POWER_ON)

}

// Password Reset droplet
func (this *Droplet) PasswordReset(dropletId string) (interface{}, error) {
	log.Println("Password Reset for Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_PASSWORD_RESET)

}

// Resize droplet
func (this *Droplet) Resize(dropletId, size string) (interface{}, error) {
	log.Println("Resize Droplet: !!: " + dropletId)
	options := fmt.Sprintf(`{"type":"%v","size":"%v"}`, DROPLETACTION_RESIZE, size)
	return this.basicDropletActionWithOptions(dropletId, options)

}

// Restore droplet
func (this *Droplet) Restore(dropletId, backupImage string) (interface{}, error) {
	log.Println("Restore Droplet: !!: " + dropletId)
	options := fmt.Sprintf(`{"type":"%v","image":"%v"}`, DROPLETACTION_RESTORE, backupImage)
	return this.basicDropletActionWithOptions(dropletId, options)

}

// Rebuild droplet
func (this *Droplet) Rebuild(dropletId, backupImage string) (interface{}, error) {
	log.Println("Rebuild Droplet: !!: " + dropletId)
	options := fmt.Sprintf(`{"type":"%v","image":"%v"}`, DROPLETACTION_REBUILD, backupImage)
	return this.basicDropletActionWithOptions(dropletId, options)

}

// Rename droplet
func (this *Droplet) Rename(dropletId, dropletName string) (interface{}, error) {
	log.Println("Rename Droplet: !!: " + dropletId)
	options := fmt.Sprintf(`{"type":"%v","name":"%v"}`, DROPLETACTION_RENAME, dropletName)
	return this.basicDropletActionWithOptions(dropletId, options)

}

// Disable Backups for droplet
func (this *Droplet) DisableBackups(dropletId string) (interface{}, error) {
	log.Println("Disable Backups for Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_DISABLE_BACKUP)

}

// Enable IP v6 droplet
func (this *Droplet) EnableIpv6(dropletId string) (interface{}, error) {
	log.Println("Enable IPv6 for Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_ENABLE_IPV6)

}

// Enable Private Networking droplet
func (this *Droplet) EnablePrivateNetworking(dropletId string) (interface{}, error) {
	log.Println("Enable Private Networking for Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_ENABLE_PRIVATE_NETWORKING)

}

// Snapshot droplet
func (this *Droplet) Snapshot(dropletId string) (interface{}, error) {
	log.Println("Snapshot for Droplet: !!: " + dropletId)
	return this.basicDropletAction(dropletId, DROPLETACTION_SNAPSHOT)

}

// Basic Action type with just type parameter to URI Endpoint

func (this *Droplet) basicDropletAction(dropletId, actionType string) (interface{}, error) {
	var dropletActionResponse *DropletActionResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/droplets/%v/%v", DIGITALOCEAN_API_URL_PREFIX, dropletId, "actions"),
		fmt.Sprintf(`{"type":"%v"}`, actionType))
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &dropletActionResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return dropletActionResponse, nil
}

// Basic Action type with just type parameter to URI Endpoint With Options

func (this *Droplet) basicDropletActionWithOptions(dropletId, options string) (interface{}, error) {
	var dropletActionResponse *DropletActionResponse
	jsonBody, err := GetJsonResponse(
		"POST",
		fmt.Sprintf("%v/droplets/%v/%v", DIGITALOCEAN_API_URL_PREFIX, dropletId, "actions"),
		options)
	if err != nil {
		return nil, err

	}
	if err := json.Unmarshal(jsonBody, &dropletActionResponse); err != nil {
		log.Println(err)
		return nil, err
	}
	return dropletActionResponse, nil
}
