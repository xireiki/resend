package email

import (
	"context"

	"github.com/xireiki/resend/types"
	"github.com/resend/resend-go/v2"
)

func SendMail(ctx context.Context, client *resend.Client, email *types.EmailEntity) (*resend.SendEmailResponse, error) {
	params, err := types.Email2Request(email)
	if err != nil {
		return nil, err
	}
	sent, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		return nil, err
	}
	return sent, nil
}

func SendMails(ctx context.Context, client *resend.Client, emails []*types.EmailEntity) (*resend.BatchEmailResponse, error) {
	params, err := types.Emails2Request(emails)
	if err != nil {
		return nil, err
	}
	sent, err := client.Batch.SendWithContext(ctx, params)
	if err != nil {
		return nil, err
	}
	return sent, nil
}

func RetrieveEmail(client *resend.Client, id string) (*resend.Email, error) {
	email, err := client.Emails.Get(id)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func UpdateEmail(client *resend.Client, id string, scheduled_at string) (*resend.UpdateEmailResponse, error) {
	updateParams := &resend.UpdateEmailRequest{
		Id:          id,
		ScheduledAt: scheduled_at,
	}
	updatedEmail, err := client.Emails.Update(updateParams)
	if err != nil {
		return nil, err
	}
	return updatedEmail, nil
}

func CancelEmail(client *resend.Client, id string) (*resend.CancelScheduledEmailResponse, error) {
	canceled, err := client.Emails.Cancel("49a3999c-0ce1-4ea6-ab68-afcd6dc2e794")
	if err != nil {
		return nil, err
	}
	return canceled, nil
}
