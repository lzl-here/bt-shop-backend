package repo

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	gc "github.com/lzl-here/bt-shop-backend/kitex_gen/goods/goodsservice"
	oc "github.com/lzl-here/bt-shop-backend/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ (RepoInterface) = (*Repo)(nil)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	Transaction(ctx context.Context, f func(ctx context.Context, tx *gorm.DB) error) error
	GetSpuListByIDs(ctx context.Context, IDs []uint64) ([]*model.GoodsSpu, error)
	GetSkuListBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.GoodsSku, error)
	GetSpecBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.Spec, error)
	GetKeywordDownList(ctx context.Context, prefix string, limit int) ([]*model.Keyword, error)
	GetAllCategoryList(ctx context.Context) ([]*model.Category, error)

	CreateSpu(ctx context.Context, spu *model.GoodsSpu) (*model.GoodsSpu, error)
	CreateSkus(ctx context.Context, skus []*model.GoodsSku) error
	CreateSpecs(ctx context.Context, specs []*model.Spec) ([]*model.Spec, error)
	CreateSpecValues(ctx context.Context, specValues []*model.SpecValue) error
	CreateAttributes(ctx context.Context, attributes []*model.Attribute) ([]*model.Attribute, error)

	UpdateOrSaveKeywordTimes(ctx context.Context, keyword string) error // 增加搜索关键词搜索次数
}

// 数据访问层实现了RepoInterface
type Repo struct {
	DB          *gorm.DB
	Cache       *redis.Client
	PayClient   pc.Client
	GoodsClient gc.Client
	OrderClient oc.Client
	UserClient  uc.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, payClient pc.Client, goodsClient gc.Client, orderClient oc.Client, userClient uc.Client) *Repo {
	return &Repo{
		DB:          db,
		Cache:       cache,
		PayClient:   payClient,
		GoodsClient: goodsClient,
		OrderClient: orderClient,
		UserClient:  userClient,
	}
}

func (r *Repo) Transaction(ctx context.Context, f func(ctx context.Context, tx *gorm.DB) error) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		return f(ctx, tx)
	})
}

func (r *Repo) GetSpuListByIDs(ctx context.Context, IDs []uint64) ([]*model.GoodsSpu, error) {
	var spus []*model.GoodsSpu
	err := r.DB.Model(&model.GoodsSpu{}).Where("id in ?", IDs).Find(spus).Error
	if err != nil {
		return nil, err
	}
	return spus, nil
}

func (r *Repo) GetSkuListBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.GoodsSku, error) {
	var skus []*model.GoodsSku
	err := r.DB.Model(&model.GoodsSku{}).Where("spu_id in ?", SpuIDs).Find(skus).Error
	if err != nil {
		return nil, err
	}
	return skus, nil
}

func (r *Repo) GetSpecBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.Spec, error) {
	var specs []*model.Spec
	err := r.DB.Model(&model.Spec{}).Where("spu_id in ?", SpuIDs).Find(specs).Error
	if err != nil {
		return nil, err
	}
	return specs, nil
}

func (r *Repo) GetKeywordDownList(ctx context.Context, prefix string, limit int) ([]*model.Keyword, error) {
	res := make([]*model.Keyword, 0)
	err := r.DB.Model(&model.Keyword{}).
		Where("value like ?", prefix+"%").
		Limit(limit).
		Order("search_times DESC").
		Find(&res).
		Error
	return res, err
}

func (r *Repo) GetAllCategoryList(ctx context.Context) ([]*model.Category, error) {
	res := make([]*model.Category, 0)
	err := r.DB.Model(&model.Category{}).Where("deleted_at is null").Find(&res).Error
	return res, err
}

func (r *Repo) CreateSpu(ctx context.Context, spu *model.GoodsSpu) (*model.GoodsSpu, error) {
	res := new(model.GoodsSpu)
	err := r.DB.Model(res).Clauses(clause.Returning{}).Create(spu).Error
	return res, err
}

func (r *Repo) CreateSkus(ctx context.Context, skus []*model.GoodsSku) error {
	return r.DB.Create(skus).Error
}
func (r *Repo) CreateSpecs(ctx context.Context, specs []*model.Spec) ([]*model.Spec, error) {
	res := make([]*model.Spec, 0)
	err := r.DB.Model(&model.Spec{}).Clauses(clause.Returning{}).Create(&specs).Error
	return res, err
}

func (r *Repo) CreateSpecValues(ctx context.Context, specValues []*model.SpecValue) error {
	return r.DB.Create(specValues).Error
}

func (r *Repo) CreateAttributes(ctx context.Context, attributes []*model.Attribute) ([]*model.Attribute, error) {
	res := make([]*model.Attribute, 0)
	err := r.DB.Model(&model.Attribute{}).Clauses(clause.Returning{}).Create(&attributes).Error
	return res, err
}

func (r *Repo) UpdateOrSaveKeywordTimes(ctx context.Context, keyword string) error {
	keywordModel := &model.Keyword{Value: keyword, SearchTimes: 1}

	// 使用 FirstOrCreate 来查找或创建记录
	result := r.DB.Where(&model.Keyword{Value: keyword}).FirstOrCreate(keywordModel)
	if result.Error != nil {
		return result.Error
	}

	// 如果记录已经存在，则更新 search_times
	if result.RowsAffected == 0 {
		// 记录已经存在，更新 search_times
		return r.DB.Model(&model.Keyword{}).Where("value = ?", keyword).Update("search_times", gorm.Expr("search_times + ?", 1)).Error
	}
	return nil
}
