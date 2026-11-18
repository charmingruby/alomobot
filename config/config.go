package config

type Config struct {
	GeminiAPIKey    string `env:"GEMINI_API_KEY"`
	GeminiModel     string `env:"GEMINI_MODEL"`
	DiscordBotToken string `env:"DISCORD_BOT_TOKEN"`
	DiscordGuildID  string `env:"DISCORD_GUILD_ID"`
}
