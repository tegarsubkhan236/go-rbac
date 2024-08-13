package test

import (
	"gorm.io/gorm"
	"testing"
)

// RunInTransaction menjalankan fungsi pengujian dalam sebuah transaksi database.
//
// Fungsi ini memulai transaksi pada database yang diberikan dan menjalankan fungsi
// pengujian yang disediakan dengan transaksi tersebut. Setelah fungsi pengujian selesai,
// transaksi akan di-rollback untuk memastikan bahwa perubahan pada database tidak
// mempengaruhi pengujian lainnya. Jika terjadi kesalahan saat memulai transaksi atau
// saat rollback, kesalahan tersebut akan dilaporkan ke dalam log pengujian.
//
// Parameter:
//   - t: *testing.T - Objek *testing.T yang digunakan untuk melaporkan kesalahan atau
//     kegagalan dalam pengujian. Ini merupakan objek yang disediakan oleh framework
//     testing Go.
//   - db: *gorm.DB - Objek *gorm.DB yang mewakili koneksi database. Ini adalah objek
//     yang digunakan untuk memulai dan mengelola transaksi database.
//   - testFunc: func(tx *gorm.DB) - Fungsi yang akan dijalankan dalam konteks transaksi.
//     Fungsi ini menerima objek *gorm.DB yang mewakili transaksi yang aktif dan dapat
//     digunakan untuk melakukan operasi database selama transaksi.
//
// Contoh Penggunaan:
//
//	func TestExample(t *testing.T) {
//	    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
//	    if err != nil {
//	        t.Fatalf("Failed to connect to the database: %v", err)
//	    }
//
//	    // Migrasi schema database
//	    db.AutoMigrate(&MyModel{})
//
//	    test_helpers.RunInTransaction(t, db, func(tx *gorm.DB) {
//	        // Buat repository dengan transaksi aktif
//	        repo := NewMyRepository(tx)
//
//	        // Operasi database menggunakan repo
//	        err := repo.Create(&MyModel{Name: "test"})
//	        if err != nil {
//	            t.Fatalf("Failed to create model: %v", err)
//	        }
//
//	        item, err := repo.FindByID(1)
//	        if err != nil {
//	            t.Fatalf("Failed to find model: %v", err)
//	        }
//
//	        if item.Name != "test" {
//	            t.Errorf("Expected name to be 'test', got '%s'", item.Name)
//	        }
//	    })
//	}
func RunInTransaction(t *testing.T, db *gorm.DB, testFunc func(tx *gorm.DB)) {
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("Failed to begin transaction: %v", tx.Error)
	}

	defer func() {
		if err := tx.Rollback().Error; err != nil {
			t.Errorf("Rollback failed: %v", err)
		}
	}()

	testFunc(tx)
}
