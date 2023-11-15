# date_changer

Homework details can be found [here](homework.md).

# Assumptions
1. A single file will be used
1. A original date format will be entered
1. A desired date format will be entered
1. Date formats are valid RFC formats

The file will be read in, and a output file will be created with the new date format used.


# Building
1. Install golang, make sure it is at least version 1.20. 
1. Pull repo from github.
1. Make sure you are in the root of the project
1. Update packages with command: `go mod tidy`
1. To run the app without building a binary: `go run main.go`. 
1. Building the binary:
    1. For building on your local machine: `go build .`
    1. For Mac Silicon: `GOOS=darwin GOARCH=arm64 go build .` 
    1. For Mac Intel: `GOOS=darwin GOARCH=amd64 go build .`
    1. For Linux amd64: `GOOS=linux GOARCH=amd64 go build .`
    1. For Linux arm: `GOOS=linux GOARCH=arm64 go build .`
    1. For Windows amd64: `GOOS=windows GOARCH=amd64 go build .`
    1. For all possiable build options checkout this [gist](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63).


# Testing
To run all the test, in the root of the project run `go test ./...`. This will run all the tests in the project

# Running
This is a command line interface application. 

Follow the build steps above, you can run the app without building, or create a binary for your OS.

```
date_changer is used to replace a date format with a new one within a file.
It will replace all matching date formats with a new specified one.

example:
date_changer ~/Download/textfile.txt -o RFC822 -n RFC850 -output newtextfile.txt

Allowed date formats: ANSIC,UnixDate,RubyDate,RFC822,RFC822Z,RFC850,RFC1123,RFC1123Z,RFC3339,RFC3339Nano,DateTime
```

A example file can be found in the example_files folder. To test with it run the following command from the root of the project: `go run main.go example_files/day_in_time.txt -o RFC3339 -n UnixDate`
