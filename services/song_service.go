package services
import "open-music/repositories"

type SongService interface {}
type songService struct { repo repositories.SongRepository }

func NewSongService(repo repositories.SongRepository) SongService {
	return &songService{repo}
}