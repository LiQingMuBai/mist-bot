package repositories

import (
	"github.com/jmoiron/sqlx"
	"homework_bot/internal/domain"
	"log"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user domain.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO tg_users (user_id, username,amount,tron_amount,tron_address, eth_address,eth_amount, associates) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	row := tx.QueryRow(query, user.UserID, user.Username, user.Amount, user.TronAmount, user.TronAddress, user.EthAddress, user.EthAmount, user.Associates)
	//row := tx.QueryRow(query, 1, user.Username, user.TronAmount, user.TronAddress, user.EthAddress, user.EthAmount, user.Associates)
	if row.Err() != nil {

		log.Println("add err:", row.Err())
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *UserRepository) Update(user domain.User) error {
	query := "UPDATE tg_users SET associates = $1, tron_amount = $2 WHERE username = $3"
	_, err := r.db.Exec(query, user.Associates, user.TronAmount, user.Username)
	return err
}

func (r *UserRepository) UpdateTimes(_times uint64, _username string) error {
	query := "UPDATE tg_users SET times = ? WHERE username = ?"
	_, err := r.db.Exec(query, _times, _username)
	return err
}

//associates VARCHAR(255),
//amount VARCHAR(255) ,
//tron_amount VARCHAR(255),
//tron_address VARCHAR(50),
//eth_address VARCHAR(50),
//eth_amount VARCHAR(255),

func (r *UserRepository) GetByUsername(_username string) (domain.User, error) {
	//query := "SELECT id,username,associates,amount, tron_amount,tron_address,eth_address,eth_amount FROM tg_users WHERE username = ? "
	//var user domain.User
	//
	//err := r.db.Get(&user, query, _username)

	jason := domain.User{}
	err := r.db.Get(&jason, "SELECT  user_id,username,amount,associates, tron_amount,tron_address,eth_address,eth_amount,times FROM tg_users WHERE username=?", _username)

	log.Println(err)
	return jason, err
}
