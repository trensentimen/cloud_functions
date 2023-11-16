package gcf

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/whatsauth/webhook"
)

func init() {
	functions.HTTP("WebHook", adaptor.FiberHandlerFunc(webhook.PostMessage))
}
