package repositories

import (
	"github.com/urosh1g/collab-editor/database"
	"github.com/urosh1g/collab-editor/models"
)

var (
	FileRepo    Repository[models.File]
	ProjectRepo Repository[models.Project]
	UserRepo    Repository[models.User]
)

func init() {
	FileRepo = NewFileRepository(database.Db)
	ProjectRepo = NewProjectRepository(database.Db)
	UserRepo = NewUserRepository(database.Db)
}
