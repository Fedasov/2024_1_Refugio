package email

import (
	"fmt"
	"sync"
)

// EmailMemoryRepository represents the implementation of EmailRepository using an in-memory storage.
type EmailMemoryRepository struct {
	mu     sync.RWMutex
	emails map[uint64]*Email
}

// NewEmailMemoryRepository creates a new instance of EmailMemoryRepository.
func NewEmailMemoryRepository() *EmailMemoryRepository {
	return &EmailMemoryRepository{
		emails: FakeEmails,
	}
}

// GetAll returns all emails from the storage.
func (repository *EmailMemoryRepository) GetAll() ([]*Email, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	emails := make([]*Email, 0, len(repository.emails))
	for _, email := range repository.emails {
		emails = append(emails, email)
	}

	return emails, nil
}

// GetByID returns an email based on its unique identifier.
func (repository *EmailMemoryRepository) GetByID(id uint64) (*Email, error) {
	repository.mu.RLock()
	defer repository.mu.RUnlock()

	email, found := repository.emails[id]
	if !found {
		return nil, fmt.Errorf("Email with id %d not found", id)
	}

	return email, nil
}

// Add adds a new email to the storage and returns the assigned unique identifier.
func (repository *EmailMemoryRepository) Add(email *Email) (uint64, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	id := uint64(len(repository.emails) + 1)
	email.ID = id
	repository.emails[id] = email

	return id, nil
}

// Update updates the data of an email in the storage based on the provided new email.
func (repository *EmailMemoryRepository) Update(newEmail *Email) (bool, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	existingEmail, found := repository.emails[newEmail.ID]
	if !found {
		return false, fmt.Errorf("Email with id %d not found", newEmail.ID)
	}

	existingEmail.Topic = newEmail.Topic
	existingEmail.Text = newEmail.Text
	existingEmail.PhotoID = newEmail.PhotoID
	existingEmail.ReadStatus = newEmail.ReadStatus
	existingEmail.Mark = newEmail.Mark
	existingEmail.Deleted = newEmail.Deleted
	existingEmail.DateOfDispatch = newEmail.DateOfDispatch
	existingEmail.ReplyToEmailID = newEmail.ReplyToEmailID
	existingEmail.DraftStatus = newEmail.DraftStatus

	return true, nil
}

// Delete removes an email from the storage based on its unique identifier.
func (repository *EmailMemoryRepository) Delete(id uint64) (bool, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	_, found := repository.emails[id]
	if !found {
		return false, fmt.Errorf("Email with id %d not found", id)
	}

	delete(repository.emails, id)

	return true, nil
}
