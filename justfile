just:
  just --list

run:
  go run main.go

test:
  go test ./...

new DAY:
  go run bin/newday.go {{DAY}}
