printf "Running Migrations: "

eval $(cat development.env) migrate -source file://config/db/migrations -database $DATABASE_URL up