docker-run:
	docker compose build && docker compose up
commit:
	@git status
	@read -p "Commit message: " msg; git add . && git commit -m "$$msg" && git push