package auth

import (
    "net/http"
    "errors"
    "strings"
)

func GetAPIKey(headers http.Header) (string, error) {
    val := headers.Get("Authorization")
    if val == "" {
        return "", errors.New("no authentication token found")
    }

    vals := strings.Split(val, " ")
    if len(vals) != 2 || vals[0] != "Bearer" {
        return "", errors.New("malformed authentication header")
    }

    return vals[1], nil    
}
