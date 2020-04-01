// +build !confonly

package porter

//go:generate errorgen

import (
	"context"
	"sync"

	"v2ray.com/core/app/porter/config"
	pcommon "v2ray.com/core/app/porter/src/common"
	"v2ray.com/core/app/porter/src/proxy"
	plog "v2ray.com/core/app/porter/src/utils/log"
	"v2ray.com/core/common"
	"v2ray.com/core/common/log"
)

// Instance is a log.Handler that handles logs.
type Instance struct {
	sync.RWMutex
	config *config.Config
	active bool
}

// New creates a new log.Instance based on the given config.
func New(ctx context.Context, config *config.Config) (*Instance, error) {
	g := &Instance{
		config: config,
		active: false,
	}
	return g, nil
}

// Type implements common.HasType.
func (*Instance) Type() interface{} {
	return (*Instance)(nil)
}

// Start implements common.Runnable.Start().
func (g *Instance) Start() error {
	config.Parameters = g.config

	pcommon.Init()

	plog.InitLog(int(g.config.LogLevel), g.config.LogDir, plog.Stdout)

	go pcommon.CheckLogFileSize()

	pcommon.StartMonitor()

	proxy.Run()

	return nil
}

// Handle implements log.Handler.
func (g *Instance) Handle(msg log.Message) {
	g.RLock()
	defer g.RUnlock()

	if !g.active {
		return
	}
}

// Close implements common.Closable.Close().
func (g *Instance) Close() error {
	newError("Logger closing").AtDebug().WriteToLog()
	g.Lock()
	defer g.Unlock()
	return nil
}

func init() {
	common.Must(common.RegisterConfig((*config.Config)(nil), func(ctx context.Context, conf interface{}) (interface{}, error) {
		return New(ctx, conf.(*config.Config))
	}))
}
