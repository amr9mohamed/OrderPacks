build:
	docker build -t orderapp .
run:
	docker run -p 8080:8080 orderapp
start :
	docker build -t orderapp . && docker run -p 8080:8080 orderapp
