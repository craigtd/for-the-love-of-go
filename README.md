# for-the-love-of-go

## Commands
Every Go project needs a go.mod file to indicate that this is a module and what the name is.
Run the following command from within the project folder:
```
go mod init <MODULE_NAME>
```

`gofmt` can be used to check for and resolve any code formatting issues 
Find formatting errors
```
gofmt -d <module_name>.go
```
Fix formatting errors
```
gofmt -w <module_name>.go
```

### Testing
Testing library: https://pkg.go.dev/testing (TODO: read)

Tests must accept a parameter with typ *testing.T
Test functions must be named in the pattern `TestXxx` and reside in a file whose name ends in `_test.go`.
Run all the tests within the current module:
```
go test
```

### Terminology

`slice` - a sequence of values of the same type (analagous to an array)