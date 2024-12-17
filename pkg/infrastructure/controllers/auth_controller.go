package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/auth"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/dto"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type AuthController struct {
	cryptoService services.CryptoService
	jwtService    services.JWTService
	userService   services.UserService
	loggerService services.LoggerService
}

func NewAuthController(
	cryptoService services.CryptoService,
	jwtService services.JWTService,
	userService services.UserService,
	loggerService services.LoggerService,
) *AuthController {
	return &AuthController{
		cryptoService: cryptoService,
		jwtService:    jwtService,
		userService:   userService,
		loggerService: loggerService,
	}
}

// GetAccessTokenHandler godoc
//
//	@Summary		Get access token (login)
//	@Description	Get token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string						false	"Language"	default(en-US)
//	@Param			User			body		dto.LoginDTO				true	"User data to be created"
//	@Success		200				{object}	serializers.TokenSerializer	"desc"
//	@Failure		401				{object}	response.Response			"desc"
//	@Failure		500				{object}	response.Response			"desc"
//	@Router			/api/v1/auth/login [post]
func (c *AuthController) GetAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	var loginDto dto.LoginDTO

	// Decodificar el cuerpo de la solicitud
	if err := json.NewDecoder(r.Body).Decode(&loginDto); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Obtener el token de acceso usando el caso de uso
	token, err := auth.GetAccessTokenUserUseCase(
		loginDto.Email,
		loginDto.Password,
		c.cryptoService,
		c.jwtService,
		c.userService,
		c.loggerService,
	)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	payload := serializers.NewTokenSerializer(token)
	response.SuccessResponse(w, http.StatusOK, payload)
}
