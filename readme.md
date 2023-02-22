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
