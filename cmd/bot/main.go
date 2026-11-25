package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmingruby/erebor/discord"
	"github.com/charmingruby/erebor/env"
	"github.com/charmingruby/erebor/o11y/logging"

	"github.com/charmingruby/alomobot/config"
	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/alomobot/internal/cmd"
	"github.com/charmingruby/alomobot/internal/core/pkmn/pokeapi"
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

	app := app.New(pokeapi.New())

	bot, err := discord.New(
		cfg.DiscordBotToken,
		discord.WithCommands(cmd.Commands(app)...),
		discord.WithCommandGuild(cfg.DiscordGuildID),
		discord.WithErrorMessage("Blublublub..."),
	)
	if err != nil {
		return fmt.Errorf("create discord bot: %w", err)
	}

	return bot.Listen(ctx)
}
