package login

import (
	"authenticationService/models"
	"authenticationService/server"
	login_services "authenticationService/server/v1/service/login_service"
	"authenticationService/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type UserLoginRequestBody struct{
	Email       string `json:"email"`
	Password    string `json:"password"`
}

var (
	EMAIL_CANNOT_BE_EMPTY=errors.New("email cannot be empty")
	PASSWORD_CANNOT_BE_EMPTY=errors.New("password cannot be empty")
	INVALID_CREDENTIALS=errors.New("invalid credentials")
)

func validateLoginRequestBody (requestBody *UserLoginRequestBody)error{
	if(requestBody.Email=="" ){
		return EMAIL_CANNOT_BE_EMPTY
	}else if(requestBody.Password=="") {
		return PASSWORD_CANNOT_BE_EMPTY
	}
	return nil
}

// type UserLoginDetails struct{
// 	User *models.User
// 	Token string
// }

func Login(server *server.Server)http.HandlerFunc{
	return func (w http.ResponseWriter,r *http.Request){
		
		var loginRequestBody *UserLoginRequestBody

		json.NewDecoder(r.Body).Decode(&loginRequestBody)
		err:=validateLoginRequestBody(loginRequestBody)

		if(err!=nil){
			utils.CreateErrorResponse(err,http.StatusOK,w)
			return
		}

		user,err:=login_services.FetchUserByEmailInDb(server.Store,loginRequestBody.Email)

		if(err!=nil){
			utils.CreateErrorResponse(err,http.StatusOK,w)
			return
		}else if(user==nil){
			utils.CreateErrorResponse(INVALID_CREDENTIALS,http.StatusUnauthorized,w)
			return
		}

		jwtToken, err := utils.CreateLoginJWTToken(user.Email, r.UserAgent())
		if err != nil {
			utils.CreateErrorResponse(err, http.StatusUnauthorized, w)
			return
		}

		userLoginResponseBody := models.UserLoginDetails{
			User:  user,
			Token: jwtToken,
		}

		utils.CreateSuccessResponse("success",userLoginResponseBody,http.StatusOK,w)
	}
}


// var (
//     oauthConfig = &oauth2.Config{
//         ClientID:     "your-client-id",
//         ClientSecret: "your-client-secret",
//         RedirectURL:  "http://localhost:8080/callback",
//         Endpoint: oauth2.Endpoint{
//             AuthURL:  "provider-auth-url",
//             TokenURL: "provider-token-url",
//         },
//         Scopes: []string{"scope1", "scope2"},
//     }
// )