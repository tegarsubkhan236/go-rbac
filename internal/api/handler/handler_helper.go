package handler

import "github.com/gofiber/fiber/v2"

type errorResponse struct {
	Code     any `json:"code"`
	Status   any `json:"status"`
	Messages any `json:"message"`
	Data     any `json:"data,omitempty"`
}

type successResponse struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	CurrentPage int    `json:"current_page,omitempty"`
	LastPage    int    `json:"last_page,omitempty"`
	TotalData   int    `json:"total_data,omitempty"`
	Data        any    `json:"data,omitempty"`
}

func ResponseOKWithPages(c *fiber.Ctx, currentPage, lastPage, totalData int, data any) error {
	return c.Status(fiber.StatusOK).JSON(&successResponse{
		Code:        fiber.StatusOK,
		Status:      "OK",
		CurrentPage: currentPage,
		LastPage:    lastPage,
		TotalData:   totalData,
		Data:        data,
	})
}

func ResponseOK(c *fiber.Ctx, data ...any) error {
	var responseData any
	if len(data) > 0 {
		responseData = data[0]
	}

	return c.Status(fiber.StatusOK).JSON(&successResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responseData,
	})
}

func ResponseCreated(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(&successResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   data,
	})
}

func ResponseBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{
		Code:     fiber.StatusBadRequest,
		Status:   "Bad Request",
		Messages: message,
	})
}

func ResponseUnauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(&errorResponse{
		Code:     fiber.StatusUnauthorized,
		Status:   "Unauthorized",
		Messages: message,
	})
}

func ResponseForbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(&errorResponse{
		Code:     fiber.StatusForbidden,
		Status:   "Forbidden",
		Messages: message,
	})
}

func ResponseNotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(&errorResponse{
		Code:     fiber.StatusNotFound,
		Status:   "Not Found",
		Messages: message,
	})
}

func ResponseInternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(&errorResponse{
		Code:     fiber.StatusInternalServerError,
		Status:   "Internal Server Error",
		Messages: message,
	})
}
