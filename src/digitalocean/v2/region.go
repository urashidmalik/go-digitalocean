package v2

import ()

// Constants Representing all supported regions
const (
	REGION_NEWYORK_1      = "nyc1"
	REGION_AMSTERDAM_1    = "ams1"
	REGION_SANFRANSISCO_1 = "sfo1"
	REGION_NEWYORK_2      = "nyc2"
	REGION_AMSTERDAM_2    = "ams2"
	REGION_LONDON_1       = "lon1"
)

// Regions Supported by Digital Ocean

type Region struct {
	Slug      string   // Can be nyc1 Human Readable name
	Name      string   // The display name of the region. This will be a full name that is used in the control panel and other interfaces. EG. New York
	Sizes     []string // TODO:Need to create All available CONST // Slice that can contain multiple values like "1024mb", "512mb"
	Features  []string // TODO:Need to create All available CONST // Slice that can contain multiple values like "ipv6","virtio", "private_networking"
	Available bool     // This is a boolean value that represents whether new Droplets can be created in this region.
}
