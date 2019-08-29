# Go REST Sample Application

## Requirements

Go 1.12.9 or later

## To Build

```sh
go build cmd/webserver/webserver.go
go test test/integration-tests/*_test.go -c -o integration-tests.test
```

## To Run Unit Tests

```sh
ginkgo -cover -r -v -skipPackage=integration-tests -outputdir=./ | tee ginkgo.output
```

## To Run Integration Tests

```sh
./integration-tests.test -test.v | go-junit-report > test.xml
```
