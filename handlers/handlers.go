package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"github.com/rojserbest/GitHubFeedsBot/handlers/private"
)

func AddHandlers(dp *ext.Dispatcher) {
	private.AddHandlers(dp)
}
