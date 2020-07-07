package Template

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

type TmpSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	tmp_repository TemplateRepository
	tmp            *model.Template
}

func (t *TmpSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, t.mock, err = sqlmock.New()
	require.NoError(t.T(), err)

	t.DB, err = gorm.Open("postgres", db)
	require.NoError(t.T(), err)

	t.DB.LogMode(true)
	t.tmp_repository = Tmp_Repository(t.DB)
}

func (t *TmpSuite) AfterSuite(_, _ string) {
	require.NoError(t.T(), t.mock.ExpectationsWereMet())
}

func TestInitPkg(t *testing.T) {
	suite.Run(t, new(TmpSuite))
}

func (t *TmpSuite) Test_Create_Temp() {
	var (
		tmp = model.Template{
			TemplateName:      "Mantap 1",
			TemplatePrice:     250000,
			UrlDemo:           "mantap1.com",
			ProductCategoryId: 1,
			CreatedDate:       time.Now(),
			CreatedBy:         "Admin",
			ModifiedDate:      time.Now(),
			ModifiedBy:        "Admin",
			DeletedDate:       time.Now(),
			DeletedBy:         "Admin",
			Active:            true,
			IsDeleted:         false,
		}
	)
	err := t.tmp_repository.CreateTemplate(&tmp)
	require.NoError(t.T(), err)
}

func (t *TmpSuite) Test_Update_Tmp() {
	var (
		tmp = model.Template{
			Id:                1,
			TemplateName:      "Mantap 1",
			TemplatePrice:     250000,
			UrlDemo:           "mantap1.com",
			ProductCategoryId: 1,
			CreatedDate:       time.Now(),
			CreatedBy:         "Admin",
			ModifiedDate:      time.Now(),
			ModifiedBy:        "Admin",
			DeletedDate:       time.Now(),
			DeletedBy:         "Admin",
			Active:            true,
			IsDeleted:         false,
		}
	)
	err := t.tmp_repository.UpdateTemplate(&tmp)
	require.NoError(t.T(), err)
}

func (t *TmpSuite) Test_Delete_Tmp() {
	var (
		tmp = model.Template{
			Id:                1,
			TemplateName:      "Mantap 1",
			TemplatePrice:     250000,
			UrlDemo:           "mantap1.com",
			ProductCategoryId: 1,
			CreatedDate:       time.Now(),
			CreatedBy:         "Admin",
			ModifiedDate:      time.Now(),
			ModifiedBy:        "Admin",
			DeletedDate:       time.Now(),
			DeletedBy:         "Admin",
			Active:            true,
			IsDeleted:         false,
		}
	)
	err := t.tmp_repository.DeleteTemplate(&tmp)
	require.NoError(t.T(), err)
}
