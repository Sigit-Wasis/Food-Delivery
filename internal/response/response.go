package response

import "github.com/gofiber/fiber/v2"

// Response adalah struktur untuk semua respons API
type Response struct {
	StatusCode int         `json:"status_code"` // Status HTTP
	Success    bool        `json:"success"`     // Indikator keberhasilan
	Message    string      `json:"message"`     // Pesan untuk klien
	Data       interface{} `json:"data,omitempty"`  // Data yang dikembalikan, jika ada
	Error      interface{} `json:"error,omitempty"` // Informasi error, jika ada
}

// NewSuccessResponse membuat respons sukses
func NewSuccessResponse(statusCode int, message string, data interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Success:    true,
		Message:    message,
		Data:       data,
		Error:      nil,
	}
}

// NewErrorResponse membuat respons error
func NewErrorResponse(statusCode int, message string, err interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Success:    false,
		Message:    message,
		Data:       nil,
		Error:      err,
	}
}

// SendResponse mengirimkan respons menggunakan Fiber
func SendResponse(c *fiber.Ctx, resp Response) error {
	return c.Status(resp.StatusCode).JSON(resp)
}
