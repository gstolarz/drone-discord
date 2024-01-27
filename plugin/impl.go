package plugin

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	build      = "build"
	buildTitle = "Build"

	deploy      = "deploy"
	deployTitle = "Deploy"
)

// Settings for the plugin.
type Settings struct {
	DiscordID    string
	DiscordToken string

	Wait bool

	Content   string
	Username  string
	AvatarURL string

	Color string

	Title       string
	Description string

	Author        string
	AuthorIconURL string

	Footer        string
	FooterIconURL string

	ShowRepoName    bool
	ShowBuildBranch bool
	ShowBuildNumber bool
	ShowBuildStatus bool
	ShowBuildEvent  bool
	ShowStageName   bool
	ShowDeployTo    bool

	UseColor     bool
	UseEmoji     bool
	UseTimestamp bool
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Validation of the settings.
	if len(p.settings.DiscordID) == 0 {
		return errors.New("no Discord ID provided")
	}

	if len(p.settings.DiscordToken) == 0 {
		return errors.New("no Discord Token provided")
	}

	if len(p.settings.Title) == 0 {
		if strings.Contains(strings.ToLower(p.pipeline.Stage.Name), strings.ToLower(build)) {
			p.settings.Title = buildTitle
		} else if strings.Contains(strings.ToLower(p.pipeline.Stage.Name), strings.ToLower(deploy)) {
			p.settings.Title = deployTitle
		} else {
			p.settings.Title = p.pipeline.Stage.Name
		}
	}

	if len(p.settings.Title) == 0 {
		return errors.New("no title provided")
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	embed := &discordgo.MessageEmbed{
		Title:       p.settings.Title,
		Description: p.settings.Description,
	}

	if p.settings.UseTimestamp {
		embed.Timestamp = time.Now().Format(time.RFC3339)
	}
	if p.settings.UseColor {
		embed.Color = p.color()
	}
	if len(p.settings.Author) > 0 || len(p.settings.AuthorIconURL) > 0 {
		embed.Author = &discordgo.MessageEmbedAuthor{
			Name:    p.settings.Author,
			IconURL: p.settings.AuthorIconURL,
		}
	}
	if len(p.settings.Footer) > 0 || len(p.settings.FooterIconURL) > 0 {
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text:    p.settings.Footer,
			IconURL: p.settings.FooterIconURL,
		}
	}

	var fields []*discordgo.MessageEmbedField

	if p.settings.ShowRepoName && len(p.pipeline.Repo.Name) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Repo name",
			Value:  "`" + p.pipeline.Repo.Name + "`",
			Inline: true,
		})
	}
	if p.settings.ShowBuildBranch && len(p.pipeline.Build.Branch) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Build branch",
			Value:  "`" + p.pipeline.Build.Branch + "`",
			Inline: true,
		})
	}
	if p.settings.ShowBuildNumber && p.pipeline.Build.Number > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Build number",
			Value:  "`" + strconv.Itoa(p.pipeline.Build.Number) + "`",
			Inline: true,
		})
	}
	if p.settings.ShowBuildStatus && len(p.pipeline.Build.Status) > 0 {
		emoji := p.emoji()
		if len(emoji) > 0 {
			emoji += " "
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Build status",
			Value:  emoji + "`" + p.pipeline.Build.Status + "`",
			Inline: true,
		})
	}
	if p.settings.ShowBuildEvent && len(p.pipeline.Build.Event) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Build event",
			Value:  "`" + p.pipeline.Build.Event + "`",
			Inline: true,
		})
	}
	if p.settings.ShowStageName && len(p.pipeline.Stage.Name) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Stage name",
			Value:  "`" + p.pipeline.Stage.Name + "`",
			Inline: true,
		})
	}
	if p.settings.ShowDeployTo && len(p.pipeline.Build.DeployTo) > 0 {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Deploy to",
			Value:  "`" + p.pipeline.Build.DeployTo + "`",
			Inline: true,
		})
	}

	if len(fields) > 0 {
		embed.Fields = fields
	}

	params := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{embed},
	}

	if len(p.settings.Content) > 0 {
		params.Content = p.settings.Content
	}
	if len(p.settings.Username) > 0 {
		params.Username = p.settings.Username
	}
	if len(p.settings.AvatarURL) > 0 {
		params.AvatarURL = p.settings.AvatarURL
	}

	session, err := discordgo.New("")
	if err != nil {
		return fmt.Errorf("error while creating Discord session: %w", err)
	}
	defer session.Close()

	_, err = session.WebhookExecute(p.settings.DiscordID, p.settings.DiscordToken, p.settings.Wait, params)
	if err != nil {
		return fmt.Errorf("error while executing Discord webhook: %w", err)
	}

	return nil
}

func (p *Plugin) emoji() string {
	var emoji string

	if p.settings.UseEmoji && len(p.settings.Color) == 0 {
		switch p.pipeline.Build.Status {
		case "success":
			emoji = ":green_square:"
		case "failure", "error", "killed":
			emoji = ":red_square:"
		default:
			emoji = ":yellow_square:"
		}
	}

	return emoji
}

func (p *Plugin) color() int {
	if len(p.settings.Color) > 0 {
		p.settings.Color = strings.Replace(p.settings.Color, "#", "", -1)
		if s, err := strconv.ParseInt(p.settings.Color, 16, 32); err == nil {
			return int(s)
		}
	}

	switch p.pipeline.Build.Status {
	case "success":
		// green
		return 0x1ac600
	case "failure", "error", "killed":
		// red
		return 0xff3232
	default:
		// yellow
		return 0xffd930
	}
}
