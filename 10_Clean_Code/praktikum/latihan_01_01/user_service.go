package main

// type user struct {

// 	id int

// 	username int

// 	password int

// }

type User struct {
	ID       int
	Username string
	Password string
}


// type userservice struct {

// 	t []user

// }

type UserService struct {
	Users []User
}



// func (u userservice) getallusers() []user {

// 	return u.t

// }

func (u UserService) GetAllUsers() []User {
	return u.Users
}


// func (u userservice) getuserbyid(id int) user {

// 	for _, r := range u.t {

// 	if id == r.id {

// 	return r

// 	}

// 	}


// 	return user{}

// }


func (u UserService) GetUserByID(id int) User {
	for _, user := range u.Users {
		if id == user.ID {
			return user
		}
	}
	return User{}
}

/*
Jawaban

1.Tidak kurang dari enam kelemahan yang terdapat dalam kode program tersebut.

2.Ini adalah beberapa kelemahan yang saya temukan dalam kode program tersebut:

- Penamaan dalam deklarasi struktur variabel user tidak mengikuti konvensi penamaan dalam Go.
- Nama-nama field dalam struktur user tidak mencerminkan informasi yang benar.
- Metode getallusers() dan getuserbyid() menggunakan nama yang tidak konsisten dan tidak sesuai dengan konvensi dalam Go.
- Variabel t dalam struktur userservice kurang deskriptif dan tidak menggambarkan isi variabel dengan jelas.
- Metode-metode dalam userservice seharusnya menggunakan receiver dengan tipe pointer (*UserService) agar perubahan dapat dilakukan di dalam metode tersebut.
- Variabel r dalam perulangan tidak memiliki nama yang memiliki makna yang jelas.

3.Ini adalah alasan untuk setiap kelemahan tersebut:

- Struktur dan field harus mengikuti konvensi penamaan dalam huruf kapital agar dapat diakses dari luar paket. Nama field juga sebaiknya mencerminkan informasi yang jelas.
- Nama metode harus dimulai dengan huruf kapital agar dapat diakses dari luar paket. Penggunaan nama yang konsisten membantu dalam memahami fungsionalitas metode.
- Variabel harus memiliki nama yang menggambarkan isinya dengan jelas agar lebih mudah dipahami.
- Penggunaan receiver dengan tipe pointer diperlukan agar metode dapat memodifikasi data pada receiver.
- Memberikan nama yang bermakna pada variabel akan membantu dalam membaca kode dengan lebih baik.
*/