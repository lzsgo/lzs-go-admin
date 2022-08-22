package example

import (
	"context"
	"github.com/lzsgo/lzs-go-admin/server/model/example"
	"github.com/lzsgo/lzs-go-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderExaFile = system.InitOrderInternal + 1

type initExaFileMysql struct{}

// auto run
func init() {
	system.RegisterInit(initOrderExaFile, &initExaFileMysql{})
}

func (i *initExaFileMysql) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&example.ExaFileUploadAndDownload{})
}

func (i *initExaFileMysql) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&example.ExaFileUploadAndDownload{})
}

func (i initExaFileMysql) InitializerName() string {
	return example.ExaFileUploadAndDownload{}.TableName()
}

func (i *initExaFileMysql) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []example.ExaFileUploadAndDownload{
		{Name: "10.png", Url: "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJ5Ozf1ziaeRmMcxDnsAuu8QOqdtAQ8OMt2BMQRldfshgpQnaga1bQwH9mJQeykZQK2cDn4IBnrmJw/132", Tag: "png", Key: "158787308910.png"},
		{Name: "logo.png", Url: "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKaf4N9IvVDRxzBr1FIeE4UhZd6kZ2Ngq3qBmu3FiaGlfhDmfnCu98larMP5MnPq8iaESZb1p0Zwsqw/132", Tag: "png", Key: "1587973709logo.png"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, example.ExaFileUploadAndDownload{}.TableName()+"表数据初始化失败!")
	}
	return ctx, nil
}

func (i *initExaFileMysql) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	lookup := example.ExaFileUploadAndDownload{Name: "logo.png", Key: "1587973709logo.png"}
	if errors.Is(db.First(&lookup, &lookup).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
