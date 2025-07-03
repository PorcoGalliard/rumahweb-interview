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

3. Jalankan go mod tidy di root directory
4. Jalankan go run cmd/main.go
5. Hit API di Postman

Untuk kodenya sebenarnya tidak sempurna, JWT nya error tidak terbaca setelah saya hit API nya yang telah saya buat di routes
Jadi yang berhasil hanya ketika saya uji register dan login. Namun saat login, token tidak direturn ke user

Terimakasih atas kesempatan yang diberikan, ini menjadi pelajaran yang sangat berharga bagi saya.