package todo

import (
	"errors"

	"github.com/absorkun/darinol/model"
	"github.com/absorkun/darinol/response"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{db}
}

// @Summary Fetch All
// @Description Fetch All
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {array} response.SuccessStruct
// @Failed 500 {object} response.FailedStruct
// @Security ApiKeyAuth
// @Router /api/v1/todos [get]
func (h *handler) GetAll(c fiber.Ctx) error {
	var todo model.Todo
	var dto []TodoGetDto

	if err := h.DB.Model(todo).Find(&dto).Error; err != nil {
		return response.InternalServerError(c, err.Error())
	}

	return response.Ok(c, dto)
}

// @Summary Fetch By Id
// @Description Fetch By Id
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessStruct
// @Failed 404 {object} response.FailedStruct
// @Failed 500 {object} response.FailedStruct
// @Param id path int true "Todo Id"
// @Security ApiKeyAuth
// @Router /api/v1/todos/{id} [get]
func (h *handler) GetById(c fiber.Ctx) error {
	var id = c.Params("id")
	var todo model.Todo
	var dto TodoGetDto
	if err := h.DB.Model(todo).First(&dto, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, "Record Todo with id "+id+" not found")
		}
		return response.InternalServerError(c, err.Error())
	}
	return response.Ok(c, dto)
}

// @Summary Create
// @Description Create
// @Tags Todo
// @Accept json
// @Produce json
// @Success 201 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Param todo body TodoCreateDto true "Input todo info"
// @Security ApiKeyAuth
// @Router /api/v1/todos [post]
func (h *handler) Create(c fiber.Ctx) error {
	var dto TodoCreateDto

	if err := c.Bind().Body(&dto); err != nil {
		return response.BadRequest(c, err.Error())
	}

	var todo = model.Todo{
		Id:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
		Finished:    dto.Finished,
		UserId:      dto.UserId,
	}

	if err := h.DB.Create(&todo).Error; err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, dto)
}

// @Summary Update
// @Description Update
// @Tags Todo
// @Accept json
// @Produce json
// @Success 204 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Failed 404 {object} response.FailedStruct
// @Param id path int true "Todo Id"
// @Param todo body TodoUpdateDto true "Input todo info"
// @Security ApiKeyAuth
// @Router /api/v1/todos/{id} [put]
func (h *handler) Update(c fiber.Ctx) error {
	var id = c.Params("id")
	var dto TodoUpdateDto

	if err := c.Bind().Body(&dto); err != nil {
		return response.BadRequest(c, err.Error())
	}

	var todo = model.Todo{
		Id:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
		Finished:    dto.Finished,
		UserId:      dto.Id,
	}

	if err := h.DB.Where("id = ?", id).Updates(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, "Record Todo with id "+id+" not found")
		}
		return response.BadRequest(c, err.Error())
	}

	return response.NoContent(c)
}

// @Summary Delete
// @Description Delete
// @Tags Todo
// @Accept json
// @Produce json
// @Success 204 {object} response.SuccessStruct
// @Failed 400 {object} response.FailedStruct
// @Failed 404 {object} response.FailedStruct
// @Param id path int true "Todo Id"
// @Security ApiKeyAuth
// @Router /api/v1/todos/{id} [delete]
func (h *handler) Delete(c fiber.Ctx) error {
	var id = c.Params("id")
	var todo model.Todo

	if err := h.DB.Unscoped().Where("id = ?", id).Delete(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NotFound(c, "Record Todo with id "+id+" not found")
		}
		return response.BadRequest(c, err.Error())
	}

	return response.NoContent(c)
}
