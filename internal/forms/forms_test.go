package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {

	r := httptest.NewRequest("POST", "/something", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when this should have been valid.")
	}
}

func TestForm_Required(t *testing.T) {

	r := httptest.NewRequest("POST", "/something", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Should not have received valid.")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/something", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Should have received valid.")
	}

}

func TestForm_IsEmail(t *testing.T) {

	r := httptest.NewRequest("POST", "/something", nil)
	form := New(r.PostForm)

	form.IsEmail("email")
	isValid := form.Valid()

	if isValid {
		t.Error("Got vaild with no email address.")
	}

	postedData := url.Values{}
	postedData.Add("email", "abc@me.com")

	r = httptest.NewRequest("POST", "/something", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	isValid = form.Valid()

	if !isValid {
		t.Error("Got invaild with email address.")
	}

	postedData = url.Values{}
	postedData.Add("email", "invalid")

	r = httptest.NewRequest("POST", "/something", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	isValid = form.Valid()

	if isValid {
		t.Error("Got vaild with an invalid email address.")
	}

}

func TestForm_MinLength(t *testing.T) {

	r := httptest.NewRequest("POST", "/something", nil)
	form := New(r.PostForm)

	form.MinLength("example", 5)
	if form.Valid() {
		t.Error("Passed for non-existant field")
	}

	isError := form.Errors.Get("example")
	if isError == ""{
		t.Error("Should have gotten an error back.")
	}

	postedData := url.Values{}
	postedData.Add("example", "example value")

	form = New(postedData)
	form.MinLength("example", 50)

	if form.Valid() {
		t.Error("Passed without enough characters")
	}

	postedData = url.Values{}
	postedData.Add("example", "more than enough characters")

	form = New(postedData)
	form.MinLength("example", 10)

	if !form.Valid() {
		t.Error("Somehow failed with enough characters")
	}

	isError = form.Errors.Get("example")
	if isError != ""{
		t.Error("Should not have gotten an error back.")
	}
}

func TestForm_Has(t *testing.T) {

	r := httptest.NewRequest("POST", "/something", nil)
	form := New(r.PostForm)

	has := form.Has("a")
	if has {
		t.Error("Form shows it has a field when it doesn't.")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)
	has = form.Has("a")

	if !has {
		t.Error("Shows form doesn't have field when it should.")
	}

}
