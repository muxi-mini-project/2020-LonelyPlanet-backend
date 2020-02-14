package middleware

import (
	"fmt"
	"github.com/2020-LonelyPlanet-backend/miniProject/handler"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
<<<<<<< HEAD
	"time"
)

type jwtClaims struct {
	jwt.StandardClaims
	Uid string	`json:"uid"`
}

var (
	key        = "miniProject" //salt
	ExpireTime = 3600          //token expire time
=======
)

var (
	key = "miniProject" //salt
	//ExpireTime = 3600          //token expire time
>>>>>>> 423d511db086ad041c2b524f496681aff9d450ec
)

func JwtAAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("token")
		if tokenStr == "" {
<<<<<<< HEAD
			c.String(401, "token invalid")
			c.Abort()
=======
			handler.ErrTokenInvalid(c, error2.TokenInvalid)
>>>>>>> 423d511db086ad041c2b524f496681aff9d450ec
			//跳转登录界面
			return
		}
		token, err := verifyToken(tokenStr)
		if token == nil || err != nil {
<<<<<<< HEAD
			c.String(401, "token invalid")
			c.Abort()
=======
			handler.ErrTokenInvalid(c, error2.TokenInvalid)
>>>>>>> 423d511db086ad041c2b524f496681aff9d450ec
			//跳转登录页面
			return
		}
		if !token.Valid {
<<<<<<< HEAD
			c.String(401, "token invalid")
			c.Abort()
=======
			handler.ErrTokenInvalid(c, error2.TokenInvalid)
>>>>>>> 423d511db086ad041c2b524f496681aff9d450ec
			//跳转登录页面
			return
		}
		claim := token.Claims
		c.Set("uid", claim.(jwt.MapClaims)["uid"])
		c.Next()
	}
}

/*
func ProduceToken(uid string) string {
	//id, _ := strconv.Atoi(uid)
	claims := &jwtClaims{
		Uid: uid,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	//fmt.Println(singedToken, err)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
*/

<<<<<<< HEAD

=======
>>>>>>> 423d511db086ad041c2b524f496681aff9d450ec
func verifyToken(verifyToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(verifyToken, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(key), nil
	})
	if err != nil {
		log.Print("verifyToken err:")
		fmt.Println(err)
		return nil, err
	}
	return token, nil
}

/*
func verifyToken(varifyToken string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(varifyToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//log.Println(ok)
			log.Panicln("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
*/

/*
func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims, err := verifyToken(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := genToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken, ", ", claims.ExpiresAt)
}
*/
