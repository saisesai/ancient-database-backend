all: clean build

build:
	go mod tidy
	go build -o ancient-database .
	mkdir ./dist
	mkdir ./dist/public
	cp ./ancient-database ./dist
	cp config.yaml ./dist

.PHONY clean:
	rm -f *.db
	rm -f *.exe
	rm -f ./model/*.db
	cd public && ls | xargs rm -rf
	rm -rf ./dist
