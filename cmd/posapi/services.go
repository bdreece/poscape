package main

import (
	"github.com/bdreece/poscape/internal/event"
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
)
