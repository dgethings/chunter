package cisco_ios

import (
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
)

var keywords = keyword.NewSet(map[string]keyword.Keyword{
	"hostname": {
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "To specify or modify the hostname for the network server, use the hostname command in global configuration mode.",
		},
		Section: "global",
	},
})
