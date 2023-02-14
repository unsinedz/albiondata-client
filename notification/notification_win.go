//go:build windows
// +build windows

package notification

import (
	"github.com/broderickhyman/albiondata-client/log"
	toast "gopkg.in/toast.v1"
)

func Push(msg string) {
	note := toast.Notification{AppID: "Albion Private Prices Client", Title: "Albion Private Prices Client", Message: msg}

	err := note.Push()
	if err != nil {
		log.Error(err)
	}
}
