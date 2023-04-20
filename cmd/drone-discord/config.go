package main

import (
    "github.com/urfave/cli/v2"

    "github.com/gstolarz/drone-discord/plugin"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
    return []cli.Flag{
        &cli.StringFlag{
            Name:        "discord.id",
            Usage:       "discord id",
            EnvVars:     []string{"PLUGIN_DISCORD_ID"},
            Destination: &settings.DiscordID,
        },
        &cli.StringFlag{
            Name:        "discord.token",
            Usage:       "discord token",
            EnvVars:     []string{"PLUGIN_DISCORD_TOKEN"},
            Destination: &settings.DiscordToken,
        },
        &cli.BoolFlag{
            Name:        "wait",
            Usage:       "wait",
            EnvVars:     []string{"PLUGIN_WAIT"},
            Destination: &settings.Wait,
            Value:       true,
        },
        &cli.StringFlag{
            Name:        "content",
            Usage:       "content",
            EnvVars:     []string{"PLUGIN_CONTENT"},
            Destination: &settings.Content,
        },
        &cli.StringFlag{
            Name:        "username",
            Usage:       "username",
            EnvVars:     []string{"PLUGIN_USERNAME"},
            Destination: &settings.Username,
        },
        &cli.StringFlag{
            Name:        "avatar.url",
            Usage:       "avatar url",
            EnvVars:     []string{"PLUGIN_AVATAR_URL"},
            Destination: &settings.AvatarURL,
        },
        &cli.StringFlag{
            Name:        "color",
            Usage:       "color",
            EnvVars:     []string{"PLUGIN_COLOR"},
            Destination: &settings.Color,
        },
        &cli.StringFlag{
            Name:        "title",
            Usage:       "title",
            EnvVars:     []string{"PLUGIN_TITLE"},
            Destination: &settings.Title,
        },
        &cli.StringFlag{
            Name:        "description",
            Usage:       "description",
            EnvVars:     []string{"PLUGIN_DESCRIPTION", "DRONE_COMMIT_MESSAGE"},
            Destination: &settings.Description,
        },
        &cli.StringFlag{
            Name:        "author",
            Usage:       "author",
            EnvVars:     []string{"PLUGIN_AUTHOR", "DRONE_COMMIT_AUTHOR_NAME", "DRONE_COMMIT_AUTHOR"},
            Destination: &settings.Author,
        },
        &cli.StringFlag{
            Name:        "author.icon.url",
            Usage:       "author icon url",
            EnvVars:     []string{"PLUGIN_AUTHOR_ICON_URL", "DRONE_COMMIT_AUTHOR_AVATAR"},
            Destination: &settings.AuthorIconURL,
        },
        &cli.StringFlag{
            Name:        "footer",
            Usage:       "footer",
            EnvVars:     []string{"PLUGIN_FOOTER"},
            Destination: &settings.Footer,
        },
        &cli.StringFlag{
            Name:        "footer.icon.url",
            Usage:       "footer icon url",
            EnvVars:     []string{"PLUGIN_FOOTER_ICON_URL"},
            Destination: &settings.FooterIconURL,
        },
        &cli.BoolFlag{
            Name:        "show.repo.name",
            Usage:       "show repo name",
            EnvVars:     []string{"PLUGIN_SHOW_REPO_NAME"},
            Destination: &settings.ShowRepoName,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "show.build.branch",
            Usage:       "show build branch",
            EnvVars:     []string{"PLUGIN_SHOW_BUILD_BRANCH"},
            Destination: &settings.ShowBuildBranch,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "show.build.number",
            Usage:       "show build number",
            EnvVars:     []string{"PLUGIN_SHOW_BUILD_NUMBER"},
            Destination: &settings.ShowBuildNumber,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "show.build.status",
            Usage:       "show build status",
            EnvVars:     []string{"PLUGIN_SHOW_BUILD_STATUS"},
            Destination: &settings.ShowBuildStatus,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "show.build.event",
            Usage:       "show build event",
            EnvVars:     []string{"PLUGIN_SHOW_BUILD_EVENT"},
            Destination: &settings.ShowBuildEvent,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "show.stage.name",
            Usage:       "show stage name",
            EnvVars:     []string{"PLUGIN_SHOW_STAGE_NAME"},
            Destination: &settings.ShowStageName,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "show.deploy.to",
            Usage:       "show deploy to",
            EnvVars:     []string{"PLUGIN_SHOW_DEPLOY_TO"},
            Destination: &settings.ShowDeployTo,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "use.color",
            Usage:       "use color",
            EnvVars:     []string{"PLUGIN_USE_COLOR"},
            Destination: &settings.UseColor,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "use.emoji",
            Usage:       "use emoji",
            EnvVars:     []string{"PLUGIN_USE_EMOJI"},
            Destination: &settings.UseEmoji,
            Value:       true,
        },
        &cli.BoolFlag{
            Name:        "use.timestamp",
            Usage:       "use timestamp",
            EnvVars:     []string{"PLUGIN_USE_TIMESTAMP"},
            Destination: &settings.UseTimestamp,
            Value:       true,
        },
    }
}
