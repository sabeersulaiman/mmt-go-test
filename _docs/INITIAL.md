# Initial Setup

## Install Go 1.18
To avoid any surprises in the project you can install go 1.18 as well, please use the commands below
to install go 1.18 in your machine:

```shell
go install golang.org/dl/go1.18@latest
go1.18 download
```

Now you can use `go1.18` command instead of your `go` command to run everything.

## Clone the project & install dependencies
Clone the project from Github and then run the following command to install dependencies.
```shell
go get
```

## Install mockgen
`mockgen` is a mock generator tool which we are going to use to generate our mocks.
To install run the following command:
```shell
 go install github.com/golang/mock/mockgen@v1.6.0
```

To then test if mockgen is working as expected, generate a mock for the `MockTestConnector` by running this command:
```shell
mockgen -destination=mocks/mock_test.go \
                -package=mocks \
                mmt.com/lolbank/ports MockTestConnector
```