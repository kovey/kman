package serv

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kovey/cli-go/app"
	"github.com/kovey/cli-go/env"
	"github.com/kovey/cli-go/util"
	"github.com/kovey/debug-go/debug"
	"github.com/kovey/kman/client/cache"
	"github.com/kovey/kman/client/etcd"
)

type serv struct {
	*app.ServBase
}

func (s *serv) Flag(a app.AppInterface) error {
	a.FlagArg("config", "config manage")
	a.FlagNonValue("e", "edit .env config", "config")
	a.FlagNonValue("l", "view .env config", "config")
	a.FlagNonValue("c", "create .env config", "config")

	a.FlagArg("cache", "cache manage")
	a.FlagArg("list", "cache list", "cache")
	a.FlagArg("get", "get config value from cache", "cache")
	a.FlagLong("name", nil, app.TYPE_STRING, "cached name", "cache", "get")
	a.FlagLong("key", nil, app.TYPE_STRING, "cached key", "cache", "get")

	return nil
}

func (s *serv) Panic(a app.AppInterface) {
	cache.Cached()
	s.ServBase.Panic(a)
}

func (s *serv) Init(app.AppInterface) error {
	return nil
}

func (s *serv) _init(app.AppInterface) error {
	if err := etcd.Init(); err != nil {
		return err
	}

	return cache.Cached()
}

func (s *serv) start(a app.AppInterface) error {
	if !env.CheckDefault() {
		return fmt.Errorf(".env config file not found")
	}

	if err := s._init(a); err != nil {
		return err
	}

	etcd.Watch()
	return cache.Cached()
}

func (s *serv) config(a app.AppInterface) error {
	if f, _ := a.Get("config", "l"); f.IsInput() {
		content, err := os.ReadFile(".env")
		if err != nil {
			return err
		}

		fmt.Println(string(content))
		return nil
	}

	if f, _ := a.Get("config", "e"); f.IsInput() {
		commands := []string{"vim", "vi", "nvim", "emacs", "gedit"}
		for _, command := range commands {
			if _, err := exec.LookPath(command); err == nil {
				cmd := exec.Command("vim", ".env")
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				return cmd.Run()
			}
		}

		return fmt.Errorf("editor not install, please install one of [%s]", strings.Join(commands, ","))
	}

	if f, _ := a.Get("config", "c"); f.IsInput() {
		if util.IsFile(".env") {
			return fmt.Errorf(".env is exists")
		}

		return os.WriteFile(".env", []byte(env_config), 0644)
	}

	return fmt.Errorf("options is empty")
}

func (s *serv) cache(a app.AppInterface) error {
	command, err := a.Arg(1, app.TYPE_STRING)
	if err != nil {
		return err
	}

	switch command.String() {
	case "list":
		namespaces := strings.Split(os.Getenv("CONFIG_NAMESPACE"), ",")
		fmt.Println(strings.Join(namespaces, "\r\n"))
		return nil
	case "get":
		name, _ := a.Get("cache", "get", "name")
		if !name.IsInput() {
			return fmt.Errorf("--name not input")
		}
		key, _ := a.Get("cache", "get", "key")
		if !key.IsInput() {
			return fmt.Errorf("--key not input")
		}

		if err := cache.LoadFromCache(name.String()); err != nil {
			return err
		}

		m := cache.Get(name.String(), key.String())
		if m == nil {
			return fmt.Errorf("%s of %s not found", key, name)
		}

		fmt.Printf("%s.%s: %s\n", name, key, m.Value)
		return nil
	}

	return fmt.Errorf("command of cache is empty")
}

func (s *serv) Run(a app.AppInterface) error {
	method, err := a.Arg(0, app.TYPE_STRING)
	if err != nil {
		s.Usage()
		return err
	}

	switch method.String() {
	case "config":
		return s.config(a)
	case "cache":
		return s.cache(a)
	default:
		return s.start(a)
	}
}

func (s *serv) Shutdown(app.AppInterface) error {
	etcd.Close()
	return nil
}

func Run() {
	cli := app.NewApp("kman-client")
	cli.SetServ(&serv{})
	if err := cli.Run(); err != nil {
		debug.Erro(err.Error())
	}
}
