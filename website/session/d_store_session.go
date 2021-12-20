package session

import (
	"errors"
	"sync"
	"time"
)

// ********** Session struct **********

type Session struct {
	FileHash  string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// ********** Session store interface **********

type SessionStore interface {
	CreateSession(sessionId string, fileHash string) error
	ReadSession(sessionId string) (Session, error)
	UpdateSession(sessionId string, session Session) error
	DeleteSession(sessionId string) error
}

type sessionStore struct {
	mu sync.RWMutex

	// map stores map[userID]Session
	store map[string]Session
}

// uses userID (user email MD5 hash) as key of session map
func (ss *sessionStore) CreateSession(sessionId, fileHash string) error {

	// write lock
	ss.mu.Lock()

	// add new session for user
	now := time.Now()
	session := Session{
		CreatedAt: time.Now(),
		ExpiresAt: now.Add(time.Minute * 20),
		FileHash:  fileHash,
	}
	ss.store[sessionId] = session

	// release write lock
	ss.mu.Unlock()

	return nil
}

func (ss *sessionStore) ReadSession(sessionId string) (session Session, err error) {

	// read lock
	ss.mu.RLock()

	// access session store
	session, exist := ss.store[sessionId]
	if !exist {
		err = errors.New("ReadSession access error")
	}

	// release read lock
	ss.mu.RUnlock()

	return session, err
}

func (ss *sessionStore) UpdateSession(sessionId string, session Session) error {

	// write lock
	ss.mu.Lock()

	// update session store
	now := time.Now()
	session.CreatedAt = now
	session.ExpiresAt = now.Add(time.Minute * 20)
	ss.store[sessionId] = session

	// release lock
	ss.mu.Unlock()

	return nil
}

func (ss *sessionStore) DeleteSession(sessionId string) error {

	// write lock
	ss.mu.Lock()

	// delete key
	delete(ss.store, sessionId)

	// release lock
	ss.mu.Unlock()

	return nil
}

func NewSessionStore() SessionStore {
	return &sessionStore{store: make(map[string]Session)}
}
