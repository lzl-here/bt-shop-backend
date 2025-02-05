package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
)

/**
* @Description: 获取商品分类列表
 */
func (h *GoodsHandler) GetCategoryList(ctx context.Context, req *ggen.GetCategoryListReq) (*ggen.GetCategoryListRsp, error) {
	categoryList, err := h.rep.GetAllCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	return buildCategoryListRsp(categoryList), nil
}

func buildCategoryListRsp(categoryList []*model.Category) *ggen.GetCategoryListRsp {
	categoryTree := buildTreeList(categoryList)

	rsp := &ggen.GetCategoryListRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.GetCategoryListRsp_GetCategoryListRspData{
			CategoryList: categoryTree,
		},
	}
	return rsp
}

// 构建目录树列表
func buildTreeList(categories []*model.Category) []*ggen.BaseCategory {
	// 创建一个map来快速查找子类别
	childrenMap := make(map[uint64][]*model.Category)
	rootCategory := make([]*model.Category, 0)
	for _, category := range categories {
		childrenMap[category.ParentID] = append(childrenMap[category.ParentID], category)
		if category.Level == 0 {
			rootCategory = append(rootCategory, category)
		}
	}

	// 构建目录树列表
	res := make([]*ggen.BaseCategory, 0)
	for _, c := range rootCategory {
		res = append(res, buildTree(c, childrenMap))
	}
	return res
}

// 根据一个根节点，构建一个目录树
func buildTree(rootCategory *model.Category, childrenMap map[uint64][]*model.Category) *ggen.BaseCategory {
	categoryRsp := &ggen.BaseCategory{
		CategoryId:   rootCategory.ID,
		CategoryName: rootCategory.Name,
		ParentId:     rootCategory.ParentID,
	}
	tree := make([]*ggen.BaseCategory, 0)
	rootID := rootCategory.ID
	for _, category := range childrenMap[rootID] {
		protoCategory := &ggen.BaseCategory{
			CategoryId:   category.ID,
			CategoryName: category.Name,
			ParentId:     category.ParentID,
			Level:        category.Level,
			Children:     buildTreeChild(category, childrenMap),
		}
		tree = append(tree, protoCategory)
	}
	categoryRsp.Children = tree
	return categoryRsp
}

func buildTreeChild(rootCategory *model.Category, childrenMap map[uint64][]*model.Category) []*ggen.BaseCategory {
	children := make([]*ggen.BaseCategory, 0)
	for _, category := range childrenMap[rootCategory.ID] {
		rsp := &ggen.BaseCategory{
			CategoryId:   category.ID,
			CategoryName: category.Name,
			ParentId:     category.ParentID,
			Level:        category.Level,
			Children:     buildTreeChild(category, childrenMap),
		}
		children = append(children, rsp)
	}
	return children
}
