package telegram_controllers

type TelegramCommand int

const (
	START TelegramCommand = iota
	HELP
	RECOMMEND
)

var Commands = map[TelegramCommand]string{
	START:     "/start",
	HELP:      "/help",
	RECOMMEND: "/recommend",
}

func (tc TelegramCommand) String() string {
	return Commands[tc]
}
