package http

import (
	"net/http"

	"github.com/ferjmc/api_ddd/user/config"
	"github.com/ferjmc/api_ddd/user/internal/models"
	"github.com/ferjmc/api_ddd/user/internal/user"
	httpErrors "github.com/ferjmc/api_ddd/user/pkg/http_errors"
	"github.com/ferjmc/api_ddd/user/pkg/logger"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

const (
	csrfHeader  = "X-CSRF-Token"
	maxFileSize = 1024 * 1024 * 10
)

type userHandlers struct {
	cfg    *config.Config
	group  *echo.Group
	userUC user.UseCase
	logger logger.Logger
}

func NewUserHandlers(
	group *echo.Group,
	userUC user.UseCase,
	logger logger.Logger,
	cfg *config.Config,
) *userHandlers {
	return &userHandlers{group: group, userUC: userUC, logger: logger, cfg: cfg}
}

// Register godoc
// @Summary Register new user
// @Tags User
// @Description register new user account, returns user data and session
// @Accept json
// @Produce json
// @Param data body models.User true "user data"
// @Success 201 {object} models.UserResponse
// @Router /user/register [post]
func (h *userHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var u models.User
		if err := c.Bind(&u); err != nil {
			h.logger.Errorf("c.Bind: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		regUser, err := h.userUC.Register(ctx, &u)
		if err != nil {
			h.logger.Errorf("userHandlers.Register.userUC.Register: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusCreated, regUser)
	}
}

// Login godoc
// @Summary Login user
// @Tags User
// @Description login user, returns user data and session
// @Accept json
// @Produce json
// @Param data body models.Login true "email and password"
// @Success 200 {object} models.UserResponse
// @Router /user/login [post]
func (h *userHandlers) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var login models.Login
		if err := c.Bind(&login); err != nil {
			h.logger.Errorf("c.Bind: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		userResponse, err := h.userUC.Login(ctx, login)
		if err != nil {
			h.logger.Errorf("userHandlers.userUC.Login: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusOK, userResponse)
	}
}

// Logout godoc
// @Summary Logout user
// @Tags User
// @Description Logout user, return no content
// @Accept json
// @Produce json
// @Success 204 ""
// @Router /user/logout [post]
func (h *userHandlers) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusNoContent)
	}
}

// GetMe godoc
// @Summary Get current user data
// @Tags User
// @Description Get current user data, required session cookie
// @Accept json
// @Produce json
// @Success 200 {object} models.UserResponse
// @Router /user/me [get]
func (h *userHandlers) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {

		var userResponse models.UserResponse

		return c.JSON(http.StatusOK, userResponse)
	}
}

// Update godoc
// @Summary Update user
// @Tags User
// @Description update user profile
// @Accept json
// @Produce json
// @Success 200 {object} models.UserResponse
// @Router /user/{id} [get]
func (h *userHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		userID := c.Param("id")
		if userID == "" {
			h.logger.Error("invalid user id param")
			return httpErrors.ErrorCtxResponse(c, httpErrors.BadRequest)
		}

		userUUID, err := uuid.FromString(userID)
		if err != nil {
			h.logger.Error("invalid user uuid")
			return httpErrors.ErrorCtxResponse(c, err)
		}

		var updUser models.UserUpdate
		if err := c.Bind(&updUser); err != nil {
			h.logger.Errorf("c.Bind: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}
		updUser.ID = userUUID

		userResponse, err := h.userUC.Update(ctx, &updUser)
		if err != nil {
			h.logger.Errorf("userHandlers.userUC.Update: %v", err)
			return httpErrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusOK, userResponse)
	}
}

func (h *userHandlers) Delete() echo.HandlerFunc {
	panic("implement me")
}

// GetUserByID godoc
// @Summary Get user by id
// @Tags User
// @Description Get user data by id
// @Accept json
// @Produce json
// @Param id path int false "user uuid"
// @Success 200 {object} models.UserResponse
// @Router /user/{id} [get]
func (h *userHandlers) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		userID := c.Param("id")
		if userID == "" {
			h.logger.Error("invalid user id param")
			return httpErrors.ErrorCtxResponse(c, httpErrors.BadRequest)
		}

		userUUID, err := uuid.FromString(userID)
		if err != nil {
			h.logger.Error("uuid.FromString")
			return httpErrors.ErrorCtxResponse(c, err)
		}

		userResponse, err := h.userUC.GetByID(ctx, userUUID)
		if err != nil {
			h.logger.Error("userUC.GetByID")
			return httpErrors.ErrorCtxResponse(c, err)
		}

		return c.JSON(http.StatusOK, userResponse)
	}
}

// UpdateAvatar godoc
// @Summary Update user avatar
// @Tags User
// @Description Upload user avatar image
// @Accept mpfd
// @Produce json
// @Param id path int false "user uuid"
// @Success 200 {object} models.UserResponse
// @Router /user/{id}/avatar [put]
