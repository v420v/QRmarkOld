package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/common"
	"github.com/v420v/qrmarkapi/controllers/services"
	"github.com/v420v/qrmarkapi/models"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	service services.UserServicer
}

func NewUserController(service services.UserServicer) *UserController {
	return &UserController{
		service: service,
	}
}

func loadRSAPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile("/home/ec2-user/QRmark/keys/private_key.pem")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the private key")
	}

	parseResult, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
	}

	privateKey, ok := parseResult.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("error!!!")
	}

	return privateKey, nil
}

func (c *UserController) LoginHandler(w http.ResponseWriter, req *http.Request) {
	type LoginInfo struct {
		Email    string
		Password string
	}

	var loginInfo LoginInfo

	if err := json.NewDecoder(req.Body).Decode(&loginInfo); err != nil {
		err = apierrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apierrors.ErrorHandler(w, req, err)
		return
	}

	user, err := c.service.SelectUserByEmailService(loginInfo.Email)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	if !user.Verified {
		apierrors.ErrorHandler(w, req, errors.New("user not verified"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"iss":  "QRmark",
			"sub":  user.ID,
			"role": user.Role,
			"exp":  jwt.NewNumericDate(time.Now().Add(time.Hour)),
		})

	privateKey, err := loadRSAPrivateKey()
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	signedToken, err := token.SignedString(privateKey)

	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    signedToken,
		Expires:  time.Now().Add(60 * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode(signedToken)

	w.WriteHeader(http.StatusOK)
}

func (c *UserController) CurrentUserHandler(w http.ResponseWriter, req *http.Request) {
	userID, err := common.GetCurrentUserID(req.Context())
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	user, err := c.service.SelectUserByIDService(userID)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (c *UserController) VerifyHandler(w http.ResponseWriter, req *http.Request) {
	token := mux.Vars(req)["token"]

	err := c.service.VerifyUserService(token)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode("user verified")
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func sendEmail(to, subject, body string) error {
	const senderEmail = "ibuki420v@gmail.com"
	const SMTPServer = "smtp.gmail.com"
	const SMTPPort = "587"

	header := make(map[string]string)
	header["From"] = senderEmail
	header["To"] = to
	header["Subject"] = subject
	header["MIME-version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"UTF-8\""

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	_ = godotenv.Load()
	auth := smtp.PlainAuth("", senderEmail, os.Getenv("GMAIL_APP_PASSWORD"), SMTPServer)

	return smtp.SendMail(
		SMTPServer+":"+SMTPPort,
		auth,
		senderEmail,
		[]string{to},
		[]byte(message),
	)
}

func (c *UserController) InsertUserHandler(w http.ResponseWriter, req *http.Request) {
	var reqUser models.User

	if err := json.NewDecoder(req.Body).Decode(&reqUser); err != nil {
		err = apierrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apierrors.ErrorHandler(w, req, err)
		return
	}

	hashedPassword, err := passwordEncrypt(reqUser.Password)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	reqUser.Password = hashedPassword
	reqUser.Email = strings.ToLower(reqUser.Email)
	reqUser.Role = "user"
	reqUser.Verified = false

	user, err := c.service.InsertUserService(reqUser)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	verification_token := models.VerificationToken{
		UserID:    user.ID,
		Token:     GenerateSecureToken(16),
		ExpiredAt: time.Now().Add(time.Minute * 5),
	}

	err = c.service.InsertVerificationTokenService(verification_token)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	_ = godotenv.Load()

	to := reqUser.Email
	subject := "QRmark アカウント有効化"
	body := fmt.Sprintf(`
		<html>
        <body style="font-family: Arial, sans-serif;">
            <div style="padding: 20px; text-align: center;">
                <h1 style="font-weight: 800;">QRmark</h1>
                <a href="https://ibukiqrmark.com/verify/%s">アカウントを有効化する</a>
            </div>
        </body>
        </html>
    `, verification_token.Token)

	err = sendEmail(to, subject, body)

	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) SelectUserDetailHandler(w http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	user, err := c.service.SelectUserByIDService(userID)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) SelectUserListHandler(w http.ResponseWriter, req *http.Request) {
	page := 0
	queryMap := req.URL.Query()

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apierrors.BadParam.Wrap(err, "page in query param must be number")
			apierrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	userList, err := c.service.SelectUserListService(page)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(userList)
}
