package protocol

import "fmt"

const (
	ChatMessage int = 1000
	UpdateOnlineUserListMessage
	NotifyUserOnlineMessage
	NotifyUserOfflineMessage
)

func FormatTextMessage(messageType int, text string) string {
	return fmt.Sprintf("%d#%s", messageType, text)
}
