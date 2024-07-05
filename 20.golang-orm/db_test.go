package golangorm

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestOpenConnection(t *testing.T) {

	db := OpenConnection()

	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	user := &Sample{
		ID:   "3",
		Name: "Roozy",
	}

	err := db.Exec("insert into sample(id,name) values(?,?)", user.ID, user.Name).Error
	assert.Nil(t, err)
}

func TestRawSQL(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	var sample Sample

	err := db.Raw("select id,name from sample where id = ?", "1").Scan(&sample).Error

	assert.Nil(t, err)
	assert.Equal(t, "1", sample.ID)
	assert.Equal(t, "Isro", sample.Name)

	var samples []Sample

	err = db.Raw("select id, name from sample").Scan(&samples).Error

	assert.Nil(t, err)
	assert.Equal(t, 3, len(samples))

	bytes, err := json.Marshal(samples)
	PanicIfError(err)

	fmt.Println(string(bytes))
}

func TestRowsSQL(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	var samples []Sample

	rows, err := db.Raw("select id, name from sample").Rows()
	PanicIfError(err)

	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		PanicIfError(err)
	}

	assert.Equal(t, 3, len(samples))

	fmt.Println(samples)

	bytes, err := json.Marshal(samples)
	PanicIfError(err)

	fmt.Println(string(bytes))

	rows2 := db.Raw("select id, name from sample where id = ?", "1").Row()
	PanicIfError(err)

	sample := Sample{}
	err = rows2.Scan(&sample.ID, &sample.Name)
	PanicIfError(err)

	fmt.Println(sample)

	bytes, err = json.Marshal(sample)
	PanicIfError(err)

	fmt.Println(string(bytes))
}

func TestCreateByModel(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	user := User{
		ID:       "roozy933",
		Password: "roozy123",
		Name: Name{
			FirstName: "Muhamad",
		},
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)

	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestCreateMany(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	users := []User{
		{
			ID:       "roozy668",
			Password: "roozy123",
			Name: Name{
				FirstName: "Muhamad",
			},
		},
		{
			ID:       "roozy669",
			Password: "roozy123",
			Name: Name{
				FirstName: "Muhamad",
			},
		},
	}

	response := db.Create(&users)
	assert.Nil(t, response.Error)

	assert.Equal(t, int64(2), response.RowsAffected)
}

func TestCreateBatch(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	users := []User{}

	for i := 11; i < 20; i++ {
		users = append(users, User{
			ID: strconv.Itoa(i),
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
			Password: "rahasia",
		})
	}

	response := db.Create(&users)

	assert.Equal(t, 9, int(response.RowsAffected))
}

func TestTransactionSuccess(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "101",
			Password: "rahasia",
			Name: Name{
				FirstName: "Roozy",
			},
		}).Error

		PanicIfError(err)

		err = tx.Create(&User{
			ID:       "103",
			Password: "rahasia",
			Name: Name{
				FirstName: "Roozy",
			},
		}).Error

		PanicIfError(err)

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionRollback(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "104",
			Password: "rahasia",
			Name: Name{
				FirstName: "Roozy",
			},
		}).Error

		PanicIfError(err)

		err = tx.Create(&User{
			ID:       "11",
			Password: "rahasia",
			Name: Name{
				FirstName: "Roozy",
			},
		}).Error

		PanicIfError(err)

		return nil
	})

	assert.NotNil(t, err)
}

func TestManualTransaction(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "400", Password: "rahasia", Name: Name{FirstName: "Ahmad"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQueryFirst(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var user = User{}

	err := db.First(&user).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, "1")
	assert.Equal(t, user.Name.FirstName, "User 1")

	user = User{}

	err = db.Last(&user).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, "roozy999")
	assert.Equal(t, user.Name.FirstName, "Muhamad")

	user = User{}

	err = db.Take(&user, "9").Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, "9")
	assert.Equal(t, user.Name.FirstName, "User 9")
}

func TestQueryInlineCondition(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var user = User{}

	result := db.First(&user, "id = ?", "9")

	assert.Nil(t, result.Error)
	assert.Equal(t, user.ID, "9")
	assert.Equal(t, user.Name.FirstName, "User 9")

	user = User{}

	result = db.Take(&user, "id = ?", "9")

	assert.Nil(t, result.Error)
	assert.Equal(t, user.ID, "9")
}

func TestQueryAllObjects(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	var users []User

	result := db.Find(&users, "id in ?", []string{"1", "2", "3", "4"})

	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
}

func TestAdvancedQuery(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	// AND OPERATOR
	var users []User

	result := db.Select("id", "first_name").Where("first_name like ?", "%User%").
		Where("password = ?", "rahasia").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 18, len(users))

	// OR OPERATOR
	var users2 []User

	result = db.Select("id", "first_name").Where("first_name like ?", "%User%").
		Or("password = ?", "rahasia").
		Find(&users2)

	assert.Nil(t, result.Error)
	assert.Equal(t, 22, len(users2))

	// NOT OPERATOR
	var users3 []User

	result = db.Select("id", "first_name").Not("first_name like ?", "%User%").
		Where("password = ?", "rahasia").
		Find(&users3)

	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users3))
}

func TestSelectFields(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var users []User

	err := db.Select("id", "first_name").
		Find(&users).Error

	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}
	assert.Equal(t, 26, len(users))
}

func TestStructCondition(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	userCondition := User{
		ID: "roozy999",
		Name: Name{
			FirstName: "Muhamad",
		},
	}

	var users []User

	result := db.Select("id", "first_name").Where(userCondition).Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
	fmt.Println(users)
}

func TestMapCondition(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	mapCondition := map[string]interface{}{
		"middle_name": "",
		"last_name":   "",
	}

	var users []User

	err := db.Where(mapCondition).Find(&users).Error

	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var users []User

	result := db.Order("id asc, first_name asc").Limit(5).Offset(5).Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, "12", users[0].ID)
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestNonModel(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var users []UserResponse

	result := db.Model(&User{}).Select("id", "first_name", "last_name").Limit(5).Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 5, len(users))
	fmt.Println(users)
}

func TestUpdate(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	db.Transaction(func(tx *gorm.DB) error {

		user := User{}

		result := db.First(&user, "id = ?", "1")

		assert.Nil(t, result.Error)

		user.Name.FirstName = "Teguh1"
		user.Name.MiddleName = "Wahyuda1"
		user.Password = "rahasia1"

		result = db.Save(&user)

		assert.Nil(t, result.Error)

		return nil
	})
}

func TestSelectedColumns(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	db.Transaction(func(tx *gorm.DB) error {

		result := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
			"middle_name": "",
			"last_name":   "Aza",
		})

		assert.Nil(t, result.Error)

		result = db.Model(&User{}).Where("id = ?", "1").Update("password", "newrahasia")

		assert.Nil(t, result.Error)

		result = db.Model(&User{}).Where("id = ?", "1").Updates(User{
			Name: Name{
				FirstName: "Test First",
				LastName:  "Test Last",
			},
		})

		assert.Nil(t, result.Error)

		return nil
	})
}

func TestAutoIncrement(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserID: 1,
			Action: "Test Action",
		}

		result := db.Create(&userLog)

		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}

func TestUpsert(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	userLog := UserLog{
		UserID: 1,
		Action: "Test Action",
	}

	err := db.Save(&userLog).Error // insert
	assert.Nil(t, err)

	userLog = UserLog{
		ID:     23,
		UserID: 1,
		Action: "Test Action Edited",
	}

	err = db.Save(&userLog).Error // update
	assert.Nil(t, err)
}

func TestConflict(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	user := User{
		ID: "1",
		Name: Name{
			FirstName: "User 99",
		},
	}

	err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user).Error // insert
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var user User

	result := db.First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)

	result = db.Delete(&user)
	assert.Nil(t, result.Error)

	result = db.Delete(&User{}, "id = ?", "101")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", "100").Delete(&User{})
	assert.Nil(t, result.Error)
}

func TestSoftDelete(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	todo := Todo{
		UserID:      1,
		Title:       "TODO 1",
		Description: "Isi TODO 1",
	}

	result := db.Create(&todo)
	assert.Nil(t, result.Error)

	result = db.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	result = db.Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))

	fmt.Println(todos)
}

func TestUnscoped(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var todo Todo

	result := db.Unscoped().First(&todo, "id = ?", "3")
	assert.Nil(t, result.Error)

	result = db.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)

	var todos []Todo
	result = db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

func TestLock(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Transaction(func(tx *gorm.DB) error {
		var user User

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "id = ?", "103").Error

		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Morro"

		return tx.Save(&user).Error
	})

	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	wallet := Wallet{
		ID:      "1",
		UserID:  "2",
		Balance: 100000,
	}

	res := db.Create(&wallet)

	assert.Nil(t, res.Error)
}

func TestRetrieveRelationPreload(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var user User

	// preload cocok untuk relasi one to many / many to many
	err := db.Model(&User{}).Preload("Wallet").First(&user, "id = ?", "2").Error

	assert.Nil(t, err)
	assert.Equal(t, "2", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	// bytes, _ := json.Marshal(&user)

	// fmt.Println(string(bytes))
}

func TestRetrieveRelationJoins(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var users []User

	// preload cocok untuk relasi one to one
	err := db.Model(&User{}).Select("users.id", "users.first_name", "users.middle_name", "users.last_name").Joins("Wallet").Take(&users, "users.id = ?", "2").Error

	assert.Nil(t, err)
	// assert.Equal(t, 25, len(users))
	// bytes, _ := json.Marshal(users)

	// fmt.Println(string(bytes))
}

func TestAutoCreateUpdate(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	user := User{
		ID:       "992",
		Password: "Rahasia",
		Name: Name{
			FirstName: "Az",
		},
		Wallet: Wallet{
			ID:      "2",
			UserID:  "992",
			Balance: 1000000,
		},
	}

	err := db.Create(&user).Error

	assert.Nil(t, err)
}

func TestOmit(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	user := User{
		ID:       "roozy10000",
		Password: "Rahasia",
		Name: Name{
			FirstName: "Az",
		},
		Wallet: Wallet{
			ID:      "3",
			UserID:  "roozy10000",
			Balance: 1000000,
		},
	}

	err := db.Omit(clause.Associations).Create(&user).Error

	assert.Nil(t, err)
}

func TestUpsertHasMany(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	user := User{
		ID:       "roozy10001",
		Password: "Rahasia",
		Name: Name{
			FirstName: "Az",
		},
		Wallet: Wallet{
			ID:      "3",
			UserID:  "roozy10001",
			Balance: 5000000,
		},
		Addresses: []Address{
			{
				UserID:  "roozy10001",
				Address: "Jalan A",
			},
			{
				UserID:  "roozy10001",
				Address: "Jalan B",
			},
		},
	}

	err := db.Create(&user).Error

	assert.Nil(t, err)
}

func TestRetrieveCombinationsRelations(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var users []User

	err := db.Model(&User{}).Select("users.id as id, users.first_name as first_name").Preload("Addresses").Joins("Wallet").Take(&users, "users.id = ?", "roozy10001").Error

	bytes, _ := json.Marshal(users)

	fmt.Println(string(bytes))

	assert.Nil(t, err)
}

func TestBelongsToAddress(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var addresses []Address

	fmt.Println("Preload")
	err := db.Model(&Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)

	assert.Equal(t, 2, len(addresses))

	fmt.Println(addresses)

	addresses = []Address{}

	fmt.Println("Joins")
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)

	assert.Equal(t, 2, len(addresses))

	fmt.Println(addresses)
}

func TestBelongsToWallet(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var wallets []Wallet

	fmt.Println("Preload")
	err := db.Model(&Wallet{}).Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	assert.Equal(t, 3, len(wallets))

	fmt.Println(wallets)

	wallets = []Wallet{}

	fmt.Println("Joins")
	err = db.Model(&Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)

	assert.Equal(t, 3, len(wallets))

	fmt.Println(wallets)
}

func TestCreateManyToMany(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var product = Product{
		ID:    "P001",
		Name:  "Contoh Product",
		Price: 1500000,
	}

	err := db.Create(&product).Error

	assert.Nil(t, err)

	err = db.Table("user_likes_product").Create(map[string]interface{}{
		"user_id":    "2",
		"product_id": "P001",
	}).Error

	assert.Nil(t, err)

	err = db.Table("user_likes_product").Create(map[string]interface{}{
		"user_id":    "3",
		"product_id": "P001",
	}).Error

	assert.Nil(t, err)
}

func TestPreloadManyToManyProduct(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var product Product

	err := db.Preload("LikedByUsers").First(&product, "id = ?", "P001").Error

	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))

	bytes, _ := json.Marshal(&product)

	fmt.Println(string(bytes))
}

func TestPreloadManyToManyUser(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var user User

	err := db.Preload("LikeProducts").Find(&user, "id = ?", "2").Error

	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))

	bytes, _ := json.Marshal(&user)

	fmt.Println(string(bytes))
}

func TestAssociationFind(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var product Product
	err := db.First(&product, "id = ?", "P001").Error

	assert.Nil(t, err)

	var users []User

	err = db.Model(&product).Association("LikedByUsers").Find(&users)

	assert.Nil(t, err)

	assert.Equal(t, 2, len(users))

	bytes, _ := json.Marshal(&users)

	fmt.Println(string(bytes))
}

func TestAssociationAppend(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var user User
	err := db.Take(&user, "id = ?", "992").Error

	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P001").Error

	assert.Nil(t, err)

	// append cocok untuk many to many
	err = db.Model(&product).Association("LikedByUsers").Append(&user)

	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Transaction(func(tx *gorm.DB) error {

		var user User
		err := tx.First(&user, "id = ?", "992").Error

		assert.Nil(t, err)

		wallet := Wallet{
			ID:      "04",
			UserID:  user.ID,
			Balance: 900000000,
		}

		assert.Nil(t, err)

		// // append cocok untuk many to many
		err = tx.Model(&user).Association("Wallet").Replace(&wallet)

		return err
	})

	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var user User

	err := db.First(&user, "id = ?", "992").Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)

	assert.Nil(t, err)

}

func TestAssociationClear(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var product Product

	err := db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear()

	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	var user User

	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Preload("Wallet", "balance > ?", 1000000000).First(&user, "id = ?", "992").Error

	fmt.Println(user)

	bytes, _ := json.Marshal(&user)

	fmt.Println(string(bytes))
	assert.Nil(t, err)
}

func TestPreloadingWithNested(t *testing.T) {
	var wallet Wallet

	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Preload("User.Addresses").Find(&wallet, "id = ?", "").Error

	assert.Nil(t, err)

	result := Stringify(wallet)

	fmt.Println(result)
}

func TestPreloadAll(t *testing.T) {

	var user User
	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Preload(clause.Associations).First(&user, "id = ?", "roozy10001").Error

	assert.Nil(t, err)

	result := Stringify(user)

	fmt.Println(result)
}

func TestJoinsQuery(t *testing.T) {
	var users []User

	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error

	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}

	err = db.Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 29, len(users))
}

func TestJoinsQueryCondition(t *testing.T) {
	var users []User

	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 5000000).Find(&users).Error

	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))

	users = []User{}

	err = db.Joins("Wallet").Where("Wallet.balance > ?", 5000000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))

	result := Stringify(&users)

	fmt.Println(result)
}

func TestCount(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	var count int64

	err := db.Model(&User{}).Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Count(&count).Error

	assert.Nil(t, err)

	assert.Equal(t, int64(3), count)

}

type AggreationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result []AggreationResult

	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Joins("User").Group("User.id").Having("sum(balance) > ?", 1000000).Find(&result).Error

	assert.Nil(t, err)

	assert.Equal(t, 2, len(result))
	fmt.Println(result)
}

func Stringify[T any](value T) string {
	bytes, _ := json.Marshal(&value)

	return string(bytes)
}

func TestContext(t *testing.T) {

	db := OpenConnection()

	defer CloseConnection(db)

	ctx := context.Background()

	var users []User

	err := db.WithContext(ctx).Find(&users).Error

	assert.Nil(t, err)
	fmt.Println(users)
}

func TestScopes(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	var wallets []Wallet

	err := db.Model(&Wallet{}).Scopes(BrokeWalletBalance).Find(&wallets).Error

	assert.Nil(t, err)
	fmt.Println(wallets)

	wallets = []Wallet{}

	err = db.Model(&Wallet{}).Scopes(SultanWalletBalance).Find(&wallets).Error

	assert.Nil(t, err)
	fmt.Println(wallets)
}

func TestMigrator(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)

	err := db.Migrator().AutoMigrate(&GuestBook{})

	assert.Nil(t, err)
}

func TestHook(t *testing.T) {
	db := OpenConnection()

	defer CloseConnection(db)
	user := User{
		Password: "Rahasia Lho",
		Name: Name{
			FirstName: "Aznew",
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)

	assert.NotEqual(t, "", user.ID)

}