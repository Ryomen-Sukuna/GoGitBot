package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"github.com/Ryomen-Sukuna/GoGitBot/handlers/private"
)

func AddHandlers(dp *ext.Dispatcher) {
	private.AddHandlers(dp)
}
