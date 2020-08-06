package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/c0re100/go-tdlib"
)

func (tg *Client) checkNewNessage() {
    fmt.Println("[Helper] New Message Receiver")
    eventFilter := func(msg *tdlib.TdMessage) bool {
        updateMsg := (*msg).(*tdlib.UpdateNewMessage)
        if updateMsg.Message.SenderUserId == tg.clientId { // Prevent users abuse your userbot :)
            return true
        }
        return false
    }

    receiver := tg.client.AddEventReceiver(&tdlib.UpdateNewMessage{}, eventFilter, 1000)
    for newMsg := range receiver.Chan {
        go func(newMsg tdlib.TdMessage) {
            updateMsg := (newMsg).(*tdlib.UpdateNewMessage)
            chatId := updateMsg.Message.ChatId
            msgId := updateMsg.Message.Id
            msgRtmId := updateMsg.Message.ReplyToMessageId

            var msgText string
            var msgEnt []tdlib.TextEntity

            switch updateMsg.Message.Content.GetMessageContentEnum() {
            case "messageText":
                msgText = updateMsg.Message.Content.(*tdlib.MessageText).Text.Text
                msgEnt = updateMsg.Message.Content.(*tdlib.MessageText).Text.Entities
            default:
                return
            }

            switch CheckCommand(msgText, msgEnt) {
            case "/id":
                msgText := tdlib.NewInputMessageText(tdlib.NewFormattedText(fmt.Sprintf("Current Group ID: %v", chatId), nil), true, false)
                tg.client.EditMessageText(chatId, msgId, nil, msgText)
            case "/pv":
                tg.setProfilePhoto(chatId, msgId, msgRtmId, msgText)
            }
        }(newMsg)
    }
}

func (tg *Client) setProfilePhoto(chatId, msgId, msgRtmId int64, msgText string) {
    if msgRtmId != 0 {
        var cId int64
        var ts float64
        var err error

        extract := strings.Split(CommandArgument(msgText), " ")
        preText := msgText + "\n\nResult: "

        if len(extract) > 1 {
            cId, err = strconv.ParseInt(extract[0], 10, 64)
            if cId >= 0 {
                msgText := tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+"Chat ID should be negative.", nil), true, false)
                tg.client.EditMessageText(chatId, msgId, nil, msgText)
                return
            }
            ts, err = strconv.ParseFloat(extract[1], 64)
        } else {
            cId = chatId
            ts, err = strconv.ParseFloat(CommandArgument(msgText), 64)
        }

        if err != nil {
            ts = 0
        }

        if ts < 0 {
            cId = int64(ts)
            ts = 0
        }

        rm, err := tg.client.GetMessage(chatId, msgRtmId)
        if err != nil {
            msgText := tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+err.Error(), nil), true, false)
            tg.client.EditMessageText(chatId, msgId, nil, msgText)
            return
        }

        if rm.Content.GetMessageContentEnum() == "messageAnimation" {
            var msgText *tdlib.InputMessageText
            if rm.Content.(*tdlib.MessageAnimation).Animation.Width == rm.Content.(*tdlib.MessageAnimation).Animation.Height &&
                rm.Content.(*tdlib.MessageAnimation).Animation.Width <= 800 && rm.Content.(*tdlib.MessageAnimation).Animation.Height <= 800 {
                f, err := tg.client.DownloadFile(rm.Content.(*tdlib.MessageAnimation).Animation.Animation.Id, 1, 0, 0, true)
                if err != nil {
                    msgText = tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+"Can't download this file - "+err.Error(), nil), true, false)
                    return
                } else {
                    // Workaround for abs path bug(?)
                    os.Rename(f.Local.Path, "pv.mp4")
                    if len(extract) > 1 || cId < 0 {
                        _, err = tg.client.SetChatPhoto(cId, tdlib.NewInputChatPhotoAnimation(tdlib.NewInputFileLocal("pv.mp4"), ts))
                    } else {
                        _, err = tg.client.SetProfilePhoto(tdlib.NewInputChatPhotoAnimation(tdlib.NewInputFileLocal("pv.mp4"), ts))
                    }
                    if err != nil {
                        msgText = tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+"Can't change profile video - "+err.Error(), nil), true, false)
                    } else {
                        msgText = tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+"Profile video is changed.", nil), true, false)
                    }
                }
            } else if rm.Content.(*tdlib.MessageAnimation).Animation.Duration > 10 {
                msgText = tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+"Duration too long. (Limit: <= 10 second)", nil), true, false)
            } else if rm.Content.(*tdlib.MessageAnimation).Animation.Animation.Size > 2*1024*1024 {
                msgText = tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+"Filesize too big. (Limit: <= 2 MB)", nil), true, false)
            } else {
                msgText = tdlib.NewInputMessageText(tdlib.NewFormattedText(preText+"Inconsistent size of animation.\nSquare video only (Limit: smaller or equal than 800*800))", nil), true, false)
            }
            tg.client.EditMessageText(chatId, msgId, nil, msgText)
        }
    }
}
