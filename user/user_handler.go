package user

import (
	"errors"

	"github.com/absorkun/darinol/model"
	"github.com/absorkun/darinol/response"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{db}
}

// @Summary Fetch all
// @Description Fetch all
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} response.SuccessStruct
// @Failed 500 {object} response.FailedStruct
// @Router /api/v1/users [get]
func (h *handler) GetAll(c fiber.Ctx) error {
	var user model.User
	var dto []UserGetDto

	if err := h.DB.Model(user).Preload("Todos").Find(&dto).Error; err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.Ok(c, dto)
}

// @Summary Fetch By Id
// @Description Fetch By Id
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessStruct
// @Failed 404 {object} response.FailedStruct
// @Failed 500 {object} response.FailedStruct
// @Param id path int true "User Id"
// @Router /api/v1/users/{id} [get]
func (h *handler) GetById(c fiber.Ctx) error {
	var id = c.Params("id")
	var user model.User
	var dto UserGetDto
	if err := h.DB.Model(user).First(&dto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, "Record User with id "+id+" not found")
		}
		return response.InternalServerError(c, err.Error())
	}
	return response.Ok(c, dto)
}

// @Summary Create
// @Description Create
// @Tags User
// @Accept json
// @Produce json
// @Success 201 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Failed 500 {object} response.FailedStruct
// @Param user body UserCreateDto true "Input user info"
// @Router /api/v1/users [post]
func (h *handler) Create(c fiber.Ctx) error {
	var dto UserCreateDto

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

// @Summary Update
// @Description Update
// @Tags User
// @Accept json
// @Produce json
// @Success 204 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Failed 404 {object} response.FailedStruct
// @Failed 500 {object} response.FailedStruct
// @Param id path int true "User Id"
// @Param user body UserUpdateDto true "Input user info"
// @Router /api/v1/users/{id} [put]
func (h *handler) Update(c fiber.Ctx) error {
	var id = c.Params("id")
	var dto UserUpdateDto

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

	if err := h.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, "Record User with id "+id+" not found")
		}
		return response.BadRequest(c, err.Error())
	}

	return response.NoContent(c)
}

// @Summary Delete
// @Description Delete
// @Tags User
// @Accept json
// @Produce json
// @Success 204 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Failed 404 {object} response.FailedStruct
// @Param id path int true "User Id"
// @Router /api/v1/users/{id} [delete]
func (h *handler) Delete(c fiber.Ctx) error {
	var id = c.Params("id")
	var user model.User

	if err := h.DB.Unscoped().Where("id = ?", id).Delete(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, "Record User with id "+id+" not found")
		}
		return response.BadRequest(c, err.Error())
	}

	return response.NoContent(c)
}
