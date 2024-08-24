package validator

import (
	"card-game/models"
	"card-game/services"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func CheckUserExist(fl validator.FieldLevel) bool {
	var user models.User
	fillModel(&user, fl)

	return services.ExistUser(user)
}

func fillModel(model any, fl validator.FieldLevel) {
	v := reflect.ValueOf(model).Elem()
	if f := v.FieldByName(fl.FieldName()); f.IsValid() && f.CanSet() {
		f.SetString(fl.Field().String())
	}
}
