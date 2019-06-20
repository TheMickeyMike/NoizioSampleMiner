package core

import (
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helpers
func TempFilename(t *testing.T) string {
	f, err := ioutil.TempFile("", "db-test-")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}

func TestDisconnect(t *testing.T) {
	tempFilename := TempFilename(t)
	defer os.Remove(tempFilename)
	db, err := sql.Open("sqlite3", tempFilename)
	if err != nil {
		t.Fatal("Failed to open database:", err)
	}

	_, err = db.Exec("drop table foo")
	_, err = db.Exec("create table foo (id integer)")
	if err != nil {
		t.Fatal("Failed to create table:", err)
	}

	store := NewStore(tempFilename)

	stmt, err := store.database.Prepare("select id from foo where id = ?")
	if err != nil {
		t.Fatal("Failed to select records:", err)
	}

	store.Disconnect()
	_, err = stmt.Exec(1)
	if err == nil {
		t.Fatal("Failed to operate closed statement")
	}
}

// Tests
func TestGetAllSounds(t *testing.T) {
	tempFilename := TempFilename(t)
	defer os.Remove(tempFilename)
	db, err := sql.Open("sqlite3", tempFilename)
	if err != nil {
		t.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE ZSOUND")
	_, err = db.Exec("CREATE TABLE ZSOUND ( Z_PK INTEGER PRIMARY KEY, ZTITLE VARCHAR, ZDATA BLOB )")
	if err != nil {
		t.Fatal("Failed to create table:", err)
	}
	res, err := db.Exec(`
		INSERT
		INTO 'ZSOUND' ('Z_PK', 'ZTITLE', 'ZDATA')
		VALUES 
			('1', 'Buy me Prada', X'7465737420736f756e64'),
			('2', 'Balenciaga', X'7465737420736f756e64')`,
	)
	if err != nil {
		t.Fatal("Failed to insert record:", err)
	}
	affected, _ := res.RowsAffected()
	if affected != 2 {
		t.Fatalf("Expected %d for affected rows, but %d:", 2, affected)
	}

	expectedSounds := Sounds{
		Sound{
			zPk:    1,
			zTitle: "Buy me Prada",
			zData:  []byte("test sound"),
		},
		Sound{
			zPk:    2,
			zTitle: "Balenciaga",
			zData:  []byte("test sound"),
		},
	}

	store := NewStore(tempFilename)

	defer store.Disconnect()

	sounds, err := store.GetAllSounds()
	assert.NoError(t, err)
	assert.Equal(t, expectedSounds, sounds)
}

func TestUpdateAllSounds(t *testing.T) {
	tempFilename := TempFilename(t)
	defer os.Remove(tempFilename)
	db, err := sql.Open("sqlite3", tempFilename)
	if err != nil {
		t.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE ZSOUND")
	_, err = db.Exec("CREATE TABLE ZSOUND ( Z_PK INTEGER PRIMARY KEY, ZTITLE VARCHAR, ZDATA BLOB )")
	if err != nil {
		t.Fatal("Failed to create table:", err)
	}
	res, err := db.Exec(`
		INSERT
		INTO 'ZSOUND' ('Z_PK', 'ZTITLE', 'ZDATA')
		VALUES 
			('1', 'Sea Waves', X'7465737420736f756e64'),
			('2', 'This sould be Thunderstorm', X'7465737420736f756e64')`,
	)
	if err != nil {
		t.Fatal("Failed to insert record:", err)
	}
	affected, _ := res.RowsAffected()
	if affected != 2 {
		t.Fatalf("Expected %d for affected rows, but %d:", 2, affected)
	}

	store := NewStore(tempFilename)

	defer store.Disconnect()

	err = store.UpdateAllSounds()
	assert.NoError(t, err)

	expectedSounds := Sounds{
		Sound{
			zPk:    1,
			zTitle: "Sea Waves",
			zData:  []byte("test sound"),
		},
		Sound{
			zPk:    2,
			zTitle: "Thunderstorm",
			zData:  []byte("test sound"),
		},
	}

	rows, err := db.Query("SELECT z_pk, ztitle, zdata FROM zsound")
	if err != nil {
		t.Fatal("Failed to select records:", err)
	}
	defer rows.Close()

	var (
		sound  Sound
		sounds Sounds
	)
	for rows.Next() {
		err = rows.Scan(&sound.zPk, &sound.zTitle, &sound.zData)
		if err != nil {
			t.Fatal("Failed to select records:", err)
		}
		sounds = append(sounds, sound)
	}
	assert.Equal(t, expectedSounds, sounds)
}
