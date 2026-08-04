package profile

import (
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	CreateProfile(userID uint, req CreateProfileRequest) (*Profile, error)
	GetProfile(userID uint) (*Profile, error)
	UpdateProfile(userID uint, req UpdateProfileRequest) (*Profile, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProfile(userID uint, req CreateProfileRequest) (*Profile, error) {

	// Prevent users from creating multiple profiles
	_, err := s.repo.GetByUserID(userID)

	if err == nil {
		return nil, errors.New("profile already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	profile := &Profile{
		UserID: userID,

		Specialization: req.Specialization,
		Hospital:       req.Hospital,
		Country:        req.Country,
		City:           req.City,

		YearsExperience: req.YearsExperience,

		LicenseNumber: req.LicenseNumber,
		Education:     req.Education,

		Languages: req.Languages,

		Bio: req.Bio,

		ProfileImageURL: req.ProfileImageURL,
	}

	if err := s.repo.Create(profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *service) GetProfile(userID uint) (*Profile, error) {
	return s.repo.GetByUserID(userID)
}

func (s *service) UpdateProfile(userID uint, req UpdateProfileRequest) (*Profile, error) {

	profile, err := s.repo.GetByUserID(userID)

	if err != nil {
		return nil, err
	}

	profile.Specialization = req.Specialization
	profile.Hospital = req.Hospital
	profile.Country = req.Country
	profile.City = req.City
	profile.YearsExperience = req.YearsExperience
	profile.LicenseNumber = req.LicenseNumber
	profile.Education = req.Education
	profile.Languages = req.Languages
	profile.Bio = req.Bio
	profile.ProfileImageURL = req.ProfileImageURL

	if err := s.repo.Update(profile); err != nil {
		return nil, err
	}

	return profile, nil
}
