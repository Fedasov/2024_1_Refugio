package user

import (
	"fmt"
	"sync"
)

// UserMemoryRepository is an in-memory implementation of UserRepository.
type UserMemoryRepository struct {
	mutex sync.RWMutex
	users map[uint32]*User
}

// NewUserMemoryRepository creates a new instance of UserMemoryRepository.
func NewInMemoryUserRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: FakeUsers,
	}
}

// GetAll returns all users from the storage.
func (repo *UserMemoryRepository) GetAll() ([]*User, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	users := make([]*User, 0, len(repo.users))
	for _, user := range repo.users {
		users = append(users, user)
	}

	return users, nil
}

// GetByID returns the user by its unique identifier.
func (repo *UserMemoryRepository) GetByID(id uint32) (*User, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	user, exists := repo.users[id]
	if !exists {
		return nil, fmt.Errorf("User with id %d not found", id)
	}

	return user, nil
}

// Add adds a new user to the storage and returns its assigned unique identifier.
func (repo *UserMemoryRepository) Add(user *User) (uint32, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	userID := uint32(len(repo.users) + 1)
	user.ID = userID
	repo.users[userID] = user

	return userID, nil
}

// Update updates the information of a user in the storage based on the provided new user.
func (repo *UserMemoryRepository) Update(newUser *User) (bool, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	_, exists := repo.users[newUser.ID]
	if !exists {
		return false, fmt.Errorf("User with id %d not found", newUser.ID)
	}

	repo.users[newUser.ID] = newUser

	return true, nil
}

// Delete removes the user from the storage by its unique identifier.
func (repo *UserMemoryRepository) Delete(id uint32) (bool, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	_, exists := repo.users[id]
	if !exists {
		return false, fmt.Errorf("User with id %d not found", id)
	}

	delete(repo.users, id)

	return true, nil
}
