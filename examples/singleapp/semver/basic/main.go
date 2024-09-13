package main

import (
	"log"

	"github.com/Masterminds/semver/v3"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// semverパッケージは、Goでセマンティックバージョンを扱う機能を提供する。
	//
	// 基本的な機能は以下。
	//
	// - セマンティックバージョンの解析
	// - セマンティック・バージョンが一連の制約の範囲内に収まるかどうかをチェック
	// - セマンティックバージョンのソート
	// - v接頭辞についての挙動
	//
	var (
		version *semver.Version
		err     error
	)

	//
	// セマンティックバージョンの解析
	//
	if version, err = semver.NewVersion("v1.2.3"); err != nil {
		return err
	}
	log.Printf("%#v", version)

	//
	// セマンティック・バージョンが一連の制約の範囲内に収まるかどうかをチェック
	//
	var (
		c1 *semver.Constraints
		c2 *semver.Constraints
	)
	if c1, err = semver.NewConstraint(">= v1.1.0"); err != nil {
		return err
	}
	if c2, err = semver.NewConstraint("<= v1.1.0"); err != nil {
		return err
	}

	log.Println(c1, c1.Check(version))
	log.Println(c2, c2.Check(version))

	return nil
}
