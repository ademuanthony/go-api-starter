package web

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID     string `json:"userId"`
	Authorized bool   `json:"authorized"`
	jwt.StandardClaims
}

func CreateToken(userid string, authorized bool) (string, error) {
	expirationTime := time.Now().Add(360 * time.Minute)
	claims := &Claims{
		UserID:     userid,
		Authorized: authorized,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	if bearToken == "" {
		bearToken = r.Header.Get("authorization")
	}
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyTelegramData() gin.HandlerFunc {
	return func(c *gin.Context) {
		telegramInitData := c.GetHeader("TelegramInitData")
		// fmt.Println("TelegramInitData", telegramInitData)
		if telegramInitData == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "telegramInitData header is missing"})
			return
		}

		if !VerifyTelegramInitData(telegramInitData) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid telegramInitData"})
			return
		}

		c.Next()
	}
}

var botToken string

func getBotToken() string {
	if botToken == "" {
		botToken = os.Getenv("TELEGRAM_TOKEN")
	}

	if botToken == "" {
		panic("BOT_TOKEN is not set in environment variables")
	}

	return botToken
}

func VerifyTelegramInitData(telegramInitData string) bool {
	urlParams, _ := url.ParseQuery(telegramInitData)
	hash := urlParams.Get("hash")
	urlParams.Del("hash")

	// Sort the parameters by key
	keys := make([]string, 0, len(urlParams))
	for key := range urlParams {
		// fmt.Println(key, urlParams.Get(key))
		keys = append(keys, key)
	}

	// // date validation
	if os.Getenv("ENV") != "local" {
		dateUnix, _ := strconv.ParseInt(urlParams.Get("auth_date"), 10, 64)
		if time.Since(time.Unix(dateUnix, 0)).Seconds() > 300 {
			return false
		}
	}
	sort.Strings(keys)

	// Build the data check string
	var dataCheckString strings.Builder
	for _, key := range keys {
		dataCheckString.WriteString(key + "=" + urlParams.Get(key) + "\n")
	}
	dataCheckStringStr := dataCheckString.String()
	dataCheckStringStr = dataCheckStringStr[:len(dataCheckStringStr)-1] // Remove the last newline

	// Create HMAC using the BOT_TOKEN as the secret
	secretKey := hmac.New(sha256.New, []byte("WebAppData"))
	secretKey.Write([]byte(getBotToken()))
	secret := secretKey.Sum(nil)

	hmacHash := hmac.New(sha256.New, secret)
	hmacHash.Write([]byte(dataCheckStringStr))
	calculatedHash := hex.EncodeToString(hmacHash.Sum(nil))

	return calculatedHash == hash
}

func GetCurrentTelegramUser(telegramInitData string) (user tgbotapi.User) {
	urlParams, _ := url.ParseQuery(telegramInitData)
	userString := urlParams.Get("user")
	json.Unmarshal([]byte(userString), &user)
	return
}
