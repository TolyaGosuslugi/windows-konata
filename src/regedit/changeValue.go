package regedit

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tolyagosuslugi/windows-konata/src/display"
	"golang.org/x/sys/windows/registry"
)

func ChangeValue(param string, value string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, fmt.Sprintf(`AppEvents\Schemes\Apps\.Default\%s\.Current`, param), registry.WRITE)
	if err != nil {
		display.Err(err)
	}

	defer k.Close()
	err = k.SetStringValue("", value)
	if err != nil {
		display.Err(err)
	}

	fmt.Println(color.GreenString("Changed \"" + param + "\" successfully!"))
}
