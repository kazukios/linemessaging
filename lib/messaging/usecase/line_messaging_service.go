package usecase

import (
	"linemessaging/lib/messaging/domain/movie/service"
	"linemessaging/lib/messaging/repository"
)

func PublishMessage() {

	// make movie
	movieGetterRepository := repository.NewFeedMovieGetterRepository()
	movieService := service.NewMovieService(movieGetterRepository)
	movie := movieService.GetBroadcastMovie()

	// broadcast movie
	messengerRepository := repository.NewLineMessengerRepository()
	messengerRepository.Broadcast(movie)
}
