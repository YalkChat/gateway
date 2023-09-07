// TODO: Move to "newchat/models/core"?
package events

import "gorm.io/gorm"

type ServerSettings struct {
	gorm.Model
	IsInitialized bool `json:"is_initialized"`
}
