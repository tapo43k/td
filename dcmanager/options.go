package dcmanager

import (
	"github.com/gotd/td/bin"
	"github.com/gotd/td/transport"
	"go.uber.org/zap"
)

type Options struct {
	UpdateHandler func(b *bin.Buffer) error
	ConfigHandler func(Config) error

	Device DeviceConfig

	Transport *transport.Transport
	Network   string

	Config *Config
	Logger *zap.Logger
}

func (o *Options) setDefaults() {
	if o.UpdateHandler == nil {
		// TODO(ccln): disable updates using tg.InvokeWithoutUpdates on primary dc?
		o.UpdateHandler = func(b *bin.Buffer) error { return nil }
	}
	if o.ConfigHandler == nil {
		o.ConfigHandler = func(c Config) error { return nil }
	}
	o.Device.setDefaults()
	if o.Logger == nil {
		o.Logger = zap.NewNop()
	}
}

// DeviceConfig is config which send when Telegram connection session created.
type DeviceConfig struct {
	// Device model
	DeviceModel string
	// Operation system version
	SystemVersion string
	// Application version
	AppVersion string
	// Code for the language used on the device's OS, ISO 639-1 standard
	SystemLangCode string
	// Language pack to use
	LangPack string
	// Code for the language used on the client, ISO 639-1 standard
	LangCode string
}

func (c *DeviceConfig) setDefaults() {
	const notAvailable = "n/a"

	if c.DeviceModel == "" {
		c.DeviceModel = notAvailable
	}
	if c.SystemVersion == "" {
		c.SystemVersion = notAvailable
	}
	if c.AppVersion == "" {
		c.AppVersion = notAvailable
	}
	if c.SystemLangCode == "" {
		c.SystemLangCode = "en"
	}
	if c.LangCode == "" {
		c.LangCode = "en"
	}
}