package protocol

import "fmt"

const (
	ChatMessage int = iota + 1000
	UpdateOnlineUserListMessage
	NotifyUserOnlineMessage
	NotifyUserOfflineMessage
)

func FormatTextMessage(username string, messageType int, text string) string {
	return fmt.Sprintf("%d#%s#%s", messageType, username, text)
}
