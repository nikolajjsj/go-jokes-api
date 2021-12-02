package Models

import (
	"context"
	"fmt"
	"github.com/nikolajjsj/go-jokes-api/api/config"
)

// GetAllJokes fetch all jokes
func GetAllJokes() ([]Joke, error) {
	var tempJokes []Joke

	if rows, err := Config.Connection.Query(context.Background(), "SELECT * FROM jokes"); err != nil {
		fmt.Println("Error while executing query", err)
		return nil, err
	} else {
		defer rows.Close()
		var tmpJoke Joke
		for rows.Next() {
			rows.Scan(&tmpJoke.ID, &tmpJoke.Joke, &tmpJoke.CreatedAt)
			fmt.Printf("%+v\n", tmpJoke)
			tempJokes = append(tempJokes, tmpJoke)
		}
		if err := rows.Err(); err != nil {
			fmt.Println("Error occur while finding joke: ", err)
			return nil, err
		}
	}

	return tempJokes, nil
}

// CreateJoke create joke

// GetJokeByID get joke by its ID
func GetJokeByID(id int) (*Joke, error) {
	var joke Joke
	if err := Config.Connection.QueryRow(context.Background(), "SELECT * FROM jokes WHERE ID=$1", id).Scan(&joke.ID, &joke.Joke, &joke.CreatedAt); err != nil {
		fmt.Println("Error occur while finding joke: ", err)
		return nil, err
	}
	fmt.Printf("Joke with id=%v is %v\n", id, joke)
	return &joke, nil
}

// GetRandomJoke random joke
func GetRandomJoke() (*Joke, error) {
	var joke Joke
	if err := Config.Connection.QueryRow(context.Background(), "SELECT * FROM jokes OFFSET floor(random() * (SELECT COUNT(*) FROM jokes)) LIMIT 1").Scan(&joke.ID, &joke.Joke, &joke.CreatedAt); err != nil {
		fmt.Println("Error occured while finding joke: ", err)
		return nil, err
	}
	fmt.Printf("joke #%v\n", joke)
	return &joke, nil
}

// UpdateJoke

// DeleteJoke
