# Practice Commerce
I create simple server using Golang framework go-fiber to make API-related products and orders.

Prerequisite:
- Go 1.18
- Docker
- Postman

What you need to do:
1. Make sure your docker is running
2. Clone the repo
3. Open terminal (bash/wsl), place your current path terminal to the repo. Then execute command → `docker compose up -d`. Docker container will running
4. Open browser (example: google chrome), type `localhost:1234` → it will open adminer. We want to check our MySQL DB. Please input:
   - System: choose `MySQL`
   - Server: `db_mysql`
   - Username: `user`
   - Password: `password`
   - Database: `commerce_db`
5. You notice commerce_db don’t have a single table. To migrate the table open the terminal, and follow one of these:
   - execute command → `go build main.go`, wait till finish and exec cmd → `./main`
   - or, you can type `go run main.go`
6. Server will be running and trigger auto migrate table. Check the db.
7. Need to input one row data merchant. Example: username: "user", password: "password"
8. You can import postman collection to start using API ([here](https://api.postman.com/collections/21766278-5219879c-093c-41fa-9eb8-bcd250cf53d2?access_key=PMAT-01H6KKWVBDYGGNDSGQREDPGWFF)). How to import: open postman → import → copy the link → import / submit

You can check this [gdocs](https://docs.google.com/document/d/1dS5xtKiAQI0YpsfBFtJbxg2gW3gbirDVuh6k3j0OCrU/edit?usp=sharing) for more detail
