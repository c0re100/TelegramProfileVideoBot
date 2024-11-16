package main

import tdlib "github.com/c0re100/gotdlib/client"

func GetSenderId(sender tdlib.MessageSender) int64 {
	if sender.MessageSenderType() == tdlib.TypeMessageSenderUser {
		return sender.(*tdlib.MessageSenderUser).UserId
	} else {
		return sender.(*tdlib.MessageSenderChat).ChatId
	}
}

func GetReplyMessageId(replyTo tdlib.MessageReplyTo) int64 {
	if replyTo == nil {
		return 0
	} else if replyTo.MessageReplyToType() == tdlib.TypeMessageReplyToMessage {
		return replyTo.(*tdlib.MessageReplyToMessage).MessageId
	}
	return 0
}

func CheckUsernameEmpty(usernames *tdlib.Usernames) string {
	if usernames != nil {
		return "@" + usernames.ActiveUsernames[0]
	} else {
		return "@None"
	}
}
