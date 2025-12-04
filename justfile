just:
  just --list

run:
  go run main.go

bench-sync:
  hyperfine -i --runs 10 'GOMAXPROCS=1 just run'

bench:
  hyperfine -i --runs 10 'just run'

test:
  go test ./...

new DAY:
  go run bin/newday.go {{DAY}}
