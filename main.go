package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/tolyagosuslugi/windows-konata/src/net"
	"github.com/tolyagosuslugi/windows-konata/src/tools"
)

func main() {
	fmt.Printf("%s\n\n", color.CyanString("Welcome to 'Windows Konata Edition' installer!"))

	// checking if os is windows
	fmt.Printf("Checking OS... ")
	if runtime.GOOS != "windows" {
		fmt.Printf("%s\n", color.RedString("This program is for Windows only. Exiting..."))
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", color.GreenString("Ok!"))
	}

	// creating folder for resources
	fmt.Println("Creating resources folder...")
	homeDir, _ := os.UserHomeDir()
	outFolder := homeDir + "/konata-sounds"
	os.Mkdir(outFolder, 0755)

	// downloading archive with sounds to folder
	fmt.Println("Downloading archive...")
	net.Download("https://github.com/tolyagosuslugi/windows-konata/raw/refs/heads/main/resources/sounds.tar.gz", outFolder)

	// unzipping downloaded archive
	fmt.Println("Extracting sounds from archive...")
	archiveFilePath := outFolder + "/sounds.tar.gz"

	archiveFile, err := os.Open(archiveFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	extractErr := tools.ExtractTarGz(archiveFile, outFolder)
	archiveFile.Close()
	if extractErr != nil {
		fmt.Println("Failed to decompress archive: " + extractErr.Error())
		archiveErr := os.Remove(archiveFilePath)
		if archiveErr != nil {
			fmt.Println(archiveErr)
		}
		return
	}
}
