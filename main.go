package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// fn "projectkp/functionProj"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// pg "project kp/struct"
type Pegawai struct {
	ID             int    `json:"Id"`
	IDSekolah      string `json:"IdSekolah"`
	IDJabatan      int    `json:"IdJabatan"`
	Inisial        string `json:"Inisial"`
	NIP            string `json:"NIP"`
	NIY            string `json:"NIY"`
	NUPTK          string `json:"NUPTK"`
	Nama           string `json:"Nama"`
	JenisKelamin   string `json:"jenisKelamin"`
	Alamat         string `json:"Alamat"`
	Telp           string `json:"Telp"`
	HP             string `json:"HP"`
	ImageType      string `json:"ImageType"`
	ImageData      string `json:"ImageData"`
	TempatLahir    string `json:"TempatLahir"`
	TanggalLahir   string `json:"TanggalLahir"`
	Pendidikan     string `json:"Pendidikan"`
	StatusGuru     string `json:"StatusGuru"`
	MKTime         string `json:"MKTime"`
	UserID         string `json:"UserId"`
	Pass           string `json:"Pass"`
	Grup           string `json:"Grup"`
	Status         string `json:"Status"`
	Skin           int    `json:"Skin"`
	Menu           string `json:"Menu"`
	Kost           int    `json:"Kost"`
	Email          string `json:"Email"`
	StatusMengajar int    `json:"StatusMengajar"`
}

//create
func createPegawai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ID := r.FormValue("id")
		IDSekolah := r.FormValue("idSekolah")
		IDJabatan := r.FormValue("idJabatan")
		Inisial := r.FormValue("inisial")
		NIP := r.FormValue("nip")
		NIY := r.FormValue("niy")
		NUPTK := r.FormValue("nuptk")
		Nama := r.FormValue("nama")
		JenisKelamin := r.FormValue("jenisKelamin")
		Alamat := r.FormValue("alamat")
		Telp := r.FormValue("telp")
		HP := r.FormValue("hp")
		ImageType := r.FormValue("imageType")
		ImageData := r.FormValue("imageData")
		TempatLahir := r.FormValue("tempatLahir")
		TanggalLahir := r.FormValue("tanggalLahir")
		Pendidikan := r.FormValue("pendidikan")
		StatusGuru := r.FormValue("statusGuru")
		MKTime := r.FormValue("mkTime")
		UserID := r.FormValue("userID")
		Pass := r.FormValue("pass")
		Grup := r.FormValue("grup")
		Status := r.FormValue("status")
		Skin := r.FormValue("skin")
		Menu := r.FormValue("menu")
		Kost := r.FormValue("kost")
		Email := r.FormValue("email")
		StatusMengajar := r.FormValue("statusMengajar")

		//insert to database
		stmt, err := db.Prepare("INSERT INTO msttbpegawai(ID,IDSekolah,IDJabatan,Inisial,NIP,NIY,NUPTK,Nama,JenisKelamin,Alamat,Telp,HP,ImageType,ImageData,TempatLahir,TanggalLahir,Pendidikan,StatusGuru,MKTime,UserID,Pass,Grup,Status,Skin,Menu,Kost,Email,StatusMengajar) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		_, err = stmt.Exec(ID, IDSekolah, IDJabatan, Inisial, NIP, NIY, NUPTK, Nama, JenisKelamin, Alamat, Telp, HP, ImageType, ImageData, TempatLahir, TanggalLahir, Pendidikan, StatusGuru, MKTime, UserID, Pass, Grup, Status, Skin, Menu, Kost, Email, StatusMengajar)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Date Created")

		}

	}
}

//getAll
func getPegawai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var msttbpegawai []Pegawai

	sql := `SELECT
				ID,
				IFNULL(IDSekolah,''),
				IFNULL(IDJabatan,'') IDJabatan,
				IFNULL(Inisial,'') Inisial,
				IFNULL(NIP,'') NIP,
				IFNULL(NIY,'') NIY,
				IFNULL(NUPTK,'') NUPTK,
				IFNULL(Nama,'') Nama ,
				IFNULL(JenisKelamin,'') JenisKelamin,
				IFNULL(Alamat,'') Alamat,
				IFNULL(Telp,'')Telp,
				IFNULL(HP,'')HP,
				IFNULL(ImageType,'')ImageType,
				IFNULL(ImageData,'')ImageData,
				IFNULL(TempatLahir,'')TempatLahir,
				IFNULL(TanggalLahir,'')TanggalLahir,
				IFNULL(Pendidikan,'')Pendidikan,
				IFNULL(StatusGuru,'')StatusGuru,
				IFNULL(MKTime,'')MKTime,
				IFNULL(UserID,'')UserID,
				IFNULL(Pass,'')Pass,
				IFNULL(Grup,'')Grup,
				IFNULL(Status,'')Status,
				IFNULL(Skin,'')Skin,
				IFNULL(Menu,'')Menu,
				IFNULL(Kost,'')Kost,
				IFNULL(Email,'')Email,
				IFNULL(StatusMengajar,'')StatusMengajar
			FROM msttbpegawai`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var pegawai Pegawai
		err := result.Scan(&pegawai.ID, &pegawai.IDSekolah, &pegawai.IDJabatan,
			&pegawai.Inisial, &pegawai.NIP, &pegawai.NIY, &pegawai.NUPTK,
			&pegawai.Nama, &pegawai.JenisKelamin, &pegawai.Alamat, &pegawai.Telp, &pegawai.HP, &pegawai.ImageType, &pegawai.ImageData, &pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.Pendidikan, &pegawai.StatusGuru, &pegawai.MKTime, &pegawai.UserID, &pegawai.Pass, &pegawai.Grup, &pegawai.Status, &pegawai.Skin, &pegawai.Menu, &pegawai.Kost, &pegawai.Email, &pegawai.StatusMengajar)

		if err != nil {
			panic(err.Error())
		}
		msttbpegawai = append(msttbpegawai, pegawai)
	}

	json.NewEncoder(w).Encode(msttbpegawai)
}

//search by id

func getPegawaiByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msttbpegawai []Pegawai
	params := mux.Vars(r)

	sql := `SELECT
				ID,
				IFNULL(IDSekolah,''),
				IFNULL(IDJabatan,'') IDJabatan,
				IFNULL(Inisial,'') Inisial,
				IFNULL(NIP,'') NIP,
				IFNULL(NIY,'') NIY,
				IFNULL(NUPTK,'') NUPTK,
				IFNULL(Nama,'') Nama ,
				IFNULL(JenisKelamin,'') JenisKelamin,
				IFNULL(Alamat,'') Alamat, 
				IFNULL(Telp,'')Telp, 
				IFNULL(HP,'')HP,
				IFNULL(ImageType,'')ImageType,
				IFNULL(ImageData,'')ImageData,
				IFNULL(TempatLahir,'')TempatLahir,
				IFNULL(TanggalLahir,'')TanggalLahir,
				IFNULL(Pendidikan,'')Pendidikan,
				IFNULL(StatusGuru,'')StatusGuru,
				IFNULL(MKTime,'')MKTime,
				IFNULL(UserID,'')UserID,
				IFNULL(Pass,'')Pass,
				IFNULL(Grup,'')Grup,
				IFNULL(Status,'')Status,
				IFNULL(Skin,'')Skin,
				IFNULL(Menu,'')Menu,
				IFNULL(Kost,'')Kost,
				IFNULL(Email,'')Email,
				IFNULL(StatusMengajar,'')StatusMengajar
			FROM msttbpegawai WHERE ID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pegawai Pegawai
	for result.Next() {

		err := result.Scan(&pegawai.ID, &pegawai.IDSekolah, &pegawai.IDJabatan,
			&pegawai.Inisial, &pegawai.NIP, &pegawai.NIY, &pegawai.NUPTK,
			&pegawai.Nama, &pegawai.JenisKelamin, &pegawai.Alamat, &pegawai.Telp, &pegawai.HP, &pegawai.ImageType, &pegawai.ImageData, &pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.Pendidikan, &pegawai.StatusGuru, &pegawai.MKTime, &pegawai.UserID, &pegawai.Pass, &pegawai.Grup, &pegawai.Status, &pegawai.Skin, &pegawai.Menu, &pegawai.Kost, &pegawai.Email, &pegawai.StatusMengajar)

		if err != nil {
			panic(err.Error())
		}

		msttbpegawai = append(msttbpegawai, pegawai)
	}

	json.NewEncoder(w).Encode(msttbpegawai)
}

//Delete Pegawai
func deletePegawai(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM msttbpegawai WHERE ID = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Pegawai with ID = %s was deleted", params["id"])
}

//Update Pegawai
func updatePegawai(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newStatusGuru := r.FormValue("statusGuru")

		stmt, err := db.Prepare("UPDATE msttbpegawai SET StatusGuru = ? WHERE ID = ?")

		_, err = stmt.Exec(newStatusGuru, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Pegawai with ID = %s was updated", params["id"])
	}
}

//get by id and school id

func getPostByIDnSchID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var msttbpegawai []Pegawai

	ID := r.FormValue("id")
	IDSekolah := r.FormValue("idSekolah")

	sql := `SELECT
				ID,
				IFNULL(IDSekolah,''),
				IFNULL(IDJabatan,'') IDJabatan,
				IFNULL(Inisial,'') Inisial,
				IFNULL(NIP,'') NIP,
				IFNULL(NIY,'') NIY,
				IFNULL(NUPTK,'') NUPTK,
				IFNULL(Nama,'') Nama ,
				IFNULL(JenisKelamin,'') JenisKelamin,
				IFNULL(Alamat,'') Alamat,
				IFNULL(Telp,'')Telp,
				IFNULL(HP,'')HP,
				IFNULL(ImageType,'')ImageType,
				IFNULL(ImageData,'')ImageData,
				IFNULL(TempatLahir,'')TempatLahir,
				IFNULL(TanggalLahir,'')TanggalLahir,
				IFNULL(Pendidikan,'')Pendidikan,
				IFNULL(StatusGuru,'')StatusGuru,
				IFNULL(MKTime,'')MKTime,
				IFNULL(UserID,'')UserID,
				IFNULL(Pass,'')Pass,
				IFNULL(Grup,'')Grup,
				IFNULL(Status,'')Status,
				IFNULL(Skin,'')Skin,
				IFNULL(Menu,'')Menu,
				IFNULL(Kost,'')Kost,
				IFNULL(Email,'')Email,
				IFNULL(StatusMengajar,'')StatusMengajar
			FROM msttbpegawai WHERE ID = ? AND IDSekolah = ?`

	result, err := db.Query(sql, ID, IDSekolah)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pegawai Pegawai

	for result.Next() {

		err := result.Scan(&pegawai.ID, &pegawai.IDSekolah, &pegawai.IDJabatan,
			&pegawai.Inisial, &pegawai.NIP, &pegawai.NIY, &pegawai.NUPTK,
			&pegawai.Nama, &pegawai.JenisKelamin, &pegawai.Alamat, &pegawai.Telp, &pegawai.HP, &pegawai.ImageType, &pegawai.ImageData, &pegawai.TempatLahir, &pegawai.TanggalLahir, &pegawai.Pendidikan, &pegawai.StatusGuru, &pegawai.MKTime, &pegawai.UserID, &pegawai.Pass, &pegawai.Grup, &pegawai.Status, &pegawai.Skin, &pegawai.Menu, &pegawai.Kost, &pegawai.Email, &pegawai.StatusMengajar)
		if err != nil {
			panic(err.Error())
		}

		msttbpegawai = append(msttbpegawai, pegawai)
	}

	json.NewEncoder(w).Encode(msttbpegawai)

}

//main
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbtesting")
	// user:pass(local)/database
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Init router
	r := mux.NewRouter()

	//handle create pegawai
	r.HandleFunc("/Pegawai", createPegawai).Methods("POST")
	r.HandleFunc("/GetPegawai", getPegawai).Methods("GET")
	r.HandleFunc("/GetPegawaiByID/{id}", getPegawaiByID).Methods("GET")
	r.HandleFunc("/DeletePegawai/{id}", deletePegawai).Methods("DELETE")
	r.HandleFunc("/UpdatePegawai/{id}", updatePegawai).Methods("PUT")
	r.HandleFunc("/getByIDnSCH", getPostByIDnSchID).Methods("POST")

	//server start
	log.Fatal(http.ListenAndServe(":8080", r))
}
