package v2

import ()

// Digital Ocean keys Structure
type Key struct {
	id          int64  // This is a unique identification number for the key. This can be used to reference a specific SSH key when you wish to embed a key into a Droplet.
	fingerprint string // This is the human-readable display name for the given SSH key. This is used to easily identify the SSH keys when they are displayed.
	publicKey   string // This attribute contains the fingerprint value that is generated from the public key. This is a unique identifier that will differentiate it from other keys using a format that SSH recognizes.
	name        string // This attribute contains the entire public key string that was uploaded. This is what is embedded into the root user's authorized_keys file if you choose to include this SSH key during Droplet creation.
}
