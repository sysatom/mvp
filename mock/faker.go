package main

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sysatom/mvp/model"
	"sync"
)

const (
	GoroutinesNumber = 10

	IssueNumber    = 10000
	CategoryNumber = 10000
	BrandNumber    = 10000

	UserNumber  = 100000 * GoroutinesNumber
	GoodsNumber = 100000 * GoroutinesNumber
)

func main() {
	db, err := sqlx.Connect("mysql", "root:123456@/mvp?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	GenerateBase(db)
	GenerateUser(db)
	GenerateGoods(db)
}

func GenerateUser(db *sqlx.DB) {
	var wg sync.WaitGroup
	for i := 1; i <= GoroutinesNumber; i++ {
		wg.Add(1)
		go insertMallUser(db, &wg)
	}

	wg.Wait()
	fmt.Println("GenerateUser Done")
}

func GenerateGoods(db *sqlx.DB) {
	var wg sync.WaitGroup
	for i := 1; i <= GoroutinesNumber; i++ {
		wg.Add(1)
		go insertMallGoods(db, &wg)
	}

	wg.Wait()
	fmt.Println("GenerateGoods Done")
}

func GenerateBase(db *sqlx.DB) {
	insertMallIssue(db)
	insertMallCategory(db)
	insertMallBrand(db)
	fmt.Println("GenerateBase Done")
}

func insertMallBrand(db *sqlx.DB) {
	var brands []*model.MallBrand
	for i := 0; i < BrandNumber; i++ {
		brand := model.MallBrand{}
		err := faker.FakeData(&brand)
		if err != nil {
			fmt.Println(err)
		}
		brand.FloorPrice = brand.FloorPrice * 1000
		brands = append(brands, &brand)
	}

	_, err := db.NamedExec("INSERT INTO mall_brand (`name`, `desc`, pic_url, sort_order, floor_price, created_at, updated_at) VALUES (:name, :desc, :picUrl, :sortOrder, :floorPrice, :createdAt, :updatedAt)", brands)
	if err != nil {
		fmt.Println(err)
	}
}

func insertMallIssue(db *sqlx.DB) {
	var issues []*model.MallIssue
	for i := 0; i < IssueNumber; i++ {
		issue := model.MallIssue{}
		err := faker.FakeData(&issue)
		if err != nil {
			fmt.Println(err)
		}
		issues = append(issues, &issue)
	}

	_, err := db.NamedExec("INSERT INTO mall_issue (`question`, `answer`, created_at, updated_at) VALUES (:question, :answer, :createdAt, :updatedAt)", issues)
	if err != nil {
		fmt.Println(err)
	}
}

func insertMallCategory(db *sqlx.DB) {
	for i := 0; i < CategoryNumber; i++ {
		category := model.MallCategory{}
		err := faker.FakeData(&category)
		if err != nil {
			fmt.Println(err)
		}
		category.Level = "L1"
		res, err := db.NamedExec("INSERT INTO mall_category (`name`, `keywords`, `desc`, pid, icon_url, pic_url, `level`, sort_order, created_at, updated_at) VALUES (:name, :keywords, :desc, :pid, :iconUrl, :picUrl, :level, :sortOrder, :createdAt, :updatedAt)", category)
		if err != nil {
			fmt.Println(err)
		}

		categoryID, err := res.LastInsertId()
		if err != nil {
			fmt.Println(err)
		}

		for j := 1; j <= 3; j++ {
			childCategory := model.MallCategory{}
			err = faker.FakeData(&childCategory)
			if err != nil {
				fmt.Println(err)
			}
			childCategory.Level = "L2"
			childCategory.Pid = int(categoryID)
			_, err := db.NamedExec("INSERT INTO mall_category (`name`, `keywords`, `desc`, pid, icon_url, pic_url, `level`, sort_order, created_at, updated_at) VALUES (:name, :keywords, :desc, :pid, :iconUrl, :picUrl, :level, :sortOrder, :createdAt, :updatedAt)", childCategory)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func insertMallUser(db *sqlx.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	var users []*model.MallUser
	for i := 0; i < UserNumber/GoroutinesNumber; i++ {
		user := model.MallUser{}
		err := faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, &user)
	}

	_, err := db.NamedExec("INSERT INTO mall_user (username, password, gender, birthday, last_login_time, last_login_ip, user_level, nickname, mobile, avatar, platform, openid, session_key, status, created_at, updated_at) VALUES (:username, :password, :gender, :birthday, :lastLoginTime, :lastLoginIp, :userLevel, :nickname, :mobile, :avatar, :platform, :openId, :sessionKey, :status, :createdAt, :updatedAt)", users)
	if err != nil {
		fmt.Println(err)
	}
}

func insertMallGoods(db *sqlx.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < GoodsNumber/GoroutinesNumber; i++ {
		goods := model.MallGoods{}
		err := faker.FakeData(&goods)
		if err != nil {
			fmt.Println(err)
		}
		gallery := []string{faker.URL(), faker.URL()}
		j, err := json.Marshal(gallery)
		if err != nil {
			fmt.Println(err)
		}
		goods.IsOnSale = true
		goods.Gallery = string(j)
		goods.CounterPrice = goods.CounterPrice * 1000
		goods.RetailPrice = goods.RetailPrice * 1000
		res, err := db.NamedExec("INSERT INTO mall_goods (goods_sn, name, category_id, brand_id, gallery, keywords, brief, is_on_sale, sort_order, pic_url, share_url, is_new, is_hot, unit, counter_price, retail_price, detail, created_at, updated_at) VALUES (:goodsSn, :name, :categoryId, :brandId, :gallery, :keywords, :brief, :isOnSale, :sortOrder, :picUrl, :shareUrl, :isNew, :isHot, :unit, :counterPrice, :retailPrice, :detail, :createdAt, :updatedAt)", goods)
		if err != nil {
			fmt.Println(err)
		}
		goodsID, err := res.LastInsertId()
		if err != nil {
			fmt.Println(err)
		}

		// Attribute
		for j := 1; j <= 3; j++ {
			attribute := model.MallGoodsAttribute{}
			err := faker.FakeData(&attribute)
			if err != nil {
				fmt.Println(err)
			}
			attribute.GoodsID = int(goodsID)
			_, err = db.NamedExec("INSERT INTO mall_goods_attribute (goods_id, attribute, value, created_at, updated_at) VALUES (:goodsId, :attribute, :value, :createdAt, :updatedAt)", attribute)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Specification
		var specificationList []string
		spec := faker.Word()
		for j := 1; j <= 3; j++ {
			specification := model.MallGoodsSpecification{}
			err := faker.FakeData(&specification)
			if err != nil {
				fmt.Println(err)
			}
			specification.Specification = spec
			specification.GoodsID = int(goodsID)
			_, err = db.NamedExec("INSERT INTO mall_goods_specification (goods_id, specification, value, pic_url, created_at, updated_at) VALUES (:goodsId, :specification, :value, :picUrl, :createdAt, :updatedAt)", specification)
			if err != nil {
				fmt.Println(err)
			}
			specificationList = append(specificationList, specification.Value)
		}

		// Product
		for _, specification := range specificationList {
			product := model.MallGoodsProduct{}
			err := faker.FakeData(&product)
			if err != nil {
				fmt.Println(err)
			}
			specifications := []string{specification}
			s, err := json.Marshal(specifications)
			if err != nil {
				fmt.Println(err)
			}
			product.Price = product.Price * 1000
			product.GoodsID = int(goodsID)
			product.Specifications = string(s)
			_, err = db.NamedExec("INSERT INTO mall_goods_product (goods_id, specifications, price, number, url, created_at, updated_at) VALUES (:goodsId, :specifications, :price, :number, :url, :createdAt, :updatedAt)", product)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
