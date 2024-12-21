package main

func main() {
	// fmt.Print("Hello World")
	// fmt.Printf("Hello %v %T", 11, 11)

	// VARIABLE
	// var kata string
	// kata = "Programming Golang"
	// var kata1, kata2 string = "Satu", "Dua"
	// var (
	// 	kata1 int = 12
	// 	kata2     = "golang"
	// )

	// fmt.Println(kata)
	// fmt.Println(kata1, kata2)

	//CONSTANT
	// const pi float32 = 3.14
	// const tanggalLahir int = 16042003
	// println(pi, tanggalLahir)

	// const (
	// 	kelamin1 = "Pria"
	// 	kelamin2 = "Wanita"
	// )
	// fmt.Println(kelamin1, kelamin2)

	//POINTER
	// var pointerSaya *string
	// var kalimat = "Hello World"
	// pointerSaya = &kalimat
	// fmt.Println(*pointerSaya)

	// kalimat = "Hallo"
	// fmt.Println(*pointerSaya)
	// fmt.Println(pointerSaya)

	// *pointerSaya = "Ini akan merubah isi var kalimat namun tidak dengan alamat yang di store"
	// fmt.Println(pointerSaya)
	// fmt.Println(kalimat)
	// fmt.Println(*pointerSaya)

	// fmt.Printf("%v", *pointerSaya)

	//FUNCTION
	// initHello("Ghufron", 20)
	// var funcCatcher = getName("Ghufron")
	// fmt.Println(funcCatcher)
	// var hasilTambah, hasilKurang = tambahKurang(10, 20)
	// fmt.Printf("Tambah %d, Kurang %d \n", hasilTambah, hasilKurang)

	// var hasilTambah2, _ = tambahKurang(1, 2)
	// fmt.Println(hasilTambah2)

	// cetakSemua("Ghufron", "Bagaskara", "Fakultas Ilmu Komputer")
	// var returnPi = func() float32 {
	// 	return 3.14
	// }
	// fmt.Println(returnPi())

	// operasiAngka(5, kuadrat)
	// operasiAngka(4, kubic)

	//CONDITIONAL STATEMENT
	// const minSaldo = 10
	// var saldo = 200
	// if saldo > minSaldo {
	// 	fmt.Println("Tarik tunai dapat dijalankan")
	// } else if saldo == minSaldo {
	// 	fmt.Println("Saldo anda pas, tarik tidak dapat dijalankan")
	// } else {
	// 	fmt.Println("Tarik tunai tidak dapat dijalankan")
	// }

	// var trafficColor = ""

	// switch trafficColor {
	// case "merah":
	// 	fmt.Println("Berhenti")
	// case "kuning":
	// 	fmt.Println("Bersiap..")
	// case "hijau":
	// 	fmt.Println("Jalan")
	// case "biru", "pink", "putih":
	// 	fmt.Println("Lampu mungkin rusak, hubungi teknisi")
	// default:
	// 	fmt.Println("Lampu sedang tidak berfungsi, hubungi teknisi")
	// }

	//ERROR
	// hasilBagiErr, err := bagiAngka(10, 0)
	// if err != nil {
	// 	fmt.Printf("Terjadi kesalahan:%s \n", err.Error())
	// 	if errors.Is(err, globErrorNol) {
	// 		fmt.Println("Menggunakan 0 akan menyebabkan hasil bagi menjadi tak hingga")
	// 	}
	// } else {
	// 	fmt.Printf("Hasil bagi anda %f \n", hasilBagiErr)
	// }

	// defer fmt.Println("Hello World")
	// defer func() {
	// 	r := recover()
	// 	fmt.Println("Panic berhasil di recover", r)
	// }()
	// var hasilBagiPanic = bagiPanic(10, 0)

	// fmt.Println(hasilBagiPanic)
	// fmt.Println("Hello World")

	// ITERATION AND ARRAYS
	// var warna = [...]string{"merah", "kuning", "hijau"}

	// for i := len(warna) - 1; i >= 0; i-- {
	// 	fmt.Println(warna[i])
	// }

	// for i, w := range warna {
	// 	fmt.Println(warna[i])
	// 	fmt.Println(w)
	// }

	// var hobby = [3]string{1: "reading"}
	// fmt.Println(hobby)

	// var foods = [...]string{100: "pecel"}
	// fmt.Println(len(foods))

	// var i int = 0
	// for i != 10 {
	// 	fmt.Println("Nilai i : ", i)
	// 	i++
	// }

	// var kalimat = "Hello World"
	// for i, k := range kalimat {
	// 	fmt.Printf("Nilai: %c, index ke: %d, tipe data: %T \n", k, i, k)
	// }

	//SLICE
	// var warna []string = []string{"merah", "kuning", "hijau"}
	// var angka1 = make([]int, 3, 5) // (tipe, panjang, kapasitas)
	// fmt.Println(angka1)

	// fmt.Println("panjang: ", len(angka1))   // panjang
	// fmt.Println("kapasitas: ", cap(angka1)) // kapasitas
	// fmt.Println(angka1[4])                  // panic:

	// angka1 = append(angka1, 1, 213, 21, 1231, 1231, 123, 5643)
	// fmt.Println(angka1)

	// fmt.Println(warna[1:3])

	//MAP
	// var hobi = map[string]string{"Ghufron": "Berenang"} // key bersifat case sensitive, tidak boleh sama
	// fmt.Println(hobi)

	// var hobi2 = make(map[string]string)
	// hobi2["Hylda"] = "Membaca"
	// fmt.Println(hobi2)

	// for kunci, nilai := range hobi {
	// 	fmt.Println("Hobi ", kunci, " adalah ", nilai)
	// }

	//STRUCT
	// var persegi = BangunRuang{10, 10}
	// var persegiPanjang = BangunRuang{panjang: 16, lebar: 4} //kita dapat meulis attributnya secara tidak berurutan

	// fmt.Println(persegi)
	// fmt.Println(persegiPanjang)

	// fmt.Println("Luas persegi dengan func : ", persegi.luas())

	// persegi.aturPanjang(100)
	// fmt.Println("panjang persegi : ", persegi.panjang)
	// fmt.Println("Luas persegi dengan func : ", persegi.luas())

}

// type BangunRuang struct {
// 	panjang int
// 	lebar   int
// }

// func (b BangunRuang) luas() int {
// 	return b.panjang * b.lebar
// }

// // struct's pointer receiver
// func (b *BangunRuang) aturPanjang(p int) {
// 	b.panjang = p
// }

// func initHello(name string, umur int) {
// 	fmt.Println("Hello ", name, ". saya berumur ", umur)
// }
// func getName(nameString string) string {
// 	return nameString
// }

// func tambahKurang(x int, y int) (int, int) {
// 	return x + y, x - y
// }
// func kuadratLima() (hasil int) {
// 	hasil = 5 * 5 // ini akan otomatis ter return
// 	return
// }

// func cetakSemua(kata ...string) {
// 	fmt.Println(kata)
// }

// func operasiAngka(angka int, fx func(int) int) {
// 	var hasil = fx(angka)
// 	fmt.Println(hasil)
// }

// func kuadrat(angka int) int {
// 	return angka * angka
// }

// func kubic(angka int) int {
// 	return angka * angka * angka
// }

// var globErrorNol = errors.New("tidak dapat membagi dengan 0")

// func bagiAngka(x float32, y float32) (float32, error) {
// 	if y == 0 {
// 		return 0, globErrorNol
// 	}

// 	if y < 0 {
// 		return 0, errors.New("Tidak mendukung pembagian dengan angka negatif")
// 	}

// 	return x / y, nil
// }

// func bagiPanic(x float32, y float32) float32 {
// 	if y == 0 {
// 		panic("tidak dapat membagi dengan 0 \n") //panic akan menghentikan program
// 	}

// 	return x / y
// }
