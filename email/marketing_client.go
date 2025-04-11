package email

import (
	"context"
	"fmt"
	"slices"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// MarketingClient handles marketing emails with unsubscribe functionality
type MarketingClient struct {
	*baseClient
	unsubGroups []int
}

// NewMarketingClient creates a new client for sending compliant marketing emails.
// apiKey is the Sendgrid API key.
// fromName and fromAddr are the display name and email address that will appear in the From field.
// unsubGroups is a list of Sendgrid unsubscribe group IDs that this client is allowed to use.
// compliancePlainTxt is the plain text compliance footer that will be appended to all marketing emails.
// renderer is used to wrap email content in a standard layout with headers, footers, etc.
func NewMarketingClient(apiKey string, fromName, fromAddr string, unsubGroups []int) *MarketingClient {
	return &MarketingClient{
		baseClient: &baseClient{
			client:   sendgrid.NewSendClient(apiKey),
			fromName: fromName,
			fromAddr: fromAddr,
		},
		unsubGroups: unsubGroups,
	}
}

// Send sends a marketing email with unsubscribe links and footer
func (c *MarketingClient) Send(ctx context.Context, content Content, unsubGroup int, opts ...Option) error {
	if unsubGroup < 0 {
		return fmt.Errorf("missing unsubscribe group")
	}
	if !slices.Contains(c.unsubGroups, unsubGroup) {
		return fmt.Errorf("unsubscribe group %d not in allowed groups %v", unsubGroup, c.unsubGroups)
	}

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

	message.SetASM(&mail.Asm{
		GroupID:         unsubGroup,
		GroupsToDisplay: c.unsubGroups,
	})

	response, err := c.client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("bad status=%d body=%s", response.StatusCode, response.Body)
	}

	return nil
}
