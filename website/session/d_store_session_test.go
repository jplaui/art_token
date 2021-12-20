package session

import (
	// "context"
	"fmt"
	"testing"
	// "time"

	"crypto/md5"
)

func TestCreateSession(t *testing.T) {

	sessionStore := &sessionStore{
		store: make(map[string]Session),
	}

	userEmail := "test@test.com"
	// pasword := "123"
	hash := md5.Sum([]byte(userEmail))
	fileHash := string(hash[:])

	fmt.Println(fileHash)

	// now := time.Now()
	// session := &Session{
	// 	CreatedAt: now,
	// 	ExpiresAt: now.Add(time.Minute * 20),
	// 	FileHash:  fileHash,
	// }

	// ctx := context.Background()

	err := sessionStore.CreateSession(userEmail, fileHash)
	if err != nil {
		t.Errorf("CreateSession failed, error: %v.", err)
	}

}
