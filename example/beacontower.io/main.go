// Copyright 2017 jungle Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// date     : 2018/1/5 11:04
// author   : caimmy@hotmail.com

package main

import (
	"github.com/caimmy/beacontower/orm"
	"fmt"
)

type Lotto_User struct {
	Id			int		`orm:"id"`
	Name 		string	`orm:"user_name"`
	Idcard 		string	`orm:"user_weixin"`
	Phone 		string	`orm:"user_phone"`

	orm.OrmModel
}

type UserModel struct {
	Id			int		`orm:"id"`
	Name 		string	`orm:"user_name"`
	Idcard 		string	`orm:"user_weixin"`
	Phone 		string	`orm:"user_phone"`

	orm.OrmModel
}

func main() {
	engine, ok := orm.NewEngine("127.0.0.1", 3306, "kefu", "abcd1234", "lotto",
		orm.ENGINE_VER_MYSQL, orm.ENGINE_CODING_UTF8)
	if ok {
		InsertdateT(engine)
	}
}

func InsertdateT(engine *orm.OrmEngine) {
	tU := Lotto_User{Name: "郭德纲", Idcard:"555555555555", Phone:"1234567890"}
	tU.SetEngine(engine)
	tU.SetInstance(&tU)

	tU.Save()
}

func FindT(engine *orm.OrmEngine) {
	ret_set := make([]interface{}, 0)

	vvv := orm.Find(&UserModel{}, "SELECT * FROM lotto_user WHERE id=?", engine, &ret_set, 6)
	if len(ret_set) > 0 {
		for _, v := range ret_set {
			m := v.(*UserModel)
			fmt.Println(m)
		}
		fmt.Println("----------------------------------")
	}
	if vvv {
		fmt.Println("ok")
	}
}
