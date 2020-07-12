package youtube

import (
	"fmt"
)

// Reload TinyProxy configuration
func Exec(isOn bool) error {
	fmt.Println("Turning Youtube access", isOn)
	return nil
}
