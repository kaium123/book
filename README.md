
step 1 : docker compose -f docker-compose-local.yml  up -d
step 2 : go run main.go entity-generate-files
step 3 : go run main.go migrate-apply
step 4 : go run main.go migrate-hash
step 5 : go run main.go migrate-diff init_db
step 6 : go run main.go migrate-diff init_db
step 7 : go run main.go seed
step 8 : go run main.go server
# book
