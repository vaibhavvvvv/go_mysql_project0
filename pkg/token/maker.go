package token

import "time"

// maker is an interface for managing tokens
type Maker interface {
	//create token creates a new username for a specific username  and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//verifyToken checks if payload valid or not
	VerifyToken(token string) (*Payload, error)
}
