package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func main() {
	// 签发
	mySingingKey := []byte("thisakeys")
	c := MyClaims{
		UserName: "mark",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "mark",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := token.SignedString(mySingingKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	// 解密
	tk, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySingingKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tk.Claims.(*MyClaims).UserName)

}
