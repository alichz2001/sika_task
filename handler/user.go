package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sika/service"
	"strconv"
)

type UserHandlerImpl struct {
	userSrv *service.UserServiceImpl
}

func NewUserHandler(app *fiber.App, userSrv *service.UserServiceImpl) *UserHandlerImpl {
	h := &UserHandlerImpl{userSrv: userSrv}

	app.Get("/users/:user_id", h.Get)
	app.Get("/users", h.GetPage)

	return h
}

func (handler *UserHandlerImpl) Get(ctx *fiber.Ctx) error {
	reqUserIDStr := ctx.Params("user_id")

	reqUserIDInt, err := strconv.Atoi(reqUserIDStr)

	if reqUserIDStr == "" || err != nil {
		return ctx.Status(http.StatusBadRequest).Send([]byte{})
	}

	user, err := handler.userSrv.GetByID(uint(reqUserIDInt))
	if err != nil {
		return ctx.Status(http.StatusNotFound).Send([]byte{})
	}

	respBody, err := json.Marshal(user)
	return ctx.Status(http.StatusOK).Send(respBody)
}

func (handler *UserHandlerImpl) GetPage(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page")
	countStr := ctx.Query("count")

	countInt, err := strconv.Atoi(countStr)
	pageInt, err := strconv.Atoi(pageStr)

	if err != nil || countInt > 100 || countInt < 10 || pageInt < 1 {
		return ctx.Status(http.StatusBadRequest).Send([]byte{})
	}

	users, err := handler.userSrv.GetPage(pageInt, countInt)
	if err != nil {
		return ctx.Status(http.StatusNotFound).Send([]byte{})
	}

	respBody, err := json.Marshal(users)
	return ctx.Status(http.StatusOK).Send(respBody)
}
