migrate_create:
	migrate create -ext sql -dir db/migrations -seq $(MIG_NAME)

migrate_up:
	migrate -source file://db/migrations -database postgres://asad:Tashkent2001@localhost:5432/simple_bank up

migrate_down:
	migrate -source file://db/migrations -database postgres://asad:Tashkent2001@localhost:5432/simple_bank down
