package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmingruby/alomobot/config"
	"github.com/charmingruby/alomobot/internal/cmd"
	"github.com/charmingruby/erebor/discord"
	"github.com/charmingruby/erebor/env"
	"github.com/charmingruby/erebor/o11y/logging"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log := logging.InitLogger()
	ctx = logging.WithLogger(ctx, log)

	if err := run(ctx); err != nil {
		log.ErrorContext(ctx, "alomobot discord bot failed", "error", err)
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
		discord.WithCommandGuild(cfg.DiscordGuildID),
		discord.WithErrorMessage("Blublublub..."),
	)
	if err != nil {
		return fmt.Errorf("create discord bot: %w", err)
	}

	if err := cmd.Register(bot); err != nil {
		return fmt.Errorf("failed to register misc commands: %w", err)
	}

	return bot.Listen(ctx)
}
