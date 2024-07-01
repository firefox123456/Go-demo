package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decodeString, _ := base64.StdEncoding.DecodeString("eyJNZXRhZGF0YSI6eyJSZWZlcmVuY2VUYWJsZXMiOltdLCJWYXJpYWJsZXMiOltdfX0=")
	fmt.Println(decodeString)
}
