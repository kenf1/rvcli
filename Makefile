.PHONY: reset_tags rgo rgot dct pct

reset_tags:
	git tag -l | xargs git tag -d

rgo:
	go run *.go

rgot:
	cd test && go test

create_env:
	touch .env userconfig.env

dct:
	cd mock && docker compose up

pct:
	cd mock && podman compose up