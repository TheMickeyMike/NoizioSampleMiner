package core

// Sound is DB entity model
type Sound struct {
	zPk    int
	zTitle string
	zData  []byte
}

// Title provides Sound title
func (s *Sound) Title() string {
	return s.zTitle
}

// Data provides Sound binary data
func (s *Sound) Data() []byte {
	return s.zData
}

// Sounds is list of Sound
type Sounds []Sound
