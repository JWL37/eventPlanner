all:
	cd ./docker && docker compose up -d #--build
stop:
	cd ./docker && docker compose down