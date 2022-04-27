package service_test

import (
	"linemessaging/lib/messaging/domain/movie/model"
	"linemessaging/lib/messaging/domain/movie/service"
	"linemessaging/lib/messaging/repository"
	"reflect"
	"testing"
)

func TestMovieService(t *testing.T) {
	movieGetterRepository := repository.NewFeedMovieGetterRepository()
	movieService := service.NewMovieService(movieGetterRepository)
	movie := movieService.GetBroadcastMovie()

	t.Log(movie)
	if movie == nil {
		t.Error("movie must not nil.")
	}
}

func TestMovieService_GetGodMovie(t *testing.T) {
	movieGetterRepository := repository.NewFeedMovieGetterRepository()
	movieService := service.NewMovieService(movieGetterRepository)
	type args struct {
		godURLList []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"return movie", args{[]string{"http://www.testtest/x"}}, "http://www.testtest/x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, _ := model.NewMovie("神曲", tt.want, nil)
			if got := movieService.GetGodMovie(tt.args.godURLList); !reflect.DeepEqual(got, want) {
				t.Errorf("MovieService.GetGodMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}
