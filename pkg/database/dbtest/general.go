package dbtest

// IsEmpty is imitation method checking empties database
func (db *DbTest) IsEmpty() bool {
	return true
}

// Clear is imitation method clearing database
func (db *DbTest) Clear() {
}

// Init is imitation method init database
func (db *DbTest) Init(_ bool) error {
	return nil
}
