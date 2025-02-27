.PHONI: gitback gitfront build run test

git:
	git add .
	git commit -m "$(m)"
	git push

gitback:
	git add backend/*
	git commit -m "$(m)"
	git push

gitfront:
	git add ui/*
	git commit -m "$(m)"
	git push

build:
	go build -o groupie ./main.go

clean:
	rm -f ./groupie

run:
	./groupie "$(port)"

test:
	go test -v ./backend/api/responses

debug:
	@make clean
	@make build
	@make run
