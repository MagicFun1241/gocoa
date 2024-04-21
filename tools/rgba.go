package tools

import "fmt"

func ParseHexRGBA(hexRGBA string) (r int, g int, b int, a int) {
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	return r, g, b, a
}
