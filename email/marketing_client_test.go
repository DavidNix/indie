package email

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/require"
)

type mockRenderer struct {
	gotWriter io.Writer
	err       error
}

func (m *mockRenderer) Render(ctx context.Context, w io.Writer) error {
	if ctx == nil {
		panic("nil context")
	}
	m.gotWriter = w
	return m.err
}

func TestMarketingClient_Send(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		apiKey := os.Getenv("SENDGRID_API_KEY")
		if apiKey == "" {
			t.Skip("SENDGRID_API_KEY not set")
		}

		client := NewMarketingClient(apiKey, "Test Sender", "test@example.com", []int{123})
		client.EnableSandbox()

		renderer := &mockRenderer{}
		content := Content{
			ToName:    "Test Recipient",
			ToEmail:   "recipient@example.com",
			Subject:   "Test Marketing Email",
			HTML:      renderer,
			PlainText: "Hello World",
		}

		err := client.Send(t.Context(), content, 123)
		require.NoError(t, err)
		require.NotNil(t, renderer.gotWriter, "renderer should have received a writer")

		trackingSettings := NewTrackingSettings()
		trackingSettings.ClickTracking = mail.NewClickTrackingSetting().SetEnable(true)
		err = client.Send(t.Context(), content, 123, WithTrackingSettings(trackingSettings))
		require.NoError(t, err)
		require.NotNil(t, renderer.gotWriter, "renderer should have received a writer")

		// With settings
		mailSettings := NewMailSettings()
		mailSettings.Footer = mail.NewFooterSetting().SetEnable(false)
		err = client.Send(t.Context(), content, 123, WithMailSettings(mailSettings))
		require.NoError(t, err)
		require.NotNil(t, renderer.gotWriter, "renderer should have received a writer")

		// With nil html
		content.HTML = nil
		err = client.Send(t.Context(), content, 123)
		require.NoError(t, err)
	})

	t.Run("negative unsubscribe group", func(t *testing.T) {
		client := NewMarketingClient("fake-key", "Test Sender", "test@example.com", []int{123})

		content := Content{
			ToName:    "Test Recipient",
			ToEmail:   "recipient@example.com",
			Subject:   "Test Marketing Email",
			HTML:      &mockRenderer{},
			PlainText: "Hello World",
		}

		err := client.Send(t.Context(), content, -1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "missing unsubscribe group")
	})

	t.Run("unsubscribe group not in allowed list", func(t *testing.T) {
		client := NewMarketingClient("fake-key", "Test Sender", "test@example.com", []int{123})

		content := Content{
			ToName:    "Test Recipient",
			ToEmail:   "recipient@example.com",
			Subject:   "Test Marketing Email",
			HTML:      &mockRenderer{},
			PlainText: "Hello World",
		}

		err := client.Send(t.Context(), content, 456)
		require.Error(t, err)
		require.Contains(t, err.Error(), "unsubscribe group 456 not in allowed groups")
	})

	t.Run("renderer fails", func(t *testing.T) {
		client := NewMarketingClient("fake-key", "Test Sender", "test@example.com", []int{123})

		content := Content{
			ToName:    "Test Recipient",
			ToEmail:   "recipient@example.com",
			Subject:   "Test Marketing Email",
			HTML:      &mockRenderer{err: fmt.Errorf("render failed")},
			PlainText: "Hello World",
		}

		err := client.Send(t.Context(), content, 123)
		require.Error(t, err)
		require.EqualError(t, err, "render html: render failed")
	})
}
