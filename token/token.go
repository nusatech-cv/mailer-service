package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var JWT_PUBLIC_KEY *rsa.PublicKey

func init() {
	pubKeyPEM, _ := base64.StdEncoding.DecodeString(os.Getenv("JWT_PUBLIC_KEY"))
	block, _ := pem.Decode(pubKeyPEM)
	if block == nil {
		log.Fatal("Failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("Failed to parse public key: " + err.Error())
	}

	var ok bool
	JWT_PUBLIC_KEY, ok = pub.(*rsa.PublicKey)
	if !ok {
		log.Fatal("Public key of unsupported type")
	}
}

type Record struct {
	Email           string  `json:"email"`
	GoogleID        string  `json:"google_id"`
	LoginTime       string  `json:"login_time,omitempty"`
	Device          string  `json:"device,omitempty"`
	IPAddress       string  `json:"ip_address,omitempty"`
	Location        string  `json:"location"`
	PaymentMethod   string  `json:"payment_method,omitempty"`
	PaymentStatus   string  `json:"payment_status,omitempty"`
	AmountPaid      float64 `json:"amount_paid,omitempty"`
	ToAccount       string  `json:"to_account,omitempty"`
	SenderAccount   string  `json:"sender_account,omitempty"`
	PaymentExpired  string  `json:"payment_expired,omitempty"`
	PaymentAt       string  `json:"payment_at,omitempty"`
}

func DecodeToken(tokenString string) (*Record, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_PUBLIC_KEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		recordBytes, err := json.Marshal(claims["record"])
		if err != nil {
			return nil, err
		}
		var record Record
		err = json.Unmarshal(recordBytes, &record)
		if err != nil {
			return nil, err
		}
		return &record, nil
	} else {
		return nil, err
	}
}
