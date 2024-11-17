package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tdlib "github.com/c0re100/gotdlib/client"
)

func (helper *Client) checkNewNessage() {
	fmt.Println("[Helper] New Message Receiver")

	receiver := helper.Client.AddEventReceiver(&tdlib.UpdateNewMessage{}, 1000)
	for newMsg := range receiver.Updates {
		go func(newMsg tdlib.Type) {
			updateMsg := newMsg.(*tdlib.UpdateNewMessage)
			chatId := updateMsg.Message.ChatId
			msgId := updateMsg.Message.Id
			senderId := GetSenderId(updateMsg.Message.SenderId)
			msgRtmId := GetReplyMessageId(updateMsg.Message.ReplyTo)

			if senderId != helper.ClientId {
				return
			}

			var msgText string
			var msgEnt []*tdlib.TextEntity

			switch updateMsg.Message.Content.MessageContentType() {
			case "messageText":
				msgText = updateMsg.Message.Content.(*tdlib.MessageText).Text.Text
				msgEnt = updateMsg.Message.Content.(*tdlib.MessageText).Text.Entities
			default:
				return
			}

			switch tdlib.CheckCommand(msgText, msgEnt) {
			case "/id":
				_, _ = helper.Client.EditMessageText(&tdlib.EditMessageTextRequest{
					ChatId:    chatId,
					MessageId: msgId,
					InputMessageContent: &tdlib.InputMessageText{
						Text: &tdlib.FormattedText{
							Text: fmt.Sprintf("Current Group ID: %v", chatId),
						},
					},
				})
			case "/pv":
				helper.setProfilePhoto(chatId, msgId, msgRtmId, msgText)
			}
		}(newMsg)
	}
}

func (helper *Client) setProfilePhoto(chatId, msgId, msgRtmId int64, text string) {
	if msgRtmId != 0 {
		var cId int64
		var ts float64
		var err error

		extract := strings.Split(tdlib.CommandArgument(text), " ")
		preText := text + "\n\nResult: "

		if len(extract) > 1 {
			cId, err = strconv.ParseInt(extract[0], 10, 64)
			if cId >= 0 {
				msgText := &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + "Chat ID should be negative."}}
				_, _ = helper.Client.EditMessageText(
					&tdlib.EditMessageTextRequest{
						ChatId:              chatId,
						MessageId:           msgId,
						InputMessageContent: msgText,
					},
				)
				return
			}
			ts, err = strconv.ParseFloat(extract[1], 64)
		} else {
			cId = chatId
			ts, err = strconv.ParseFloat(tdlib.CommandArgument(text), 64)
		}

		if err != nil {
			ts = 0
		}

		if ts < 0 {
			cId = int64(ts)
			ts = 0
		}

		rm, err := helper.Client.GetMessage(&tdlib.GetMessageRequest{
			ChatId:    chatId,
			MessageId: msgRtmId,
		})
		if err != nil {
			msgText := &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + err.Error()}}
			_, _ = helper.Client.EditMessageText(&tdlib.EditMessageTextRequest{
				ChatId:              chatId,
				MessageId:           msgId,
				InputMessageContent: msgText,
			})
			return
		}

		if rm.Content.MessageContentType() == tdlib.TypeMessageAnimation {
			var msgText *tdlib.InputMessageText
			if rm.Content.(*tdlib.MessageAnimation).Animation.Width == rm.Content.(*tdlib.MessageAnimation).Animation.Height &&
				rm.Content.(*tdlib.MessageAnimation).Animation.Width <= 1200 && rm.Content.(*tdlib.MessageAnimation).Animation.Height <= 1200 {
				f, dErr := helper.Client.DownloadFile(
					&tdlib.DownloadFileRequest{
						FileId:      rm.Content.(*tdlib.MessageAnimation).Animation.Animation.Id,
						Priority:    1,
						Offset:      0,
						Limit:       0,
						Synchronous: true,
					})
				if dErr != nil {
					msgText = &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + "Can't download this file - " + dErr.Error()}}
					_, _ = helper.Client.EditMessageText(&tdlib.EditMessageTextRequest{
						ChatId:              chatId,
						MessageId:           msgId,
						InputMessageContent: msgText,
					})
					return
				} else {
					// Workaround for abs path bug(?)
					_ = os.Rename(f.Local.Path, "pv.mp4")
					if len(extract) > 1 || cId < 0 {
						_, err = helper.Client.SetChatPhoto(&tdlib.SetChatPhotoRequest{
							ChatId: cId,
							Photo: &tdlib.InputChatPhotoAnimation{
								Animation: &tdlib.InputFileLocal{
									Path: "pv.mp4",
								},
								MainFrameTimestamp: ts,
							},
						})
					} else {
						_, err = helper.Client.SetProfilePhoto(&tdlib.SetProfilePhotoRequest{
							Photo: &tdlib.InputChatPhotoAnimation{
								Animation: &tdlib.InputFileLocal{
									Path: "pv.mp4",
								},
								MainFrameTimestamp: ts,
							},
						})
					}
					if err != nil {
						msgText = &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + "Can't change profile video - " + err.Error()}}
					} else {
						msgText = &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + "Profile video is changed."}}
					}
				}
			} else if rm.Content.(*tdlib.MessageAnimation).Animation.Duration > 10 {
				msgText = &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + "Duration too long. (Limit: <= 10 second)"}}
			} else if rm.Content.(*tdlib.MessageAnimation).Animation.Animation.Size > 2*1024*1024 {
				msgText = &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + "Filesize too big. (Limit: <= 2 MB)"}}
			} else {
				msgText = &tdlib.InputMessageText{Text: &tdlib.FormattedText{Text: preText + "Inconsistent size of animation.\nSquare video only (Limit: smaller or equal than 1200*1200))"}}
			}
			_, _ = helper.Client.EditMessageText(&tdlib.EditMessageTextRequest{
				ChatId:              chatId,
				MessageId:           msgId,
				InputMessageContent: msgText,
			})
		}
	}
}
