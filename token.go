package r6

import (
	"encoding/base64"
	"fmt"
)

// createTokenByEmail создает uplay token по email/password
func createTokenByEmail(email, password string) string {

	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", email, password)))
}
