package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/delta/codecharacter-lsp-2023/models"
)

func StreamReader(ws *models.WebsocketConnection) {
	if ws.LSPServer.Process == nil {
		return
	}
	for {
		var responseMessageBytes []byte
		reader := bufio.NewReader(ws.LSPServer.Stdout)
		contentLengthLine, err := reader.ReadBytes('\n')
		var contentLengthHeader string
		contentLengthHeader = string(contentLengthLine)
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			continue
		}
		const prefix = "Content-Length: "
		const suffix = "\r\n"
		contentLengthHeader = strings.TrimPrefix(contentLengthHeader, prefix)
		contentLengthHeader = strings.TrimSuffix(contentLengthHeader, suffix)
		contentLength, err := strconv.Atoi(contentLengthHeader)
		if err != nil {
			continue
		}
		_, err = reader.ReadBytes('\n')
		if err != nil {
			continue
		}
		for i := 0; i < contentLength; i++ {
			currbyte, err := reader.ReadByte()
			if err != nil {
				continue
			}
			responseMessageBytes = append(responseMessageBytes, currbyte)
		}
		var responseMessage map[string]interface{}
		err = json.Unmarshal(responseMessageBytes, &responseMessage)
		if err != nil {
			continue
		}
		_ = SendMessage(ws, responseMessage)
	}
}

func handleJSONRPC(ws *models.WebsocketConnection, requestMessageBytes []byte) error {
	contentLength := len(requestMessageBytes)
	header := fmt.Sprintf("Content-Length: %d\n\n", contentLength)
	requestMessage := header + string(requestMessageBytes)
	_, err := ws.LSPServer.Stdin.Write([]byte(requestMessage))
	if err != nil {
		return err
	}
	return nil
}
