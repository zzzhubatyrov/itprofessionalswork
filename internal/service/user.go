package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io/ioutil"
	"ipw-clean-arch/internal/model"
	"ipw-clean-arch/internal/repository"
	"ipw-clean-arch/internal/utils"
	"log"
	"strconv"
)

type UserServices struct {
	repo repository.UserHandler
}

func NewUserService(repo repository.UserHandler) *UserServices {
	return &UserServices{repo: repo}
}

func (u *UserServices) GetUser(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error) {
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims, err := utils.GetToken(token)
	if err != nil {
		return nil, fmt.Errorf("неверный формат токена")
	}
	user, err := u.repo.GetUser(data, claims)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServices) UpdateUser(data model.User, secretKey string, c *fiber.Ctx) (*model.User, error) {
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	updateUser, err := u.repo.UpdateUser(data, claims)
	if err != nil {
		return nil, err
	}
	//if err := elasticsearch.IndexUser(*updateUser); err != nil {
	//	return nil, err
	//}
	return updateUser, nil
}

func (u *UserServices) UploadPhoto(secretKey string, c *fiber.Ctx) (*model.User, error) {
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	file, err := c.FormFile("photo")
	if err != nil {
		return nil, err
	}
	fileBytes, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileBytes.Close()
	if err != nil {
		log.Fatal(err)
	}
	photoData, err := ioutil.ReadAll(fileBytes)
	if err != nil {
		return nil, err
	}
	uploadPhoto, err := u.repo.UploadPhoto(claims, photoData)
	if err != nil {
		return nil, err
	}
	return uploadPhoto, nil
}

// GetAllUsers TODO Add check for token
func (u *UserServices) GetAllUsers(data []model.User) ([]model.User, error) {
	getAllUsers, err := u.repo.GetAllUsers(data)
	if err != nil {
		return nil, err
	}
	return getAllUsers, nil
}

func (u *UserServices) CreateResume(data model.Resume, secretKey string, c *fiber.Ctx) (*model.Resume, error) {
	var user model.User
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	getUser, err := u.repo.GetUser(user, claims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	resume := &model.Resume{
		UserID:      getUser.ID,
		Email:       getUser.Email,
		Name:        getUser.Name,
		Gender:      getUser.Gender,
		Tag:         getUser.Tag,
		Direction:   data.Direction,
		Level:       data.Level,
		Salary:      data.Salary,
		Location:    data.Location,
		Status:      data.Status,
		Description: data.Description,
		Skills:      data.Skills,
	}
	createResume, err := u.repo.CreateResume(resume)
	if err != nil {
		return nil, err
	}
	return createResume, nil
}

func (u *UserServices) UpdateResume(data model.Resume, id, secretKey string, c *fiber.Ctx) (*model.Resume, error) {
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	resume := &model.Resume{
		Email:       data.Email,
		Name:        data.Name,
		Tag:         data.Tag,
		Direction:   data.Direction,
		Level:       data.Level,
		Salary:      data.Salary,
		Location:    data.Location,
		Status:      data.Status,
		Description: data.Description,
		Skills:      data.Skills,
	}
	updateResume, err := u.repo.UpdateResume(resume, id)
	if err != nil {
		return nil, err
	}
	return updateResume, nil
}

func (u *UserServices) GetResume() {
	//TODO implement me
	panic("implement me")
}

func (u *UserServices) GetResumeByID(id int) (*model.Resume, error) {
	getResumeByID, err := u.repo.GetResumeByID(id)
	if err != nil {
		return nil, err
	}
	return getResumeByID, nil
}

func (u *UserServices) GetAllResumes(data []model.Resume) ([]model.Resume, error) {
	getAllResumes, err := u.repo.GetAllResumes(data)
	if err != nil {
		return nil, err
	}
	return getAllResumes, nil
}

func (u *UserServices) DeleteResume(id string) error {
	return u.repo.DeleteResume(id)
}

func (u *UserServices) CreateResponse(data model.Vacancy, secretKey string, c *fiber.Ctx) (*model.Response, error) {
	var user model.User
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	getUser, err := u.repo.GetUser(user, claims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	response := &model.Response{
		UserID:      getUser.ID,
		VacancyID:   data.ID,
		CompanyName: data.CompanyName,
		JobTitle:    data.Level + " " + data.Direction,
		Applicant:   getUser.Name,
		Email:       getUser.Email,
		Phone:       getUser.Number,
	}
	createResponse, err := u.repo.CreateResponse(response)
	if err != nil {
		return nil, err
	}
	return createResponse, nil
}

func (u *UserServices) CreateCompany(company model.Company, secretKey string, c *fiber.Ctx) (*model.Company, error) {
	var user model.User
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	getUser, err := u.repo.GetUser(user, claims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	if claims.Valid() != nil {
		return nil, fmt.Errorf("невалидный токен: %v", claims.Valid())
	}
	file, err := c.FormFile("photo")
	if err != nil {
		return nil, err
	}
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()
	img, _, err := image.Decode(fileContent)
	if err != nil {
		return nil, err
	}
	img = resize.Resize(800, 600, img, resize.Lanczos3)
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, nil)
	if err != nil {
		return nil, err
	}
	usrID, _ := strconv.Atoi(claims.Issuer)
	data := &model.Company{
		UserID:      usrID,
		Photo:       buf.Bytes(),
		Name:        company.Name,
		Tag:         company.Tag,
		Email:       company.Email,
		Phone:       company.Phone,
		Location:    company.Location,
		Description: company.Description,
		CompanySize: company.CompanySize,
		WebSite:     company.WebSite,
	}
	if err != nil {
		return nil, err
	}
	createCompany, err := u.repo.CreateCompany(data, getUser, claims)
	if err != nil {
		return nil, err
	}
	return createCompany, nil
}

func (u *UserServices) UpdateCompanyData(company model.Company, secretKey string, c *fiber.Ctx) (*model.Company, error) {
	var user model.User
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	getUser, err := u.repo.GetUser(user, claims)
	if err != nil {
		return nil, errors.New("у вас нет компании")
	}
	usrID, _ := strconv.Atoi(claims.Issuer)
	updateCompany := &model.Company{
		ID:          getUser.Company.ID,
		UserID:      usrID,
		Photo:       getUser.Company.Photo,
		Name:        company.Name,
		Tag:         company.Tag,
		Email:       company.Email,
		Phone:       company.Phone,
		Location:    company.Location,
		Description: company.Description,
		CompanySize: company.CompanySize,
		WebSite:     company.WebSite,
		Vacancy:     getUser.Company.Vacancy,
	}
	updateCompanyData, err := u.repo.UpdateCompanyData(updateCompany, getUser)
	if err != nil {
		return nil, err
	}
	return updateCompanyData, nil
}

//updateCompanyData, err := u.repo.UpdateCompanyData(*updateCompany, getUser)
//if err != nil {
//	return nil, err
//}

func (u *UserServices) GetAllCompanies(company []model.Company) ([]model.Company, error) {
	getAllCompanies, err := u.repo.GetAllCompanies(company)
	if err != nil {
		return nil, err
	}
	return getAllCompanies, nil
}

func (u *UserServices) UpdateRoleByUserID(userID string, roleID int) error {
	updateRoleByUserID := u.repo.UpdateRoleByUserID(userID, roleID)
	return updateRoleByUserID
}

func (u *UserServices) GetCompanyByID(id string) (*model.Company, error) {
	getCompanyByID, err := u.repo.GetCompanyByID(id)
	if err != nil {
		return nil, err
	}
	return getCompanyByID, nil
}

func (u *UserServices) CreateVacancy(data model.Vacancy, secretKey string, c *fiber.Ctx) (*model.Vacancy, error) {
	cookie := c.Cookies("ipw")
	token, err := utils.ParseToken(cookie, secretKey)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthenticated")
	}
	if !token.Valid {
		return nil, errors.New("недействительный JWT токен")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("неверный формат токена")
	}
	createVacancy, err := u.repo.CreateVacancy(data, claims)
	if err != nil {
		return nil, err
	}
	return createVacancy, nil
}

func (u *UserServices) GetAllVacancy(data []model.Vacancy) ([]model.Vacancy, error) {
	getAllVacancy, err := u.repo.GetAllVacancy(data)
	if err != nil {
		return nil, err
	}
	return getAllVacancy, nil
}

func (u *UserServices) GetVacancyByID(id string) (*model.Vacancy, error) {
	getVacancyByID, err := u.repo.GetVacancyByID(id)
	if err != nil {
		return nil, err
	}
	return getVacancyByID, nil
}

func (u *UserServices) UpdateVacancy() {
	//TODO implement me
	panic("implement me")
}

func (u *UserServices) DeleteVacancy() {
	//TODO implement me
	panic("implement me")
}
