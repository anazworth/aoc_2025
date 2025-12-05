just:
  just --list

run:
  go run main.go

build:
  go build -o aoc main.go 

bench-sync: build
  hyperfine -i -w 1 --runs 10 'GOMAXPROCS=1 ./aoc'

bench: build
  hyperfine -i -w 1 --runs 10 './aoc'

bench-compare: build
  hyperfine -i -w 1 --runs 10 './aoc' 'GOMAXPROCS=1 ./aoc'

test:
  go test ./...

new DAY:
  go run bin/newday.go {{DAY}}
