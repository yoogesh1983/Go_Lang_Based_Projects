package data

import (
	"math/rand"
	"time"

	"github.com/hashicorp/go-hclog"
)

type JwtToken struct {
	Log   hclog.Logger
	Token string
}

func NewToken(l hclog.Logger) (*JwtToken, error) {
	t := &JwtToken{Log: l, Token: ""}
	err := t.GetToken()
	return t, err
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes
func (e *JwtToken) MonitorJWTStatus(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates

				// is this a postive or negative change
				direction := rand.Intn(1)
				newTokenValue := ""

				if direction == 0 {
					newTokenValue = "7777777777777"
				} else {
					newTokenValue = "99999999999999"
				}

				// modify the rate
				e.Token = newTokenValue

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()

	return ret
}

func (e *JwtToken) GetToken() error {
	e.Token = "152207"
	return nil
}
