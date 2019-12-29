package app

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/elcompadre/webapi/utils"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func (next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		notAuth := []string{"/api/user/new","/api/user/login"}
		requestPath := request.URL.Path

		for _, value := range notAuth{
			if value == requestPath {
				next.ServeHTTP(responseWriter, request)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := request.Header.Get("Authorization")

		//Check if contains token
		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			responseWriter.WriteHeader(http.StatusForbidden)
			utils.Respond(responseWriter, response)
			return
		}

		//Check format of token
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			responseWriter.WriteHeader(http.StatusForbidden)
			utils.Respond(responseWriter, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = utils.Message(false, "Malformed authentication token")
			responseWriter.WriteHeader(http.StatusForbidden)
			utils.Respond(responseWriter, response)
			return
		}

		fmt.Sprintf("User %s", tk.Username)
		ctx := context.WithValue(request.Context(), "user", tk.UserId)
		request = request.WithContext(ctx)
		next.ServeHTTP(responseWriter, request)
	})
}