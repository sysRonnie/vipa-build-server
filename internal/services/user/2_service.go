package user

import (
	"context"
	"go-tailwind-test/internal/config"
	"go-tailwind-test/internal/services/auth"
	"log"

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

	session, err :=
	s.store.ValidateRefreshToken(
		ctx,
		refreshToken,
		)

	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	email, isAdmin, err :=
	s.store.GetUserAuthByUUID(
		ctx,
		session.UserID,
	)
	if err != nil {
		return nil, ErrDatabaseFailure
	}

	accessToken, err :=
		auth.GenerateAccessToken(
			email,
			isAdmin,
			session.SessionID,
		)

	if err != nil {
		return nil, ErrLoginFailed
	}

	newRefreshToken := auth.GenerateRefreshToken()

	refreshTokenHash :=
		auth.HashRefreshToken(
			newRefreshToken,
		)

	refreshExpiry :=
		auth.GenerateRefreshTokenExpiry()

	err = s.store.RotateAuthSession(
		ctx,
		session.SessionID,
		refreshTokenHash,
		refreshExpiry,
		userAgent,
		ip,
	)

	if err != nil {
		log.Println(
			"failed to insert auth session",
			err,
		)
		return nil, ErrDatabaseFailure
	}

	return &auth.AuthResponse{
		AccessToken: accessToken,
		RefreshToken: newRefreshToken,
		RefreshTokenExpiry: refreshExpiry,
		IsAdmin: isAdmin,
		Email: email,
	}, nil
}

func (s *Service) Login(params LoginServiceParams) (*auth.AuthResponse, error) {
	log.Println("validating id token")
	cfg := config.Envs

	log.Println("os variable for google client id", cfg.GoogleClientIDWeb)


	payload, err := idtoken.Validate(
		params.ctx,
		params.GoogleIDToken,
		cfg.GoogleClientIDWeb,
	)

	if err != nil {
		return nil, ErrLoginFailed
	}

	emailValue, ok := payload.Claims["email"]

	if !ok {
		return nil, ErrLoginFailed
	}

	email, ok := emailValue.(string)

	if !ok {
		return nil, ErrLoginFailed
	}

	isAllowed, isAdmin, err := s.store.IsApprovedEmail(params.ctx,email)

	if err != nil {
		return nil, ErrDatabaseFailure
	}

	if !isAllowed {
		return nil, ErrUserNotAuthorized
	}

	uuid, err := s.store.GetOrCreateUser(params.ctx, email, payload.Subject)
	
	if err != nil {
		return nil, ErrDatabaseFailure
	}

	log.Println("new user created", uuid)

	refreshToken :=
	auth.GenerateRefreshToken()

refreshTokenHash :=
	auth.HashRefreshToken(
		refreshToken,
	)

	refreshExpiry := auth.GenerateRefreshTokenExpiry()

	sessionID, err := s.store.CreateAuthSession(
			params.ctx,
			uuid,
			refreshTokenHash,
			refreshExpiry,
			params.UserAgent,
			params.IP,
		)

	if err != nil {

		log.Println(
			"failed to create auth session",
			err,
		)

		return nil, ErrDatabaseFailure
	}

	accessToken, err :=
		auth.GenerateAccessToken(
			email,
			isAdmin,
			sessionID,
		)

	if err != nil {
		return nil, ErrLoginFailed
	}

	return &auth.AuthResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		RefreshTokenExpiry: refreshExpiry,
		IsAdmin: isAdmin,
		Email:   email,
	}, nil
}