package auth

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
    Id uuid.UUID `json:"id"`
    Name string `json:"name"`
    jwt.RegisteredClaims
}

func SetToken(w http.ResponseWriter, token string) {
    cookie := http.Cookie{
        Name: "AccessToken",
        Value: token,
        MaxAge: 36000,
        Path: "/",
        HttpOnly: true,
    }

    http.SetCookie(w, &cookie)
}

func RemoveToken(w http.ResponseWriter) {
    cookie := http.Cookie{
        Name: "AccessToken",
        Value: "",
        MaxAge: 0,
        Path: "/",
        HttpOnly: true,
    }

    http.SetCookie(w, &cookie)
}

func GetToken(r *http.Request) (string, error) {
    cookie, err := r.Cookie("AccessToken")
    if err != nil {
        if errors.Is(err, http.ErrNoCookie) {
            return "", errors.New("cookie not found")
        }
        return "", err
    }

    return cookie.Value, nil
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return bytes, err
}

func CheckPasswordWithHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func CreateNewJWT(id uuid.UUID, name string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims {
        Id: id,
        Name: name,
    })

    key := GetHMACKey()
    
    return token.SignedString([]byte(key))
}

func ParseJWT(token string) (*jwt.Token, error) {
    parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(GetHMACKey()), nil
    })

    return parsedToken, err
}

func GetIdFromJWT(token string) (uuid.UUID, error) {
    parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(GetHMACKey()), nil
    })
    if err != nil {
        return uuid.New(), err
    }
        
    id := parsedToken.Claims.(*CustomClaims).Id

    return id, nil
}

func GetHMACKey() string {
    godotenv.Load(".env")

    key := os.Getenv("HMAC_KEY")
    if key == "" {
        log.Fatal("HMAC Secret key is missing")
    }

    return key
}
