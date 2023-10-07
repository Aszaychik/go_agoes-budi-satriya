# Pengantar Deployment

## Sistem & Penyebaran Perangkat Lunak

Deployment adalah proses mendistribusikan aplikasi/produk yang dikembangkan oleh para pengembang, seringkali untuk mengubahnya dari status sementara menjadi permanen. Penyebaran dapat dilakukan melalui berbagai metode tergantung pada jenis aplikasinya, dengan aplikasi web & API dideploy ke server sementara aplikasi mobile ke Play Store/App Store.

## AWS (Amazon Web Services)

Amazon Web Services (AWS) adalah penyedia layanan komputasi awan terkemuka secara global, yang disediakan oleh Amazon. AWS menyediakan infrastruktur dan berbagai layanan awan untuk membantu bisnis dan individu mengelola aplikasi, data, dan sumber daya komputasi mereka secara efisien. Untuk memberikan penjelasan yang rinci, jelas, dan menarik tentang AWS, mari kita bahas beberapa aspek kunci:

1. **Infrastruktur Global:**
   AWS memiliki pusat data di seluruh dunia, memungkinkan pelanggan menjalankan aplikasi dan menyimpan data mereka di berbagai lokasi. Ini memungkinkan akses cepat dan andal ke layanan AWS dari hampir mana saja di dunia.

2. **Komputasi:**
   Amazon EC2 menyediakan berbagai jenis mesin virtual yang dapat dikonfigurasi sesuai kebutuhan. Pengguna dapat menjalankan mesin virtual dengan spesifikasi yang berbeda, mulai dari mesin kecil hingga mesin dengan daya komputasi tinggi, sesuai dengan tuntutan aplikasi mereka.

3. **Penyimpanan dan Basis Data:**
   AWS menyediakan layanan penyimpanan, termasuk Amazon S3 untuk penyimpanan objek, Amazon RDS untuk basis data SQL, dan Amazon DynamoDB untuk basis data NoSQL. Ini memungkinkan penyimpanan data yang aman dan skalabilitas tinggi.

4. **Jaringan dan Keamanan:**
   AWS memiliki jaringan global yang kuat dan menyediakan alat-alat keamanan canggih untuk melindungi data dan aplikasi. Layanan seperti Amazon Virtual Private Cloud (VPC) memungkinkan isolasi sumber daya dalam jaringan yang aman.

5. **Kecerdasan Buatan (AI) dan Pembelajaran Mesin (ML):**
   AWS menyediakan alat dan infrastruktur untuk pengembangan dan implementasi model pembelajaran mesin dan kecerdasan buatan. Pengguna dapat memanfaatkan TensorFlow (perpustakaan pembelajaran mesin open-source) atau AutoML untuk membuat dan melatih model AI dengan mudah.

6. **Layanan Khusus Industri:**
   AWS juga menawarkan solusi industri khusus, seperti Healthcare API untuk perusahaan kesehatan, Retail Solutions untuk peritel, dan lainnya. Ini memungkinkan perusahaan mengintegrasikan solusi awan yang disesuaikan dengan kebutuhan mereka.

7. **Manajemen Layanan:**
   AWS Management Console menyediakan antarmuka pengguna yang mudah digunakan untuk mengelola sumber daya dan layanan AWS. Selain itu, AWS juga menawarkan antarmuka baris perintah (Command Line Interface) untuk otomatisasi dan skrip.

8. **Harga yang Kompetitif:**
   AWS bersaing secara agresif dalam hal harga dengan penyedia awan lainnya. Mereka menawarkan model penagihan yang fleksibel, termasuk penagihan berdasarkan penggunaan (pay-as-you-go), sehingga pengguna hanya membayar untuk sumber daya yang benar-benar digunakan.

9. **Ekosistem Mitra:**
   AWS memiliki ekosistem mitra yang luas, termasuk berbagai penyedia layanan manajemen awan, penyedia perangkat lunak independen, dan konsultan. Ini membantu pelanggan mengintegrasikan dan mengelola solusi awan mereka lebih baik.

Amazon EC2 dan RDS menyediakan fleksibilitas dan kontrol penuh bagi pengembang untuk menjalankan aplikasi mereka di lingkungan awan. Dengan fitur-fitur seperti otomatisasi, skalabilitas, dan integrasi yang kuat dengan layanan AWS lainnya, EC2 dan RDS menjadi pilihan populer di kalangan pengembang.

## Amazon EC2 (Elastic Compute Cloud)

Amazon EC2 (Elastic Compute Cloud) adalah layanan web yang menyediakan kapasitas komputasi yang dapat disesuaikan di awan. Ini dirancang untuk membuat komputasi awan skala web menjadi lebih mudah bagi pengembang. Mari jelajahi beberapa aspek kunci dari Amazon EC2:

1. **Server Virtual (Instances):**
   EC2 memungkinkan pengguna membuat dan menjalankan server virtual, yang dikenal sebagai instances, di awan. Pengguna memiliki kontrol penuh atas instances, termasuk pilihan sistem operasi, tipe instance, dan pengaturan keamanan.

2. **Konfigurasi Instance:**
   Pengguna dapat mengkonfigurasi instance EC2 sesuai kebutuhan aplikasi mereka. Mereka dapat memilih jumlah CPU, jumlah memori, dan tipe instance yang sesuai.

3. **Otomatisasi Infrastruktur:**
   Amazon EC2 memungkinkan otomatisasi infrastruktur melalui penggunaan Amazon Machine Images (AMI) untuk dengan cepat dan efisien menyiapkan instances.

4. **Skalabilitas Otomatis:**
   EC2 menyediakan skalabilitas horizontal otomatis melalui Auto Scaling, memungkinkan penyesuaian otomatis kapasitas instance berdasarkan beban kerja.

5. **Keamanan:**
   Pengguna dapat mengelola keamanan instance menggunakan grup keamanan EC2, termasuk menetapkan aturan firewall untuk mengontrol lalu lintas masuk dan keluar.

6. **Manajemen Versi Aplikasi:**
   Pengguna dapat membuat Amazon Machine Images (AMI) yang berisi konfigurasi dan aplikasi mereka, memfasilitasi manajemen versi yang efisien.

7. **Integrasi dengan Layanan AWS Lainnya:**
   EC2 dapat dengan mudah diintegrasikan dengan layanan AWS lainnya, seperti Amazon S3 untuk penyimpanan objek dan Amazon RDS untuk basis data relasional.

8. **Pemantauan dan Logging:**
   AWS menyediakan alat pemantauan, termasuk Amazon CloudWatch, untuk melacak kinerja instance EC2 dan menganalisis log untuk pemecahan masalah.

9. **Penagihan Berbasis Penggunaan:**
   EC2 mengikuti model penagihan pay-as-you-go, di mana pengguna hanya membayar untuk sumber daya komputasi yang mereka gunakan.

Amazon EC2 menyediakan fleksibilitas dan kontrol penuh kepada pengembang untuk menjalankan aplikasi mereka di awan. Dengan fitur-fitur seperti otomatisasi, skalabilitas, dan integrasi yang kuat dengan layanan AWS lainnya, EC2 menjadi pilihan populer di kalangan pengembang.

## Amazon RDS (Relational Database Service)

Amazon RDS (Relational Database Service) adalah layanan basis data relasional yang dikelola sepenuhnya oleh AWS. Layanan ini menyederhanakan operasi dan pemeliharaan basis data SQL. Beberapa aspek kunci dari Amazon RDS adalah:

1. **Dukungan untuk Basis Data SQL Populer:**
   Amazon RDS mendukung beberapa sistem basis data SQL populer, termasuk MySQL, PostgreSQL, dan SQL Server.

2. **Otomatisasi Infrastruktur:**
   AWS mengelola semua aspek infrastruktur secara otomatis, termasuk manajemen server, pembaruan perangkat lunak, dan pencadangan data.

3. **Skalabilitas:**
   Amazon RDS memungkinkan pengguna dengan mudah menyesuaikan kapasitas CPU, memori, dan penyimpanan basis data berdasarkan kebutuhan mereka.

4. **Keamanan:**
   AWS memberikan perhatian khusus pada keamanan data di Amazon RDS. Ini termasuk enkripsi data saat berpindah dan saat istirahat, firewall, pengaturan otentikasi yang kuat, dan pemantauan keamanan.

5. **Pemulihan Bencana:**
   Amazon RDS secara otomatis melakukan pencadangan data secara teratur, memungkinkan pengguna mengembalikan basis data ke titik waktu tertentu jika terjadi kegagalan atau bencana.

6. **Ketersediaan Data Global:**
   Pengguna dapat meng-host basis data Amazon RDS di berbagai zona dan wilayah AWS, memastikan pengalaman aplikasi yang cepat di seluruh dunia.

7. **Integrasi dengan Alat AWS Lainnya:**
   RDS dapat dengan mudah diintegrasikan dengan layanan AWS lainnya, seperti Amazon EC2, Amazon S3, dan AWS Lambda, untuk membangun solusi komprehensif di awan.

8. **Manajemen Melalui Antarmuka Web dan Baris Perintah:**
   Pengguna dapat mengelola basis data Amazon RDS melalui antarmuka web yang mudah digunakan atau melalui baris perintah menggunakan AWS Command Line Interface (CLI).

9. **Penagihan Berbasis Penggunaan:**
   Amazon RDS mengikuti model penagihan pay-as-you-go, di mana pengguna hanya membayar untuk sumber daya basis data yang mereka gunakan.

Amazon RDS adalah solusi kuat untuk mengelola basis data SQL di awan. Ini menghilangkan banyak beban administrasi yang terkait dengan basis data, memungkinkan pengembang dan bisnis fokus pada pengembangan aplikasi tanpa khawatir tentang infrastruktur basis data yang rumit. Dengan fitur keamanan dan skalabilitas yang kuat, Amazon RDS membantu menjaga basis data Anda aman dan responsif.

## Tautan untuk Deployment Restful API untuk Tugas ini

http://18.141.231.208:8080
