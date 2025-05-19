package access

import (
	"os"
	"time"

	"github.com/kovey/cli-go/env"
	"github.com/kovey/kman/web/module/libs/token"
	"github.com/kovey/kow/jwt"
)

var t *jwt.Jwt[*token.Ext]

func init() {
	if env.HasEnv() {
		env.LoadDefault(time.Now())
	}
	expired, _ := env.GetInt("JWT_AUTH_EXPIRE")
	t = jwt.NewJwt[*token.Ext](os.Getenv("JWT_AUTH_KEY"), int32(expired))
}

func Token(ext *token.Ext) (string, error) {
	return t.Encode(ext)
}

func Decode(token string) (*token.Ext, error) {
	return t.Decode(token)
}
