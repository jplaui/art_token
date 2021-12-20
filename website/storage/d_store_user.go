package storage

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/badoux/checkmail"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/crypto/bcrypt"
)

// ******** User struct **********

type customTime struct {

	// using time.RFC1123 as layout: "Mon, 02 Jan 2006 15:04:05 MST"
	time.Time
}

func (c customTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", c.Format(time.RFC1123))), nil
}

func (c customTime) UnmarshalJSON(v []byte) error {
	var err error
	c.Time, err = time.Parse(time.RFC1123, strings.ReplaceAll(string(v), "\"", ""))
	if err != nil {
		return err
	}
	return nil
}

// important: userId is the md5 hash of User.Email
type User struct {
	UUID      string     `json:"uuid"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt customTime `json:"created_at"`
}

type UserReturn struct {
	UserId    string     `json:"uuid"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	CreatedAt customTime `json:"created_at"`
}

// Validate user input
func (u *User) Validate(action string) error {

	switch strings.ToLower(action) {

	// minimal information requirement for user
	case "login":

		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil

	// all fields are required
	default:

		if u.FirstName == "" {
			return errors.New("FirstName is required")
		}
		if u.LastName == "" {
			return errors.New("LastName is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	}

}

// Prepare strips user input of any white spaces
// Hashes password of user
func (u *User) Prepare() error {

	password := strings.TrimSpace(u.Password)
	hashedpassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	u.Password = string(hashedpassword)
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	return nil
}

// ******* User store interface *********

var UserStoreError = errors.New("User store error")

type UserStoreConfig struct {
	UsersPath string
}

type UserStore interface {
	WriteUser(user User) error
	ReadUser(email string) (User, error)
	// UpdateUser(ctx context.Context, user User) error
	DeleteUser(email string) error
}

type userStore struct {
	mu     sync.RWMutex
	config UserStoreConfig
	logger log.Logger
}

func (us *userStore) WriteUser(user User) error {

	// log level
	logger := log.With(us.logger, "method", "WriteUser")

	// load user storage file
	fileLocation := getUserPath(user.Email, us.config.UsersPath, ".json")
	f, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		level.Error(logger).Log("OpenFile", err)
		return err
	}
	defer f.Close()

	// us.mu.Lock()
	err = json.NewEncoder(f).Encode(user)
	// us.mu.Unlock()

	if err != nil {
		level.Error(logger).Log("NewEncoder", err)
		return err
	}

	return nil
}

func (us *userStore) ReadUser(email string) (user User, err error) {

	// log level
	logger := log.With(us.logger, "method", "ReadUser")

	// access user json file
	path := getUserPath(email, us.config.UsersPath, ".json")
	f, err := os.Open(path)
	if err != nil {
		level.Error(logger).Log("os.Open:", err)
		return user, err
	}
	defer f.Close()

	// read json and decode into user struct
	if err := json.NewDecoder(f).Decode(user); err != nil {
		level.Error(logger).Log("json.NewReader(f).Decode(user):", err)
		return user, err
	}

	return user, nil
}

// func (us *userStore) UpdateUser(user User) error {
// 	// same as WriteUser
// }

func (us *userStore) DeleteUser(email string) error {

	// log level
	logger := log.With(us.logger, "method", "DeleteUser")

	// delete user json file
	path := getUserPath(email, us.config.UsersPath, ".json")
	err := os.Remove(path)
	if err != nil {
		level.Error(logger).Log("os.Remove:", err)
		return err
	}

	return nil
}

func NewUserStore(config UserStoreConfig, logger log.Logger) UserStore {
	return &userStore{
		config: config,
		logger: logger,
	}
}

// ******* utils functions ********

// hashes a password using bycrypt
func HashPassword(password string) (string, error) {

	// 11 is the cost for hashing the password. bcrypt.DefaultCost is 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	return string(bytes), err
}

// compares password with another password hash
func CheckPasswordHash(password, hash string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}

	return nil
}

func getUserPath(email, usersPath, ending string) string {
	var b bytes.Buffer
	hash := md5.Sum([]byte(email))
	fileHash := string(hash[:])
	b.WriteString(usersPath)
	b.WriteString(fileHash)
	b.WriteString(ending)
	return b.String()
}
