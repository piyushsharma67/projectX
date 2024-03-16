package jwt_v1

import (
	"authenticationService/server"
	"authenticationService/utils"
	"errors"
	"net/http"
	"strings"
)

var(
	FORBIDDEN=errors.New("forbidden!!")
)

func ValidateJwtController(server *server.Server) http.HandlerFunc{
	return func (w http.ResponseWriter,req *http.Request){
		tokenFromHeader:=req.Header.Get("Authorization")
	
		if(tokenFromHeader==""){
			utils.CreateErrorResponse(FORBIDDEN,http.StatusForbidden,w)
			return
		}
	
		tokenString:=strings.Split(tokenFromHeader," ")
	
		if(tokenString[1]==""){
			utils.CreateErrorResponse(FORBIDDEN,http.StatusForbidden,w)
			return
		}
	
		isValid,userID:=utils.VerifyToken(tokenString[1])
	
		if(!isValid){
			utils.CreateErrorResponse(FORBIDDEN,http.StatusForbidden,w)
			return
		}
	
		utils.CreateSuccessResponse("Successful",userID,http.StatusOK,w)
		
	}
}