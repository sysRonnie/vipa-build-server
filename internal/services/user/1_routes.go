package user

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/network"
	"log"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service UserService
	store UserStore
}

func NewHandler(service UserService, store UserStore) *Handler {
	return &Handler{
		service: service,
		store: store,
	}
}



func (h *Handler) RegisterUserRoutes(g *echo.Group) {
	g.POST("/login-user", h.HandleLoginRequest)
	g.POST("/handle-refresh-token", h.HandleRefreshTokenRequest)
	g.POST("/handle-logout", h.HandleLogoutRequest, auth.Middleware)
	g.GET("/test-auth", h.HandleTestAuth, auth.Middleware)
}

func (h *Handler) HandleLogoutRequest(c echo.Context,) error {

	claims := c.Get(
		string(auth.ClaimsContextKey),
	).(*auth.Claims)

	log.Println(
		"LOGOUT SESSION ID:",
		claims.SessionID,
	)

	err := h.store.RevokeAuthSession(
		c.Request().Context(),
		claims.SessionID,
	)
	if err != nil {
		log.Println("error revoking auth session:",err,)
		return network.Fail(c,network.SandboxResponse{
				StatusCode: 500,
				Message: "Logout failed",
			},
		)
	}

	return network.Success(c,network.SandboxResponse{
			StatusCode: 200,
			Message: "Successfully logged out",
		},
	)
}

func (h *Handler) HandleRefreshTokenRequest(c echo.Context,) error {

	accessHeader := c.Request().Header.Get("Authorization")
	refreshToken := c.Request().Header.Get("X-Refresh-Token")
	if refreshToken == "" {
		return network.Fail(c,network.SandboxResponse{
				StatusCode: 401,
				Message: "Missing refresh token",
			},
		)
	}

	log.Println("REFRESH ACCESS HEADER:",accessHeader)
	log.Println("REFRESH TOKEN:",refreshToken,)

	ctx := c.Request().Context()
	ip := c.RealIP()
	ua := c.Request().UserAgent()

	res, err := h.service.RefreshAccessToken(ctx, refreshToken, ua, ip)

	if err != nil {
		log.Println("Error refreshing access token:", err)
		return network.Fail(c, network.SandboxResponse{
			StatusCode: 401,
			Message: "Invalid refresh token",
		})
	}

	log.Println("--- refreshed jwt for client -----")
	log.Println("[HandleRefreshTokenRequest] new jwt = ", res.AccessToken)
	log.Println("[HandleRefreshTokenRequest] new refreshToken = ", res.RefreshToken)


	return network.Success(
		c,
		network.SandboxResponse{
			StatusCode: 200,
			Message: "Refresh token endpoint",
			Data: res,
		},
	)
}

func (h *Handler) HandleTestAuth(c echo.Context) error {
	log.Println("hit endpoint")
	claims := c.Get(string(auth.ClaimsContextKey),).(*auth.Claims)
	log.Println(claims.UserEmail)
	log.Println("did it work?")
	log.Println("This expires at:")
	log.Println(claims.ExpiresAt.Time)

	return network.Success(c, network.SandboxResponse{
		StatusCode: 200,
		Message: "Authenticated!",
	})
}

func (h *Handler) HandleLoginRequest(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return network.Fail(c, network.SandboxResponse{
			StatusCode: 400,
			Message: "Invalid Request",
		})
	}

	ip := c.RealIP()
	ua := c.Request().UserAgent()
	ctx := c.Request().Context()

	res, err := h.service.Login(LoginServiceParams{
		ctx: ctx,
		GoogleIDToken: req.Token,
		IP: ip,
		UserAgent: ua,
	})
	
	if err != nil {
		log.Println("login error", err)
		return network.Fail(c, network.SandboxResponse{
			StatusCode: 500,
			Message: "Login failed",
		})
	}

	log.Println("--- issuing a new jwt to client -----")
	log.Println("[HandleLoginRequest] jwt = ", res.AccessToken)
	log.Println("[HandleLoginRequest] refreshToken = ", res.RefreshToken)

	return network.Success(c, network.SandboxResponse{
		StatusCode: 200,
		Message: "Successful Login!",
		Data: res,
	})
}