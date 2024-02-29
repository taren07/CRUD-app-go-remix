# Start containers
up:
    docker compose up -d

# Stop containers
down:
    docker compose down

# Remove container images, volumes, and orphan containers
destroy:
    docker compose down --rmi all --volumes --remove-orphans

# Build containers
build:
    docker compose build --no-cache --force-rm

# Shell login to Go container
sh-go:
    docker compose exec go bash

sh-migrate:
    docker-compose exec golang sh -c "GO_ENV=dev go run migrate/migrate.go"

# Shell login to Remix container
sh-remix:
    docker compose exec remix bash

# Shell login to MySQL container
sh-mysql:
    docker compose exec mysql bash

# Run migrations
migrate:
    docker compose exec go go run migrate/migrate.go

# Run go mod tidy
tidy:
    docker compose exec go go mod tidy

# Run npm install
npm-i:
    docker compose exec remix npm install

# Run npm run dev
npm-d:
    docker compose exec remix npm run dev
