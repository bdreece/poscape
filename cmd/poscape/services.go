package main

import (
	"github.com/bdreece/poscape/pkg/event"
	"github.com/bdreece/poscape/pkg/tui"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"
	"go.uber.org/fx"
)

var Services = fx.Module("services",
	fx.Provide(
		fx.Annotate(
			event.NewHub[string](),
			fx.ResultTags(`name:"messages"`),
		),
		fx.Annotate(
			event.NewHub[string](),
			fx.ResultTags(`name:"ssh"`),
		),
	),
	fx.Provide(
		activeterm.Middleware,
		logging.Middleware,
		tui.Middleware,
	),
)
