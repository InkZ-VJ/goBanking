package token

import "time"

type Maker interface {
	CreateToken(username string, durtion time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
