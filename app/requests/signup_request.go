package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}

	return govalidator.New(opts).ValidateStruct()

}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email": []string{"required", "email", "min:4", "max:30"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
	}

	return govalidator.New(opts).ValidateStruct()

}
