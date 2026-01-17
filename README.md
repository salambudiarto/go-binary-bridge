# go-binary-bridge

A high-performance, zero-dependency REST API prototype demonstrating raw binary data transfer (Uint8Array) over HTTP using Golang and Vanilla JS.

---

## Deskripsi Proyek

Proyek ini mendemonstrasikan implementasi komunikasi data **low-level binary** menggunakan **Golang** sebagai backend dan **Vanilla JavaScript** sebagai client. Alih-alih menggunakan format teks berat seperti JSON, sistem ini mentransfer byte mentah (`Uint8Array`) melalui protokol HTTP untuk efisiensi maksimal.

---

## Nilai Jual & Arsitektur Teknis

Bagi para pengembang profesional dan perekrut, proyek ini menunjukkan kompetensi dalam:

1. **Zero-Dependency Architecture**: Menggunakan *standard library* Go (`net/http`) tanpa framework pihak ketiga. Menunjukkan pemahaman mendalam tentang fundamental bahasa.
2. **Memory Efficiency**: Dengan mentransfer data binary (`application/octet-stream`), sistem meminimalisir beban CPU pada proses *serialization/deserialization* (JSON overhead).
3. **Low-Level Data Handling**: Implementasi penggunaan `io.ReadAll` dan manipulasi `[]byte` secara langsung di sisi server, serta `ArrayBuffer` di sisi client.
4. **CORS & Preflight Mastery**: Penanganan manual metode `OPTIONS` dan header keamanan untuk memastikan integritas komunikasi lintas domain.

---

## Instalasi & Penggunaan

### 1. Persiapan Server

Pastikan Go sudah terinstal, kemudian jalankan:

```bash
# Clone repository ini
git clone [https://github.com/salambudiarto/go-binary-bridge.git](https://github.com/salambudiarto/go-binary-bridge.git)

# Masuk ke direktori
cd go-binary-bridge

# Jalankan server
go run main.go

```

Server akan aktif di `http://localhost:8080`.

### 2. Penggunaan Client

1. Buka `index.html` di browser.
2. Buka **Developer Console** (F12).
3. Gunakan wrapper fungsi instan yang telah disediakan:

```javascript
// Contoh pemanggilan dari console
await kirim("Halo dari Client");

```

---

## Potensi Implementasi Industri

Arsitektur minimalis ini sangat relevan untuk skenario berikut:

* **IoT & Edge Computing**: Komunikasi perangkat dengan sumber daya terbatas (RAM/CPU) yang memerlukan efisiensi byte-level.
* **High-Frequency Data Streaming**: Pengiriman data koordinat, status sensor, atau log mentah dengan latensi rendah.
* **Custom Binary Protocol**: Langkah awal sebelum mengimplementasikan protokol yang lebih kompleks seperti Protocol Buffers (Protobuf) atau MessagePack.

---

## Struktur Kode

### Server (`main.go`)

Menggunakan pola *stateless* untuk menangani stream byte dan menggabungkan prefix respons secara efisien menggunakan `append` pada slice of bytes.

### Client (`index.html`)

Menyediakan interface `window.kirim` sebagai alat pengujian (debugging tool) yang mengubah string menjadi `Uint8Array` melalui `TextEncoder`.

---

## Roadmap Pengembangan

* [ ] Implementasi **Endianness** handling untuk transfer data numerik.
* [ ] Integrasi **Gzip Compression** pada level binary stream.
* [ ] Penambahan **Header-based Authentication**.

---

## Lisensi

Didistribusikan di bawah Lisensi MIT. Lihat file `LICENSE` untuk informasi lebih lanjut.
