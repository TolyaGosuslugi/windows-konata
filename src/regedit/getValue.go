package regedit

import (
	"fmt"

	"github.com/tolyagosuslugi/windows-konata/src/display"
	"golang.org/x/sys/windows/registry"
)

func GetValue(param string) string {
	k, err := registry.OpenKey(registry.CURRENT_USER, fmt.Sprintf(`AppEvents\Schemes\Apps\.Default\%s`, param), registry.READ)
	if err != nil {
		display.Err(err)
	}

	defer k.Close()
	val, _, err := k.GetStringValue("")
	if err != nil {
		display.Err(err)
	}

	return val
}
