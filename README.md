# rumahweb-interview

untuk menjalankan aplikasi

1. Jalankan docker-compose up
2. Buka pgAdmin docker dengan credential yang ada di dalam docker pgAdmin
3. Register DB baru dengan nama users
4. Buat table users dengan field

   id bigserial primary key,
   name varchar(255) not null,
   email varchar(255) unique not null,
   password varchar(255) not null,

5. Jalankan go mod tidy di root directory
6. Jalankan go run cmd/main.go
7. Hit API di Postman

Untuk kodenya sebenarnya tidak sempurna, JWT nya error tidak terbaca setelah saya hit API nya yang telah saya buat di routes
Jadi yang berhasil hanya ketika saya uji register dan login. Namun saat login, token tidak direturn ke user

Terimakasih atas kesempatan yang diberikan, ini menjadi pelajaran yang sangat berharga bagi saya.

Link Postman => https://lively-moon-385293.postman.co/workspace/e5c02471-b626-4bc0-8b28-a2cfd67424f6/documentation/23530905-87b54629-e082-4da0-903f-a9e9b03c9f5e
