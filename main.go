package main

import (
	"flag"
	"fmt"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/random"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"log"
	"os"
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
	image, err := random.Image(1, 1)
	if err != nil {
		return err
	}

	nameRef, err := name.ParseReference(repo, name.WeakValidation)
	if err != nil {
		return err
	}

	keychainOpt := remote.WithAuthFromKeychain(authn.DefaultKeychain)
	if err := remote.Write(nameRef, image, keychainOpt); err != nil {
		return err
	}

	digest, err := image.Digest()
	if err != nil {
		return err
	}

	digestRef, err := name.NewDigest(fmt.Sprintf("%s@%s", nameRef.Context().Name(), digest.String()), name.WeakValidation)
	if err != nil {
		return err
	}

	if err := remote.Delete(digestRef, keychainOpt); err != nil {
		return err
	}

	return nil
}
