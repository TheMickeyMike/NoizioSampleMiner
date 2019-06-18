package pkg

import (
	"fmt"
	"os"

	"howett.net/plist"

	log "github.com/sirupsen/logrus"
)

type InfoPlist struct {
	BundleShortVersion string `plist:"CFBundleShortVersionString"`
}

func Load() {
	infoPlistFile, err := os.Open("/Applications/Noizio.app/Contents/Info.plist")
	if err != nil {
		log.Fatalln(err)
	}
	defer infoPlistFile.Close()
	log.Infoln("Successfully Opened Info.plist")

	var data InfoPlist
	decoder := plist.NewDecoder(infoPlistFile)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("%+v", data)
}
