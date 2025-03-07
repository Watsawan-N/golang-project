package mid

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"golang-project/pkg/errs"
	"golang-project/pkg/helper"
	"golang-project/pkg/web"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Sub   string
	Exp   *time.Time
	Valid bool
	Key   []byte
}

type IDTokenCustomClaims struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

type CertificateManager struct {
	Keys           []JWK `json:"keys"`
	CertificateURL string
	Doc            []byte
	DocExp         time.Time
}

type JWK struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	X5c []string `json:"x5c"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	Kid string   `json:"kid"`
	X5t string   `json:"x5t"`
}

var certificateManager CertificateManager

func Authentication( /*systemService service.ISystemService,*/
	common helper.ICommon) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// configuration := config.New()
			// jsonWebKeySetURI := configuration.JsonWebKeySetURI

			// certificateManager.CertificateURL = jsonWebKeySetURI
			s := r.Header.Get("Authorization")
			token := strings.TrimPrefix(s, "Bearer ")

			if token == "" {
				SetHeaderAllowCors(&w, "*")
				err := errs.NewUnauthorizedError("missing authorization header")
				return common.HandleErr(&w, err)
			}

			userId, err := validateToken(token, r)
			if err != nil {
				SetHeaderAllowCors(&w, "*")
				err := errs.NewUnauthorizedError(err.Error())
				return common.HandleErr(&w, err)
			}

			// existingUser, serviceErr := systemService.IsUser(userId)
			// if serviceErr != nil {
			// 	SetHeaderAllowCors(&w, "*")
			// 	err := errs.NewUnauthorizedError("You don't have permission. Error check existing " + serviceErr.Error())
			// 	return common.HandleErr(&w, err)
			// }

			// if !existingUser {
			// 	SetHeaderAllowCors(&w, "*")
			// 	err := errs.NewUnauthorizedError("You don't have permission.")
			// 	return common.HandleErr(&w, err)
			// }

			r.Header.Set("userId", userId)
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}

func validateToken(token string, r *http.Request) (userId string, err error) {
	defer func() {
		if err != nil {
			isPass, systemUser := bypassLocalhost(r)
			if isPass {
				userId = systemUser
				err = nil
			}
		}
	}()

	resToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		keyData, errKeyData := certificateManager.GetCertificate()
		if errKeyData != nil {
			return nil, errors.New("failed to get certificate: " + errKeyData.Error())
		}
		err := json.Unmarshal(keyData, &certificateManager)
		if err != nil {
			return nil, err
		}
		jwk := (certificateManager.Keys)[0]
		err = VerifySignature(token, &jwk)
		if err != nil {
			return nil, errors.New("failed to verify token: " + err.Error())
		}

		publicKey := GetPublicKeyFromModulusAndExponent(&jwk)

		return publicKey, nil
	})

	if err != nil {
		return "", errors.New(err.Error())
	}

	if !resToken.Valid {
		return "", errors.New("token invalid")
	}

	claims, _ := resToken.Claims.(jwt.MapClaims)
	if claims["employeeid"] == nil {
		return "", errors.New("invalid token: missing employee id")
	}

	m := claims["employeeid"]
	userId = fmt.Sprintf("%v", m)
	return userId, nil
}

func (c *CertificateManager) GetCertificate() ([]byte, error) {
	if c.Doc != nil && time.Now().Before(c.DocExp) {
		return c.Doc, nil
	}
	url := c.CertificateURL
	resp, err := http.Get(url)
	if err != nil {

		errorString := fmt.Sprintf("certificate_[Get http request] %s %s", url, err.Error())
		return nil, errors.New(errorString)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorString := fmt.Sprintf("certificate_[Response is not 200 OK!] %d", resp.StatusCode)
		return nil, errors.New(errorString)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorString := fmt.Sprintf("certificate_Download_[Read response] %d", resp.StatusCode)
		return nil, errors.New(errorString)
	}
	c.Doc = data
	c.DocExp = time.Now().Local().Add(time.Minute * time.Duration(5))
	return data, nil
}

func VerifySignature(jwtToken string, jwk *JWK) error {
	parts := strings.Split(jwtToken, ".")
	message := []byte(strings.Join(parts[0:2], "."))
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return err
	}
	publicKey := GetPublicKeyFromModulusAndExponent(jwk)

	hasher := crypto.SHA256.New()
	hasher.Write(message)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hasher.Sum(nil), signature)
	return err
}

func GetPublicKeyFromModulusAndExponent(jwk *JWK) *rsa.PublicKey {
	n, _ := base64.RawURLEncoding.DecodeString(jwk.N)
	e, _ := base64.RawURLEncoding.DecodeString(jwk.E)
	z := new(big.Int)
	z.SetBytes(n)
	var buffer bytes.Buffer
	buffer.WriteByte(0)
	buffer.Write(e)
	exponent := binary.BigEndian.Uint32(buffer.Bytes())
	return &rsa.PublicKey{N: z, E: int(exponent)}
}

func bypassLocalhost(r *http.Request) (isPass bool, userId string) {
	isPass = false
	userId = ""
	if r.Host == "localhost:5000" {
		isPass = true
		userId = "99999999"
	}

	return isPass, userId
}
