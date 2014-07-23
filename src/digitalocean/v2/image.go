package v2

import ()

// Digital Ocean images Structure

type Image struct {
	Id           int64    // A unique integer that can be used to identify and reference a specific image.
	Name         string   // The display name that has been given to an image. This is what is shown in the control panel and is generally a descriptive title for the image in question.
	Distribution string   // This attribute describes the base distribution used for this image.
	Slug         string   // A uniquely identifying string that is associated with each of the DigitalOcean-provided public images. These can be used to reference a public image as an alternative to the numeric id.
	Public       bool     // This is a boolean value that indicates whether the image in question is public or not. An image that is public is available to all accounts. A non-public image is only accessible from your account.
	Regions      []string // This attribute is an array of the regions that the image is available in. The regions are represented by their identifying slug values.
	Created_at   string   // A time value given in ISO8601 combined date and time format that represents when the image was created.
}
