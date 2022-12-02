.PHONY: help run-test run-fuzz-test run-fuzz-test5s install-gotip

run-test:
	gotip test .

run-fuzz-test:
	gotip test -fuzz .

run-fuzz-test5s:
	gotip test -fuzz . -fuzztime 5s .

install-gotip:
	@go install golang.org/dl/gotip@latest
	@gotip download