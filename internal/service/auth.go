package service

import (
	"errors"
	"fmt"
	"main/internal/models"
	"main/internal/repository"
	"math/rand"
	"time"
)

type AuthService struct{
	userRepo repository.UserRepository
	authRepo repository.AuthRepository
}

func NewAuthService(userRepo repository.UserRepository, authRepo repository.AuthRepository) *AuthService{
	return &AuthService{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

func (s *AuthService) RequestForgotPassword(email string) error{
	user, err := s.userRepo.GetByEmail(email)
	if  err != nil {
		return err
	}
	if user == nil {
		return errors.New("email not found")
	}

	otpCode := fmt.Sprintf("%06d", rand.Intn(1000000))
	
	data := models.ForgotPassword{
		Email: email,
		OTPCode: otpCode,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(5*time.Minute),
	}

	err = s.authRepo.CreateForgotPassword(data)
	if  err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ResetPassword(email string, code string, newPassword string) error{
	data, err := s.authRepo.GetDataByEmailCode(email, code)
	if err != nil {
		return  err
	}

	if  time.Now().After(data.ExpiredAt) {
		return errors.New("OTP expired!!!")
	}

	err = s.userRepo.UpdatePasswordByEmail(email, newPassword)
	if err != nil {
		return  err
	}

	err = s.authRepo.DeleteDataByCode(code)
	if err != nil{
		return err
	}

	return  nil
}