package middlewares

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/common"
)

func loadRSAPublicKey() (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile("/home/ec2-user/QRmark/keys/public_key.pem")

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the public key")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA public key: %v", err)
	}

	publicKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return publicKey, nil
}

func (m *QrmarkAPIMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("token")
		fmt.Println(cookie)
		fmt.Println(cookie.Value)
		if err != nil {
			err = apierrors.RequiredAuthorization.Wrap(err, "invalid cookie")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		authorization := cookie.Value

		idToken := authorization

		if idToken == "" {
			err := apierrors.RequiredAuthorization.Wrap(errors.New("invalid req header"), "invalid header")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		pubKey, err := loadRSAPublicKey()
		if err != nil {
			apierrors.ErrorHandler(w, req, err)
			return
		}

		token, err := jwt.Parse(
			idToken,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					err = apierrors.Unauthorizated.Wrap(err, "invalid token")
					return nil, err
				}
				return pubKey, nil
			},
		)
		if err != nil {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		var exp float64
		var id float64
		var ok bool

		if exp, ok = claims["exp"].(float64); !ok {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		if id, ok = claims["sub"].(float64); !ok {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		if time.Now().After(time.Unix(int64(exp), 0)) {
			err = apierrors.Unauthorizated.Wrap(err, "token expired")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		ctx := context.WithValue(req.Context(), common.UserKey{}, int(id))

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func (m *QrmarkAPIMiddleware) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("token")
		if err != nil {
			err = apierrors.RequiredAuthorization.Wrap(err, "invalid cookie")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		authorization := cookie.Value

		idToken := authorization

		if idToken == "" {
			err := apierrors.RequiredAuthorization.Wrap(errors.New("invalid req header"), "invalid header")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		pubKey, err := loadRSAPublicKey()
		if err != nil {
			apierrors.ErrorHandler(w, req, err)
			return
		}

		token, err := jwt.Parse(
			idToken,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					err = apierrors.Unauthorizated.Wrap(err, "invalid token")
					return nil, err
				}
				return pubKey, nil
			},
		)

		if err != nil {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		var exp float64
		var ok bool

		if exp, ok = claims["exp"].(float64); !ok {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		if time.Now().After(time.Unix(int64(exp), 0)) {
			err = apierrors.Unauthorizated.Wrap(err, "token expired")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		var role string

		if role, ok = claims["role"].(string); !ok {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		if role != "admin" {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		var id float64
		if id, ok = claims["sub"].(float64); !ok {
			err = apierrors.Unauthorizated.Wrap(err, "invalid token")
			apierrors.ErrorHandler(w, req, err)
			return
		}

		ctx := context.WithValue(req.Context(), common.UserKey{}, int(id))

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
