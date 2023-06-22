package utils

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Trans(key string, data map[string]interface{}) string {
	//return key
	param := &i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
	}
	message := ginI18n.MustGetMessage(param)
	if message == "" {
		message = key
	}
	return message
}
