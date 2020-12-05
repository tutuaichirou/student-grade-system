package util

import (
	"bufio"
	"os"
	"strings"
)

func KeyboardInput()string{
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	res = strings.TrimSpace(res)
	if strings.ToUpper(res)=="Q"{return "q"}
	return res
}
