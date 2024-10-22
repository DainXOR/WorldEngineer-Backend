package db

import (
	"dainxor/we/configs"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"
	"strconv"
)

func GetUserByID(id string) types.Result[models.UserDB, models.ErrorResponse] {
	var user models.UserDB
	configs.DB.First(&user, id)

	err := models.ErrorNotFound(
		"User not found",
		"User with ID "+id+" not found",
	)
	return types.ResultOf(user, err, user.ID != 0)
}
func GetUserByEmail(email string) types.Result[models.UserDB, models.ErrorResponse] {
	var user models.UserDB
	configs.DB.Where("email = ?", email).First(&user)

	err := models.ErrorNotFound(
		"User not found",
		"User with email "+email+" not found",
	)

	return types.ResultOf(user, err, user.ID != 0)
}
func GetUserByNameTag(nameTag string) types.Result[models.UserDB, models.ErrorResponse] {
	var user models.UserDB
	configs.DB.Where("name_tag = ?", nameTag).First(&user)

	err := models.ErrorNotFound(
		"User not found",
		"User with name tag "+nameTag+" not found",
	)

	return types.ResultOf(user, err, user.ID != 0)
}

func GetAllUsers() types.Result[[]models.UserDB, models.ErrorResponse] {
	var users []models.UserDB
	configs.DB.Find(&users)

	err := models.ErrorNotFound(
		"No users found",
	)

	return types.ResultOf(users, err, len(users) > 0)
}
func GetAllFilteredUsers(predicate types.Predicate[models.UserDB]) types.Result[[]models.UserDB, models.ErrorResponse] {
	result := GetAllUsers()

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
func GetAllUsersByUsername(username string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return GetAllFilteredUsers(func(u models.UserDB) bool { return u.Username == username })
}
func GetAllUsersByStatusID(id string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return GetAllFilteredUsers(func(u models.UserDB) bool {
		intID, err := strconv.Atoi(id)
		return err != nil && u.IDStatus == uint(intID)
	})
}

func CreateUser(user models.UserCreate) types.Result[models.UserDB, models.ErrorResponse] {
	if GetUserByEmail(user.Email).IsOk() {
		return types.ResultErr[models.UserDB](models.Error(
			types.Http.Conflict(),
			"conflict",
			"Email is already in use",
		))
	}

	var newUser models.UserDB
	newUser.NameTag = utils.GenerateUserTag(user.Username)
	newUser.Username = user.Username
	newUser.Email = user.Email
	newUser.CreatedAt = configs.DB.NowFunc()
	newUser.UpdatedAt = configs.DB.NowFunc()
	newUser.IDStatus = 1
	newUser.StatusTimeStamp = configs.DB.NowFunc()

	configs.DB.Create(&newUser)

	err := models.Error(
		types.Http.InternalServerError(),
		"internal",
		"User not created",
	)

	return types.ResultOf(newUser, err, newUser.ID != 0)
}

func UpdateUserByID(id string, user models.UserUpdate) types.Result[models.UserDB, models.ErrorResponse] {
	result := GetUserByID(id)

	if result.IsErr() {
		return result
	} else {
		result := result.Value()
		configs.DB.Model(&result).Updates(user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not updated",
		)

		return types.ResultOf(result, err, result.ID != 0)
	}
}
func UpdateUserByEmail(email string, user models.UserUpdate) types.Result[models.UserDB, models.ErrorResponse] {
	result := GetUserByEmail(email)

	if result.IsErr() {
		return result
	} else {
		u := result.Value()
		configs.DB.Model(&u).Updates(user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not updated",
		)

		return types.ResultOf(u, err, u.ID != 0)
	}
}
func UpdateUserByNameTag(nameTag string, user models.UserUpdate) types.Result[models.UserDB, models.ErrorResponse] {
	result := GetUserByNameTag(nameTag)

	if result.IsErr() {
		return result
	} else {
		u := result.Value()
		configs.DB.Model(&u).Updates(user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not updated",
		)

		return types.ResultOf(u, err, u.ID != 0)
	}
}

func UpdateAllUsers(user models.UserUpdate) types.Result[[]models.UserDB, models.ErrorResponse] {
	configs.DB.Model(&models.UserDB{}).Updates(user)

	return GetAllUsers()
}
func UpdateAllFilteredUsers(user models.UserUpdate, predicate types.Predicate[models.UserDB]) types.Result[[]models.UserDB, models.ErrorResponse] {
	result := GetAllFilteredUsers(predicate)

	if result.IsErr() {
		return result
	} else {
		var errored []models.UserDB
		users := result.Value()

		for _, u := range users {
			configs.DB.Model(&u).Updates(user)

			if u.ID == 0 {
				errored = append(errored, u)
			}
		}

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Users not updated",
			types.Reduce(errored, func(r string, u models.UserDB) string { return r + string(u.ID) }, ""),
		)

		return types.ResultOf(users, err, len(errored) == 0)
	}
}
func UpdateAllUsersByUsername(username string, user models.UserUpdate) types.Result[[]models.UserDB, models.ErrorResponse] {
	return UpdateAllFilteredUsers(user, func(u models.UserDB) bool { return u.Username == username })
}
func UpdateAllUsersByStatusID(id string, user models.UserUpdate) types.Result[[]models.UserDB, models.ErrorResponse] {
	return UpdateAllFilteredUsers(user, func(u models.UserDB) bool {
		intID, err := strconv.Atoi(id)
		return err != nil && u.IDStatus == uint(intID)
	})
}

func DeleteUserByID(id string) types.Result[models.UserDB, models.ErrorResponse] {
	result := GetUserByID(id)

	if result.IsErr() {
		return result
	} else {
		user := result.Value()
		configs.DB.Delete(&user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not deleted",
		)

		return types.ResultOf(user, err, user.DeletedAt.Valid)
	}
}
func DeleteUserByEmail(email string) types.Result[models.UserDB, models.ErrorResponse] {
	result := GetUserByEmail(email)

	if result.IsErr() {
		return result
	} else {
		user := result.Value()

		configs.DB.Delete(&user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not deleted",
		)

		return types.ResultOf(user, err, user.DeletedAt.Valid)
	}
}
func DeleteUserByNameTag(nameTag string) types.Result[models.UserDB, models.ErrorResponse] {
	result := GetUserByNameTag(nameTag)

	if result.IsErr() {
		return result
	} else {
		user := result.Value()

		configs.DB.Delete(&user)

		err := models.Error(
			types.Http.InternalServerError(),
			"internal",
			"User not deleted",
		)

		return types.ResultOf(user, err, user.DeletedAt.Valid)
	}
}

func DeleteAllUsers() types.Result[[]models.UserDB, models.ErrorResponse] {
	result := GetAllUsers()

	if result.IsErr() {
		return result
	}

	configs.DB.Delete(&models.UserDB{})

	if GetAllUsers().IsErr() {
		return types.ResultOk[[]models.UserDB, models.ErrorResponse](result.Value())
	} else {
		return types.ResultErr[[]models.UserDB](models.Error(
			types.Http.InternalServerError(),
			"internal",
			"Users not deleted",
		))
	}
}
func DeleteAllFilteredUsers(predicate types.Predicate[models.UserDB]) types.Result[[]models.UserDB, models.ErrorResponse] {
	result := GetAllFilteredUsers(predicate)

	if result.IsErr() {
		return result
	}

	errored := []models.UserDB{}

	for _, u := range result.Value() {
		configs.DB.Delete(&u)

		if !u.DeletedAt.Valid {
			errored = append(errored, u)
		}
	}

	err := models.Error(
		types.Http.InternalServerError(),
		"internal",
		"Users not deleted",
		types.Reduce(errored, func(r string, u models.UserDB) string { return r + string(u.ID) }, ""),
	)

	return types.ResultOf(result.Value(), err, len(errored) == 0)
}
func DeleteAllUsersByUsername(username string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return DeleteAllFilteredUsers(func(u models.UserDB) bool { return u.Username == username })
}
func DeleteAllUsersByStatusID(id string) types.Result[[]models.UserDB, models.ErrorResponse] {
	return DeleteAllFilteredUsers(func(u models.UserDB) bool {
		intID, err := strconv.Atoi(id)
		return err != nil && u.IDStatus == uint(intID)
	})
}
