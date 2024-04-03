package token

import "time"

type Maker interface {
	CreateToken(username string, role string, durtion time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
