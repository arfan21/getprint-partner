build-dev:
	docker build -f dev.Dockerfile -t getprint-service-partner-dev .

build-prod:
	docker build -f prod.Dockerfile -t getprint-service-partner-prod .