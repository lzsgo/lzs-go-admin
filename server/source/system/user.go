package system

import (
	"context"
	sysModel "github.com/lzsgo/lzs-go-admin/server/model/system"
	"github.com/lzsgo/lzs-go-admin/server/service/system"
	"github.com/lzsgo/lzs-go-admin/server/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const initOrderUser = initOrderAuthority + 1

type initUser struct{}

// auto run
func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysUser{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	password := utils.BcryptHash("lishier")
	adminPassword := utils.BcryptHash("lishier")

	entities := []sysModel.SysUser{
		{
			UUID:        uuid.NewV4(),
			Username:    "admin",
			Password:    adminPassword,
			NickName:    "李十二",
			HeaderImg:   "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJ5Ozf1ziaeRmMcxDnsAuu8QOqdtAQ8OMt2BMQRldfshgpQnaga1bQwH9mJQeykZQK2cDn4IBnrmJw/132",
			AuthorityId: 888,
			Phone:       "18518561648",
			Email:       "yize1346697@163.com",
		},
		{
			UUID:        uuid.NewV4(),
			Username:    "a303176530",
			Password:    password,
			NickName:    "lishier_plus_User",
			HeaderImg:   "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKaf4N9IvVDRxzBr1FIeE4UhZd6kZ2Ngq3qBmu3FiaGlfhDmfnCu98larMP5MnPq8iaESZb1p0Zwsqw/132",
			AuthorityId: 9528,
			Phone:       "18500582119",
			Email:       "1398978641@qq.com"},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Authorities").Replace(authorityEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Authorities").Replace(authorityEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "a303176530").
		Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].AuthorityId == 888
}
