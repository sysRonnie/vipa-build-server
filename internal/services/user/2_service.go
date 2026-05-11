package user

import (
	"context"
	"go-tailwind-test/internal/config"
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"google.golang.org/api/idtoken"
)

type Service struct {
	store UserStore
}

func NewService(
	store UserStore,
) *Service {

	return &Service{
		store: store,
	}
}


type RefreshAccessTokenParams struct {
	Context context.Context
	AccessToken string
	RefreshToken string
	UserAgent string
	IP string
}


type UserService interface {
	Login(LoginServiceParams) (*auth.AuthResponse, error)
	RefreshAccessToken(ctx context.Context, refreshToken string, userAgent string, ip string) (*auth.AuthResponse, error)
}

func (s *Service) RefreshAccessToken(ctx context.Context, refreshToken string, userAgent string, ip string) (*auth.AuthResponse, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("validating refresh token")
	session, err := s.store.ValidateRefreshToken(ctx,refreshToken)

	if err != nil {
		advisor.Error("failed_to_validate_refresh_token", err)
		return nil, ErrInvalidRefreshToken
	}

	email, isAdmin, err := s.store.GetUserAuthByUUID(ctx,session.UserID)
	if err != nil {
		advisor.Error("failed_to_get_user_auth_by_uuid", err)
		return nil, ErrDatabaseFailure
	}

	accessToken, err := auth.GenerateAccessToken(email,isAdmin,session.SessionID)

	if err != nil {
		advisor.Error("failed_to_generate_access_token", err)
		return nil, ErrLoginFailed
	}

	newRefreshToken := auth.GenerateRefreshToken()

	refreshTokenHash := auth.HashRefreshToken(newRefreshToken)

	refreshExpiry := auth.GenerateRefreshTokenExpiry()

	err = s.store.RotateAuthSession(
		ctx,
		session.SessionID,
		refreshTokenHash,
		refreshExpiry,
		userAgent,
		ip,
	)

	if err != nil {
		advisor.Error("failed_to_rotate_auth_session", err)
		return nil, ErrDatabaseFailure
	}

	advisor.Log("refresh token validated, returning auth response")
	return &auth.AuthResponse{
		AccessToken: accessToken,
		RefreshToken: newRefreshToken,
		RefreshTokenExpiry: refreshExpiry,
		IsAdmin: isAdmin,
		Email: email,
	}, nil
}

func (s *Service) Login(params LoginServiceParams) (*auth.AuthResponse, error) {
	cfg := config.Envs
	advisor := advisor.FromContext(params.ctx)

	advisor.Log("validating google id token")
	payload, err := idtoken.Validate(
		params.ctx,
		params.GoogleIDToken,
		cfg.GoogleClientIDWeb,
	)

	if err != nil {
		advisor.Error("failed_to_validate_google_id_token", err)
		return nil, ErrLoginFailed
	}

	emailValue, ok := payload.Claims["email"]

	if !ok {
		advisor.Error("email_claim_not_found_in_google_id_token", nil)
		return nil, ErrLoginFailed
	}

	email, ok := emailValue.(string)

	if !ok {
		advisor.Error("email_claim_in_google_id_token_is_not_a_string", nil)
		return nil, ErrLoginFailed
	}

	advisor.Log("google id token validated, checking if email is in allowed users list")
	isAllowed, isAdmin, err := s.store.IsApprovedEmail(params.ctx,email)

	if err != nil {
		advisor.Error("failed to check if email is in allowed users list", err)
		return nil, ErrDatabaseFailure
	}

	if !isAllowed {
		advisor.Log("email is not in allowed users list, rejecting login")
		return nil, ErrUserNotAuthorized
	}

	advisor.Log("email is allowed, checking if user exists or needs to be created")
	uuid, err := s.store.GetOrCreateUser(params.ctx, email, payload.Subject)
	if err != nil {
		advisor.Error("failed to get or create user", err)
		return nil, ErrDatabaseFailure
	}


	advisor.Log("new user created or existing user found, creating auth session and tokens")
	refreshToken := auth.GenerateRefreshToken()
	refreshTokenHash := auth.HashRefreshToken(refreshToken)
	refreshExpiry := auth.GenerateRefreshTokenExpiry()

	advisor.Log("auth session and tokens created, inserting auth session into database")

	sessionID, err := s.store.CreateAuthSession(
			params.ctx,
			uuid,
			refreshTokenHash,
			refreshExpiry,
			params.UserAgent,
			params.IP,
		)

	if err != nil {
		advisor.Error(
			"failed to create auth session",
			err,
		)

		return nil, ErrDatabaseFailure
	}

	advisor.Log("auth session inserted into database, generating access token")
	accessToken, err :=
		auth.GenerateAccessToken(
			email,
			isAdmin,
			sessionID,
		)

	if err != nil {
		advisor.Error("failed to generate access token",err)
		return nil, ErrLoginFailed
	}

	advisor.Log("access token generated, returning auth response")

	return &auth.AuthResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		RefreshTokenExpiry: refreshExpiry,
		IsAdmin: isAdmin,
		Email:   email,
	}, nil
}