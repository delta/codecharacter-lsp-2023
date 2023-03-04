package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/delta/codecharacter-lsp-2023/models"
)

// func ended(){
// 	fmt.Println("ended")
// }

func SendMessageFunc(ws *models.WebsocketConnection) {
	// defer ended()
	fmt.Println("starting routine")
	for {
		var responseMessageBytes []byte
		reader := bufio.NewReader(ws.LSPServer.Stdout)
		// byttttteeees,_ := reader.Peek(100)
		// fmt.Println(string(byttttteeees))
		line1, err := reader.ReadBytes('\n')
		// fmt.Println("READ EROR", err)
		if err != nil && err.Error() == "EOF" {
			fmt.Println(err)
			return
		}
		line1 = line1[16:]
		line1 = line1[:len(line1)-2]
		length, err := strconv.Atoi(string(line1))
		fmt.Println(length)
		_, err = reader.ReadBytes('\n')
		// fmt.Println(line2)
		for i := 0; i < length; i++ {
			currbyte, err := reader.ReadByte()
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println(err)
			responseMessageBytes = append(responseMessageBytes, currbyte)
		}
		// TODO: Read Content Length
		// for i := 0; i < 20; i++ {
		// 	gg,err := reader.ReadString('}')
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return err
		// 	}
		// 	fmt.Println(gg)
		// 	// fmt.Println("JSONRPC Response : ", len(responseMessageBytes), " with ID : ", ws.ID)
		// }
		// fmt.Println(respone)
		// fmt.Println("came here")
		// fmt.Println(string(responseMessageBytes))
		var responseMessage map[string]interface{}
		err = json.Unmarshal(responseMessageBytes, &responseMessage)
		// if err != nil {
		// 	fmt.Println("Error in unmarshalling JSONRPC response : ", err)
		// 	return err
		// }
		SendMessage(ws, responseMessage)
	}
}

func handleJSONRPC(ws *models.WebsocketConnection, requestMessageBytes []byte) error {
	u := fmt.Sprintf("Content-Length: %d", len(requestMessageBytes))
	result := u + "\n\n" + string(requestMessageBytes)
	// fmt.Println(result)
	_, err := ws.LSPServer.Stdin.Write([]byte(result))
	if err != nil {
		fmt.Println("error here", err)
		// fmt.Println(err)
		// return err
	}

	// buff := bufio.NewScanner(ws.LSPServer.Stdout)
	// // var allText []string

	// for buff.Scan() {
	// 	fmt.Println([]byte(buff.Text()))
	// }
	return nil
}
