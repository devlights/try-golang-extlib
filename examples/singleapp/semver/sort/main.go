package main

import (
	"log"
	"os"
	"sort"

	"github.com/Masterminds/semver/v3"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdin)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		rawVersions = []string{
			"v1.0.0",
			"1.0.3",
			"1.0.4-beta.1",
			"v1.2.3",
			"v1.2.3-rc.3+build65536",
			"v1.1.4-alpha.3",
			"v0.9.2",
			"v0.8.4",
			"0.2.3-beta.4+build10",
		}

		versions = make([]*semver.Version, len(rawVersions))
	)

	for i, v := range rawVersions {
		var (
			ver *semver.Version
			err error
		)

		if ver, err = semver.NewVersion(v); err != nil {
			return err
		}

		versions[i] = ver
	}

	printVers(versions)
	hr()
	sort.Sort(semver.Collection(versions))
	printVers(versions)

	return nil
}

func printVers(versions []*semver.Version) {
	for _, v := range versions {
		log.Println(v)
	}
}

func hr() {
	log.Println("------------------------------")
}
