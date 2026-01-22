package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

// Khởi tạo instance dùng chung để tiết kiệm tài nguyên
var validate = validator.New()

// ErrorResponse định nghĩa cấu trúc lỗi trả về cho Client
type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateStruct thực hiện kiểm tra và format lại lỗi
func ValidateStruct(s interface{}) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(s)
	
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = strings.ToLower(err.Field())
			element.Message = msgForTag(err.Tag(), err.Param())
			errors = append(errors, element)
		}
	}
	return errors
}

// msgForTag ánh xạ các tag lỗi sang thông báo thân thiện
func msgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "Trường này không được để trống"
	case "email":
		return "Email không hợp lệ"
	case "min":
		return fmt.Sprintf("Độ dài tối thiểu là %s ký tự", param)
	case "gte":
		return fmt.Sprintf("Giá trị phải lớn hơn hoặc bằng %s", param)
	}
	return "Dữ liệu không hợp lệ"
}