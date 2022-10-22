package utils

import (
	"fmt"
	"net/http"
)

func setBearerOauth(r *http.Request, bearerToken string) {
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
}
