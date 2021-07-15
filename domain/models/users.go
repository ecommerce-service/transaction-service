package models

import (
	"database/sql"
	"time"
)

type Users struct {
	id            string
	firstName     string
	lastName      string
	email         string
	userName      string
	password      string
	address       sql.NullString
	phoneNumber   string
	depositAmount sql.NullFloat64
	roleId        int64
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     sql.NullTime

	Role *Roles
}

func (model *Users) Id() string {
	return model.id
}

func (model *Users) SetId(id string) *Users {
	model.id = id

	return model
}

func (model *Users) FirstName() string {
	return model.firstName
}

func (model *Users) SetFirstName(firstName string) *Users {
	model.firstName = firstName

	return model
}

func (model *Users) LastName() string {
	return model.lastName
}

func (model *Users) SetLastName(lastName string) *Users {
	model.lastName = lastName

	return model
}

func (model *Users) Email() string {
	return model.email
}

func (model *Users) SetEmail(email string) *Users {
	model.email = email

	return model
}

func (model *Users) UserName() string {
	return model.userName
}

func (model *Users) SetUserName(userName string) *Users {
	model.userName = userName

	return model
}

func (model *Users) Password() string {
	return model.password
}

func (model *Users) SetPassword(password string) *Users {
	model.password = password

	return model
}

func (model *Users) Address() sql.NullString {
	return model.address
}

func (model *Users) SetAddress(address string) *Users {
	model.address.String = address

	return model
}

func (model *Users) PhoneNumber() string {
	return model.phoneNumber
}

func (model *Users) SetPhoneNumber(phoneNumber string) *Users {
	model.phoneNumber = phoneNumber

	return model
}

func (model *Users) DepositAmount() sql.NullFloat64 {
	return model.depositAmount
}

func (model *Users) SetDepositAmount(depositAmount float64) *Users {
	model.depositAmount.Float64 = depositAmount

	return model
}

func (model *Users) RoleId() int64 {
	return model.roleId
}

func (model *Users) SetRoleId(roleId int64) *Users {
	model.roleId = roleId

	return model
}

func (model *Users) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Users) SetCreatedAt(createdAt time.Time) *Users {
	model.createdAt = createdAt

	return model
}

func (model *Users) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Users) SetUpdatedAt(updatedAt time.Time) *Users {
	model.updatedAt = updatedAt

	return model
}

func (model *Users) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Users) SetDeletedAt(deletedAt time.Time) *Users {
	model.deletedAt.Time = deletedAt

	return model
}

func NewUserModel() *Users {
	return &Users{}
}

const (
	UserSelectStatement = `SELECT u.id,u.first_name,u.last_name,u.email,u.username,u.password,u.address,u.phone_number,` +
		`u.deposit_amount,u.created_at,u.updated_at,r.id,r.name FROM users u`
	UserSelectCountStatement  = `SELECT COUNT(u.id) FROM users u `
	UserJoinStatement         = `INNER JOIN roles r ON r.id = u.role_id`
	UserDefaultWhereStatement = `WHERE u.deleted_at IS NULL`
)

func (model *Users) ScanRows(rows *sql.Rows) (interface{}, error) {
	model.Role = NewRoleModel()
	err := rows.Scan(&model.id, &model.firstName, &model.lastName, &model.email, &model.userName, &model.password, &model.address, &model.phoneNumber, &model.depositAmount, &model.createdAt,
		&model.updatedAt, &model.Role.id, &model.Role.name)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Users) ScanRow(row *sql.Row) (interface{}, error) {
	model.Role = NewRoleModel()
	err := row.Scan(&model.id, &model.firstName, &model.lastName, &model.email, &model.userName, &model.password, &model.address, &model.phoneNumber, &model.depositAmount, &model.createdAt,
		&model.updatedAt, &model.Role.id, &model.Role.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
