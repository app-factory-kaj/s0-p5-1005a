// Package greeting builds the Greeting response shape for the /hello endpoint.
package greeting

import "fmt"

// Greeting is the JSON response shape returned by GET /hello.
type Greeting struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message"`
}

// For returns a Greeting addressed to name, or a generic greeting when name is empty.
func For(name string) Greeting {
	if name == "" {
		return Greeting{Message: "Hello, World!"}
	}
	return Greeting{Name: name, Message: fmt.Sprintf("Hello, %s!", name)}
}
