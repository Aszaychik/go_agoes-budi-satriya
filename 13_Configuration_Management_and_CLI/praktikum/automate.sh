#!/bin/bash

# Membuat folder utama dengan nama gabungan dari argumen pertama dan tanggal
folder_name="$1 at $(date +'%a %b %d %T %Z %Y')"
mkdir -p "$folder_name"

# Membuat struktur folder dan file-file yang diperlukan
mkdir -p "$folder_name/about_me/personal"
mkdir -p "$folder_name/about_me/professional"
mkdir -p "$folder_name/my_friends"
mkdir -p "$folder_name/my_system_info"
touch "$folder_name/about_me/personal/facebook.txt"
touch "$folder_name/about_me/professional/linkedin.txt"
touch "$folder_name/my_friends/list_of_my_friends.txt"
touch "$folder_name/my_system_info/about_this_laptop.txt"
touch "$folder_name/my_system_info/internet_connection.txt"

# Mengisi file facebook.txt dan linkedin.txt dengan URL sesuai argumen kedua dan ketiga
echo "https://www.facebook.com/$2" > "$folder_name/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$folder_name/about_me/professional/linkedin.txt"

# Mengisi file list_of_my_friends.txt dengan URL pastebin
curl "https://pastebin.com/raw/ZnH9qdBs" > "$folder_name/my_friends/list_of_my_friends.txt"

# Mengisi file about_this_laptop.txt dengan nama user dan informasi uname -a
echo "My Username: $(whoami)" > "$folder_name/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$folder_name/my_system_info/about_this_laptop.txt"

# Mengisi file internet_connection.txt dengan hasil keluaran ping ke google.com sebanyak 3 kali
ping -c 3 google.com > "$folder_name/my_system_info/internet_connection.txt"

# Menampilkan struktur folder setelah dibuat
tree "$folder_name"