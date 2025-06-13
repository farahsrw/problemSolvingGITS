# Deskripsi
Repository ini diunggah sebagai syarat pelaksanaan tes problem solving pada seleksi magang GITS.id. Terdapat 3 soal yang mana ketiga soal tersebut disolusikan menggunakan code dengan bahasa Go.

# Output soal 1:
![output_soal1](https://github.com/user-attachments/assets/4a7d576e-549d-42fa-b130-020ec9b776a8)

# Output soal 2:
<img width="491" alt="output_soal2" src="https://github.com/user-attachments/assets/85a138a5-04d3-4c5c-80d5-ba9fd06b3c61" />

# Output soal 3:
![output_soal3](https://github.com/user-attachments/assets/63a0291c-9745-4b45-939f-e976eb70ed2a)

# Kompleksitas soal 3:
Fungsi memiliki kompleksitas sebesar O(n), di mana n adalah panjang string input. Secara garis besar, fungsi ini bekerja dengan memeriksa setiap karakter dalam string satu per satu menggunakan perulangan. Setiap karakter dicek apakah merupakan kurung buka ((, [, {) atau kurung tutup (), ], }). Jika itu kurung buka, maka karakter disimpan dalam struktur stack, dan jika itu kurung tutup, maka program akan memeriksa apakah stack kosong (tidak ada pasangan buka) atau apakah urutan buka-tutup sesuai. Jika tidak cocok, maka fungsi langsung mengembalikan “NO”. Jika perulangan berjalan sampai panjang input menjadi 0 dan stack habis, maka fungsi akan mengembalikan "YES". Kasus terburuk adalah jika seluruh karakter dalam input adalah kurung buka, maka semua karakter akan dimasukkan ke dalam stack, menghasilkan penggunaan ruang sebesar O(n).
