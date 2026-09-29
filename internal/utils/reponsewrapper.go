package utils

import "github.com/Teepee/internal/global"

type reponsepayload struct {
	Statuscode   uint   `json:"status"`
	MessageError string `json:"message_error"`
	Data         any    `json:"data"`
}

func WrapPayload(payload any, status uint, err error) reponsepayload {
	var messageErr string
	if err != nil {
		messageErr = err.Error()
	} else {
		messageErr = ""
	}
	payloadwrapped := reponsepayload{
		Statuscode:   status,
		MessageError: messageErr,
		Data:         payload,
	}
	return payloadwrapped
}

func CheckAvailablity(s string) bool {
	if s == global.AvailableAvailable {
		return true
	}
	if s == global.AvailableNot_available {
		return true
	}
	if s == global.AvailableOut_of_stock {
		return true
	}
	return false
}
