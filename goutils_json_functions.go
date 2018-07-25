package goutils

// Return JSON format
type JSONSerializer interface {
  ToJSON() ([]byte, error)
}

// Read JSON data
type JSONDeserializer interface {
  FromJSON([]byte) error
}
