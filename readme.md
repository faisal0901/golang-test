### Pengaturan Environment untuk Menghubungkan ke Database MS SQL Server

Sebelum dapat menggunakan aplikasi ini, pastikan Anda telah melakukan pengaturan environment untuk menghubungkan aplikasi dengan database MS SQL Server. Langkah-langkah yang perlu dilakukan adalah sebagai berikut:

1. Pastikan Anda telah memiliki instance MS SQL Server yang dapat digunakan. Jika belum, silakan mengunduh dan menginstal instance MS SQL Server dari situs resmi Microsoft.

2. Buka file .env pada direktori root aplikasi,DB_link

### Menjalankan Aplikasi Menggunakan Perintah go run

1. Ketikkan perintah `go run main.go` pada terminal atau command prompt, dan tekan enter. Aplikasi akan mulai dijalankan.

2. Buka browser dan akses http://localhost:8080/ untuk mengakses aplikasi. Jika tidak ada masalah, maka tampilan halaman utama aplikasi akan muncul.

# Nama API

API ini menyediakan layanan untuk ...

## Endpoint

### POST /api/register

Deskripsi: Endpoint untuk melakukan proses registrasi

**Parameter:**

- `Name`: (string) Nama pengguna yang akan didaftarkan
- `Email`: (string) Email pengguna yang akan didaftarkan
- `Password`: (string) Password pengguna yang akan digunakan untuk login
- `Phone`: (string) Nomor telepon pengguna yang akan didaftarkan
- `Address`: (string) Alamat pengguna yang akan didaftarkan

### POST /api/login

```
POST /api/register HTTP/1.1
Host: example.com
Content-Type: application/json

{
"Name":"faisal",
"Email":"exmuhammadfaisal042205@gmail.com",
"Password":"faisal123",
"Phone" :"0895115461161",
"Address":"jalan swadaya"
}

```

**Contoh response:**

```
{
    "id": 13,
    "name": "faisal",
    "email": "exmuhammadfaisal042205@gmail.com",
    "phone": "089516464646",
    "address": "jalan swadaya",
    "created_at": "2023-02-22T12:54:29.1083068+07:00",
    "updated_at": "2023-02-22T12:54:29.1083068+07:00"
}
```

Deskripsi: Endpoint untuk melakukan proses login

**Parameter:**

- `Email`: (string) Email pengguna yang digunakan untuk login
- `Password`: (string) Password pengguna yang digunakan untuk login

**Contoh request:**

```
POST /api/login HTTP/1.1

Content-Type: application/json

{
"Email":"exmuhammadfaisal0402@gmail.com",
"Password":"faisal123"
}

```

**Contoh response:**

```
{
    "id": 2,
    "email": "exmuhammadfaisal0402@gmail.com",
    "Password": "$2a$10$vF4Z4HufECrRGPsvFhYLpuLcPK.kOaCC25lMZwNvO9AhGczfSh9Gi",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzcwNzAzODIsInVzZXJfaWQiOjJ9.tW1hhvhoiwLx0vg9aQiTb20_VKTLL2jPt7oNeq7DRyY"
}
```

**Keterangan:**

- Jika email dan password sesuai dengan data yang terdaftar, maka akan menghasilkan response dengan status 200 OK, token JWT yang digunakan sebagai authentikasi, dan informasi user yang berhasil login.
- Jika email atau password tidak sesuai, maka akan menghasilkan response dengan status 401 Unauthorized, dan pesan error "Invalid email or password".

### POST /api/index/transaction

Deskripsi: Endpoint untuk melakukan transaksi pembelian produk

**Authorization:**

Untuk melakukan akses ke API ini, dibutuhkan token JWT yang diberikan setelah proses login.

**Parameter:**

- `Qty`: (integer) Jumlah produk yang akan dibeli
- `ProductID`: (integer) ID produk yang akan dibeli

**Contoh request:**

```
POST /api/index/transaction HTTP/1.1
Host: example.com
Authorization: Bearer token_jwt
Content-Type: application/json

{
"Qty":3,
"ProductID":1
}
```

**Contoh response:**

```
  {
    "ID": 20,
    "Qty": 3,
    "Price": 15000,
    "CustomerID": 2,
    "ProductID": 1,
    "CreatedAt": "2023-02-22T12:59:28.6526276+07:00",
    "UpdatedAt": "2023-02-22T12:59:28.6526276+07:00"
}
```

### POST /api/index/logout

Deskripsi: Endpoint untuk melakukan proses logout dari aplikasi

**Authorization:**

Untuk melakukan akses ke API ini, dibutuhkan token JWT yang diberikan setelah proses login.

**Contoh request:**

```
POST /api/index/logout HTTP/1.1
Host: example.com
Authorization: Bearer token_jwt

```

### GET /api/index/transaction

Deskripsi: Endpoint untuk melihat daftar transaksi yang dilakukan oleh pengguna

**Authorization:**

Untuk melakukan akses ke API ini, dibutuhkan token JWT yang diberikan setelah proses login.

**Contoh request:**

```
GET /api/index/transaction HTTP/1.1
Host: example.com
Authorization: Bearer token_jwt
```

**Contoh response:**

```
[
    {
        "id": 2,
        "name": "faisal",
        "email": "exmuhammadfaisal0402@gmail.com",
        "phone": "089516543215",
        "address": "jalan swadaya",
        "transactions": [
            {
                "id": 1,
                "qty": 3,
                "price": 15000,
                "productId": 1
            },
            {
                "id": 2,
                "qty": 3,
                "price": 15000,
                "productId": 1
            },
    },
    {
        "id": 3,
        "name": "faisal",
        "email": "exmuhammadfaisal0401@gmail.com",
        "phone": "089516543215",
        "address": "jalan swadaya",
        "transactions": null
    },

]
```

### POST /api/index/product

Deskripsi: Endpoint untuk menambahkan produk baru ke dalam database

**Authorization:**

Untuk melakukan akses ke API ini, dibutuhkan token JWT yang diberikan setelah proses login.

**Contoh request:**

```
POST /api/index/product HTTP/1.1
Host: example.com
Content-Type: application/json
Authorization: Bearer token_jwt

{
"Name":"ayam",
"MerchantID":1,
"Price":28000,
"Category":"makanan",
"Description":"makanan yang enak banget"
}

```

**Contoh response:**

```
{
    "id": 6,
    "name": "ayam",
    "description": "makanan yang enak banget",
    "merchant_id": 1,
    "price": 28000,
    "created_at": "2023-02-22T13:09:58.2069911+07:00",
    "updated_at": "2023-02-22T13:09:58.2069911+07:00"
}

```

### GET /api/index/merchant

Deskripsi: Endpoint untuk melihat daftar merchant beserta produk yang dimiliki oleh masing-masing merchant.

**Authorization:**

Untuk melakukan akses ke API ini, dibutuhkan token JWT yang diberikan setelah proses login.

**Contoh request:**

```
GET /api/index/merchant HTTP/1.1
Host: example.com
Authorization: Bearer token_jwt

```

**Contoh response:**

```
[
    {
        "DeletedAt": null,
        "ID": 1,
        "Name": "toko kue cisalak",
        "Address": "jalan cisalak",
        "Phone": "089516543215\r\n",
        "Products": [
            {
                "ID": 1,
                "Name": "kue kering",
                "Description": "kue kering isi coklat",
                "MerchantID": 1,
                "Price": 5000,
                "CreatedAt": "2023-02-20T18:49:50Z",
                "UpdatedAt": "2023-02-20T18:49:51Z"
            },

        ],
        "CreatedAt": "2023-02-20T18:50:16Z",
        "UpdatedAt": "2023-02-20T18:50:16Z"
    }
]

```

### POST /api/index/merchant

Deskripsi: Endpoint untuk mendaftarkan merchant

**Authorization:**

Untuk melakukan akses ke API ini, dibutuhkan token JWT yang diberikan setelah proses login.

**Contoh request:**

```
POST /api/index/merchant HTTP/1.1
Host: example.com
Authorization: Bearer token_jwt

{
    "Name":"ayam",
    "Address":"jalan raaya bogor",
    "Phone":"08951654315"

}

```

**Contoh response:**

```
{
    "ID": 2,
    "Name": "toko maknaan enak",
    "Address": "jalan raaya bogor",
    "Phone": "08951654315",
    "CreatedAt": "2023-02-22T21:19:25.1807918+07:00",
    "UpdatedAt": "2023-02-22T21:19:25.1807918+07:00"
}
```
