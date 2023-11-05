package providers

import (
	"fmt"

	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/slack-go/slack"
)

func SendSlackAlert(text string) {
	slackConfig, _ := config.GetConfig[config.SlackConfig]("Slack")

	api := slack.New(slackConfig.AccessToken)

	channelID, timestamp, err := api.PostMessage(
		slackConfig.Channel,
		slack.MsgOptionText(text, false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

/*

func SendSlackAlert(taggedSlackUsers []string, commentText string, base64Image string) error {
    slackConfig, _ := config.GetConfig[config.SlackConfig]("Slack")
	slackClient := slack.New(slackConfig.AccessToken)

	// For every tagged user, join the channel and send the message.
	for _, slackUser := range taggedSlackUsers {

		if slackUser.WebhookChannelID != nil {
			// The Slack API handles:
			// 1. Joining a channel the bot is already a member of
			// 2. Joining a Slack user
			// Because of this, we can skip checking for this in our application code.
			_, _, _, err := slackClient.JoinConversation(*slackUser.WebhookChannelID)
			if err != nil {
				log.Error(e.Wrap(err, "failed to join slack channel"))
			}

			_, _, err = slackClient.PostMessage(*slackUser.WebhookChannelID, slack.MsgOptionBlocks())
			if err != nil {
				return e.Wrap(err, "error posting slack message via slack bot")
			}
		}
	}

	// We need to write the base64 image as a png on disk to upload to Slack.
	// We create a unique file name for the image.
	uploadedFileKey := fmt.Sprintf("slack-image-%d.png", time.Now().UnixNano())

	dec, err := base64.StdEncoding.DecodeString(*base64Image)
	if err != nil {
		log.Error(e.Wrap(err, "Failed to decode base64 image"))
	}
	f, err := os.Create(uploadedFileKey)
	if err != nil {
		log.Error(e.Wrap(err, "Failed to create file on disk"))
	}
	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		log.Error(e.Wrap(err, "Failed to write file on disk"))
	}
	if err := f.Sync(); err != nil {
		log.Error("Failed to sync file on disk")
	}

	// We need to write the base64 image to disk, read the file, then upload it to Slack.
	// We can't send Slack a base64 string.
	fileUploadParams := slack.FileUploadParameters{
		Filetype: "image/png",
		Filename: "Upload.png",
		// These are the channels that will have access to the uploaded file.
		Channels: channels,
		File:     uploadedFileKey,
	}
	_, err = slackClient.UploadFile(fileUploadParams)

	if err != nil {
		log.Error(e.Wrap(err, "failed to upload file to Slack"))
	}

	if uploadedFileKey != "" {
		if err := os.Remove(uploadedFileKey); err != nil {
			log.Error(e.Wrap(err, "Failed to remove temporary session screenshot"))
		}
	}
}
*/
