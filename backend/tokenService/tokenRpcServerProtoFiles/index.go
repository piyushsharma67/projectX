package tokenRpcServerProtoFiles

import (
	context "context"
	"fmt"
	"tokenService/utils"
)

type TokenServices struct{}

// mustEmbedUnimplementedTokenServiceServer implements TokenServiceServer.
func (s *TokenServices) mustEmbedUnimplementedTokenServiceServer() {
	panic("unimplemented")
}

// mustEmbedUnimplementedTokenValidationServer implements TokenValidationServer.
func (s *TokenServices) mustEmbedUnimplementedTokenValidationServer() {
	panic("unimplemented")
}

func (s *TokenServices) GenerateToken(ctx context.Context, req *GenerateTokenRequest) (*TokenResponse, error) {
	email := req.EmailId

	fmt.Println("email id is", email)

	jwt, err := utils.CreateLoginJWTToken(email)

	if err != nil {
		return &TokenResponse{
			Response: &TokenResponse_Error{
				Error: &ErrorResponse{
					StatusCode:   401,
					ErrorMessage: err.Error(),
				},
			},
		}, nil
	}
	return &TokenResponse{
		Response: &TokenResponse_Success{
			Success: &SuccessResponse{
				StatusCode: 200,
				Message:    "Token validated successfully",
				Data:       jwt,
				// You can add additional data if needed
			},
		},
	}, nil

}

func (s *TokenServices) ValidateToken(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	token := req.Token

	isValidToken, err, value := utils.VerifyToken(token)

	if isValidToken {
		// Token is valid, construct and return a success response
		return &TokenResponse{
			Response: &TokenResponse_Success{
				Success: &SuccessResponse{
					StatusCode: 200,
					Message:    "Token validated successfully",
					Data:       value,
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

func (s *TokenServices) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return &PingResponse{Message: "Server is running"}, nil
}
