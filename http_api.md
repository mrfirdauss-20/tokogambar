# HTTP API

Dokumen ini berisi spesifikasi mengenai API yang digunakan pada sistem ini.

> **Note:**
>
> Kamu diperbolehkan untuk merubah spesifikasi dari API ini.

## Mencari Gambar Serupa

POST: `/similars`

Endpoint ini digunakan untuk mencari gambar serupa dengan gambar yang di-upload.

**Body Fields:**

- `data`, String => berisi blob data dari gambar yang di-upload.

**Example Request:**

```
POST /similars
Content-Type: application/json

{
    "data": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD...."
}
```

**Success Response:**

```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": [
        {
            "filename": "image_1.jpg",
            "similarity_score": 100.0
        },
        {
            "filename": "image_2.jpg",
            "similarity_score": 90.0
        }
    ]
}
```

**Error Responses:**

- Bad Request (`400`)

    ```
    HTTP/1.1 400 Bad Request
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_BAD_REQUEST",
        "msg": "missing `data`"
    }
    ```

- Internal Server Error (`500`)

    ```
    HTTP/1.1 500 Bad Request
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_INTERNAL_ERROR",
        "msg": "unable to connect to database"
    }
    ```