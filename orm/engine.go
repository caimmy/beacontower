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

// date     : 2018/1/11 9:43
// author   : caimmy@hotmail.com

package orm

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ENGINE_VER_MYSQL 		= "mysql"
	ENGINE_CODING_UTF8		= "utf8"
)

/**
	数据库操作引擎
 */
type OrmEngine struct {
	host 		string
	port 		int
	user 		string
	pass 		string
	dbname		string
	engineVer	string			// 数据库引擎版本
	charset		string

	database	*sql.DB
	rawSql		string
}

/**
user:password@tcp(localhost:5555)/dbname?charset=utf8
 */
func (engine *OrmEngine) Init() error {
	db, err := sql.Open(engine.engineVer,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", engine.user, engine.pass, engine.host, engine.port,
			engine.dbname, engine.charset))
	if err == nil {
		engine.database = db
		return nil
	} else {
		log.Print(err)
		return err
	}
}

func (engine *OrmEngine) Close() error {
	return engine.database.Close()
}

/**
执行原始sql语句
 */
func (engine *OrmEngine) Raw(sql string) (*OrmEngine) {
	engine.rawSql = sql
	return engine
}

/**
 执行数据库原生语句的update 和 delete 操作

 */
func (engine *OrmEngine) Exec(args ...interface{}) (int64, error) {
	var ret_affected_rows int64
	result, err := engine.database.Exec(engine.rawSql, args...)
	if err == nil {
		af_rows, qe := result.RowsAffected()
		if qe == nil {
			ret_affected_rows = af_rows
		}
	}
	return ret_affected_rows, err
}

/**
执行原生sql语句的插入操作
 */
func (engine *OrmEngine) Insert(args ...interface{}) (int64, error) {
	var ret_last_id int64
	result, err := engine.database.Exec(engine.rawSql, args...)
	if err == nil {
		last_id, ie := result.LastInsertId()
		if ie == nil {
			ret_last_id = last_id
		}
	}
	return ret_last_id, err
}

func (engine *OrmEngine) FetchResults(args ...interface{}) (*sql.Rows, error) {
	if 0 == len(args) {
		return engine.database.Query(engine.rawSql)
	} else {
		return engine.database.Query(engine.rawSql, args...)
	}
}

func (engine *OrmEngine) RegisterModel(interface{}) error {
	return nil
}

/**
构造并初始化数据库引擎
 */
func NewEngine(host string, port int, user string, pwd string, db string, enginever string, charset string) (*OrmEngine, bool) {
	ori_engine := OrmEngine{host, port, user, pwd, db, enginever,
		charset, nil, ""}
	err := ori_engine.Init()
	if err == nil {
		return &ori_engine, true
	} else {
		return nil, false
	}
}