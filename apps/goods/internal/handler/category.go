package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
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
	categoryTree := buildCategoryTree(categoryList)

	rsp := &ggen.GetCategoryListRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.GetCategoryListRsp_GetCategoryListRspData{
			CategoryList: categoryTree,
		},
	}
	return rsp
}

// 列表转树状结构
func buildCategoryTree(categories []*model.Category) []*ggen.Category {
	// 创建一个map来快速查找子类别
	childrenMap := make(map[uint64][]*model.Category)
	for _, category := range categories {
		childrenMap[category.ParentID] = append(childrenMap[category.ParentID], category)
	}

	// 递归构建树结构
	var buildTree func(parentID uint64) []*ggen.Category
	buildTree = func(parentID uint64) []*ggen.Category {
		var tree []*ggen.Category
		for _, category := range childrenMap[parentID] {
			protoCategory := &ggen.Category{
				CategoryId:   category.ID,
				CategoryName: category.Name,
				ParentId:     category.ParentID,
				Children:     buildTree(category.ID),
			}
			tree = append(tree, protoCategory)
		}
		return tree
	}

	// 从根节点开始构建树
	return buildTree(0) // 假设根节点的ParentID为0
}
