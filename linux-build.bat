set GOOS=linux
set CGO_ENABLED=0
del bin/demo1
go build -o bin/demo1 services/demo1/main.go
REM pause