package service

import (
	"fmt"
	"go-task/client"
	"go-task/models"
	"log"
)

func getName() (*models.Person, error) {
	person := new(models.Person)
	err := client.Get("https://names.mcquay.me/api/v0", "", nil, person)

	if err != nil {
		return new(models.Person), err
	}

	return person, nil
}

func getJoke(firstName string, lastName string) (*models.JokeResponse, error) {
	joke := new(models.JokeResponse)
	err := client.Get(fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=nerdy", firstName, lastName), "", nil, joke)

	if err != nil {
		return new(models.JokeResponse), err
	}

	return joke, nil
}

func GetOrientedJoke() string {
	personName, personNameError := getName()

	if personNameError != nil {
		log.Print(personNameError)
		return ""
	}

	joke, jokeError := getJoke(personName.FirstName, personName.LastName)

	if jokeError != nil {
		log.Print(jokeError)
		return ""
	}

	return joke.Value.Joke
}
