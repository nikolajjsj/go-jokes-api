package Models

import (
	"github.com/nikolajjsj/golang-jokes-api/Config"
)

// GetAllJokes fetch all jokes
func GetAllJokes(jokes *[]Joke) (err error) {
	if err = Config.DB.Find(jokes).Error; err != nil {
		return err
	}
	return nil
}

// CreateJoke

// GetJokeByID

// UpdateJoke

// DeleteJoke
