//go:build disable_extras
// +build disable_extras

package federation

type UserDTO struct {
	Name  string
	Email string
}

type CompanyDTO struct {
	UUID  string
	Email string
}
