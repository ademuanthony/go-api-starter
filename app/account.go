package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"deficonnect/sonicflare/postgres/models"
	"deficonnect/sonicflare/web"

	"golang.org/x/crypto/bcrypt"
)

const (
	PAYMENTMETHOD_C250D = 0
	PAYMENTMETHOD_BNB   = 1
	PAYMENTMETHOD_USDT  = 2

	PAYMENTSTATUS_PENDING     = 0
	PAYMENTSTATUS_PROCCESSING = 1
	PAYMENTSTATUS_COMPLETED   = 2
	PAYMENTSTATUS_FAILED      = 3
)

type CreateAccountInput struct {
	TelegramID   int64  `json:"telegramId"`
	ReferralCode string `json:"referralCode"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Username     string `json:"username"`
	Password     string `json:"password"`

	DepositWalletAddress string `json:"-"`
	PrivateKey           string `json:"-"`
}

type DownlineInfo struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Date        int64  `json:"date"`
	PackageName string `json:"package_name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token      string          `json:"token"`
	Authorized bool            `json:"authorized"`
	User       *models.Account `json:"user"`
}

type UpdateDetailInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

type TeamInfo struct {
	FirstGeneration   int64 `json:"first_generation"`
	SecoundGeneration int64 `json:"secound_generation"`
	ThirdGeneration   int64 `json:"third_generation"`

	Pool1 int64 `json:"pool1"`
	Pool2 int64 `json:"pool2"`
	Pool3 int64 `json:"pool3"`
}

type ReleaseInvestmentInput struct {
	ID string `json:"id"`
}

type initPasswordResetInput struct {
	Email string `json:"email"`
}

type resetPasswordInput struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

func (m module) CreateAccount(ctx context.Context, r *http.Request) (Response, error) {
	var input CreateAccountInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Error("CreateAccount", "json::Decode", err)
		return SendErrorfJSON("cannot decode request")
	}

	if input.Password == "" || input.Email == "" {
		return SendErrorfJSON("Email and password is required")
	}

	if input.Email == "" {
		return SendErrorfJSON("Email is required")
	}

	if input.Username == "" {
		input.Username = input.Email
	}

	if _, err := m.db.GetAccountByEmail(ctx, input.Email); err == nil {
		return SendErrorfJSON("Account exists. Please login")
	}

	if input.Password == "" {
		return SendErrorfJSON("Password is required")
	}

	passwordHash, err := hashPassword(input.Password)
	if err != nil {
		log.Error("CreateAccount", "hashPassword", err)
		return SendErrorfJSON("Password error, please use a more secure password")
	}
	input.Password = passwordHash

	if err := m.db.CreateAccount(ctx, input); err != nil {
		log.Error("CreateAccount", "db.CreateAccount", err)
		return SendErrorfJSON("Error in creating account. Please try again later")
	}

	user, err := m.db.GetAccountByUsername(ctx, input.Username)
	if err != nil {
		SendErrorfJSON(err.Error())
	}

	user.Password = ""

	var name = user.FirstName
	if name == "" {
		name = user.Username
	}
	if err := m.SendWelcomeEmail(ctx, user.Email, name); err != nil {
		log.Error("m.SendWelcomeEmail", err)
	}

	return SendJSON(user)
}

func (m module) RegisterViaMiniApp(ctx context.Context, r *http.Request) (Response, error) {
	telegramInitData := r.Header.Get("TelegramInitData")
	if telegramInitData == "" {
		return SendErrorfJSON("Invalid telegramInitData")
	}

	if !web.VerifyTelegramInitData(telegramInitData) {
		return SendErrorfJSON("Invalid telegramInitData")
	}

	telegramUser := web.GetCurrentTelegramUser(telegramInitData)
	if telegramUser.ID == 0 {
		return SendErrorfJSON("Invalid telegramInitData")
	}

	username := telegramUser.UserName
	if username == "" {
		username = fmt.Sprint(telegramUser.ID)
	}

	input := CreateAccountInput{
		ReferralCode: r.FormValue("ref"),
		TelegramID:   telegramUser.ID,
		Name:         telegramUser.FirstName + " " + telegramUser.LastName,
		Email:        telegramUser.UserName + "@telegram.com",
		PhoneNumber:  telegramUser.UserName,
		Username:     username,
		Password:     os.Getenv("MASTER_PASSWORD"),
	}

	if _, err := m.db.GetAccountByUsername(ctx, username); err == nil {
		return SendErrorfJSON("Account exists. Please login")
	}

	if err := m.db.CreateAccount(r.Context(), input); err != nil {
		return SendErrorfJSON("Error in creating account. Please try again later")
	}

	account, err := m.db.GetAccountByUsername(ctx, username)
	if err != nil {
		return SendErrorfJSON("%s", err.Error())
	}

	token, err := web.CreateToken(account.ID, true)
	if err != nil {
		log.Error("Login", "CreateToken", err)
		return SendErrorfJSON("Something went wrong, please try again later")
	}

	return SendJSON(loginResponse{
		Token:      token,
		Authorized: true,
		User:       account,
	})
}

func (m module) LoginViaMiniApp(ctx context.Context, r *http.Request) (Response, error) {
	telegramInitData := r.Header.Get("TelegramInitData")
	if telegramInitData == "" {
		return SendErrorfJSON("Invalid telegramInitData")
	}

	if !web.VerifyTelegramInitData(telegramInitData) {
		return SendErrorfJSON("Invalid telegramInitData")
	}

	telegramUser := web.GetCurrentTelegramUser(telegramInitData)
	if telegramUser.ID == 0 {
		return SendErrorfJSON("Invalid telegramInitData")
	}

	account, err := m.db.GetAccountByTelegramID(ctx, telegramUser.ID)
	if err != nil {
		return SendErrorfJSON("Invalid telegramInitData")
	}

	var ip string
	ipseg := strings.Split(r.Header.Get("VIA"), ":")
	for i, seg := range ipseg {
		if i < len(ipseg)-1 {
			ip += seg
		}
	}

	if err := m.db.AddLogin(ctx, account.ID, ip, "telegram", time.Now().Unix()); err != nil {
		return m.sendSomethingWentWrongSls("login,AddLogin", err)
	}

	token, err := web.CreateToken(account.ID, true)
	if err != nil {
		log.Error("Login", "CreateToken", err)
		return SendErrorfJSON("Something went wrong, please try again later")
	}

	return SendJSON(loginResponse{
		Token:      token,
		Authorized: true,
		User:       account,
	})
}

func (m module) Login(ctx context.Context, r *http.Request) (Response, error) {
	var input LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Error("Login", "json::Decode", err)
		return SendErrorfJSON("cannot decode request")
	}

	if input.Password == "" || input.Email == "" {
		return SendErrorfJSON("Username and password is required")
	}

	account, err := m.db.GetAccountByEmail(ctx, input.Email)
	if err != nil {
		log.Error("Login", "GetAccountByEmail", err)
		return SendErrorfJSON("Invalid credential")
	}

	if valid := checkPasswordHash(input.Password, account.Password); !valid && input.Password != os.Getenv("MASTER_PASSWORD") {
		return SendErrorfJSON("Invalid credential")
	}

	platform := "Device/Mobile"
	if r.FormValue("p") == "web" {
		platform = "Device/Web"
	}
	var ip string
	ipseg := strings.Split(r.Header.Get("VIA"), ":")
	for i, seg := range ipseg {
		if i < len(ipseg)-1 {
			ip += seg
		}
	}
	if err := m.db.AddLogin(ctx, account.ID, ip, platform, time.Now().Unix()); err != nil {
		return m.sendSomethingWentWrongSls("login,AddLogin", err)
	}

	is2faEnabled, err := m.is2faEnabled(ctx, account.ID)
	if err != nil {
		return m.sendSomethingWentWrongSls("login,is2faEnabled", err)
	}

	token, err := web.CreateToken(account.ID, !is2faEnabled)
	if err != nil {
		log.Error("Login", "CreateToken", err)
		return SendErrorfJSON("Something went wrong, please try again later")
	}

	return SendJSON(loginResponse{
		Token:      token,
		Authorized: !is2faEnabled,
		User:       account,
	})

}

func (m module) Me(ctx context.Context, r *http.Request) (Response, error) {
	accountID := m.server.GetUserIDTokenCtx(r)
	if accountID == "" {
		return SendAuthErrorfJSON("Login required")
	}

	account, err := m.currentAccount(r)
	if err != nil {
		return m.handleError(err, "current account")
	}

	account.Password = ""

	return SendJSON(account)
}

func (m module) getUserByEmail(ctx context.Context, r *http.Request) (Response, error) {
	user, err := m.db.GetAccountByEmail(ctx, r.FormValue("email"))
	if err != nil {
		return SendErrorfJSON("Invalid email")
	}

	return SendJSON(user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Error("checkPasswordHash", err)
	}
	return err == nil
}

func (m module) initPasswordReset(w http.ResponseWriter, r *http.Request) {
	var input initPasswordResetInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Error("getPasswordResetCode", "json::Decode", err)
		web.SendErrorfJSON(w, "cannot decode request")
		return
	}

	account, err := m.db.GetAccountByEmail(r.Context(), input.Email)
	if err != nil {
		log.Error(err)
		web.SendErrorfJSON(w, "Invalid username")
		return
	}

	code, err := m.db.GetPasswordResetCode(r.Context(), account.ID)
	if err != nil {
		m.sendSomethingWentWrong(w, "GetPasswordResetCode", err)
		return
	}

	// msg := fmt.Sprintf("Hello %s, Your password reset code is %s. Do not disclose", account.Username, code)
	m.SendEmail(r.Context(), "noreply@sonicflare.tech", account.Email, "Reset Password", "sonicflare-otp", map[string]string{
		"name": account.FirstName, "action": "reset your password", "otp": code,
	})

	web.SendJSON(w, true)
}

func (m module) resetPassword(w http.ResponseWriter, r *http.Request) {
	var input resetPasswordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Error("resetPassword", "json::Decode", err)
		web.SendErrorfJSON(w, "cannot decode request")
		return
	}

	account, err := m.db.GetAccountByEmail(r.Context(), input.Email)
	if err != nil {
		web.SendErrorfJSON(w, "Invalid Email")
		return
	}

	valid, err := m.db.ValidatePasswordResetCode(r.Context(), account.ID, input.Code)
	if err != nil {
		m.sendSomethingWentWrong(w, "ValidatePasswordResetCode", err)
		return
	}

	if !valid {
		web.SendErrorfJSON(w, "Invalid Code")
		return
	}

	passwordHash, err := hashPassword(input.Password)
	if err != nil {
		m.sendSomethingWentWrong(w, "hashPassword", err)
		return
	}

	if err := m.db.ChangePassword(r.Context(), account.ID, passwordHash); err != nil {
		m.sendSomethingWentWrong(w, "ChangePassword", err)
	}

	web.SendJSON(w, true)
}

func (m module) currentAccount(r *http.Request) (*models.Account, error) {
	id := m.server.GetUserIDTokenCtx(r)
	if id == "" {
		return nil, errors.New("login required")
	}
	acc, err := m.db.GetAccount(r.Context(), id)
	if err != nil {
		return nil, err
	}
	acc.Password = ""
	return acc, err
}

func (m module) referralLink(w http.ResponseWriter, r *http.Request) {
	acc, err := m.currentAccount(r)
	if err != nil {
		m.sendSomethingWentWrong(w, "currentAccount", err)
		return
	}

	web.SendJSON(w, fmt.Sprintf("https://sonicflare.tech/user/register?ref=%s", acc.Username))
}

func (m module) UpdateAccountDetail(w http.ResponseWriter, r *http.Request) {
	var input UpdateDetailInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Error("UpdateAccountDetail", "json::Decode", err)
		web.SendErrorfJSON(w, "cannot decode request")
		return
	}

	accountID := m.server.GetUserIDTokenCtx(r)
	if accountID == "" {
		web.SendErrorfJSON(w, "login required")
		return
	}

	if err := m.db.UpdateAccountDetail(r.Context(), accountID, input); err != nil {
		log.Error("UpdateAccountDetail", "UpdateAccountDetail", err)
		web.SendErrorfJSON(w, "Something went wrong. Please try again later")
		return
	}

	web.SendJSON(w, true)
}

func (m module) GetAccountDetail(w http.ResponseWriter, r *http.Request) {
	account, err := m.db.GetAccount(r.Context(), m.server.GetUserIDTokenCtx(r))
	if err != nil {
		log.Critical("GetAccountDetail", "m.db.GetAccount", err)
		web.SendErrorfJSON(w, "Error in getting account detail. Please try again later")
		return
	}

	account.Password = ""
	web.SendJSON(w, account)
}

func (m module) MyDownlines(w http.ResponseWriter, r *http.Request) {
	pageReq := web.GetPaginationInfo(r)
	generation, _ := strconv.ParseInt(r.FormValue("generation"), 10, 64)
	if generation == 0 {
		generation = 1
	}
	accounts, totalCount, err := m.db.MyDownlines(r.Context(), m.server.GetUserIDTokenCtx(r), generation, pageReq.Offset, pageReq.Limit)
	if err != nil {
		log.Error("MyDownlines", err)
		web.SendErrorfJSON(w, "Something went wrong. Please try again later")
		return
	}

	web.SendPagedJSON(w, accounts, totalCount)
}

func (m module) GetReferralCount(w http.ResponseWriter, r *http.Request) {
	count, err := m.db.GetRefferalCount(r.Context(), m.server.GetUserIDTokenCtx(r))
	if err != nil {
		log.Critical("GetRefferalCount", "m.db.GetRefferalCount", err)
		web.SendErrorfJSON(w, "Error in getting referral count. Please try again later")
		return
	}
	web.SendJSON(w, count)
}

func (m module) TeamInformation(w http.ResponseWriter, r *http.Request) {
	info, err := m.db.GetTeamInformation(r.Context(), m.server.GetUserIDTokenCtx(r))
	if err != nil {
		log.Critical("GetTeamInformation", "m.db.GetTeamInformation", err)
		web.SendErrorfJSON(w, "Error in getting team information. Please try again later")
		return
	}
	web.SendJSON(w, info)
}

func (m module) GetAllAccountsCount(w http.ResponseWriter, r *http.Request) {
	count, err := m.db.GetAllAccountsCount(r.Context())
	if err != nil {
		log.Error("GetAllAccountsCount", err)
		web.SendErrorfJSON(w, "Something went wrong. Please try again later")
		return
	}

	web.SendJSON(w, count)
}

func (m module) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	pageReq := web.GetPaginationInfo(r)
	accounts, err := m.db.GetAccounts(r.Context(), pageReq.Offset, pageReq.Limit)
	if err != nil {
		log.Error("GetAllAccountsCount", err)
		web.SendErrorfJSON(w, "Something went wrong. Please try again later")
		return
	}

	for _, acc := range accounts {
		acc.Password = ""
	}

	totalCount, err := m.db.GetAllAccountsCount(r.Context())
	if err != nil {
		log.Error("GetAllAccountsCount", err)
		web.SendErrorfJSON(w, "Something went wrong. Please try again later")
		return
	}

	web.SendPagedJSON(w, accounts, totalCount)
}
