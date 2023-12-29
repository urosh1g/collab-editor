package services

import "github.com/urosh1g/collab-editor/repositories"

var (
	FilesService    *FileService
	ProjectsService *ProjectService
	UsersService    *UserService
)

func init() {
	FilesService = NewFileService(repositories.FileRepo, repositories.ProjectRepo)
	ProjectsService = NewProjectService(repositories.ProjectRepo)
	UsersService = NewUserService(repositories.UserRepo)
}
