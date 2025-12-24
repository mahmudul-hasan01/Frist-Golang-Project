package main

import (
	"back-end/cmd"
	"back-end/util"
	"fmt"
	// "encoding/base64"
)

func main() {
	cmd.Server()

	// var s string
	// s = "Hello world"

	//   byteAry := []byte(s)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64Str := enc.EncodeToString(byteAry)

	// fmt.Println(b64Str)

	// decStr, err := enc.DecodeString(b64Str)
	// if err != nil {
	// 	fmt.Println("Error decoding base64 string:", err)
	// 	return
	// }
	// fmt.Println(decStr)

	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         "1234567890",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		IsShopOwner: true,
		Role:        "admin",
	})
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return
	}
	fmt.Println(jwt)
}
