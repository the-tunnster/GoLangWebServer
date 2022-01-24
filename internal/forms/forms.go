package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//Custom form struct and embeds a url.values object
type Form struct {
	url.Values
	Errors errors
}

//Initialises a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank.")
		return false
	}
	return true
}

//Valid returns true if there are no errors in the form
func (f *Form) Valid() bool {
	return (len(f.Errors) == 0)
}

//Checks for minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if(len(x)<length){
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long.", length))
		return false
	}
	return true
}
