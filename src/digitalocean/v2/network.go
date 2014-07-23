package v2

import ()

// Digital Ocean sizes Structure

type Network struct {
	V4 []V4
	V6 []V6
}
type V4 struct {
	Ip_address string
	Netmask    string
	Gateway    string
	Type       string
}
type V6 struct {
}
