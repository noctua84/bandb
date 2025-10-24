package forms

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/mail"
	"net/url"
	"strings"
)

// Form struct embeds a url.Values object and includes a map to hold form errors
type Form struct {
	url.Values
	Errors errors
}

// New initializes a Form struct with the provided url.Values and an empty errors map
func New(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

// Valid checks if there are any errors in the form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Has checks if a specific field is present and not empty in the form data
// If the field is empty, it adds an error message to the Errors map
func (f *Form) Has(field string, r *http.Request) bool {
	value := r.Form.Get(field)
	if value == "" {
		return false
	}
	return true
}

// Required checks if specific fields are present and not empty in the form data
// If any passed field is empty, it adds an error message to the Errors map
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			log.Printf("Required field '%s' is missing", field)
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks if a specific field meets a minimum length requirement
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	value := r.Form.Get(field)
	if len(value) < length {
		log.Printf("Required field '%s' is too short", field)
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}

	return true
}

// IsEmail checks if a specific field contains a valid email address
func (f *Form) IsEmail(field string) bool {
	value := f.Get(field)

	// 1. Basic format check using net/mail
	m, err := mail.ParseAddress(value)
	if err != nil || m.Address == "" {
		log.Printf("Invalid email format for field '%s': %v", field, err)
		f.Errors.Add(field, "Invalid email address")
	}

	// 2. basic sanity checks (length and presence of '@')
	if len(m.Address) > 254 {
		log.Printf("Email address too long for field '%s': %v", field, len(m.Address))
		f.Errors.Add(field, "Email address is too long")
	}

	parts := strings.Split(m.Address, "@")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		log.Printf("Email address must contain a single '@' character for field '%s'", field)
		f.Errors.Add(field, "Invalid email address format")
	}

	// 3. Further checks on local and domain parts
	local, domain := parts[0], parts[1]
	// Local part checks
	if len(local) > 64 {
		log.Printf("Local part of email address too long: %v", len(local))
		f.Errors.Add(field, "Local part of email address is too long")
	}

	if strings.HasPrefix(local, ".") || strings.HasSuffix(local, ".") || strings.Contains(local, "..") {
		log.Printf("Invalid local part in email address: %s", local)
		f.Errors.Add(field, "Invalid local part in email address")
	}

	// Domain part checks
	if len(domain) > 253 {
		log.Printf("Domain part of email address too long: %v", len(domain))
		f.Errors.Add(field, "Domain part of email address is too long")
	}

	if strings.HasPrefix(domain, "-") || strings.HasSuffix(domain, "-") || strings.Contains(domain, "..") {
		log.Printf("Invalid domain part in email address: %s", domain)
		f.Errors.Add(field, "Invalid domain part in email address")
	}

	if !strings.Contains(domain, ".") {
		log.Printf("Domain part must contain a dot: %s", domain)
		f.Errors.Add(field, "Domain part must contain a dot")
	}

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		log.Printf("No MX records found for domain '%s': %v", domain, err)
		f.Errors.Add(field, "Domain does not have valid MX records")
	}

	_, err = net.LookupHost(domain)
	if err != nil {
		log.Printf("Domain '%s' does not resolve: %v", domain, err)
		f.Errors.Add(field, fmt.Sprintf("Domain does not resolve, exists or cannot receive emails: %s", err))
	}

	errorCount := len(f.Errors.Get(field))

	if errorCount != 0 {
		return false
	}

	return true
}
