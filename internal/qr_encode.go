package internal

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

type IkarusQRCode struct {
	// QRCode is the QR code string
	Content string `json:"content"`
	Size    int    `json:"size"`
}

func (code *IkarusQRCode) Generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}
