package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

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

	fmt.Printf("Enter number:\n%s\n%s\n:", color.CyanString("0. Install Konata's sounds"), color.CyanString("1. Return sounds to Windows Default"))
	num, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	num = strings.TrimSpace(num)

	fmt.Printf("\n")

	switch num {
	case "0":
		//////////////////////////  Install Konata's sounds  //////////////////////////
		homeDir, _ := os.UserHomeDir()
		outFolder := homeDir + "/konata-sounds"

		// deleting folder with previous sounds
		if _, err := os.Stat(outFolder); !os.IsNotExist(err) {
			fmt.Printf("Deleting old folder... ")
			errRm := os.RemoveAll(outFolder)
			if errRm != nil {
				display.Err(errRm)
			}
			fmt.Printf("%s\n", color.GreenString("Ok!"))
		}

		// creating new folder for resources
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
		regedit.ChangeValue(".Default\\.Current", outFolder+"/DefaultBeep.wav")
		regedit.ChangeValue("SystemExclamation\\.Current", outFolder+"/Exclamation.wav")
		regedit.ChangeValue("SystemHand\\.Current", outFolder+"/CriticalStop.wav")
		//regedit.ChangeValue("WindowsLogoff\\.Current", outFolder+"/Logoff.wav")
		//regedit.ChangeValue("SystemExit\\.Current", outFolder+"/Logoff.wav")
		//regedit.ChangeValue("WindowsLogon\\.Current", outFolder+"/Logon.wav")
		regedit.ChangeValue("WindowsUAC\\.Current", outFolder+"/UserAccControl.wav")

	case "1":
		//////////////////////////  Return sounds to Windows Default  //////////////////////////
		fmt.Printf("Returning sounds to Windows Default...\n")
		regedit.ChangeValue(".Default\\.Current", regedit.GetValue(".Default\\.Default"))
		regedit.ChangeValue("SystemExclamation\\.Current", regedit.GetValue("SystemExclamation\\.Default"))
		regedit.ChangeValue("SystemHand\\.Current", regedit.GetValue("SystemHand\\.Default"))
		//regedit.ChangeValue("WindowsLogoff\\.Current", regedit.GetValue("WindowsLogoff\\.Default"))
		//regedit.ChangeValue("SystemExit\\.Current", regedit.GetValue("SystemExit\\.Default"))
		//regedit.ChangeValue("WindowsLogon\\.Current", regedit.GetValue("WindowsLogon\\.Default"))
		regedit.ChangeValue("WindowsUAC\\.Current", regedit.GetValue("WindowsUAC\\.Default"))
	default:
		fmt.Printf("Please, enter 0 or 1!")
		os.Exit(0)
	}
}
