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

type UserModel struct {
	Id			int		`orm:"hhh"`
	Name 		string	`orm:"eee"`
	Idcard 		string	`orm:"bbb"`
	Phone 		string	`orm:"abc"`

	orm.OrmModel
}

func main() {
	engine, ok := orm.NewEngine("127.0.0.1", 3306, "root", "abcd1234", "gov_info",
		orm.ENGINE_VER_MYSQL, orm.ENGINE_CODING_UTF8)
	if ok {
		ret_set := make([]interface{}, 0)
		m := orm.Find(&UserModel{}, "SELECT * FROM mgr_user", engine, &ret_set, nil)
		e := ret_set[0].(UserModel)
		fmt.Println("*****************************")
		fmt.Println(e)
		if m {
			fmt.Println("ok")
		}
	}


}
