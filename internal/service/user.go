package service

import (
	"errors"
	"main/internal/lib"
	"main/internal/models"
	"main/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() ([]models.User, error){
	user, err := s.repo.GetUser()
	if err != nil{
		return  nil, err
	}
	
	return user, nil
}

func (s *UserService) GetById(id int) (*models.User, error){
	user, err := s.repo.GetById(id)
		if err != nil{
		return  nil, err
	}
	
	return user, nil
}

func (s *UserService) GetByEmail(email string) (*models.User, error){
	user, err := s.repo.GetByEmail(email)
		if err != nil{
		return  nil, err
	}
	
	return user, nil
}

func (s *UserService) Register(req *models.CreateUserRequest) error{
	existingUser, _ := s.repo.GetByEmail(req.Email)

	if existingUser != nil{
		return  errors.New("Email was registered!")
	}

	hashed, err := lib.HashPassword(req.Password)
	if err != nil{
		return err
	}

	newUser := models.User{
		Fullname: req.Fullname,
		Email: req.Email,
		Password: hashed,
	}

	return  s.repo.Create(newUser)
}

func (s *UserService) Login(req models.LoginUserRequest) (string, error){
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil{
		return "", errors.New("Invalid email/password!")
	}

	ok := lib.VerifyPassword(req.Password, user.Password)
	if !ok{
		return "", errors.New("invalid email/password!")
	}

	if ok{
		token, err := lib.GenerateToken(user.ID)
		
		if err != nil {
			return  "", err
		}

		return  token, nil
	}

	return  "", err
}

func (s *UserService) Update(email string, req *models.UpdateUserRequest) (*models.User, error){
	user, err := s.repo.GetByEmail(email)
	if  err != nil {
		return  nil, errors.New("User not found")
	}

	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}

	if  req.Email != "" {
		user.Email = req.Email
	}

	if  req.Password != "" {
		user.Password = req.Password
	}

	err = s.repo.UpdateUser(*user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(email string) error{
	user, err := s.repo.GetByEmail(email)
	if err != nil{
		return errors.New("user not found!")
	}

	return s.repo.DeleteUser(user.ID)
}