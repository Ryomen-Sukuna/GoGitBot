package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"FeedsBot/handlers/private"
)

func AddHandlers(dp *ext.Dispatcher) {
	private.AddHandlers(dp)
}
