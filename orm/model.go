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
	"log"
	"strconv"
	"strings"
)

func parseModelRows(i interface{}) (map[string] RawDatasetDef, error) {
	elem := reflect.TypeOf(i).Elem()
	retElemDef := make(map[string] RawDatasetDef)
	if elem.NumField() > 1 {
		for i := 0; i < elem.NumField(); i++ {
			itemField := elem.Field(i)
			if "OrmModel" != itemField.Name {
				tagName := itemField.Tag.Get("orm")
				setName := tagName
				if tagName == "" {
					setName = itemField.Name
				}
				rawDatasetDef := RawDatasetDef{itemField.Type.Kind(), itemField.Name}
				retElemDef[setName] = rawDatasetDef
			}
		}
		return retElemDef, nil
	} else {
		return retElemDef, errors.New("length of Model is 0")
	}
}

type RawDatasetDef struct {
	ValType 		reflect.Kind
	ValName			string
}

type OrmModel struct {
	db_engine			*OrmEngine
	instance			interface{}
	__tablename__		string

	rows_def 			map[string] RawDatasetDef
	bUpdateRecord		bool
}

func (model *OrmModel) SetTable(tablename string) {
	model.__tablename__ = tablename
}

func (model *OrmModel) parseModelInfor() {
	model.parseModelRows()
	model.parseTablename()
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

func (model *OrmModel) parseTablename() {
	if model.__tablename__ == "" {
		rawStructName := reflect.TypeOf(model.instance).String()
		splited_str := strings.Split(rawStructName, ".")
		model.__tablename__ = strings.ToLower(splited_str[len(splited_str) - 1])
	}
}

func (model *OrmModel)Query() sql.Rows  {
	return sql.Rows{}
}

func (model *OrmModel)Insert() int64 {
	return 0
}

func (model *OrmModel)Update() int64 {
	return 0
}

func (model *OrmModel)Delete() int64{
	return 0
}

func (model *OrmModel)Save() bool {
	model.parseModelInfor()
	if model.bUpdateRecord {
		fmt.Println("脏数据，做更新操作！")
	} else {
		fmt.Println("新数据，做插入操作！")
	}
	fmt.Println(fmt.Sprintf("%v", model.instance))

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
	modelElementDefs, _ := parseModelRows(model)
	dataset, err := engine.Raw(sql).FetchResults(args...)
	if err == nil {
		query_cols, qce := dataset.Columns()
		if err == nil && qce == nil {
			instanceType := reflect.Indirect(reflect.ValueOf(model)).Type()
			// 构造接收数据表字段的参数
			tmpGabageArrPrt := make([] interface{}, len(query_cols))
			tmpGabageArr := make([] interface{}, len(query_cols))

			for i := range tmpGabageArr {
				tmpGabageArrPrt[i] = &tmpGabageArr[i]
			}
			for dataset.Next() {
				extractInstance := reflect.New(instanceType)
				extractDataset := extractInstance.Elem()

				scan_err := dataset.Scan(tmpGabageArrPrt...)
				if scan_err == nil {
					for i_pos, v := range query_cols {
						preDefineItem, ok := modelElementDefs[v]
						if ok {
							fieldItem := extractDataset.FieldByName(preDefineItem.ValName)
							if fieldItem.IsValid() && fieldItem.CanSet() {

								switch fieldItem.Type().Kind() {
								case reflect.String:
									fieldStr := string(tmpGabageArr[i_pos].([]byte))
									fieldItem.SetString(fieldStr)
								case reflect.Int:
									switch tmpGabageArr[i_pos].(type) {
									case int64:
										fieldVal := tmpGabageArr[i_pos].(int64)
										fieldItem.SetInt(fieldVal)
									case []byte:
										fieldVal := string(tmpGabageArr[i_pos].([]byte))
										tmpInt, atoiErr := strconv.Atoi(fieldVal)
										if atoiErr == nil {
											fieldItem.SetInt(int64(tmpInt))
										}
									}
								}
							}
						}
					}
				} else {
					log.Println(scan_err)
				}

				fieldMethodSetInstance := extractInstance.MethodByName("SetInstance")
				setInstanceParams := make([]reflect.Value, 1)
				setInstanceParams[0] = extractInstance
				if fieldMethodSetInstance.IsValid() {
					fieldMethodSetInstance.Call(setInstanceParams)
				}

				fieldMethodSetEngine := extractInstance.MethodByName("SetEngine")
				setEngineParams := make([]reflect.Value, 1)
				setEngineParams[0] = reflect.ValueOf(engine)
				if fieldMethodSetEngine.IsValid() {
					fieldMethodSetEngine.Call(setEngineParams)
				}

				*ret = append(*ret, extractInstance.Interface())
			}

		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}


	return true
}
