set GOOS=linux
set CGO_ENABLED=0
del bin/gate,bin/live
go build -o bin/gate services/gate/main.go
go build -o bin/live services/live/main.go
REM pause