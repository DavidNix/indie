package email

import (
	"context"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// TransactionalClient handles transactional emails like account verification, password reset, etc.
type TransactionalClient struct {
	baseClient
}

// NewTransactionalClient creates a new client for sending transactional emails
func NewTransactionalClient(apiKey string, fromName, fromAddr string) *TransactionalClient {
	return &TransactionalClient{
		baseClient: baseClient{
			client:   sendgrid.NewSendClient(apiKey),
			fromName: fromName,
			fromAddr: fromAddr,
		},
	}
}

// Send sends a transactional email (like account verification, password reset, etc.)
// These emails don't have unsubscribe links.
func (c *TransactionalClient) Send(ctx context.Context, content Content, opts ...Option) error {
	from := mail.NewEmail(c.fromName, c.fromAddr)
	toEmail := mail.NewEmail(content.ToName, content.ToEmail)

	html, err := c.renderHTML(ctx, content.HTML)
	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(from, content.Subject, toEmail, content.PlainText, html)

	var opt option
	for _, apply := range opts {
		apply(&opt)
	}
	message.TrackingSettings = c.buildTrackingSettings(opt)
	message.MailSettings = c.buildMailSettings(opt)

	response, err := c.client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("bad status=%d body=%s", response.StatusCode, response.Body)
	}

	return nil
}
