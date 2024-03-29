package xdgspec

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/adrg/xdg"
	"github.com/devlights/gomy/output"
)

// FileOp は、[xdg](https://github.com/adrg/xdg) を利用して
// XDGの規定に従った場所にファイルを配置したり検索したりしています.
func FileOp() error {
	var (
		dataDir = filepath.Join(xdg.DataHome, "try-golang")
	)

	output.Stdoutl("[OS]", runtime.GOOS)
	output.Stdoutl("[dataDir]", dataDir)

	// 処理する前に既にファイルが存在してたら消す
	if _, err := os.Stat(dataDir); err == nil {
		output.Stdoutl("[exists?]", "存在する --> 削除")

		if err = os.RemoveAll(dataDir); err != nil {
			output.Stderrl("[os.Remove]", err)
			return err
		}
	} else {
		output.Stdoutl("[exists?]", "存在しない")
	}

	// xdg.DataFile() に アプリ名/ファイル名 で渡すとXDGの規定に従ったパスを
	// 生成して返してくれる. xdg.DataFile() にすると $XDG_DATA_HOME
	// xdg.ConfigFile() にすると $XDG_CONFIG_HOME がベースとなる
	//
	// このとき、ファイルの親ディレクトリ（つまりアプリ名の部分）が
	// 存在しない場合は、ディレクトリを作成してくれる
	dataFile, err := xdg.DataFile("try-golang/mydata.txt")
	if err != nil {
		output.Stderrl("[xdg.DataFile]", err)
		return err
	}

	output.Stdoutl("[xdg.DataFile]", dataFile)

	if _, err = os.Stat(dataDir); err == nil {
		output.Stdoutl("[exists?]", "存在する")
	} else {
		output.Stdoutl("[exists?]", "存在しない")
	}

	if err = os.WriteFile(dataFile, []byte("helloworld\n"), 0644); err != nil {
		output.Stderrl("[ioutil.WriteFile]", err)
		return err
	}

	if bytes, readErr := os.ReadFile(dataFile); readErr == nil {
		output.Stdoutl("[ioutil.ReadFile]", string(bytes))
	}

	// xdg.SearchDataFile() を利用すると、指定した アプリ名/ファイル名 を探してくれる.
	// 存在しない場合は、err に値が入る.
	// xdg.DataFile() と違い、こちらは親ディレクトリを作ったりはしてくれない.
	// 既にファイルが存在する場合に利用する.
	dataFile2, searchErr := xdg.SearchDataFile("try-golang/mydata.txt")
	if searchErr != nil {
		output.Stderrl("[xdg.SearchDataFile]", searchErr)
		return searchErr
	}

	output.Stdoutl("[xdg.SearchDataFile]", dataFile2)

	// 後始末
	_ = os.RemoveAll(dataDir)

	return nil
}
