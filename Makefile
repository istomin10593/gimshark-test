up:
	sudo docker compose up --build

down:
	sudo docker compose down

server-run:
	cd server && make run

test:
	cd server && make test

run-server:
	cd server && make run