package v2

import ()

// Digital Ocean image action Structure

type ImageAction struct {
	slug      string   // Can be nyc1 Human Readable name
	name      string   // The display name of the region. This will be a full name that is used in the control panel and other interfaces. EG. New York
	sizes     []string // TODO:Need to create All available CONST // Slice that can contain multiple values like "1024mb", "512mb"
	features  []string // TODO:Need to create All available CONST // Slice that can contain multiple values like "ipv6","virtio", "private_networking"
	available bool     // This is a boolean value that represents whether new Droplets can be created in this region.
}
