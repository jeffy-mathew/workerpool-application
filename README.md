## Workerpool application


</br>
This application uses worker pool with go channels to call urls passed as arguments to the application concurrently.


The number of concurrent workers can be controlled with ```parallel``` flag.


</br>

## Prerequesites:

 - Go 1.15 [link](https://golang.org/doc/go1.16)



Steps to run application:

Download / `git pull` the soure code locally and run following commands from the project root directory

Download dependencies - `go mod download`

Run unit tests - `go test ./...`

Build application - `go build -o workerpool cmd/main.go`

Run application - ```./workerpool -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com ```


</br>

NOTE: The application runs wth default configuration on 5 concurrent workers, it can be overridden with flag `parallel`.
