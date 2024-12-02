package fiberserver

import (
	"encoding/xml"
	"github.com/Juminiy/kube/pkg/netserver/http/stdserver"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/Juminiy/kube/pkg/util/safe_validator"
	"github.com/gofiber/fiber/v3"
)

func New() *fiber.App {
	return fiber.New(fiber.Config{
		ServerHeader:                 "",
		StrictRouting:                true,
		CaseSensitive:                true,
		Immutable:                    false,
		UnescapePath:                 false,
		BodyLimit:                    128 * util.Mi,
		Concurrency:                  1 * util.M,
		Views:                        nil,
		ViewsLayout:                  "",
		PassLocalsToViews:            false,
		ReadTimeout:                  0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		ReadBufferSize:               0,
		WriteBufferSize:              0,
		CompressedFileSuffixes:       nil,
		ProxyHeader:                  "",
		GETOnly:                      false,
		ErrorHandler:                 nil,
		DisableKeepalive:             false,
		DisableDefaultDate:           false,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		AppName:                      "",
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: true,
		ReduceMemoryUsage:            true,
		JSONEncoder:                  safe_json.SafeConfig().Marshal,
		JSONDecoder:                  safe_json.SafeConfig().Unmarshal,
		XMLEncoder:                   xml.Marshal,
		EnableTrustedProxyCheck:      false,
		TrustedProxies:               nil,
		EnableIPValidation:           false,
		ColorScheme:                  fiber.Colors{},
		StructValidator:              safe_validator.Strict(),
		RequestMethods:               stdserver.LawHTTPMethod,
		EnableSplittingOnParsers:     false,
	})
}
