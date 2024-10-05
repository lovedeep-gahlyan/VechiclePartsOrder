package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "log"
)


// Define a secret key (in a real application, keep this secret and secure)
var jwtSecret = []byte("your_secret_key")

// GenerateJWT generates a new JWT token
func GenerateJWT(username string, role string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 72).Unix(), // Token valid for 72 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateJWT validates the provided JWT token and returns the parsed token
func ValidateJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorMalformed)
        }
        return jwtSecret, nil
    })

    return token, err
}


// HashPassword hashes the user's password using bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatal(err)
        return "", err
    }
    return string(bytes), nil
}

// CheckPasswordHash compares the hashed password with the plain password
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CheckPassword(plainPassword, storedPassword string) bool {
    return plainPassword == storedPassword
}

