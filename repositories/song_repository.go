package repositories
import "gorm.io/gorm"

type SongRepository interface {}
type songRepository struct { db *gorm.DB }

func NewSongRepository(db *gorm.DB) SongRepository {
	return &songRepository{db}
}