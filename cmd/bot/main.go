package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmingruby/erebor/discord"
	"github.com/charmingruby/erebor/discord/command"
	"github.com/charmingruby/erebor/discord/mention"
	"github.com/charmingruby/erebor/env"
	"github.com/charmingruby/erebor/o11y/logging"
	"github.com/charmingruby/slowbot/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log := logging.InitLogger()
	ctx = logging.WithLogger(ctx, log)

	if err := run(ctx); err != nil {
		log.ErrorContext(ctx, "slowbot discord bot failed", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := env.Load[config.Config]()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	bot, err := discord.NewBot(
		cfg.DiscordBotToken,
		discord.WithErrorMessage("Slowbot está na má vontade, rs..."),
		discord.WithMentionReplyHandler(func(ctx context.Context, m mention.Message) (string, error) {
			return "kkkk", nil
		}),
	)
	if err != nil {
		return fmt.Errorf("create discord bot: %w", err)
	}

	if err := bot.RegisterCommand(command.Definition{
		Name:        "flip",
		Description: "Executa o flip do Slowbot.",
		Handler: func(ctx context.Context, i command.Interaction) (string, error) {
			for _, opt := range i.Options {
				if opt.Name == "type" {
					return opt.StringValue(), nil
				}
			}

			return "Normal", nil
		},
		Parameters: []*command.Parameter{
			command.String(command.Parameter{
				Name:        "type",
				Description: "Tipo de flip.",
				Required:    true,
			}),
		},
	}); err != nil {
		return fmt.Errorf("register flip command: %w", err)
	}

	return bot.Listen(ctx)
}
