#!/bin/bash
#set -o errexit -o pipefail -o nounset

docker run --rm --name test-reg --publish 5555:5000 -e REGISTRY_STORAGE_DELETE_ENABLED=true registry:2 2>&1 >/dev/null &
go run github.com/google/go-containerregistry/cmd/registry -port 6666 2>/dev/null &
GR_PID=$!
trap "docker rm -f test-reg; pkill -fl 'port 6666'" EXIT

curl --silent --retry 99 --retry-connrefused localhost:5555
curl --silent --retry 99 --retry-connrefused localhost:6666

go run main.go -repo localhost:5555/repo/name:tag
crane manifest localhost:5555/repo/name:tag

go run main.go -repo localhost:6666/repo/name:tag
crane manifest localhost:6666/repo/name:tag

go run main.go -repo demo.goharbor.io/delete-test/repo/name:tag
crane manifest demo.goharbor.io/delete-test/repo/name:tag

go run main.go -repo gcr.io/dotnet-build-gcp/delete-test/repo/name:tag
crane manifest gcr.io/dotnet-build-gcp/delete-test/repo/name:tag


