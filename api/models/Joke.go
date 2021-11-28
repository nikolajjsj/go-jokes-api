package Models

import (
	"api/Config"
	"context"
	"fmt"
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
			rows.Scan(&tmpJoke.ID, &tmpJoke.Title, &tmpJoke.Joke, &tmpJoke.CreatedAt)
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
func GetJokeByID(id int) (*Joke, error) {
	var joke Joke
	if err := Config.Connection.QueryRow(context.Background(), "SELECT * WHERE ID=$1", id).Scan(&joke.ID, &joke.Title, &joke.Joke, &joke.CreatedAt); err != nil {
		fmt.Println("Error occur while finding joke: ", err)
		return nil, err
	}
	fmt.Printf("Joke with id=%v is %v\n", id, joke)
	return &joke, nil
}

// GetJokeByID

// UpdateJoke

// DeleteJoke
