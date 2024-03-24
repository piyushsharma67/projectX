package v1

import (
	"authenticationService/models"
	"authenticationService/server"
	authenticate_service "authenticationService/server/services/authenticate/v1"
	"authenticationService/tokenRpcServerProtoFiles"
	tokenServiceConnPackage "authenticationService/tokenServiceConn"
	"authenticationService/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	EMAIL_CANNOT_BE_EMPTY=errors.New("email cannot be empty")
	PASSWORD_CANNOT_BE_EMPTY=errors.New("password cannot be empty")
	INVALID_CREDENTIALS=errors.New("invalid credentials")
)

func ValidateRequestBody(user *models.User)error{
	var err error
	
	if(user.Name==""){
		err=errors.New("name cannot be empty")
	}else if(user.Email==""){
		err=errors.New("email cannot be empty")
	}else if(user.Mobile==""){
		err=errors.New("mobile cannot be empty")
	}else if(user.Password==""){
		err=errors.New("password cannot be empty")
	}

	return err
}

func SignUpController(server *server.Server,tsc *tokenServiceConnPackage.TokenServiceConn) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){

		var err error
		var user *models.User
		var jwt string

		err=json.NewDecoder(r.Body).Decode(&user)
		if(err!=nil){
			utils.CreateErrorResponse(err,http.StatusInternalServerError,w)
			return
		}

		if err=ValidateRequestBody(user); err!=nil{
			utils.CreateErrorResponse(err,http.StatusOK,w)
			return
		}

		if(user.Email!=""){
			if err=authenticate_service.ValidateIsEmailExists(server.Store,user.Email); err!=nil{
				utils.CreateErrorResponse(err,http.StatusOK,w)
				return
			}
		}

		if err=authenticate_service.SaveUserInDb(server.Store,user); err!=nil{
			utils.CreateErrorResponse(err,http.StatusOK,w)
			return
		}

		users,err:=authenticate_service.FetchUserInDb(server.Store,user.Email)

		if(err!=nil){
			utils.CreateErrorResponse(err,http.StatusOK,w)
			return
		}else if(users==nil){
			utils.CreateErrorResponse(INVALID_CREDENTIALS,http.StatusUnauthorized,w)
			return
		}

		request := &tokenRpcServerProtoFiles.GenerateTokenRequest{
			EmailId: user.Email,
		}

		response,err:=tsc.TokenClient.GenerateToken(context.Background(),request)

		switch payload := response.Response.(type) {
		case *tokenRpcServerProtoFiles.TokenResponse_Success:
			jwt=payload.Success.Data
		case *tokenRpcServerProtoFiles.TokenResponse_Error:
			utils.CreateErrorResponse(errors.New(payload.Error.ErrorMessage),http.StatusUnauthorized,w)
			return
		}

		if err != nil {
			utils.CreateErrorResponse(err, http.StatusUnauthorized, w)
			return
		}

		userLoginResponseBody := models.UserLoginDetails{
			User:  user,
			Token: jwt,
		}

		userLoginResponseBody.User.Password=""
		userLoginResponseBody.User.HashedPassword=""

		utils.CreateSuccessResponse("successful",userLoginResponseBody,http.StatusOK,w)
	}
}

type UserLoginRequestBody struct{
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func validateLoginRequestBody (requestBody *UserLoginRequestBody)error{
	if(requestBody.Email=="" ){
		return EMAIL_CANNOT_BE_EMPTY
	}else if(requestBody.Password=="") {
		return PASSWORD_CANNOT_BE_EMPTY
	}
	return nil
}

func LoginController(server *server.Server,tsc *tokenServiceConnPackage.TokenServiceConn)http.HandlerFunc{
	return func (w http.ResponseWriter,r *http.Request){
		
		var loginRequestBody *UserLoginRequestBody
		var jwt string

		json.NewDecoder(r.Body).Decode(&loginRequestBody)
		err:=validateLoginRequestBody(loginRequestBody)

		if(err!=nil){
			
			utils.CreateErrorResponse(err,http.StatusOK,w)
			return
		}

		user,err:=authenticate_service.FetchUserByEmailInDb(server.Store,loginRequestBody.Email)
		
		if(err!=nil){
			utils.CreateErrorResponse(err,http.StatusOK,w)
			return
		}else if(user==nil){
			utils.CreateErrorResponse(INVALID_CREDENTIALS,http.StatusUnauthorized,w)
			return
		}

		request := &tokenRpcServerProtoFiles.GenerateTokenRequest{
			EmailId: user.Email,
		}

		response,err:=tsc.TokenClient.GenerateToken(context.Background(),request)

		if response == nil || response.Response == nil {
			
			fmt.Println("Response is nil")
			return
		}

		switch payload := response.Response.(type) {
		case *tokenRpcServerProtoFiles.TokenResponse_Success:
			
			successResponse := payload.Success
			jwt=successResponse.Data
			
		case *tokenRpcServerProtoFiles.TokenResponse_Error:
			utils.CreateErrorResponse(errors.New(payload.Error.ErrorMessage),http.StatusUnauthorized,w)
			return
		}
		if err != nil {
			utils.CreateErrorResponse(err, http.StatusUnauthorized, w)
			return
		}
		user.Password=""

		userLoginResponseBody := models.UserLoginDetails{
			User:  user,
			Token: jwt,
		}

		utils.CreateSuccessResponse("success",userLoginResponseBody,http.StatusOK,w)
	}
}


func FetchUserByEmailController(server *server.Server) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){

		var err error
		var user *models.User

		err=json.NewDecoder(r.Body).Decode(&user)
		if(err!=nil){
			utils.CreateErrorResponse(err,http.StatusInternalServerError,w)
			return
		}

		if(user.Email==""){
			utils.CreateErrorResponse(EMAIL_CANNOT_BE_EMPTY,http.StatusOK,w)
			return
		}

		result,errors:=authenticate_service.FetchUserByEmailInDb(server.Store,user.Email); 
		
		if errors!=nil{
			utils.CreateErrorResponse(errors,http.StatusOK,w)
			return
		}

		utils.CreateSuccessResponse("successful",result,http.StatusOK,w)
	}
}
