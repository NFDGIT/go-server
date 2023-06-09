package dataAccess

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	gomysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gdb *gorm.DB

type User struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

func ConnectViaGORM() {
	dsn := "root:feng6368@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gdb, err = gorm.Open(gomysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		// 错误处理
	}
	fmt.Println("Connected!")
}

// 创建记录
func CreateRecord(record interface{}) error {
	result := gdb.Create(record)
	return result.Error
}

// 查询单个记录
func GetRecord(id uint, record interface{}) error {
	result := gdb.First(record, id)
	return result.Error
}

// 查询全部记录
func GetAllRecord(records []interface{}) error {
	result := gdb.Find(&records)
	return result.Error
}

// 更新记录
func UpdateRecord(record interface{}, field string, value interface{}) error {
	result := gdb.Model(record).Update(field, value)
	return result.Error
}

// 删除记录
func DeleteRecord(record interface{}) error {
	result := gdb.Delete(record)
	return result.Error
}

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

func Connect() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "feng6368",
		// User:   os.Getenv("DBUSER"),
		// Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// albID, err := addAlbum(Album{
	// 	Title:  "John Coltrane",
	// 	Artist: "Betty Carter",
	// 	Price:  49.99,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albID)

	// albums, err := albumsByArtist("Betty Carter")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Albums found: %v\n", albums)

	// // Hard-code ID 2 here to test the query.
	// alb, err := albumByID(25)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Album found: %v\n", alb)
}

// albumsByArtist queries for albums that have the specified artist name.
func AlbumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func AlbumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func AddAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
