# gostackdriver
Redirecting stdin to Google Stackdriver for UNIX/Linux pipe

## Build

```
$ go build -o gostackdriver main.go
```

### for Raspberry Pi

```
$ GOOS=linux GOARCH=arm go build -o gostackdriver main.go
```

## Usage

```
$ ./gostackdriver -h
Usage:
  gostackdriver [OPTIONS]

Application Options:
  -l, --log-id=     LOG_ID
  -p, --project-id= PROJECT_ID

Help Options:
  -h, --help        Show this help message
```

### Example

```
$ tail -F /var/log/application.log | GOOGLE_APPLICATION_CREDENTIALS=~/keyfile.json ./gostackdriver -l app -p your-project-1234
```

## Setup a Service Account key file

https://cloud.google.com/iam/docs/creating-managing-service-account-keys
