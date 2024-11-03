package db

import (
	"dainxor/we/base/configs"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"

	"math/rand"
	"strconv"
)

type userType struct{}

var User userType

func (userType) GetUserByID(id string) types.Result[models.UserDB, models.ErrorResponse] {
	var user models.UserDB
	configs.DataBase.First(&user, id)

	err := models.ErrorNotFound(
		"User not found",
		"User with ID "+id+" not found",
	)
	return types.ResultOf(user, err, user.ID != 0)
}
func (userType) GetUserByEmail(email string) types.Result[models.UserDB, models.ErrorResponse] {
	var user models.UserDB
	configs.DataBase.Where("email = ?", email).First(&user)

	err := models.ErrorNotFound(
		"User not found",
		"User with email "+email+" not found",
	)

	return types.ResultOf(user, err, user.ID != 0)
}
func (userType) GetUserByNameTag(nameTag string) types.Result[models.UserDB, models.ErrorResponse] {
	var user models.UserDB
	configs.DataBase.Where("name_tag = ?", nameTag).First(&user)

	err := models.ErrorNotFound(
		"User not found",
		"User with name tag "+nameTag+" not found",
	)

	return types.ResultOf(user, err, user.ID != 0)
}

func (userType) GetAllUsers() types.Result[[]models.UserDB, models.ErrorResponse] {
	var users []models.UserDB
	configs.DataBase.Find(&users)

	err := models.ErrorNotFound(
		"No users found",
	)

	return types.ResultOf(users, err, len(users) > 0)
}
func (userType) GetAllFilteredUsers(predicate types.Predicate[models.UserDB]) types.Result[[]models.UserDB, models.ErrorResponse] {
	result := User.GetAllUsers()

	if result.IsOk() {
		users := result.Value()
		users = types.Filter(users, predicate)

		err := models.ErrorNotFound(
			"No users found",
		)

		return types.ResultOf(users, err, len(users) > 0)
	} else {
		return result
	}
}
func (userType) GetAllUsersByUsername(username string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return User.GetAllFilteredUsers(func(u models.UserDB) bool { return u.Username == username })
}
func (userType) GetAllUsersByStatusID(id string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return User.GetAllFilteredUsers(func(u models.UserDB) bool {
		intID, err := strconv.Atoi(id)
		return err != nil && u.IDStatus == uint(intID)
	})
}

func (userType) CreateUser(user models.UserCreate) types.Result[models.UserDB, models.ErrorResponse] {
	if User.GetUserByEmail(user.Email).IsOk() {
		return types.ResultErr[models.UserDB](models.Error(
			types.Http.Conflict(),
			"conflict",
			"Email is already in use",
		))
	}

	var newUser models.UserDB
	newUser.NameTag = user.NameTag
	newUser.Username = user.Username
	newUser.Email = user.Email
	newUser.CreatedAt = configs.DataBase.NowFunc()
	newUser.UpdatedAt = configs.DataBase.NowFunc()
	newUser.IDStatus = 1
	newUser.StatusTimeStamp = configs.DataBase.NowFunc()

	configs.DataBase.Create(&newUser)

	err := models.Error(
		types.Http.InternalServerError(),
		"internal",
		"User not created",
	)

	return types.ResultOf(newUser, err, newUser.ID != 0)
}

func (userType) UpdateUserByID(id string, user models.UserUpdate) types.Result[models.UserDB, models.ErrorResponse] {
	result := User.GetUserByID(id)

	if result.IsErr() {
		return result
	} else {
		result := result.Value()
		configs.DataBase.Model(&result).Updates(user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not updated",
		)

		return types.ResultOf(result, err, result.ID != 0)
	}
}
func (userType) UpdateUserByEmail(email string, user models.UserUpdate) types.Result[models.UserDB, models.ErrorResponse] {
	result := User.GetUserByEmail(email)

	if result.IsErr() {
		return result
	} else {
		u := result.Value()
		configs.DataBase.Model(&u).Updates(user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not updated",
		)

		return types.ResultOf(u, err, u.ID != 0)
	}
}
func (userType) UpdateUserByNameTag(nameTag string, user models.UserUpdate) types.Result[models.UserDB, models.ErrorResponse] {
	result := User.GetUserByNameTag(nameTag)

	if result.IsErr() {
		return result
	} else {
		u := result.Value()
		configs.DataBase.Model(&u).Updates(user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not updated",
		)

		return types.ResultOf(u, err, u.ID != 0)
	}
}

func (userType) UpdateAllUsers(user models.UserUpdate) types.Result[[]models.UserDB, models.ErrorResponse] {
	configs.DataBase.Model(&models.UserDB{}).Updates(user)

	return User.GetAllUsers()
}
func (userType) UpdateAllFilteredUsers(user models.UserUpdate, predicate types.Predicate[models.UserDB]) types.Result[[]models.UserDB, models.ErrorResponse] {
	result := User.GetAllFilteredUsers(predicate)

	if result.IsErr() {
		return result
	} else {
		var errored []models.UserDB
		users := result.Value()

		for _, u := range users {
			configs.DataBase.Model(&u).Updates(user)

			if u.ID == 0 {
				errored = append(errored, u)
			}
		}

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Users not updated",
			types.Reduce(errored, func(r string, u models.UserDB) string { return r + strconv.Itoa(int(u.ID)) }, ""),
		)

		return types.ResultOf(users, err, len(errored) == 0)
	}
}
func (userType) UpdateAllUsersByUsername(username string, user models.UserUpdate) types.Result[[]models.UserDB, models.ErrorResponse] {
	return User.UpdateAllFilteredUsers(user, func(u models.UserDB) bool { return u.Username == username })
}
func (userType) UpdateAllUsersByStatusID(id string, user models.UserUpdate) types.Result[[]models.UserDB, models.ErrorResponse] {
	return User.UpdateAllFilteredUsers(user, func(u models.UserDB) bool {
		intID, err := strconv.Atoi(id)
		return err != nil && u.IDStatus == uint(intID)
	})
}

func (userType) DeleteUserByID(id string) types.Result[models.UserDB, models.ErrorResponse] {
	result := User.GetUserByID(id)

	if result.IsErr() {
		return result
	} else {
		user := result.Value()
		configs.DataBase.Delete(&user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not deleted",
		)

		return types.ResultOf(user, err, user.DeletedAt.Valid)
	}
}
func (userType) DeleteUserByEmail(email string) types.Result[models.UserDB, models.ErrorResponse] {
	result := User.GetUserByEmail(email)

	if result.IsErr() {
		return result
	} else {
		user := result.Value()

		configs.DataBase.Delete(&user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not deleted",
		)

		return types.ResultOf(user, err, user.DeletedAt.Valid)
	}
}
func (userType) DeleteUserByNameTag(nameTag string) types.Result[models.UserDB, models.ErrorResponse] {
	result := User.GetUserByNameTag(nameTag)

	if result.IsErr() {
		return result
	} else {
		user := result.Value()

		configs.DataBase.Delete(&user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not deleted",
		)

		return types.ResultOf(user, err, user.DeletedAt.Valid)
	}
}

func (userType) DeleteAllUsers() types.Result[[]models.UserDB, models.ErrorResponse] {
	result := User.GetAllUsers()

	if result.IsErr() {
		return result
	}

	configs.DataBase.Delete(&models.UserDB{})

	if User.GetAllUsers().IsErr() {
		return types.ResultOk[[]models.UserDB, models.ErrorResponse](result.Value())
	} else {
		return types.ResultErr[[]models.UserDB](models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Users not deleted",
		))
	}
}
func (userType) DeleteAllFilteredUsers(predicate types.Predicate[models.UserDB]) types.Result[[]models.UserDB, models.ErrorResponse] {
	result := User.GetAllFilteredUsers(predicate)

	if result.IsErr() {
		return result
	}

	errored := []models.UserDB{}

	for _, u := range result.Value() {
		configs.DataBase.Delete(&u)

		if !u.DeletedAt.Valid {
			errored = append(errored, u)
		}
	}

	err := models.Error(
		types.Http.InternalServerError(),
		"internal",
		"Users not deleted",
		types.Reduce(errored, func(r string, u models.UserDB) string { return r + strconv.Itoa(int(u.ID)) }, ""),
	)

	return types.ResultOf(result.Value(), err, len(errored) == 0)
}
func (userType) DeleteAllUsersByUsername(username string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return User.DeleteAllFilteredUsers(func(u models.UserDB) bool { return u.Username == username })
}
func (userType) DeleteAllUsersByStatusID(id string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return User.DeleteAllFilteredUsers(func(u models.UserDB) bool {
		intID, err := strconv.Atoi(id)
		return err != nil && u.IDStatus == uint(intID)
	})
}

func (userType) GenerateUsername() string {
	username := utils.Usernames()[rand.Intn(utils.UsernamesCount())]

	return username
}
func (userType) GenerateNameTag(username string) string {
	randomNumber := utils.FillZeros(rand.Intn(99999), 5)
	nameTag := username + "#" + randomNumber

	return nameTag
}
func (userType) AvailableUsername(username string) bool {
	if len(username) < 2 || len(username) > 24 {
		return false
	}

	// ^[^'"¨´`#\\]{2,24}$
	// ^[a-zA-Z0-9_\-+*|$&<>~!¡[\]@?¿^.:,;]{2,24}$
	// namePattern := regexp.MustCompile("^[^'\"¨´`#\\]{2,24}$")
	// followsPattern := namePattern.MatchString(nameTag)
	err := Mail.VerifyEmailAddress(username)

	return err.IsPresent() // && followsPattern
}
func (userType) AvailableNameTag(nameTag string) bool {
	return User.GetUserByNameTag(nameTag).IsErr()
}

func (userType) CreateNameTag(username string) types.Result[string, models.ErrorResponse] {
	maxAttemps := 10
	nameTag := User.GenerateNameTag(username)

	for !User.AvailableNameTag(nameTag) || maxAttemps > 0 {
		nameTag = User.GenerateNameTag(username)
		maxAttemps--
	}

	if maxAttemps == 0 {
		return types.ResultErr[string](models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Failed to generate name tag",
		))
	}

	return types.ResultOk[string, models.ErrorResponse](nameTag)
}
