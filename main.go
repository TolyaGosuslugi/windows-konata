package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/tolyagosuslugi/windows-konata/src/display"
	"github.com/tolyagosuslugi/windows-konata/src/net"
	"github.com/tolyagosuslugi/windows-konata/src/regedit"
	"github.com/tolyagosuslugi/windows-konata/src/tools"
)

func main() {
	fmt.Printf("%s\n\n", color.CyanString("Welcome to 'Windows Konata Edition' installer :3"))

	// checking if os is windows
	fmt.Printf("Checking OS... ")
	if runtime.GOOS != "windows" {
		display.Err(errors.New("This program is for Windows only. Exiting..."))
	} else {
		fmt.Printf("%s\n", color.GreenString("Ok!"))
	}

	homeDir, _ := os.UserHomeDir()
	outFolder := homeDir + "/konata-sounds"

	if _, err := os.Stat(outFolder); !os.IsNotExist(err) {
		errRm := os.RemoveAll(outFolder)
		if errRm != nil {
			panic(errRm)
		}
	}

	// creating folder for resources
	fmt.Printf("Creating resources folder... ")
	err := os.Mkdir(outFolder, 0755)
	if err != nil {
		display.Err(err)
	} else {
		fmt.Printf("%s\n", color.GreenString("Ok!"))
	}

	// downloading archive with sounds to folder
	fmt.Printf("Downloading archive... ")
	err = net.Download("https://github.com/tolyagosuslugi/windows-konata/raw/refs/heads/main/resources/sounds.tar.gz", outFolder+"/sounds.tar.gz")
	if err != nil {
		display.Err(err)
	} else {
		fmt.Printf("%s\n", color.GreenString("Ok!"))
	}

	// unzipping downloaded archive
	fmt.Printf("Extracting sounds from archive... ")
	archiveFilePath := outFolder + "/sounds.tar.gz"

	archiveFile, err := os.Open(archiveFilePath)
	if err != nil {
		display.Err(err)
	}

	err = tools.ExtractTarGz(archiveFile, outFolder+"/")
	archiveFile.Close()
	archiveErr := os.Remove(archiveFilePath)
	if archiveErr != nil {
		display.Err(err)
	}
	if err != nil {
		display.Err(err)
	} else {
		fmt.Printf("%s\n", color.GreenString("Ok!"))
	}

	//setting sounds in regedit
	fmt.Printf("Changing sounds in regedit...\n")
	regedit.ChangeValue(".Default", outFolder+"/DefaultBeep.wav")
	regedit.ChangeValue("SystemExclamation", outFolder+"/Exclamation.wav")
	regedit.ChangeValue("SystemHand", outFolder+"/CriticalStop.wav")
	//regedit.ChangeValue("WindowsLogoff", outFolder+"/Logoff.wav")
	//regedit.ChangeValue("SystemExit", outFolder+"/Logoff.wav")
	//regedit.ChangeValue("WindowsLogon", outFolder+"/Logon.wav")
	regedit.ChangeValue("WindowsUAC", outFolder+"/UserAccControl.wav")
}
