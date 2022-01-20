package forms

import (
	"net/http"
	"net/url"
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

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if(x==""){
		return false
	}
	return true
}
