package main

import (
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
		ver *semver.Version
		c1  *semver.Constraints
		c2  *semver.Constraints
		c3  *semver.Constraints
		err error
	)
	if ver, err = semver.NewVersion("v1.2.3"); err != nil {
		return err
	}
	if c1, err = semver.NewConstraint(">= v1.2.2"); err != nil {
		return err
	}
	if c2, err = semver.NewConstraint(">= v1.2.4"); err != nil {
		return err
	}
	if c3, err = semver.NewConstraint("<= v1.2.4"); err != nil {
		return err
	}

	var (
		constraints = []*semver.Constraints{c1, c2, c3}
	)
	for _, c := range constraints {
		log.Printf("%s compare %s ==> %v", ver, c, c.Check(ver))
	}

	return nil
}
