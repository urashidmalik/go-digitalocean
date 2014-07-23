// Digital Ocean API version 2 implementation
package v2

import ()

const (
	DROPLETACTION_REBOOT                    = "reboot"
	DROPLETACTION_POWER_CYCLE               = "power_cycle"
	DROPLETACTION_SHUTDOWN                  = "shutdown"
	DROPLETACTION_POWER_OFF                 = "power_off"
	DROPLETACTION_POWER_ON                  = "power_on"
	DROPLETACTION_PASSWORD_RESET            = "password_reset"
	DROPLETACTION_RESIZE                    = "resize"
	DROPLETACTION_RESTORE                   = "restore"
	DROPLETACTION_REBUILD                   = "rebuild"
	DROPLETACTION_RENAME                    = "rename"
	DROPLETACTION_ENABLE_IPV6               = "enable_ipv6"
	DROPLETACTION_DISABLE_BACKUP            = "disable_backups"
	DROPLETACTION_ENABLE_PRIVATE_NETWORKING = "enable_private_networking"
	DROPLETACTION_SNAPSHOT                  = "snapshot"
)

// Digital ocean actions are events.

type Action struct {
	Id           int64  // A unique numeric ID that can be used to identify and reference an action.
	Status       string // The current status of the action. This can be "in-progress", "completed", or "errored".
	Type         string // This is the type of action that the object represents. For example, this could be "transfer" to represent the state of an image transfer action.
	StartedAt    string // A time value given in ISO8601 combined date and time format that represents when the action was initiated.
	CompletedAt  string // A time value given in ISO8601 combined date and time format that represents when the action was completed.
	ResourceId   int64  // A unique identifier for the resource that the action is associated with.
	ResourceType string // The type of resource that the action is associated with.
	Region       string // A slug representing the region where the action occurred.
}

type ActionsResponse struct {
	Actions []Action
}
type DropletActionResponse struct {
	Action Action
}
type DropletActionsResponse ActionsResponse
