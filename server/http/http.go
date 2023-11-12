package http

import (
	"fmt"

	"github.com/mixarchitecture/microp/server/http/error_handler"

	"github.com/goccy/go-json"
	"github.com/mixarchitecture/i18np"

	i18nHttp "github.com/mixarchitecture/microp/server/http/i18n"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host           string
	Port           int
	Group          string
	AppName        string
	CreateHandler  func(router fiber.Router) fiber.Router
	I18n           *i18np.I18n
	BodyLimit      int
	ReadBufferSize int
}

func RunServer(cfg Config) {
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	RunServerOnAddr(addr, cfg)
}

func RunServerOnAddr(addr string, cfg Config) {
	if cfg.AppName == "" {
		cfg.AppName = "Mix Architecture"
	}
	if cfg.BodyLimit == 0 {
		cfg.BodyLimit = 5 * 1024 * 1024
	}
	if cfg.ReadBufferSize == 0 {
		cfg.ReadBufferSize = 5 * 1024 * 1024
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: error_handler.New(error_handler.Config{
			// DfMsgKey: "error_internal_server_error",
			I18n: cfg.I18n,
		}),
		JSONEncoder:    json.Marshal,
		JSONDecoder:    json.Unmarshal,
		CaseSensitive:  true,
		AppName:        cfg.AppName,
		ServerHeader:   cfg.AppName,
		BodyLimit:      cfg.BodyLimit,
		ReadBufferSize: cfg.ReadBufferSize,
	})
	group := app.Group(fmt.Sprintf("/%v", cfg.Group))
	setGlobalMiddlewares(app, cfg)
	cfg.CreateHandler(group)

	logrus.Infof("Starting server on %v", addr)
	if err := app.Listen(addr); err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func setGlobalMiddlewares(router fiber.Router, cfg Config) {
	router.Use(recover.New(recover.ConfigDefault), compress.New(compress.Config{}), i18nHttp.New(*cfg.I18n))
}
