package interfaces

import (
	"homework_bot/internal/domain"
	"time"
)

type IHomeworkRepository interface {
	Create(homework domain.Homework) (int, error)
	GetByTags(tags []string) ([]domain.HomeworkToGet, error)
	GetByName(name string) ([]domain.HomeworkToGet, error)
	GetByWeek() ([]domain.HomeworkToGet, error)
	GetById(id int) (domain.HomeworkToGet, error)
	GetAll() ([]domain.HomeworkToGet, error)
	GetByToday() ([]domain.HomeworkToGet, error)
	GetByTomorrow() ([]domain.HomeworkToGet, error)
	GetByDate(date time.Time) ([]domain.HomeworkToGet, error)
	Update(homeworkToUpdate domain.HomeworkToUpdate) (domain.Homework, error)
	Delete(id int) error
}

type IUserRepository interface {
	Create(user domain.User) error
	Update(user domain.User) error
	GetByUsername(username string) (domain.User, error)
	UpdateTimes(_times uint64, _username string) error
	FetchNewestAddress() ([]domain.User, error)
	NotifyTronAddress() ([]domain.User, error)
	NotifyEthereumAddress() ([]domain.User, error)
	BindTronAddress(_address string, _username string) error
	BindEthereumAddress(_address string, _username string) error
	BindChat(_associates string, _username string) error
	DisableTronAddress(_address string) error
}
