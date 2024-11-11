docker run -d --name my-postgres-irsyaan -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -v my-pg-volume-irsya
an:/var/lib/postgresql/data -p 5435:5432 postgres

docker run -d --name my-postgres-v2-irsyaan -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -v my-pg-volume-ir
syaan:/var/lib/postgresql/data -p 5435:5432 postgres