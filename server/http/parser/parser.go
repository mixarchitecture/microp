package parser

import (
	i18nHttp "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/mixarchitecture/microp/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
)

func ParseBody(c *fiber.Ctx, v validator.Validator, i i18np.I18n, d interface{}) {
	l, a := i18nHttp.GetLanguagesInContext(i, c)
	if err := c.BodyParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_body", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func ParseQuery(c *fiber.Ctx, v validator.Validator, i i18np.I18n, d interface{}) {
	l, a := i18nHttp.GetLanguagesInContext(i, c)
	if err := c.QueryParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_query", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func ParseParams(c *fiber.Ctx, v validator.Validator, i i18np.I18n, d interface{}) {
	l, a := i18nHttp.GetLanguagesInContext(i, c)
	if err := c.ParamsParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_params", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func validateStruct(d interface{}, v validator.Validator, i i18np.I18n, l, a string) {
	if errors := v.ValidateStruct(d, l, a); len(errors) > 0 {
		panic(result.ErrorDetail(i.Translate("error_validation_failed", l, a), errors, fiber.StatusBadRequest))
	}
}
