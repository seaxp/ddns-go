package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenExplorer 打开本地浏览器
func OpenExplorer(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	default: // Linux
		cmd = exec.Command("xdg-open", url)
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("请手动打开浏览器并访问 %s 进行配置\n", url)
	} else {
		fmt.Println("成功打开浏览器, 请在网页中进行配置")
	}
}
