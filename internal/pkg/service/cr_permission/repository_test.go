package cr_permission

import (
	"github.com/stretchr/testify/assert"
	"github.com/tegarsubkhan236/redis-jwt-auth/internal/pkg/model"
	"github.com/tegarsubkhan236/redis-jwt-auth/internal/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCrPermissionRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.Nil(t, err, "failed to connect to the database")

	err = db.AutoMigrate(&model.CrPermission{})
	assert.Nil(t, err, "failed to migrate database")

	t.Run("Create CrPermission", func(t *testing.T) {
		util.RunInTransaction(t, db, func(tx *gorm.DB) {
			repo := NewCrPermissionRepository(tx)
			data := &model.CrPermission{
				Name: "admin",
			}
			err := repo.Create(data)
			assert.Nil(t, err)
			assert.NotZero(t, data.ID)
		})
	})

	t.Run("Fetch All CrPermissions", func(t *testing.T) {
		util.RunInTransaction(t, db, func(tx *gorm.DB) {
			repo := NewCrPermissionRepository(tx)
			data1 := &model.CrPermission{Name: "admin"}
			data2 := &model.CrPermission{Name: "user"}
			err := repo.Create(data1)
			assert.Nil(t, err)
			err = repo.Create(data2)
			assert.Nil(t, err)

			items, err := repo.FindAll()
			assert.Nil(t, err)
			assert.Len(t, items, 2)
		})
	})

	t.Run("Find CrPermission By ID", func(t *testing.T) {
		util.RunInTransaction(t, db, func(tx *gorm.DB) {
			repo := NewCrPermissionRepository(tx)
			data := &model.CrPermission{
				Name: "admin",
			}
			err := repo.Create(data)
			assert.Nil(t, err)

			item, err := repo.FindByID(data.ID)
			assert.Nil(t, err)
			assert.Equal(t, data.ID, item.ID)
		})
	})

	t.Run("Update CrPermission", func(t *testing.T) {
		util.RunInTransaction(t, db, func(tx *gorm.DB) {
			repo := NewCrPermissionRepository(tx)
			data := &model.CrPermission{
				Name: "admin",
			}
			err := repo.Create(data)
			assert.Nil(t, err)

			item, err := repo.FindByID(data.ID)
			assert.Nil(t, err)

			item.Name = "admin updated"
			err = repo.Update(item)
			assert.Nil(t, err)

			updatedItem, err := repo.FindByID(item.ID)
			assert.Nil(t, err)
			assert.Equal(t, "admin updated", updatedItem.Name)
		})
	})

	t.Run("Delete CrPermission", func(t *testing.T) {
		util.RunInTransaction(t, db, func(tx *gorm.DB) {
			repo := NewCrPermissionRepository(tx)
			data := &model.CrPermission{
				Name: "admin",
			}
			err := repo.Create(data)
			assert.Nil(t, err)

			result, err := repo.FindByID(data.ID)
			assert.Nil(t, err)

			err = repo.Delete(result)
			assert.Nil(t, err)

			_, err = repo.FindByID(data.ID)
			assert.NotNil(t, err)
		})
	})
}
