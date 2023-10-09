# Ringkasan Markdown: CI/CD

## Continuous Integration and Continuous Deployment (CI/CD)

Continuous Integration (CI) dan Continuous Deployment (CD) adalah praktik pengembangan perangkat lunak yang bertujuan meningkatkan proses pengembangan, meningkatkan kolaborasi, dan memberikan perangkat lunak berkualitas tinggi secara lebih efisien. CI melibatkan pengintegrasian otomatis perubahan kode dari berbagai kontributor ke dalam repositori bersama, memungkinkan deteksi dini masalah integrasi. CD memperluas proses ini dengan mengotomatiskan implementasi kode ke lingkungan produksi.

## Alat CI/CD

Alat CI/CD memfasilitasi implementasi praktik CI/CD. Alat ini mengotomatiskan tugas seperti integrasi kode, pengujian, dan implementasi. Mereka meningkatkan kolaborasi, mengurangi kesalahan manual, dan mempercepat pengiriman perangkat lunak. Alat CI/CD populer termasuk Jenkins, Travis CI, CircleCI, GitLab CI/CD, dan GitHub Actions.

## GitHub Actions

GitHub Actions adalah alat CI/CD dan otomatisasi yang disediakan oleh GitHub. Ini memungkinkan Anda mendefinisikan alur kerja menggunakan file YAML langsung di repositori kode Anda. Alur kerja dipicu oleh peristiwa seperti dorongan kode atau permintaan tarik. GitHub Actions memberikan lingkungan yang fleksibel dan dapat disesuaikan untuk membangun, menguji, dan mendeploy kode.

## Contoh Github Action

Berikut adalah contoh alur kerja GitHub Actions untuk CI Docker:

```yaml
name: Docker Image CI

on:
  push:
    branches: 'main'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Login to Docker Hub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          file: ./Dockerfile
          push: true
          tags: ${{ secrets.DOCKERHUB_USERNAME }}/go-simple-clean-rest-api:latest

      - name: Connect SSH
        uses: appleboy/ssh-action@v0.1.9
        with:
          host: ${{ secrets.SSH_HOST }}
          username: ${{ secrets.SSH_USERNAME }}
          key: ${{ secrets.SSH_KEY }}
          port: ${{ secrets.SSH_PORT }}
          script: |
            docker stop ${{ secrets.CONTAINER_NAME }}
            docker rm ${{ secrets.CONTAINER_NAME }}
            docker rmi ${{ secrets.DOCKERHUB_USERNAME }}/go-simple-clean-rest-api
            docker run --name ${{ secrets.CONTAINER_NAME }} -p 8080:8080 -d -e DB_HOST=${{secrets.DB_HOST}} -e DB_PORT=3306 -e DB_NAME=${{secrets.DB_NAME}} -e DB_USERNAME=${{secrets.DB_USERNAME}} -e DB_PASSWORD=${{secrets.DB_PASSWORD}} -e JWT_SECRET=${{secrets.JWT_SECRET}} ${{ secrets.DOCKERHUB_USERNAME }}/go-simple-clean-rest-api
```

## Alur kerja GitHub Actions ini melakukan langkah-langkah berikut:

- Checkout kode dari repositori.
- Login ke Docker Hub.
- Menyiapkan Docker Buildx.
- Membangun dan mendorong gambar Docker.
- Terhubung ke server jarak jauh menggunakan SSH.
- Menghentikan dan menghapus kontainer Docker yang ada.
- Menghapus gambar Docker sebelumnya.
- Menjalankan kontainer Docker baru dengan kode dan variabel lingkungan yang diperbarui.
