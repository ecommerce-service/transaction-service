package usecases

import (
	"booking-car/domain/models"
	"booking-car/domain/requests"
	"booking-car/domain/usecases"
	"booking-car/domain/view_models"
	"booking-car/pkg/functioncaller"
	"booking-car/pkg/hashing"
	"booking-car/pkg/logruslogger"
	"booking-car/pkg/messages"
	"booking-car/repository/commands"
	"booking-car/repository/queries"
	"errors"
	"time"
)

type UserUseCase struct {
	*UseCaseContract
}

func NewUserUseCase(useCaseContract *UseCaseContract) usecases.IUserUseCase {
	return &UserUseCase{UseCaseContract: useCaseContract}
}

func (uc UserUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.UserVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	users, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-browse")
		return res, pagination, err
	}
	for _, user := range users.([]*models.Users) {
		res = append(res, view_models.NewUserVm(user))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc UserUseCase) GetByID(id string) (res view_models.UserVm, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())

	user, err := q.ReadBy("u.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-readByID")
		return res, err
	}
	res = view_models.NewUserVm(user.(*models.Users))

	return res, nil
}

func (uc UserUseCase) Edit(req *requests.UserEditRequest, id string) (res string, err error) {
	now := time.Now().UTC()
	var password string

	uc.Config.DB.GetDbInstance()

	isDuplicate, err := uc.CheckDuplication(req.Email, req.Username, req.PhoneNumber, id)
	if err != nil && isDuplicate {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-checkDuplication")
		return res, err
	}

	model := models.NewUserModel().SetFirstName(req.FirstName).SetLastName(req.LastName).SetEmail(req.Email).SetUserName(req.Username).
		SetAddress(req.Address).SetPhoneNumber(req.PhoneNumber).SetRoleId(req.RoleID).SetUpdatedAt(now).SetId(id).SetPassword(password)
	if req.Password != "" {
		password, _ := hashing.HashAndSalt(req.Password)
		model.SetPassword(password)
	}
	cmd := commands.NewUserCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-user-edit")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) Add(req *requests.UserAddRequest) (res string, err error) {
	now := time.Now().UTC()

	isDuplicate, err := uc.CheckDuplication(req.Email, req.Username, req.PhoneNumber, "")
	if err != nil && isDuplicate {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-checkDuplication")
		return res, err
	}

	password, _ := hashing.HashAndSalt(req.Password)
	model := models.NewUserModel().SetFirstName(req.FirstName).SetLastName(req.LastName).SetEmail(req.Email).SetUserName(req.Username).
		SetPassword(password).SetAddress(req.Address).SetPhoneNumber(req.PhoneNumber).SetRoleId(req.RoleID).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewUserCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-user-add")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("u.id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-CountByID")
		return err
	}
	if count > 0 {
		model := models.NewUserModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewUserCommand(uc.Config.DB.GetDbInstance(), model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-user-delete")
			return err
		}
	}

	return nil
}

func (uc UserUseCase) Count(search string) (res int, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())

	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-count")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-countBy")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) CheckDuplication(email, username, phoneNumber, id string) (bool, error) {
	count, err := uc.CountBy("u.email", "=", id, email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-countByEmail")
		return true, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.EmailAlreadyExist, functioncaller.PrintFuncName(), "query-user-countByEmail")
		return true, errors.New(messages.EmailAlreadyExist)
	}

	count, err = uc.CountBy("u.username", "=", id, username)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-countByUserName")
		return true, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.UserNameAlreadyExist, functioncaller.PrintFuncName(), "query-user-countByUserName")
		return true, errors.New(messages.UserNameAlreadyExist)
	}

	count, err = uc.CountBy("u.phone_number", "=", id, phoneNumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-countByPhoneNumber")
		return true, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.PhoneAlreadyExist, functioncaller.PrintFuncName(), "query-user-countByPhoneNumber")
		return true, errors.New(messages.PhoneAlreadyExist)
	}

	return false, err
}
