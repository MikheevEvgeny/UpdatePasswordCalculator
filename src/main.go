package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"props"
	"strings"
)

func main() {
	propFilePath := "build.prop"
	if len(os.Args) >= 2 {
		propFilePath = os.Args[1]
	}

	buildPropsFile, err := os.Open(propFilePath)
	if err != nil {
		fmt.Println("Can't open file on path ", propFilePath)
	} else {
		buildProps := props.NewProperties()
		_ = buildProps.Load(buildPropsFile)
		_ = buildPropsFile.Close()

		if buildProps.Get("ro.hardware.static_password") == "true" {
			fmt.Println("+Ekfrl51Qkshsk#@zpdhkdWkd~-f")
		} else {
			password := GetSha512Hash(GetSha512Hash("ro.product.model=" + buildProps.Get("ro.product.model") +
				"ro.product.brand=" + buildProps.Get("ro.product.brand") +
				"ro.product.name=" + buildProps.Get("ro.product.name") +
				"ro.product.device=" + buildProps.Get("ro.product.device") +
				"ro.product.board=" + buildProps.Get("ro.product.board") +
				"ro.product.cpu.abi=" + buildProps.Get("ro.product.cpu.abi") +
				"ro.product.cpu.abi2=" + buildProps.Get("ro.product.cpu.abi2") +
				"ro.product.manufacturer=" + buildProps.Get("ro.product.manufacturer") +
				"ro.product.locale.language=" + buildProps.Get("ro.product.locale.language") +
				"ro.product.locale.region=" + buildProps.Get("ro.product.locale.region")))[10:38]

			fmt.Println(password)
		}

		//if runtime.GOOS == "windows" {
		//	C.system(C.CString("pause"))
		//} else {
		//	fmt.Println("Press the Enter Key to terminate the console screen!")
		//	_, _ = fmt.Scanln()
		//}
	}
	fmt.Println("Press the Enter Key to terminate the console screen!")
	_, _ = fmt.Scanln()
}

func GetSha512Hash(text string) string {
	hash := sha512.Sum512([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}
