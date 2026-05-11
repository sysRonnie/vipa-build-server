package user

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
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

func (h *Handler) HandleTest(c echo.Context) error {

	return network.Success(c, network.SandboxResponse{
		StatusCode: 200,
		Message: "This is a test endpoint for the user service",
	})
}

func (h *Handler) HandleLogoutRequest(c echo.Context,) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("handling_logout_request")
	claims := auth.GetClaimsFromContext(c)

	err := h.store.RevokeAuthSession(c.Request().Context(),claims.SessionID)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.Success(c,network.SandboxResponse{
			StatusCode: 200,
			Message: "Successfully logged out",
		},
	)
}

func (h *Handler) HandleRefreshTokenRequest(c echo.Context,) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("handling_refresh_token_request")
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
		return network.FailFromError(c, err)
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
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("handling_login_request")
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		advisor.Error("failed_to_bind_login_request", err)
		return network.Fail(c, network.ErrInvalidPayload)
	}

	ip, ua := network.ExtractIPandAgent(c)
	ctx := c.Request().Context()
	
	advisor.Log("extracted_ip_and_user_agent_beginning_login_service")

	res, err := h.service.Login(LoginServiceParams{
		ctx: ctx,
		GoogleIDToken: req.Token,
		IP: ip,
		UserAgent: ua,
	})
	
	if err != nil {
		advisor.Error("login_service_failed", err)
		return network.FailFromError(c, err)
	}

	advisor.Log("--- issuing a new jwt to client -----")

	return network.Success(c, network.SandboxResponse{
		StatusCode: 200,
		Message: "Successful Login!",
		Data: res,
	})
}