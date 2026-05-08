package user

import (
	"context"
	"database/sql"
	"errors"
	"go-tailwind-test/internal/services/auth"

	"github.com/google/uuid"
)




type Store struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

type UserStore interface {
	IsApprovedEmail(ctx context.Context, email string) (bool, bool, error)
	GetOrCreateUser(ctx context.Context, email string, googleSub string) (uuid.UUID, error)
	ValidateRefreshToken(ctx context.Context, refreshToken string) (*auth.AuthSession, error)
	GetUserAuthByUUID(ctx context.Context, userID uuid.UUID) (string, bool, error)
	RotateAuthSession(ctx context.Context, sessionID uuid.UUID, refreshTokenHash string, expiresAt int64, userAgent string, ipAddress string) error
	CreateAuthSession(ctx context.Context, userID uuid.UUID, refreshTokenHash string, expiresAt int64, userAgent string, ipAddress string) (uuid.UUID, error)
	RevokeAuthSession(ctx context.Context, sessionID uuid.UUID) error
}

func (s *Store) RevokeAuthSession(ctx context.Context, sessionID uuid.UUID) error {
	_, err := s.db.ExecContext(
		ctx,
		`
		UPDATE user_auth_session
		SET revoked_at = NOW(), revoked_reason = 'user_logout'
		WHERE session_id = $1
		AND revoked_at IS NULL
		`,
		sessionID,
	)

	return err
}
func (s *Store) CreateAuthSession(
	ctx context.Context,
	userID uuid.UUID,
	refreshTokenHash string,
	expiresAt int64,
	userAgent string,
	ipAddress string,
) (uuid.UUID, error) {

	var sessionID uuid.UUID

	err := s.db.QueryRowContext(
		ctx,
		`
		INSERT INTO user_auth_session (
			user_id,
			refresh_token_hash,
			user_agent,
			ip_address,
			expires_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			to_timestamp($5)
		)
		RETURNING session_id
		`,
		userID,
		refreshTokenHash,
		userAgent,
		ipAddress,
		expiresAt,
	).Scan(&sessionID)

	if err != nil {
		return uuid.Nil, err
	}

	return sessionID, nil
}

func (s *Store) RotateAuthSession(
	ctx context.Context,
	sessionID uuid.UUID,
	refreshTokenHash string,
	expiresAt int64,
	userAgent string,
	ipAddress string,
) error {

	// find old auth sessions and revoke? 

	_, err := s.db.ExecContext(
		ctx,
		`
		UPDATE user_auth_session
		SET
			refresh_token_hash = $1,
			expires_at = to_timestamp($2),
			user_agent = $3,
			ip_address = $4,
			last_used_at = NOW(),
			updated_at = NOW()
		WHERE session_id = $5
		AND revoked_at IS NULL
		`,
		refreshTokenHash,
		expiresAt,
		userAgent,
		ipAddress,
		sessionID,
	)

	return err
}

func (s *Store) GetUserAuthByUUID(ctx context.Context, userID uuid.UUID) (string, bool, error) {
	var email string
	var isAdmin bool
	
	err := s.db.QueryRowContext(
		ctx,
		`
		SELECT A.email, COALESCE(B.is_admin, FALSE)
		FROM user_auth A
		LEFT JOIN user_auth_approved_email B ON lower(A.email) = lower(B.email)
		WHERE A.uuid = $1
		`,
		userID,
	).Scan(&email, &isAdmin)
	
	if err != nil {
		return "", false, err
	}

	return email, isAdmin, nil
}

func (s *Store) ValidateRefreshToken(
	ctx context.Context,
	refreshToken string,
) (*auth.AuthSession, error) {

	refreshTokenHash :=
		auth.HashRefreshToken(
			refreshToken,
		)

	var session auth.AuthSession

	err := s.db.QueryRowContext(
		ctx,
		`
		SELECT
			session_id,
			user_id
		FROM user_auth_session
		WHERE refresh_token_hash = $1
		AND expires_at > NOW()
		AND revoked_at IS NULL
		`,
		refreshTokenHash,
	).Scan(
		&session.SessionID,
		&session.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}




func (s *Store) GetOrCreateUser(
	ctx context.Context,
	email string,
	googleSub string,
) (uuid.UUID, error) {

	var userID uuid.UUID

	err := s.db.QueryRowContext(
		ctx,
		`
		INSERT INTO user_auth (
			email,
			google_sub
		)
		VALUES ($1, $2)

		ON CONFLICT (email)
		DO UPDATE SET
			google_sub = EXCLUDED.google_sub

		RETURNING uuid
		`,
		email,
		googleSub,
	).Scan(&userID)

	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func (s *Store) IsApprovedEmail(ctx context.Context, email string) (bool, bool, error) {
	var isAdmin bool

	err := s.db.QueryRowContext(
		ctx,
		`
		SELECT is_admin
		FROM user_auth_approved_email
		WHERE lower(email) = lower($1)
		`,
		email,
	).Scan(&isAdmin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, false, nil
		}

		return false, false, err
	}

	return true, isAdmin, nil
}