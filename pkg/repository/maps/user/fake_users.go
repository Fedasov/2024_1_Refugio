package user

import "mail/pkg/repository/models"

var FakeUsers = map[uint32]*models.User{
	1: {
		ID:        1,
		FirstName: "Sergey",
		Surname:   "Fedasov",
		Login:     "sergey@mail.ru",
		Password:  "$2a$10$4PcooWbEMRjvdk2cMFumO.ajWaAclawIljtlfu2.2f5/fV8LkgEZe",
		AvatarID:  "image",
	},
	2: {
		ID:        2,
		FirstName: "Ivan",
		Surname:   "Karpov",
		Login:     "ivan@mail.ru",
		Password:  "$2a$10$4PcooWbEMRjvdk2cMFumO.ajWaAclawIljtlfu2.2f5/fV8LkgEZe",
		AvatarID:  "image",
	},
	3: {ID: 3,
		FirstName: "Max",
		Surname:   "Frelich",
		Login:     "max@mail.ru",
		Password:  "$2a$10$4PcooWbEMRjvdk2cMFumO.ajWaAclawIljtlfu2.2f5/fV8LkgEZe",
		AvatarID:  "image",
	},
}
