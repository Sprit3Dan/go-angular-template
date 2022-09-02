package entities

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId        string        `json:"userId"`
	Email         string        `json:"email"`
	Picture       string        `json:"picture"`
	Subscriptions Subscriptions `json:"subscriptions"`
}

func CreateNewUser(db *sql.DB, email, userId, picture string) (*User, error) {
	r, err := db.Query("INSERT INTO users (email, id, picture) VALUES (?, ?, ?) RETURNING *;", email, userId, picture)
	if err != nil {
		return nil, err
	}

	if !r.Next() {
		return nil, fmt.Errorf("something gone terribly wrong")
	}

	r.Next()

	u := &User{}
	r.Scan(&u.UserId, &u.Email)

	return u, nil
}

func CreateNewSubsciptionForUser(db *sql.DB, u *User, s *Subscription) (*UserSubscription, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	createdSubscription := &Subscription{}
	if r, err := tx.Query("INSERT INTO subscriptions (cityId) VALUES (?) RETURNING *", s.CityId); err != nil {
		return nil, err
	} else {
		r.Next()
		r.Scan(&createdSubscription.Id, &createdSubscription.CityId)
	}

	userSub := &UserSubscription{}
	if r, err := tx.Query(`
		INSERT INTO users_subscriptions (userId, subscriptionId) VALUES (?, ?) RETURNING *`,
		u.UserId, createdSubscription.Id); err != nil {
		return nil, err
	} else {
		r.Next()
		if err := r.Scan(&userSub.UserId, &userSub.SubscriptionId); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return userSub, nil
}

func LookupUser(db *sql.DB, id string) (*User, error) {
	r, err := db.Query("SELECT * from users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	if !r.Next() {
		return nil, fmt.Errorf("User with id %v doesn't exist", id)
	}

	u := &User{}
	r.Scan(&u.UserId, &u.Email, &u.Picture)

	return u, nil
}
