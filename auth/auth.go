package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/arasdenizhan/go-store/constants"
	"github.com/golang-jwt/jwt/v5"
)

func CheckCookie(r *http.Request) error {
	cookie, err := r.Cookie(constants.COOKIE_NAME)
	if err != nil {
		return errors.New("User not authenticated!")
	}

	token, err := jwtValidate(cookie.Value)
	if err != nil {
		return errors.New("Cookie is not valid!")
	}

	issuedAt, err := token.Claims.GetIssuedAt()
	if err != nil {
		return errors.New(err.Error())
	}

	duration, err := time.ParseDuration(constants.TOKEN_VALIDITY_MS)
	if err != nil {
		return errors.New(err.Error())
	}
	expiration := issuedAt.Add(duration)

	if expiration.Before(time.Now()) {
		return errors.New("Cookie is expired!")
	}
	return nil
}

func GetUserId(r *http.Request) (string, error){
	cookie, err := r.Cookie(constants.COOKIE_NAME)
	if(err != nil){
		return "nil", errors.New("User not authenticated!")
	}
	
	token, err := jwtValidate(cookie.Value)
	if(err != nil){
		return "nil", errors.New("Cookie is not valid!")
	}

	claims := token.Claims.(jwt.MapClaims)
	userIdString := fmt.Sprintf("%v", claims["sub"].(float64))
	if(userIdString == ""){
		return "nil", errors.New("UserId could not be fetched from token!")
	}
	return userIdString, nil
}

func jwtValidate(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
	return []byte(constants.SECRET_KEY), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	return token, nil
}