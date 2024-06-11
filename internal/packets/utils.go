package packets

import "fmt"

func PrintFormattedHex(data []byte) {
	bytesPerLine := 16

	for i := 0; i < len(data); i += bytesPerLine {
		// Print the offset
		fmt.Printf("%04x  ", i)

		// Print hex bytes
		for j := 0; j < bytesPerLine; j++ {
			if i+j < len(data) {
				fmt.Printf("%02x ", data[i+j])
			} else {
				fmt.Print("   ")
			}
		}

		// Print ASCII characters
		fmt.Print(" |")
		for j := 0; j < bytesPerLine; j++ {
			if i+j < len(data) {
				b := data[i+j]
				if b >= 32 && b <= 126 {
					fmt.Printf("%c", b)
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
}
