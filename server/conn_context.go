package server

import "C"
import "github.com/gofrs/uuid/v5"

type ConnectID interface {
	ID() uuid.UUID
}

type ConnContext struct {
	ConnectID
}
