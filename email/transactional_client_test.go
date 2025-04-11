package email

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/require"
)

func TestTransactionalClient_Send(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		apiKey := os.Getenv("SENDGRID_API_KEY")
		if apiKey == "" {
			t.Skip("SENDGRID_API_KEY not set")
		}

		renderer := &mockRenderer{}
		client := NewTransactionalClient(apiKey, "Test Sender", "test@example.com")
		client.EnableSandbox()

		content := Content{
			ToName:    "Test Recipient",
			ToEmail:   "recipient@example.com",
			Subject:   "Test Transactional Email",
			HTML:      renderer,
			PlainText: "Hello World",
		}

		err := client.Send(context.Background(), content)
		require.NoError(t, err)
		require.NotNil(t, renderer.gotWriter, "renderer should have received a writer")

		trackingSettings := NewTrackingSettings()
		trackingSettings.ClickTracking = mail.NewClickTrackingSetting().SetEnable(true)
		err = client.Send(context.Background(), content, WithTrackingSettings(trackingSettings))
		require.NoError(t, err)
		require.NotNil(t, renderer.gotWriter, "renderer should have received a writer")

		mailSettings := NewMailSettings()
		mailSettings.Footer = mail.NewFooterSetting().SetEnable(false)
		err = client.Send(context.Background(), content, WithMailSettings(mailSettings))
		require.NoError(t, err)
		require.NotNil(t, renderer.gotWriter, "renderer should have received a writer")

		// With nil html
		content.HTML = nil
		err = client.Send(context.Background(), content)
		require.NoError(t, err)
	})

	t.Run("renderer fails", func(t *testing.T) {
		client := NewTransactionalClient("fake-key", "Test Sender", "test@example.com")

		content := Content{
			ToName:    "Test Recipient",
			ToEmail:   "recipient@example.com",
			Subject:   "Test Transactional Email",
			HTML:      &mockRenderer{err: fmt.Errorf("render failed")},
			PlainText: "Hello World",
		}

		err := client.Send(context.Background(), content)
		require.Error(t, err)
		require.EqualError(t, err, "render html: render failed")
	})
}
