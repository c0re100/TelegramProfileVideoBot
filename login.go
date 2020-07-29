package main

import (
    "fmt"
    "syscall"

    "github.com/c0re100/go-tdlib"
    "golang.org/x/crypto/ssh/terminal"
)

func (tg *Client) login() {
    for {
        currentState, _ := tg.client.Authorize()
        if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
            fmt.Print("Enter phone: ")
            var number string
            fmt.Scanln(&number)
            _, err := tg.client.SendPhoneNumber(number)
            if err != nil {
                fmt.Printf("Error sending phone number: %v", err)
            }
        } else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
            fmt.Print("Enter code: ")
            var code string
            fmt.Scanln(&code)
            _, err := tg.client.SendAuthCode(code)
            if err != nil {
                fmt.Printf("Error sending auth code : %v", err)
            }
        } else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPasswordType {
            fmt.Print("Enter Password: ")
            bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
            if err != nil {
                fmt.Println(err)
            }
            _, err = tg.client.SendAuthPassword(string(bytePassword))
            if err != nil {
                fmt.Printf("Error sending auth password: %v", err)
            }
        } else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
            me, err := tg.client.GetMe()
            if err != nil {
                fmt.Println(err)
                return
            }
            tg.clientId = me.Id
            fmt.Println("Hello!", me.FirstName, me.LastName, "("+me.Username+")")
            break
        }
    }
}
