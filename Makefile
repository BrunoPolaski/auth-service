docker-run:
	docker compose build && docker compose up
commit:
	git add . && git commit -m "$(message)" && git push