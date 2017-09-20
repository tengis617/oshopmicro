package Profile

import "context"
import "github.com/jinzhu/gorm"

// Service is a CRUD interface to handle Profiles
type Service interface {
	GetProfile(ctx context.Context, id string) (Profile, error)
	CreateProfile(ctx context.Context, p Profile) error
	DeleteProfile(ctx context.Context, id string) error
	UpdateProfile(ctx context.Context, id string, p Profile) error
}

// Profile represents a single user profile
type Profile struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type service struct {
	db *gorm.DB
}

// NewService creates a new profile service
func NewService() Service {
	return &service{}
}

func (s *service) GetProfile(ctx context.Context, id string) (Profile, error) {
	var p Profile
	if err := s.db.First(&p, id).Error; err != nil {
		return Profile{}, err
	}
	return p, nil
}

func (s *service) CreateProfile(ctx context.Context, p Profile) error            {}
func (s *service) DeleteProfile(ctx context.Context, id string) error            {}
func (s *service) UpdateProfile(ctx context.Context, id string, p Profile) error {}
