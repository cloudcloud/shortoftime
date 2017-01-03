
local-coverage: INT?=0
local-coverage: OUT?=../shortcov
local-coverage:
	RUNINTEGRATION=$(INT) gocov test $(shell go list ./... | grep -v /vendor/) > "$(OUT).json" && gocov-html "$(OUT).json" > "$(OUT).html"

local-test:
	@go test $(shell go list ./... | grep -v vendor)
