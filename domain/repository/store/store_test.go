package store

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"muju-frontstore-go/domain/model"
	"testing"
	"time"
)

type StoreSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	store_repository StoreRepository
	store            *model.Store
}

func (s *StoreSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.store_repository = Store_Repository(s.DB)
}

func (s *StoreSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitStore(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}

func (s *StoreSuite) Test_Create_Suite() {
	var (
		store = model.Store{
			StoreName:         "Toko 2",
			StoreDomain:       "www.toko2.muju.com",
			ProductCategoryId: 1,
			CountryId:         2,
			ProvinceId:        11,
			CityId:            1101,
			PostalCode:        "12312",
			CreatedBy:         "Admin",
			CreatedDate:       time.Now(),
			ModifiedBy:        "Admin",
			ModifiedDate:      time.Now(),
			DeletedBy:         "Admin",
			DeletedDate:       time.Now(),
			Active:            true,
			IsDeleted:         false,
		}
	)
	err := s.store_repository.CreateStores(&store)
	require.NoError(s.T(), err)
}

func (s *StoreSuite) Test_Update_Store() {
	var (
		store = model.Store{
			Id:                1,
			StoreName:         "Toko 2",
			StoreDomain:       "www.toko2.muju.com",
			ProductCategoryId: 1,
			CountryId:         2,
			ProvinceId:        11,
			CityId:            1101,
			PostalCode:        "12312",
			CreatedBy:         "Admin",
			CreatedDate:       time.Now(),
			ModifiedBy:        "Admin",
			ModifiedDate:      time.Now(),
			DeletedBy:         "Admin",
			DeletedDate:       time.Now(),
			Active:            true,
			IsDeleted:         false,
		}
	)
	err := s.store_repository.UpdateStores(&store)
	require.NoError(s.T(), err)
}

func (s *StoreSuite) Test_Delete_Store() {
	var (
		store = model.Store{
			Id:                3,
			StoreName:         "Toko 2",
			StoreDomain:       "www.toko2.muju.com",
			ProductCategoryId: 1,
			CountryId:         2,
			ProvinceId:        11,
			CityId:            1101,
			PostalCode:        "12312",
			CreatedBy:         "Admin",
			CreatedDate:       time.Now(),
			ModifiedBy:        "Admin",
			ModifiedDate:      time.Now(),
			DeletedBy:         "Admin",
			DeletedDate:       time.Now(),
			Active:            true,
			IsDeleted:         false,
		}
	)
	err := s.store_repository.DeleteStores(&store)
	require.NoError(s.T(), err)
}
