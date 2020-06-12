package auth

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

const SessionName = "auth-session"

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(User{})
	return nil
}
