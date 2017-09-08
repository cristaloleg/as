package as

// Value interface for any value
type Value interface {
	// Get value
	Get() interface{}

	// HasValue returns true if value is avaliable
	HasValue() bool
}
