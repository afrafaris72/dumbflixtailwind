package handlers

import (
	dto "dumbflix/dto/result"
	usersdto "dumbflix/dto/users"
	"dumbflix/models"
	"dumbflix/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})
}

func (h *handler) GetUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)})
}

func (h *handler) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")

	user, err := h.UserRepository.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)})
}

func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})

}

func (h *handler) UpdateUserPicture(c echo.Context) error {

	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	request := usersdto.UpdateUserPicture{
		Email:     c.FormValue("email"),
		Fullname:  c.FormValue("fullname"),
		Gender:    c.FormValue("gender"),
		Phone:     c.FormValue("phone"),
		Address:   c.FormValue("address"),
		Thumbnail: dataFile,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Thumbnail != "" {
		user.Thumbnail = request.Thumbnail
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Address != "" {
		user.Address = request.Address
	}

	data, err := h.UserRepository.UpdateUserPicture(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:        u.ID,
		Fullname:  u.Fullname,
		Email:     u.Email,
		Password:  u.Password,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Address:   u.Address,
		Subscribe: u.Subscribe,
		Role:      u.Role,
	}
}
