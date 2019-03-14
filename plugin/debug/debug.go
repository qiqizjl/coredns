package debug

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("debug", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	config := dnsserver.GetConfig(c)
	mem := true

	for c.Next() {
		args := c.RemainingArgs()
		if len(args) != 0 {
			return plugin.Error("debug", c.ArgErr())
		}

		config.Debug = true
		for c.NextBlock() {
			switch c.Val() {
			case "memory":
				if len(c.RemainingArgs()) != 0 {
					return plugin.Error("debug", c.ArgErr())
				}
				mem = true
			default:
				return plugin.Error("debug", c.ArgErr())
			}
		}
	}
	if mem {
		stop := make(chan struct{})
		c.OnStartup(func() error { reportMemory(stop); return nil })
		c.OnRestart(func() error { stop <- struct{}{}; return nil })
		c.OnFinalShutdown(func() error { stop <- struct{}{}; return nil })
	}

	return nil
}
