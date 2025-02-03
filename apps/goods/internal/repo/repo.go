package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
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
	GetAllBrandList(ctx context.Context) ([]*model.Brand, error)

	CreateSpu(ctx context.Context, spu *model.GoodsSpu) (*model.GoodsSpu, error)
	CreateSkus(ctx context.Context, skus []*model.GoodsSku) error
	CreateSpecs(ctx context.Context, specs []*model.Spec) ([]*model.Spec, error)
	CreateSpecValues(ctx context.Context, specValues []*model.SpecValue) error
	CreateAttributes(ctx context.Context, attributes []*model.Attribute) ([]*model.Attribute, error)

	UpdateOrSaveKeywordTimes(ctx context.Context, keyword string) error // 增加搜索关键词搜索次数

	AddSpuToES(ctx context.Context, spu *model.GoodsSpu) error // 添加spu到es
	SearchSpu(ctx context.Context, req *ggen.SearchSpuListReq) ([]*model.GoodsSpu, error)
	SearchAttr(ctx context.Context, req *ggen.SearchSpuListReq) ([]*model.Attribute, error)
}

// 数据访问层实现了RepoInterface
type Repo struct {
	DB          *gorm.DB
	ES          *elasticsearch.TypedClient
	Cache       *redis.Client
	PayClient   pc.Client
	GoodsClient gc.Client
	OrderClient oc.Client
	UserClient  uc.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, es *elasticsearch.TypedClient, payClient pc.Client, goodsClient gc.Client, orderClient oc.Client, userClient uc.Client) *Repo {
	return &Repo{
		DB:          db,
		Cache:       cache,
		ES:          es,
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
	err := r.DB.Model(&model.GoodsSpu{}).Where("id in ?", IDs).Find(&spus).Error
	if err != nil {
		return nil, err
	}
	return spus, nil
}

func (r *Repo) GetSkuListBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.GoodsSku, error) {
	var skus []*model.GoodsSku
	err := r.DB.Model(&model.GoodsSku{}).Where("spu_id in ?", SpuIDs).Find(&skus).Error
	if err != nil {
		return nil, err
	}
	return skus, nil
}

func (r *Repo) GetSpecBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.Spec, error) {
	var specs []*model.Spec
	err := r.DB.Model(&model.Spec{}).Where("spu_id in ?", SpuIDs).Find(&specs).Error
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

func (r *Repo) GetAllBrandList(ctx context.Context) ([]*model.Brand, error) {
	res := make([]*model.Brand, 0)
	err := r.DB.Model(&model.Brand{}).Where("deleted_at is null").Find(&res).Error
	return res, err
}

func (r *Repo) CreateSpu(ctx context.Context, spu *model.GoodsSpu) (*model.GoodsSpu, error) {
	err := r.DB.Model(&model.GoodsSpu{}).Clauses(clause.Returning{}).Create(spu).Error
	return spu, err
}

func (r *Repo) CreateSkus(ctx context.Context, skus []*model.GoodsSku) error {
	return r.DB.Create(skus).Error
}
func (r *Repo) CreateSpecs(ctx context.Context, specs []*model.Spec) ([]*model.Spec, error) {
	err := r.DB.Model(&model.Spec{}).Clauses(clause.Returning{}).Create(&specs).Error
	return specs, err
}

func (r *Repo) CreateSpecValues(ctx context.Context, specValues []*model.SpecValue) error {
	return r.DB.Create(specValues).Error
}

func (r *Repo) CreateAttributes(ctx context.Context, attributes []*model.Attribute) ([]*model.Attribute, error) {
	err := r.DB.Model(&model.Attribute{}).Clauses(clause.Returning{}).Create(&attributes).Error
	return attributes, err
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

func (r *Repo) AddSpuToES(ctx context.Context, spu *model.GoodsSpu) error {
	// 将 spu 结构体转换为 JSON 格式
	data, err := json.Marshal(spu)
	if err != nil {
		return err
	}

	// 创建一个 Elasticsearch 请求体
	req := esapi.IndexRequest{
		Index:      "goods_spu",
		DocumentID: strconv.FormatUint(spu.ID, 10),
		Body:       strings.NewReader(string(data)),
	}

	// 执行请求
	res, err := req.Do(ctx, r.ES)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// 检查响应状态码
	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}
	return nil
}

func (r Repo) SearchAttr(ctx context.Context, req *ggen.SearchSpuListReq) ([]*model.Attribute, error) {
	// TODO
	return nil, nil
}

func (r *Repo) SearchSpu(ctx context.Context, req *ggen.SearchSpuListReq) ([]*model.GoodsSpu, error) {
	// 构建 bool 查询
	boolQuery := types.BoolQuery{
		Must: []types.Query{},
		Should: []types.Query{
			{
				Match: map[string]types.MatchQuery{
					"spu_name": {
						Query: req.Keyword,
					},
				},
			},
			{
				Match: map[string]types.MatchQuery{
					"spu_desc": {
						Query: req.Keyword,
					},
				},
			},
		},
		MinimumShouldMatch: "1", // 至少匹配一个 should 子句
	}

	// 添加 category_ids 精准匹配
	if len(req.CategoryIds) > 0 {
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Terms: &types.TermsQuery{
				TermsQuery: map[string]types.TermsQueryField{
					"category_id": req.CategoryIds,
				},
			},
		})
	}

	// 添加 brand_ids 精准匹配
	if len(req.BrandIds) > 0 {
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Terms: &types.TermsQuery{
				TermsQuery: map[string]types.TermsQueryField{
					"brand_id": req.BrandIds,
				},
			},
		})
	}

	// 构建搜索请求
	searchReq := r.ES.Search().
		Index(constant.ESIndexGoodsSpu).
		Query(&types.Query{
			Bool: &boolQuery,
		}).
		From(int((req.PageNo - 1) * req.PageSize)).
		Size(int(req.PageSize))

	// 执行请求
	resp, err := searchReq.Do(ctx)
	if err != nil {
		return nil, err
	}

	spus := make([]*model.GoodsSpu, 0)
	for _, h := range resp.Hits.Hits {
		var spu model.GoodsSpu
		bytes, err := h.Source_.MarshalJSON()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bytes, &spu)
		if err != nil {
			return nil, err
		}
		spus = append(spus, &spu)
	}
	return spus, nil
}
