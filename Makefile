
run-app:
	docker-compose -f docker-compose.yml up

clean-app:
	docker-compose -f docker-compose.yml down && docker volume rm go-ecomrce_database_mysql