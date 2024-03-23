package authenticationServiceRpcServer

import (
	"authenticationService/utils"
	"context"
)

type ValidationServerGrpcStruct struct{}

// mustEmbedUnimplementedTokenValidationServer implements TokenValidationServer.
func (s *ValidationServerGrpcStruct) mustEmbedUnimplementedTokenValidationServer() {
	panic("unimplemented")
}

func (s *ValidationServerGrpcStruct) ValidateToken(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	// Your token validation logic goes here
	token := req.Token
    
	isValidToken, err,value := utils.VerifyToken(token)
    
	if isValidToken {
		// Token is valid, construct and return a success response
		return &TokenResponse{
			Response: &TokenResponse_Success{
				Success: &SuccessResponse{
					StatusCode: 200,
					Message:    "Token validated successfully",
                    Data:value ,
					// You can add additional data if needed
				},
			},
		}, nil
	} else {
		// Token is not valid, construct and return an error response
		return &TokenResponse{
			Response: &TokenResponse_Error{
				Error: &ErrorResponse{
					StatusCode:   401,
					ErrorMessage: err,
				},
			},
		}, nil
	}
}
