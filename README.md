## Test Majoo Backend

Sebelum running project ini, langkah-langkah yang harus dilakukan :
- setelah melakukan git clone/ download project ini, running `go mod vendor` atau `go mod tidy`
- buat file bernama `.env`
- copy isi file `.env.example` ke dalam `.env`
- isi hal-hal penting di dalam `.env` seperti `DB_DATABASE`,`DB_USERNAME`,`DB_PASSWORD`
- running projectnya dengan cara ketik `go run main.go`
- buka postman 
- list url nya: 



# Migration

**to Run migration**

`DBEVENT=migrate go run main.go`

**to rollback**

`DBEVENT=rollback go run main.go`

**to rollback and migration**

`DBEVENT=rollback_migrate go run main.go`

**for dockerfile** 

add `ENV DBEVENT=rollback_migrate`

if you only want to run code, you just have to execute the command `go run main.go`