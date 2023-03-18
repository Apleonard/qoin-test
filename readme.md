# Soal 1

```
Rancangkanlah diagram database untuk aplikasi rumah makan.
Jelaskan teknologi yang akan dipakai untuk aplikasi ini dan mengapa anda memilih teknologi
tersebut.
Kebutuhan:
1. Aplikasi ini bisa memasukkan pesanan-pesanan makanan pelanggan
2. Aplikasi ini bisa mengeluarkan struk pembelian
3. Aplikasi ini bisa mengeluarkan laporan penghasilan mingguan dan bulanan
4. Aplikasi ini bisa mengeluarkan laporan stok
Selain kebutuhan pokok diatas, silahkan tambahkan ide original anda untuk membuat aplikasi
lebih baik.
(catatan: Soal ini tidak membutuhkan pengambil test untuk membuat aplikasi (coding). )
```

## Teknologi yang digunakan

aplikasi yang akan saya buat menggunakan database `postgre`
aplikasi ini menggunakan `gorm` untuk memudahkan development dalam migrasi dan membuat relational antar table
untuk models untuk gorm migrate dapat dilihat code dan gambar diagramnya di:
folder `soal1`

1. memasukkan pesanan create ke table `orders` dan `order_details` dengan transaction
2. struk pembelanjaan dapat diambil dari `order_details` dengan `where order_id = order_id`
3. rencana saya adalah dengan menggunakan cron job untuk input ke table `weekly_reports` dan `monthly_reports`
4. laporan stok untuk masing-masing menu ada pada table `menus`

Selain kebutuhan pokok di atas, saya menambahkan kolom `most_popular_dish` dan `most_frequent_cust` pada table `weekly_reports` dan `monthly_reports`.
`most_frequent_cust` berguna supaya restaurant memiliki data customer tersering membeli di restaurantnya dan dapat digunakan untuk memberi promo, voucher, dll.
`most_frequent_cust` berguna supaya restaurant memiliki data makanan yang tersering dibeli guna menaikkan harga makanan tersebut, memberi promo, dll.

# Soal 2

terdapat di dalam folder soal2
`cd soal2`
