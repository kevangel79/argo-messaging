#!/bin/bash

bash -lc "mkdir /tmp/go && cd /tmp/go && export GOPATH=\$PWD && export PATH=\$GOPATH/bin:\$PATH && \
          mkdir -p src/github.com/ARGOeu && cp -Rp ~/ src/github.com/ARGOeu/argo-messaging/ && \
          cd src/github.com/ARGOeu/argo-messaging/ && go get github.com/axw/gocov/... && \
          go get github.com/AlekSi/gocov-xml && gocov test \$(go list ./... | grep -v /vendor/) | gocov-xml > ~/coverage.xml && \
          go get -u github.com/jstemmer/go-junit-report && go test \$(go list ./... | grep -v /vendor/) -v | go-junit-report > ~/junitresults.xml "