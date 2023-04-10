# Penjelasan Tahapan dan Sourcecodenya

Hal pertama yang perlu dilakukan adalah menyelesaikan file boiler plate, backend dan frontend yang ada pada tugas minggu sebelumnya. Pada UTS kita diminta untuk menampilkan data backend ke template frontend yang baru. Hal-hal yang perlu di perhatikan seperti berikut :

- Pada file controller.go/backend ditambahkan perintah berikut pada setiap collection yang dipanggil, contohnya pada perwalian:

```go
	func GetAllPerwalian(db *mongo.Database, col string) (data []model.Perwalian) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
```

- Pada file dev_test.go/backend ditambahkan perintah berikut, contohnya pada perwalian:

```go
func TestGetAll(t *testing.T) {
	data :=module.GetAllPerwalian(module.MongoConn, "perwalian")
	fmt.Println(data)
}
```

Pada file url.go/boiler plate ditambahkan perintah berikut :

```go
	page.Get("/perwalian", controller.GetAllPerwalian)
	page.Get("/allmahasiswa", controller.GetAllMahasiswa)
	page.Get("/alldosen", controller.GetAllDosen)
	page.Get("/alllocation", controller.GetAllLocation)
	page.Get("/allwaktu", controller.GetAllWaktu)
	page.Get("/allruangan", controller.GetAllRuangan)
}
```

# URL HEROKU

- https://alnovianti.herokuapp.com/perwalian
- https://alnovianti.herokuapp.com/allmahasiswa
- https://alnovianti.herokuapp.com/alldosen
- https://alnovianti.herokuapp.com/alllocation
- https://alnovianti.herokuapp.com/allwaktu
- https://alnovianti.herokuapp.com/allruangan

# URL Github Pages Frontend

- https://alnoviantirs.github.io/uts_perwalian/template/

# Screenshot Setiap Collection dan Isi Pada MongoDB

- MongoDB Dosen
  ![Mongo1](https://user-images.githubusercontent.com/125644091/230994878-e48f5102-c3ad-48ee-91e2-acf77b5ec030.png)
- MongoDB Location
  ![Mongo2](https://user-images.githubusercontent.com/125644091/230994898-9d958c1d-7e77-4cb9-9f5e-a06dd02a1e5a.png)
- MongoDB Mahasiswa
  ![Mongo3](https://user-images.githubusercontent.com/125644091/230994911-8d740a32-e300-4fa4-ab0e-da2d8c13f1b3.png)
- MongoDB Perwalian
  ![Mongo4](https://user-images.githubusercontent.com/125644091/230994924-67846412-d6cb-43da-9532-2f9ae7ac34df.png)
- MongoDB Ruangan
  ![Mongo5](https://user-images.githubusercontent.com/125644091/230994934-a524a145-ced4-48ea-ae2c-28f6386a7284.png)
- MongoDB Waktu
  ![Mongo6](https://user-images.githubusercontent.com/125644091/230994952-944a6a7c-7207-45ed-8eca-4873ac47d6c1.png)

# Screenshot Pada Postman

- Postman Perwalian
  ![Postman1](https://user-images.githubusercontent.com/125644091/230995402-a5f9fe8b-09c0-499b-a7cc-9391b3705af9.png)
- Postman Mahasiswa
  ![Postman2](https://user-images.githubusercontent.com/125644091/230995418-220275e4-8cf5-406b-a17c-870fcc3538f7.png)
- Postman Dosen
  ![Postman3](https://user-images.githubusercontent.com/125644091/230995431-08f453da-42bd-4dd2-a6d9-64922cbd5b98.png)
- Postman Location
  ![Postman4](https://user-images.githubusercontent.com/125644091/230995442-ebaee5cf-78c2-4d0a-bc68-93a6088a018e.png)
- Postman Waktu
  ![Postman5](https://user-images.githubusercontent.com/125644091/230995468-1bda9009-2025-46ff-8a15-7ceb267543de.png)
- Postman Ruangan
  ![Postman6](https://user-images.githubusercontent.com/125644091/230995489-35432679-e284-4180-ae14-f5c074b1650e.png)

# Screenshot Hasil Frontend

- Tabel Perwalian
  ![Frontend1](https://user-images.githubusercontent.com/125644091/230999457-b75431c3-f9b6-4cb0-a7d7-9ae419272a0e.png)
- Tabel Mahasiswa
  ![Frontend2](https://user-images.githubusercontent.com/125644091/230999486-1b4fae05-fde2-4562-a693-de13eeb89aad.png)
- Tabel Dosen
  ![Frontend3](https://user-images.githubusercontent.com/125644091/230999497-0a884732-2d5f-4e0c-8ad8-fd5df1ff0ab1.png)
- Tabel Waktu
  ![Frontend4](https://user-images.githubusercontent.com/125644091/230999523-45d785c0-a601-4995-babb-2ec8bac1b997.png)
- Tabel Ruangan
  ![Frontend5](https://user-images.githubusercontent.com/125644091/230999543-928b7b3a-b9ec-416d-a6a0-45b0dcc155f5.png)
- Tabel Location
  ![Frontend6](https://user-images.githubusercontent.com/125644091/230999555-1859c334-1d5f-4781-bac4-197db6e59768.png)
