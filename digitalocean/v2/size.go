package v2

import ()

// Digital Ocean sizes Structure
type Size struct {
	Slug         string   // Can be 512mb Human Readable name
	Memory       int64    // The amount of RAM available to Droplets created with this size. This value is given in megabytes.
	Vcpus        int64    // The number of virtual CPUs that are allocated to Droplets with this size.
	Disk         int64    // This is the amount of disk space set aside for Droplets created with this size. The value is given in gigabytes.
	Transfer     int64    // The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.
	PriceMonthly string   // This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars.
	PriceHourly  string   // This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.
	Regions      []string // TODO:Need to create All available CONST // An array that contains the region slugs where this size is available for Droplet creates. "nyc1", "br1", "sfo1", "ams4

}
