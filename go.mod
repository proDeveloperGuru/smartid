module github.com/proDeveloperGuru/smartid

go 1.23

require (
	github.com/go-resty/resty/v2 v2.16.5
	github.com/stretchr/testify v1.10.0
	go.uber.org/mock v0.5.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
    v0.1.2 // Security vulnerability discovered.
)