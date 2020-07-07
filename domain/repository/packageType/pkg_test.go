package packageType

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"muju-frontstore-go/domain/model"
	"testing"
	"time"
)

type PkgSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	pkg_repository PackageRepository
	pkg            *model.PackageType
}

func (s *PkgSuite) SetupSuite() {
	var (
		db *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.pkg_repository = Pkg_Repository(s.DB)
}

func (p *PkgSuite) AfterTest(_, _ string) {
	require.NoError(p.T(), p.mock.ExpectationsWereMet())
}

func TestInitStore(t *testing.T) {
	suite.Run(t, new(PkgSuite))
}

func (p *PkgSuite) Test_Create_Pkg(){
	var(
		pkg = model.PackageType{
			PackageName:  "Silver",
			PackagePrice: 200000,
			CreatedDate:  time.Now(),
			CreatedBy:    "Admin",
			ModifiedDate: time.Now(),
			ModifiedBy:   "Admin",
			DeletedDate:  time.Now(),
			DeletedBy:    "Admin",
			Active:       true,
			IsDeleted:    false,
		}
	)
	err := p.pkg_repository.CreatePackage(&pkg)
	require.NoError(p.T(), err)
}

func (p *PkgSuite) Test_Update_Pkg(){
	var (
		pkg = model.PackageType{
			Id:           1,
			PackageName:  "Silver",
			CreatedDate:  time.Now(),
			CreatedBy:    "Admin",
			ModifiedDate: time.Now(),
			ModifiedBy:   "Admin",
			DeletedDate:  time.Now(),
			DeletedBy:    "Admin",
			Active:       true,
			IsDeleted:    false,
		}
	)
	err := p.pkg_repository.UpdatePackage(&pkg)
	require.NoError(p.T(), err)
}

func (p *PkgSuite) Test_Delete_Pkg(){
	var (
		pkg = model.PackageType{
			Id:           1,
			PackageName:  "Silver",
			CreatedDate:  time.Now(),
			CreatedBy:    "Admin",
			ModifiedDate: time.Now(),
			ModifiedBy:   "Admin",
			DeletedDate:  time.Now(),
			DeletedBy:    "Admin",
			Active:       true,
			IsDeleted:    false,
		}
	)
	err := p.pkg_repository.DeletePackage(&pkg)
	require.NoError(p.T(), err)
}



