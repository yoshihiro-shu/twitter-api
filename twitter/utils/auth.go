package utils

import (
	"net/http"
)

func setBearerOauth(r *http.Request, bearerToken string) {
	r.Header.Set("Authorization", bearerToken)
}
