package pkg

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

// Store provides DB methods
type Store struct {
	database *sql.DB
}

// NewStore provides Store
func NewStore(path string) *Store {
	var store Store
	var err error
	log.Infof("Open DB file %s\n", path)
	store.database, err = sql.Open("sqlite3", path+"?mode=ro")
	if err != nil {
		log.Fatalln(err)
	}
	return &store
}

// Disconnect closes DB connection
func (s *Store) Disconnect() {
	defer s.database.Close()
}

// GetAllSounds returns all rescords from 'zsound' table
func (s *Store) GetAllSounds() Sounds {
	var (
		sounds         Sounds
		sound          Sound
		allSoundsQuery = "select z_pk, ztitle, zdata from zsound"
	)
	rows, err := s.database.Query(allSoundsQuery)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		err = rows.Scan(&sound.zPk, &sound.zTitle, &sound.zData)
		if err != nil {
			log.Fatalln(err)
		}
		sounds = append(sounds, sound)
	}
	return sounds
}
