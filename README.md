# Basic Website

Super simple website using responsive CSS and a GoLang server. 

## Requirements

Since the server I am using is a go server, you need to make sure you computer is capable of [building GO](https://golang.org/doc/install).

Currently, this is SUPER simple, so we build the binary and the binary expects the same file structure as it was build in. 

## To Build

Run the following commands to build and run the server

```shell
go build server/main.go;
./main;
```

Then you can go to `localhost:8080` to see the basic webpage.

