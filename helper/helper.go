package helper

import (
	"IM/define"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im")

// GenerateToken
// 生成 token
func GenerateToken(identity, email string) (string, error) {
	UserClaim := &UserClaims{
		Identity:         identity,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// SendCode 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <1639601246@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.Send("smtp.qq.com:587",
		smtp.PlainAuth("", "1639601246@qq.com", define.MailPassword, "smtp.qq.com"))

}

// GetCode 生成验证码
func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// GetUUID 生成唯一码
func GetUUID() string {
	u := uuid.New()
	return fmt.Sprintf("%x", u)
}
