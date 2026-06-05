// Codes by vision
package banner

import (
	"fmt"
)

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
)

func Display() {
banner := `
   _____         _ __       __    ____             __      
  / ___/      __(_) /______/ /_  / __ \____  __  __/ /____ 
  \__ \ | /| / / / __/ ___/ __ \/ /_/ / __ \/ / / / __/ _ \
 ___/ / |/ |/ / / /_/ /__/ / / / _, _/ /_/ / /_/ / /_/  __/
/____/|__/|__/_/\__/\___/_/ /_/_/ |_|\____/\__,_/\__/\___/
`
	fmt.Printf("%s%s%s\n", ColorGreen, banner, ColorReset)
	fmt.Printf("%s                    IP Rotation Tool - Professional CLI Edition%s\n", ColorCyan, ColorReset)
	fmt.Printf("%s                    GitHub: https://github.com/vision-dev1%s\n\n", ColorBlue, ColorReset)
}

func PrintSuccess(message string) {
	fmt.Printf("%s✓ %s%s\n", ColorGreen, message, ColorReset)
}

func PrintError(message string) {
	fmt.Printf("%s✗ %s%s\n", ColorRed, message, ColorReset)
}

func PrintInfo(message string) {
	fmt.Printf("%s→ %s%s\n", ColorCyan, message, ColorReset)
}

func PrintWarning(message string) {
	fmt.Printf("%s⚠ %s%s\n", ColorYellow, message, ColorReset)
}

func PrintActiveIP(ip string) {
	fmt.Printf("%s🔄 Active IP: %s%s\n", ColorYellow, ip, ColorReset)
}
