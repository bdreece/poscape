package ssh

import (
	"fmt"
	"net"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"go.uber.org/fx"
)

type Options struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	KeyPath string `yaml:"key_path"`
}

type Params struct {
	fx.In

	Middleware []wish.Middleware `group:"middleware"`
	Options    *Options
}

func NewServer(p Params) (*ssh.Server, error) {
	return wish.NewServer(
		wish.WithAddress(net.JoinHostPort(p.Options.Host, fmt.Sprint(p.Options.Port))),
		wish.WithHostKeyPath(p.Options.KeyPath),
		wish.WithMiddleware(p.Middleware...),
	)
}
