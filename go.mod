module github.com/00bpa/advent_of_code

go 1.21.4

require github.com/stretchr/testify v1.8.4

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require internal/set v1.0.0

replace internal/set => ./internal/set

require (
	gonum.org/v1/gonum v0.14.0
	internal/stack v1.0.0
)

replace internal/stack => ./internal/stack
