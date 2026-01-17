# go-binary-bridge

A high-performance, zero-dependency REST API prototype demonstrating raw binary data transfer (Uint8Array) over HTTP using Golang and Vanilla JS.

![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Web%20%7C%20IoT-orange?style=for-the-badge)
![Stability](https://img.shields.io/badge/Stability-Stable-brightgreen?style=for-the-badge)

---

## Deskripsi Proyek

Proyek ini mendemonstrasikan implementasi komunikasi data **low-level binary** menggunakan **Golang** sebagai backend dan **Vanilla JavaScript** sebagai client. Alih-alih menggunakan format teks berat seperti JSON, sistem ini mentransfer byte mentah (`Uint8Array`) melalui protokol HTTP untuk efisiensi maksimal.
> **Key Insight**: Menghindari overhead parsing JSON secara signifikan mengurangi latensi dan penggunaan memori pada sistem skala besar.

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

## Kontribusi & Masukan

Proyek ini merupakan purwarupa eksploratif yang akan terus dikembangkan. Saya sangat terbuka terhadap segala bentuk masukan, kritik membangun, maupun *insight* teknis untuk meningkatkan performa dan fungsionalitas sistem ini.

Beberapa rencana pengembangan ke depan meliputi:
- [ ] Implementasi **Endianness** handling untuk akurasi transfer data numerik.
- [ ] Integrasi **Gzip Compression** untuk optimasi *bandwidth* pada *stream* yang lebih besar.
- [ ] Penambahan mekanisme **Header-based Authentication** untuk keamanan tingkat lanjut.

Jika Anda memiliki ide atau ingin berdiskusi mengenai arsitektur ini, silakan buka [Issue](https://github.com/salambudiarto/go-binary-bridge/issues) atau hubungi saya melalui profil GitHub saya.

---

## Lisensi

Didistribusikan di bawah Lisensi MIT. Lihat file `LICENSE` untuk informasi lebih lanjut.

---

## 📩 Mari Terhubung!

Saya sangat terbuka untuk diskusi teknis, peluang kolaborasi, atau sekadar berbagi *insight* seputar pengembangan perangkat lunak. Silakan hubungi saya melalui kanal berikut:

[![Website](https://img.shields.io/badge/Website-sba.web.id-0055ff?style=for-the-badge&logo=google-chrome&logoColor=white)](https://sba.web.id)
[![WhatsApp](https://img.shields.io/badge/WhatsApp-Chat%20Me-25D366?style=for-the-badge&logo=whatsapp&logoColor=white)](https://wa.me/628981389448)
[![Email](https://img.shields.io/badge/Email-salambudiarto%40gmail.com-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:salambudiarto@gmail.com)
[![Instagram](https://img.shields.io/badge/Instagram-@salambudiarto-E4405F?style=for-the-badge&logo=instagram&logoColor=white)](https://instagram.com/salambudiarto)

---
<p align="center">
  Optimized for performance & built with precision. <br>
  <b>Salam Budiarto</b> © 2026
</p>
