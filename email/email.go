package email

import (
	"fmt"
	"regexp"
	"time"
)

// Regular expression that matches 99% of the email addresses in use today.
var validBasicEmail = regexp.MustCompile(`(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`)

// IDs are set only for Emails that are saved by a EmailManager.
type Email struct {
	ID        int64
	Address   string
	CreatedAt time.Time
}

func NewEmail(address string) (*Email, error) {
	if !isValid(address) {
		return nil, fmt.Errorf("invalid email address")
	}
	return &Email{0, address, time.Now()}, nil
}

// EmailManager manages a list of emails in memory.
type EmailManager struct {
	emails []*Email
	lastID int64
}

// NewEmailManager returns an empty EmailManager.
func NewEmailManager() *EmailManager {
	return &EmailManager{}
}

// Save saves the given Email in the EmailManager.
func (m *EmailManager) Save(email *Email) error {
	if email.ID == 0 {
		m.lastID++
		email.ID = m.lastID
		m.emails = append(m.emails, cloneEmail(email))
		return nil
	}

	for i, t := range m.emails {
		if t.ID == email.ID {
			m.emails[i] = cloneEmail(email)
			return nil
		}
	}
	return fmt.Errorf("unknown email")
}

// Creates and returns a deep copy of the given Email.
func cloneEmail(e *Email) *Email {
	c := *e
	return &c
}

// All returns the list of all the Emails in the EmailManager.
func (m *EmailManager) All() []*Email {
	return m.emails
}

// Find returns the Email with the given id in the EmailManager and a boolean
// indicating if the id was found.
func (m *EmailManager) Find(ID int64) (*Email, bool) {
	for _, e := range m.emails {
		if e.ID == ID {
			return e, true
		}
	}
	return nil, false
}

func isValid(address string) bool {
	return validBasicEmail.MatchString(address)
}
