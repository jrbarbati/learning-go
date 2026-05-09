package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrInvalidID = errors.New("invalid id")

type ErrMissingField struct {
	Field string
}

func (e ErrMissingField) Error() string {
	return fmt.Sprintf("missing field: %s", e.Field)
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v\n", count, emp)
		if err != nil {
			switch e := err.(type) {
			case interface{ Unwrap() []error }:
				allErrors := e.Unwrap()
				var messages []string

				for _, e := range allErrors {
					messages = append(messages, processError(e, emp))
				}

				message += " all errors: " + strings.Join(messages, ", ")
			default:
				message += " error: " + processError(e, emp)
			}
		}

		fmt.Println(message)
	}
}

/*
processError is a method for converting an error into a message. It has special handling for ErrInvalidID and
EmptyFieldError.
*/
func processError(err error, emp Employee) string {
	var fieldErr ErrMissingField

	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("%s: %s", err.Error(), emp.ID)
	} else if errors.As(err, &fieldErr) {
		return fieldErr.Error()
	}

	return fmt.Sprintf("%v", err)
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	var allErrors []error

	if len(e.ID) == 0 {
		allErrors = append(allErrors, ErrMissingField{Field: "ID"})
	} else if !validID.MatchString(e.ID) {
		allErrors = append(allErrors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		allErrors = append(allErrors, ErrMissingField{Field: "FirstName"})
	}
	if len(e.LastName) == 0 {
		allErrors = append(allErrors, ErrMissingField{Field: "LastName"})
	}
	if len(e.Title) == 0 {
		allErrors = append(allErrors, ErrMissingField{Field: "Title"})
	}

	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}
