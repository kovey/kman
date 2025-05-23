package refresh

import (
	"os"
	"time"

	"github.com/kovey/cli-go/env"
	"github.com/kovey/kman/kman-web/module/libs/token"
	"github.com/kovey/kow/jwt"
)

var r *jwt.Jwt[*token.Ext]

func init() {
	if env.HasEnv() {
		env.LoadDefault(time.Now())
	}

	expired, _ := env.GetInt("JWT_REFRESH_EXPIRE")
	r = jwt.NewJwt[*token.Ext](os.Getenv("JWT_REFRESH_KEY"), int32(expired))
}

func Token(ext *token.Ext) (string, error) {
	return r.Encode(ext)
}

func Decode(token string) (*token.Ext, error) {
	return r.Decode(token)
}
