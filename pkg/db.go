package core

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

var (
	allSoundsQuery  = "SELECT z_pk, ztitle, zdata FROM zsound"
	updateAllSounds = `UPDATE zsound
						        SET ztitle='Thunderstorm'
	                  WHERE ztitle NOT IN ('Campfire', 'October Rain', 'Sea Waves', 'Sunny Day', 'Thunderstorm')`
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

// UpdateAllSounds updates sounds title to `Thunderstorm`
func (s *Store) UpdateAllSounds() error {
	tx, err := s.database.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec(updateAllSounds)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Debugf("Result: query executed successfully. %d rows affected", rowsAffected)
	return nil
}

// GetAllSounds returns all rescords from 'zsound' table
func (s *Store) GetAllSounds() (Sounds, error) {
	var (
		sounds Sounds
		sound  Sound
	)
	rows, err := s.database.Query(allSoundsQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&sound.zPk, &sound.zTitle, &sound.zData)
		if err != nil {
			return nil, err
		}
		sounds = append(sounds, sound)
	}
	return sounds, nil
}
