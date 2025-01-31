package controller

import (
	"github.com/chinlee1523/go-admin/modules/auth"
	"github.com/chinlee1523/go-admin/modules/db"
	"github.com/chinlee1523/go-admin/modules/db/dialect"
)

func NewManager(dataList map[string][]string) {

	if dataList["name"][0] == "" ||
		dataList["username"][0] == "" ||
		dataList["password"][0] == "" {
		panic("账号密码不能为空")
	}

	// 更新管理员表
	id, _ := db.Table("yunyun_users").
		Insert(dialect.H{
			"username": dataList["username"][0],
			"password": auth.EncodePassword([]byte(dataList["password"][0])),
			"name":     dataList["name"][0],
			"avatar":   dataList["avatar"][0],
		})

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			_, _ = db.Table("yunyun_role_users").
				Insert(dialect.H{
					"role_id": dataList["role_id[]"][i],
					"user_id": id,
				})
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			_, _ = db.Table("yunyun_user_permissions").
				Insert(dialect.H{
					"permission_id": dataList["permission_id[]"][i],
					"user_id":       id,
				})
		}
	}
}

func EditManager(dataList map[string][]string) {

	if dataList["name"][0] == "" ||
		dataList["username"][0] == "" ||
		dataList["password"][0] == "" {
		panic("账号密码不能为空")
	}

	// 更新管理员表
	_, _ = db.Table("yunyun_users").
		Where("id", "=", dataList["id"][0]).
		Update(dialect.H{
			"username": dataList["username"][0],
			"password": auth.EncodePassword([]byte(dataList["password"][0])),
			"name":     dataList["name"][0],
			"avatar":   dataList["avatar"][0],
		})

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			checkRole, _ := db.Table("yunyun_role_users").
				Where("role_id", "=", dataList["role_id[]"][i]).
				Where("user_id", "=", dataList["id"][0]).
				First()

			if checkRole == nil {
				_, _ = db.Table("yunyun_role_users").
					Insert(dialect.H{
						"role_id": dataList["role_id[]"][i],
						"user_id": dataList["id"][0],
					})
			}
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			checkPermission, _ := db.Table("yunyun_user_permissions").
				Where("permission_id", "=", dataList["permission_id[]"][i]).
				Where("user_id", "=", dataList["id"][0]).
				First()

			if checkPermission == nil {
				_, _ = db.Table("yunyun_user_permissions").
					Insert(dialect.H{
						"permission_id": dataList["permission_id[]"][i],
						"user_id":       dataList["id"][0],
					})
			}
		}
	}
}

func NewRole(dataList map[string][]string) {
	// 更新管理员角色表
	id, _ := db.Table("yunyun_roles").
		Insert(dialect.H{
			"name": dataList["name"][0],
			"slug": dataList["slug"][0],
		})

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			_, _ = db.Table("yunyun_role_permissions").
				Insert(dialect.H{
					"permission_id": dataList["permission_id[]"][i],
					"role_id":       id,
				})
		}
	}
}

func EditRole(dataList map[string][]string) {
	// 更新管理员角色表
	_, _ = db.Table("yunyun_roles").
		Where("id", "=", dataList["id"][0]).
		Update(dialect.H{
			"name": dataList["name"][0],
			"slug": dataList["slug"][0],
		})

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			checkPermission, _ := db.Table("yunyun_role_permissions").
				Where("permission_id", "=", dataList["permission_id[]"][i]).
				Where("role_id", "=", dataList["id"][0]).
				First()

			if checkPermission == nil {
				_, _ = db.Table("yunyun_role_permissions").
					Insert(dialect.H{
						"permission_id": dataList["permission_id[]"][i],
						"role_id":       dataList["id"][0],
					})
			}
		}
	}
}
