package email

import "testing"

func newEmailOrFatal(t *testing.T, address string) *Email {
	email, err := NewEmail(address)
	if err != nil {
		t.Fatalf("new email: %v", err)
	}
	return email
}

func TestNewEmail(t *testing.T) {
	address := "user@email.com"
	email := newEmailOrFatal(t, address)
	if email.Address != address {
		t.Errorf("expected address qq, got %q", address, email.Address)
	}
}

func TestNewEmailInvalidAddress(t *testing.T) {
	_, err := NewEmail("bogus")
	if err == nil {
		t.Errorf("expected 'invalid address', got nil")
	}
}

func TestSaveEmailAndRetrieve(t *testing.T) {
	email := newEmailOrFatal(t, "new@email.com")

	m := NewEmailManager()
	m.Save(email)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 email, got %v", len(all))
	}
	if *all[0] != *email {
		t.Errorf("expected %v, got %v", email, all[0])
	}
}

func TestSaveAndRetrieveTwoEmails(t *testing.T) {
	email1 := newEmailOrFatal(t, "user1@email.com")
	email2 := newEmailOrFatal(t, "user2@email.com")

	m := NewEmailManager()
	m.Save(email1)
	m.Save(email2)

	all := m.All()
	if len(all) != 2 {
		t.Errorf("expected 2 emails, got %v", len(all))
	}
	if *all[0] != *email1 && *all[1] != *email1 {
		t.Errorf("missing email: %v", email1)
	}
	if *all[0] != *email2 && *all[1] != *email2 {
		t.Errorf("missing email: %v", email2)
	}
}

func TestSaveModifyAndRetrieve(t *testing.T) {
	email := newEmailOrFatal(t, "new@email.com")
	m := NewEmailManager()
	m.Save(email)

	email.Address = "newer@email.com"
	if m.All()[0].Address == "newer@email.com" {
		t.Errorf("updated email wasn't saved")
	}
}

func TestSaveTwiceAndRetrieve(t *testing.T) {
	email := newEmailOrFatal(t, "new@email.com")
	m := NewEmailManager()
	m.Save(email)
	m.Save(email)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 email, got %v", len(all))
	}
	if *all[0] != *email {
		t.Errorf("expected email %v, got %v", email, all[0])
	}
}

func TestSaveAndFind(t *testing.T) {
	email := newEmailOrFatal(t, "new@email.com")
	m := NewEmailManager()
	m.Save(email)

	ne, ok := m.Find(email.ID)
	if !ok {
		t.Errorf("didn't find email")
	}
	if *email != *ne {
		t.Errorf("expected %v, got %v", email, ne)
	}
}
