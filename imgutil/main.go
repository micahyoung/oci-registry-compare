package main

import (
	"flag"
	"github.com/google/go-containerregistry/pkg/authn"
	"log"
	"os"

	"github.com/buildpacks/imgutil/remote"
)

var repoFlag = flag.String("repo", "", "repo name for test")

func main() {
	flag.Parse()
	if *repoFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*repoFlag); err != nil {
		log.Fatalln(err)
	}
}

func run(repo string) error {
	image, err := remote.NewImage(repo, authn.DefaultKeychain)
	if err != nil {
		return err
	}

	if err := image.Delete(); err != nil {
		return err
	}

	return nil
}
