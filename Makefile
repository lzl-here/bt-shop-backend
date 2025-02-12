deploy:
	go run deploy/es/es.go -host=http://localhost:9200

.PHONY: deploy