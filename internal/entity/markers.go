package entity

type Markers []Marker

// Save stores the markers in the database.
func (m Markers) Save(fileUID string) error {
	for _, marker := range m {
		if fileUID != "" {
			marker.FileUID = fileUID
		}

		if _, err := UpdateOrCreateMarker(&marker); err != nil {
			return err
		}
	}

	return nil
}

// Contains returns true if a marker at the same position already exists.
func (m Markers) Contains(m2 Marker) bool {
	const d = 0.07

	for _, m1 := range m {
		if m2.X > (m1.X-d) && m2.X < (m1.X+d) && m2.Y > (m1.Y-d) && m2.Y < (m1.Y+d) {
			return true
		}
	}

	return false
}

// FaceCount returns the number of valid face markers.
func (m Markers) FaceCount() (faces int) {
	for _, marker := range m {
		if !marker.MarkerInvalid && marker.MarkerType == MarkerFace {
			faces++
		}
	}

	return faces
}

// FindMarkers returns up to 1000 markers for a given file uid.
func FindMarkers(fileUID string) (Markers, error) {
	m := Markers{}

	err := Db().
		Where(`file_uid = ?`, fileUID).
		Order("marker_uid").
		Offset(0).Limit(1000).
		Find(&m).Error

	return m, err
}
