# SlowLoris
This is an implemetention of the SlowLoris attack in Golang

## Usage 
```bash
go run main.go -t 127.0.0.1:port -r [optional; default is 1000]
or 
go build main.go
./main -t 127.0.0.1:port -r [optional; default is 1000]
```

## Test the program
```bash
docker build -t name-of-the-image .
docker run -d --name container-name -p 8800:80 name-of-the-image
```

This program was created for educational purposes. Please use it at your own risk.