package app

import (
	"time"

	"github.com/kovey/cli-go/env"
)

var data = env.LoadDefault(time.Now())
