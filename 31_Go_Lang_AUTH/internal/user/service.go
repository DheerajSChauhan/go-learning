package user

import (
	"context"
	"time"
	"errors"
	"go-auth/internal/auth"
	"strings"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repo
	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{
		repo: repo,
		jwtSecret: jwtSecret,
	}
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResult struct{
	Token string `json:"token"`
	User PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	// Check if user already exists
	email := strings.ToLower(strings.TrimSpace(input.Email))
	pass := strings.TrimSpace(strings.TrimSpace(input.Password))

	if email == "" || pass == "" {
		return AuthResult{}, errors.New("email and password are required")
	}

	if len(pass) < 6 {
		return AuthResult{}, errors.New("password must be at least 6 characters")
	}
	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, errors.New("user already exists")
	}
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return AuthResult{}, errors.New("error occurred while checking user")
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, errors.New("error occurred while hashing password")
	}
	
	now := time.Now().UTC()

	u := User{
		Email: email,
		PasswordHash: string(hashBytes),
		Role: "user",
		CreatedAt: now,
		UpdatedAt: now,
	}

	created, err := s.repo.Create(ctx, u)
	if err != nil {
		return AuthResult{}, errors.New("error occurred while creating user")
	}
	token, err := auth.CreateToken(s.jwtSecret, created.ID.Hex(), created.Role)
	if err != nil {
		return AuthResult{}, errors.New("error occurred while creating token")
	}
	return AuthResult{
		Token: token,
		User: ToPublicUser(created),
	}, nil
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	// Normalize email:
	// - Remove leading/trailing spaces
	// - Convert to lowercase so TEST@gmail.com and test@gmail.com are treated the same
	email := strings.ToLower(strings.TrimSpace(input.Email))

	// Remove leading/trailing spaces from password
	pass := strings.TrimSpace(input.Password)

	// Validate required fields
	if email == "" || pass == "" {
		return AuthResult{}, errors.New("email and password are required")
	}

	// Look up the user in the database by email
	user, err := s.repo.FindByEmail(ctx, email)

	// If user is not found (or another DB error occurs),
	// don't reveal whether the email exists for security reasons
	if err != nil {
		return AuthResult{}, errors.New("invalid email or password")
	}

	// Compare the stored bcrypt hash with the password entered by the user
	// CompareHashAndPassword returns nil only when they match
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(pass),
	); err != nil {
		return AuthResult{}, errors.New("invalid email or password")
	}

	// Generate a JWT token containing:
	// - User ID (stored in the "sub" claim)
	// - User role
	// - Expiration time
	token, err := auth.CreateToken(
		s.jwtSecret,
		user.ID.Hex(),
		user.Role,
	)
	if err != nil {
		return AuthResult{}, errors.New("error occurred while creating token")
	}

	// Return the JWT token and public user information
	// Never return PasswordHash to the client
	return AuthResult{
		Token: token,
		User:  ToPublicUser(user),
	}, nil
}