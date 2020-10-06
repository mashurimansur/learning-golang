package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "Mashuri Mansur"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded :", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodedString = string(decodeByte)
	fmt.Println("decoded :", decodedString)

	enc, dec := decode1(data)
	fmt.Println("encoding :", enc)
	fmt.Println("decoding :", dec)

	var url = "https://kalipare.com/"
	urlenc, urldec := encodeDecodeURL(url)
	fmt.Println("encoding :", urlenc)
	fmt.Println("decoding :", urldec)

}

func decode1(data string) (string, string) {
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	return encodedString, decodedString
}

func encodeDecodeURL(data string) (string, string) {
	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	return encodedString, decodedString
}
