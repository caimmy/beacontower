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

// date     : 2018/1/11 10:07
// author   : caimmy@hotmail.com

package orm

import (
	"database/sql"
	"reflect"
	"fmt"
	"errors"
)

func parseModelRows(i interface{}) (map[string] reflect.Type, error) {
	elem := reflect.TypeOf(i).Elem()
	if elem.NumField() > 1 {
		ret_elem_def := make(map[string] reflect.Type)
		for i := 0; i < elem.NumField(); i++ {
			item_name := elem.Field(i).Name
			if "OrmModel" != item_name {
				tag_name := elem.Field(i).Tag.Get("orm")
				if tag_name == "" {
					tag_name = item_name
				}
				ret_elem_def[tag_name] = elem.Field(i).Type
			}
		}
		return ret_elem_def, nil
	} else {
		return nil, errors.New("length of Model is 0")
	}
}

type OrmModel struct {
	db_engine			*OrmEngine
	instance			interface{}

	rows_def 			map[string] reflect.Type
}

func (model *OrmModel) parseModelRows() bool {
	res, err := parseModelRows(model.instance)
	if err == nil {
		model.rows_def = res
		return true
	} else {
		return false
	}
}

func (model * OrmModel) Query() sql.Rows  {
	return sql.Rows{}
}

func (model *OrmModel)Update() int64 {
	return 0
}

func (model *OrmModel)Delete() int64{
	return 0
}

func (model *OrmModel)Save() bool {
	fmt.Println(fmt.Sprintf("%v", model.instance))
	model.parseModelRows()
	return true
}

func (model *OrmModel)SetInstance(v interface{}) {
	model.instance = v
}

func (model *OrmModel) SetEngine(engine *OrmEngine) {
	model.db_engine = engine
}

type IDbModel interface {
	SetInstance(v interface{})
	SetEngine(e *OrmEngine)
	Query() sql.Rows
	Update() int64
	Delete() int64
	Save() bool
}

func NewDbModel(m interface{}, engine *OrmEngine) interface{} {
	t := reflect.ValueOf(m).Type()
	v := reflect.New(t).Interface()
	v.(IDbModel).SetInstance(v)
	v.(IDbModel).SetEngine(engine)
	return v
}

func Find(model interface{}, sql string, engine *OrmEngine, ret *[]interface{}, args ...interface{}) bool {
	model_elem_def, e := parseModelRows(model)
	v := reflect.ValueOf(model)
	t := reflect.Indirect(v).Type()
	add_one := reflect.New(t).Elem()
	add_one.FieldByName("Name").SetString("abcd1234")
	add_one.FieldByName("Idcard").SetString("oop~~~0123456789")

	fmt.Println(add_one.Interface())
	fmt.Println("-----------------------------")
	*ret = append(*ret, add_one)
	if e == nil {
		fmt.Println(model_elem_def)
		/*
		fmt.Println(model_elem_def)
		data_rows, err := engine.Raw(sql).FetchResults(args...)
		if err == nil {
			cols, _ := data_rows.Columns()

			for data_rows.Next(){

			}
		}
		*/



	}
	return true
}
