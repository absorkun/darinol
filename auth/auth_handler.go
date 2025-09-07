package auth

import (
	"errors"
	"os"
	"time"

	"github.com/absorkun/darinol/model"
	"github.com/absorkun/darinol/response"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{db}
}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Failed 404 {object} response.FailedStruct
// @Failed 500 {object} response.FailedStruct
// @Param Auth body LoginDto true "Info Auth Login info"
// @Router /api/v1/auth/login [post]
func (h *handler) Login(c fiber.Ctx) error {
	var user model.User
	var dto LoginDto

	if err := c.Bind().Body(&dto); err != nil {
		return response.BadRequest(c, err.Error())
	}

	if err := h.DB.Model(user).First(&user, "email = ?", dto.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, "Email is invalid")
		}
		return response.InternalServerError(c, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return response.BadRequest(c, "Password is invalid")
	}

	var claims = jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}
	var jwt = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := jwt.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.Ok(c, fiber.Map{"access_token": t})
}

// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Failed 500 {object} response.FailedStruct
// @Param Auth body RegisterDto true "Info Auth Register info"
// @Router /api/v1/auth/register [post]
func (h *handler) Register(c fiber.Ctx) error {
	var dto RegisterDto

	if err := c.Bind().Body(&dto); err != nil {
		return response.BadRequest(c, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.InternalServerError(c, err.Error())
	}

	var user = model.User{
		Id:       dto.Id,
		Email:    dto.Email,
		Password: string(hashedPassword),
		Role:     dto.Role,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, dto)
}
