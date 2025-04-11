package email

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var bufpool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func NewTrackingSettings() *mail.TrackingSettings {
	return mail.NewTrackingSettings()
}

func NewMailSettings() *mail.MailSettings {
	return mail.NewMailSettings()
}

// Option is a function type that modifies the baseClient
type Option func(*option)

type option struct {
	trackingSettings *mail.TrackingSettings
	mailSettings     *mail.MailSettings
}

// WithTrackingSettings sets custom tracking settings
// Overrides all settings (does not merge with defaults)
func WithTrackingSettings(settings *mail.TrackingSettings) Option {
	return func(o *option) {
		o.trackingSettings = settings
	}
}

// WithMailSettings sets custom mail settings
// Overrides all settings (does not merge with defaults)
func WithMailSettings(settings *mail.MailSettings) Option {
	return func(o *option) {
		o.mailSettings = settings
	}
}

type Renderer interface {
	// Render renders HTML to the writer
	Render(ctx context.Context, w io.Writer) error
}

type baseClient struct {
	client   *sendgrid.Client
	fromName string
	fromAddr string
	sandbox  bool
}

func (c *baseClient) buildTrackingSettings(o option) *mail.TrackingSettings {
	settings := o.trackingSettings
	if settings == nil {
		settings = mail.NewTrackingSettings()
		// turn off some tracking to improve deliverability
		settings.GoogleAnalytics = mail.NewGaSetting().SetEnable(false)
		settings.OpenTracking = mail.NewOpenTrackingSetting().SetEnable(false)
		// Prevents Sendgrid from injecting default unsubscribe
		settings.SubscriptionTracking = mail.NewSubscriptionTrackingSetting().SetEnable(false)
	}
	settings.SandboxMode = &mail.SandboxModeSetting{Enable: &c.sandbox}
	return settings
}

func (c *baseClient) buildMailSettings(o option) *mail.MailSettings {
	settings := o.mailSettings
	if settings == nil {
		settings = mail.NewMailSettings()
	}
	settings.SandboxMode = &mail.Setting{Enable: &c.sandbox}
	return settings
}

func (c *baseClient) renderHTML(ctx context.Context, html Renderer) (string, error) {
	if html == nil {
		return "", nil
	}
	buf := bufpool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufpool.Put(buf)

	if err := html.Render(ctx, buf); err != nil {
		return "", fmt.Errorf("render html: %w", err)
	}
	result := buf.String()
	return result, nil
}

func (c *baseClient) EnableSandbox() {
	c.sandbox = true
}

type Content struct {
	ToName    string
	ToEmail   string
	Subject   string
	HTML      Renderer
	PlainText string
}
