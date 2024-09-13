package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"

	"github.com/Masterminds/semver/v3"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		version *semver.Version
		err     error
	)

	if version, err = semver.NewVersion("v1.2.3-beta-10+build456"); err != nil {
		return err
	}
	log.Printf("%#v", version)

	var (
		b []byte
	)

	if b, err = json.Marshal(version); err != nil {
		return err
	}
	log.Println(string(b))

	if b, err = xml.Marshal(version); err != nil {
		return err
	}
	log.Println(string(b))

	log.Printf("Major: %v", version.Major())
	log.Printf("Mivor: %v", version.Minor())
	log.Printf("Patch: %v", version.Patch())
	log.Printf("Pre  : %v", version.Prerelease())
	log.Printf("Meta : %v", version.Metadata())

	return nil
}
