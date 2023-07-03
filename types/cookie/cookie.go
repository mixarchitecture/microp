package cookie

type Base struct {
	HttpOnly bool   `env:"COOKIE_HTTP_ONLY"`
	Secure   bool   `env:"COOKIE_SECURE"`
	SameSite string `env:"COOKIE_SAME_SITE"`
	Domain   string `env:"COOKIE_DOMAIN"`
}
