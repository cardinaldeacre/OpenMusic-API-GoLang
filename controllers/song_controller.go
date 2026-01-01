package controllers
import (
	"open-music/services"
	"github.com/gofiber/fiber/v2"
)

type SongController struct { Service services.SongService }

func NewSongController(s services.SongService) *SongController {
	return &SongController{Service: s}
}