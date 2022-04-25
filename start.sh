docker run --name postgres -e POSTGRES_PASSWORD=root -d -p 5432:5432 postgres 
docker build -t app1 .
docker run --name app1 -d -p 8000:8000 app1:latest 
docker exec -i postgres psql -U postgres -d postgres < /home/merlins/Projeler/to-do-api/db/migrations/create_comments_table.sql
docker exec -i postgres psql -U postgres -d postgres < /home/merlins/Projeler/to-do-api/db/migrations/create_taskstable.sql