package main

import (
	"log"
	"os"

	"github.com/mattn/go-isatty"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(">>> ")

	// isatty(3) は、指定したファイルディスクリプタが端末（tty）を参照しているかをチェックする関数.
	// そのディスクリプタが端末を指している場合は 1, それ以外は 0 を返す.
	// go-isattyパッケージは、Goにてtty判定を利用する場合に最も利用されているパッケージ.
	//
	// - ディスクリプタ自体が無効な場合（開かれていない/存在しない）は0を返し、errnoにEBADFが設定
	// - ディスクリプタが存在するが「端末ではない」（例えば通常のファイルやパイプなど）場合は0を返し、errnoにENOTTYが設定
	//
	// OSがLinuxの場合、内部で unix.IoctlGetTermios(int(fd), unix.TCGETS) が呼び出されている.
	// ioctl(2)にTCGETSを渡すと指定のディスクリプタが指す端末情報を取得できる.
	//
	// 	#include <sys/ioctl.h>
	// 	#include <termios.h>
	//
	// 	struct termios tty = {};
	// 	ioctl(1, TCGETS, &tty);
	//
	// ちなみに、isattyという名前は is-atty ではなく、is-a-tty という意味。
	//
	// REF: https://man7.org/linux/man-pages/man3/isatty.3.htm
	switch isatty.IsTerminal(os.Stdout.Fd()) {
	case true:
		// 端末に接続されている
		log.Println("Terminal")
	default:
		// 端末に接続されていない
		log.Println("Pipe or Redirect")
	}
}
