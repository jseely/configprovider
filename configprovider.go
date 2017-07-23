package configprovider

// ConfigurationStore interface defines the functions required to be used
// as a store for configuration data.
// ConfigurationStore implementers should handle contextual expansion of "__ref__"
// keys. References will be a json object with the "type" key specifying the type of
// reference it is, and then key value pairs for any addition information that type
// may need. Implementers should return a not implemented exception for all types that are
// not explicitly implemented. The following reference types should be supported:
// "key"		=> 	References another key within the store. "path" key will also be provided
//					as an array of strings containing the path to the referenced key.
// "keyvault" 	=>	References a secret stored in keyvault. "subscription", "vault",
//					"secret" will also be provided as strings to perform the lookup
type ConfigurationStore interface {
	// Get returns the fully expanded object at the given path, if expansion fails at any
	// point Get should return the unexpanded object and the corresponding error
	Get(path ...string) (interface{}, error)
	// GetRaw returns the unexpanded object at the given path
	GetRaw(path ...string) (interface{}, error)
	// Set sets the given value at the specified path
	Set(value interface{}, path ...string) error
}
