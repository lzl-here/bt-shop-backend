package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// 运行 go run . -host=http://localhost:9200

// 初始化es索引
func main() {

	host := flag.String("host", "http://127.0.0.1:9200", "Elasticsearch host address")
	flag.Parse()

	cfg := elasticsearch.Config{
		Addresses: []string{
			*host,
		},
	}
	cli, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}

	// 确保 goods_spu 索引存在
	err = ensureIndexExists(context.Background(), cli, "goods_spu", goodsSpuMapping)
	if err != nil {
		panic(err)
	}
}

// 检查并创建索引
func ensureIndexExists(ctx context.Context, cli *elasticsearch.TypedClient, indexName string, mapping string) error {
	// 检查索引是否存在
	exists, err := indexExists(ctx, cli, indexName)
	if err != nil {
		return err
	}

	if !exists {
		// 创建索引
		req := esapi.IndicesCreateRequest{
			Index: indexName,
			Body:  strings.NewReader(mapping),
		}

		res, err := req.Do(ctx, cli)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return fmt.Errorf("error creating index: %s", res.Status())
		}
		fmt.Printf("Index %s created successfully\n", indexName)
	}
	fmt.Println("索引已存在")
	return nil
}

// 检查索引是否存在
func indexExists(ctx context.Context, cli *elasticsearch.TypedClient, indexName string) (bool, error) {
	req := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}

	res, err := req.Do(ctx, cli)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

var goodsSpuMapping = `{
    "mappings": {
        "properties": {
            "id": { "type": "keyword" },
            "spu_name": { "type": "text" },
            "spu_desc": { "type": "text" },
            "category_id": { "type": "keyword" },
            "category_name": { "type": "text" },
            "brand_id": { "type": "keyword" },
            "brand_name": { "type": "text" },
            "price": { "type": "text" },
            "enabled": { "type": "boolean" },
            "img_url": { "type": "text" },
            "attribute_ids": { "type": "text" },
            "spec_ids": { "type": "text" }
        }
    }
}`
