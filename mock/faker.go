package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sysatom/mvp/model"
	"sync"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:123456@/mvp?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//GenerateUser(db)
	InsertMallIssue(db)
}

func GenerateUser(db *sqlx.DB) {
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go InsertMallUser(db, &wg)
	}

	wg.Wait()
}

func InsertMallUser(db *sqlx.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	var users []*model.MallUser
	for i := 0; i < 1000; i++ {
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

func InsertMallBrand(db *sqlx.DB) {
	var brands []*model.MallBrand
	for i := 0; i < 1000; i++ {
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

func InsertMallIssue(db *sqlx.DB) {
	var issues []*model.MallIssue
	for i := 0; i < 1000; i++ {
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
