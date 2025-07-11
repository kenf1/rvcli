.PHONY: reset_tags rgo

reset_tags:
	git tag -l | xargs git tag -d

rgo:
	go run *.go