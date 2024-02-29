package mail

import (
	"testing"

	"github.com/VatJittiprasert/goBanking/utils"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := utils.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EMAIL_SENDER_NAME, config.EMAIL_SENDER_ADDRESS, config.EMAIL_SENDER_PASSWORD)

	subject := "A test emil"
	content := `
  <h1>Hello world</h1>
  <p>This is a test message from <a href="http://techschool.guru"> Tech School<a/></p>
  `

	to := []string{config.EMAIL_TO_ADDRESS}
	attachFiles := []string{"../readme.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
