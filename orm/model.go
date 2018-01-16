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
	"bytes"
)

const (
	// 默认主键的键名，理论上每个数据模型都应该存在id字段，并且是数据表的主键
	DEFAULT_PRIMARY_KEY 			= "Id"
)

func parseModelRows(i interface{}) (map[string] RawDatasetDef, error) {
	elem := reflect.TypeOf(i).Elem()
	retElemDef := make(map[string] RawDatasetDef)
	if elem.NumField() > 1 {
		for i := 0; i < elem.NumField(); i++ {
			itemField := elem.Field(i)
			if "OrmModel" != itemField.Name {
				tagName := itemField.Tag.Get("orm")
				autoIncreament := false
				if "1" == itemField.Tag.Get("autoincreament") {
					autoIncreament = true
				}
				setName := tagName
				if tagName == "" {
					setName = itemField.Name
				}
				rawDatasetDef := RawDatasetDef{itemField.Type.Kind(), itemField.Name, autoIncreament}
				retElemDef[setName] = rawDatasetDef
			}
		}
		return retElemDef, nil
	} else {
		return retElemDef, errors.New("length of Model is 0")
	}
}

type RawDatasetDef struct {
	ValType 			reflect.Kind
	ValName				string
	ValAutoincreament	bool
}

type OrmModel struct {
	db_engine			*OrmEngine
	instance			interface{}
	__tablename__		string

	rows_def     		map[string] RawDatasetDef
	UpdateRecord 		bool
}

func (model *OrmModel) SetTable(tablename string) {
	model.__tablename__ = tablename
}

// 获取模型的主键值，供update操作时使用
func (model *OrmModel) getPkValue() int64 {
	model_values := reflect.ValueOf(model.instance).Elem()
	if model_values.IsValid() {
		primary_field := model_values.FieldByName(DEFAULT_PRIMARY_KEY)
		if primary_field.IsValid() {
			return primary_field.Int()
		}
	}
	return 0
}

func (model *OrmModel) parseModelInfor() {
	model.parseModelRows()
	model.parseTablename()
}

func (model *OrmModel) parseModelRows() bool {
	if model.rows_def == nil {
		res, err := parseModelRows(model.instance)
		if err == nil {
			model.rows_def = res
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func (model *OrmModel) parseTablename() {
	if model.__tablename__ == "" {
		rawStructName := reflect.TypeOf(model.instance).String()
		splited_str := strings.Split(rawStructName, ".")
		model.__tablename__ = strings.ToLower(splited_str[len(splited_str) - 1])
	}
}

/**
内省数据模型自身的属性、字段信息
 */
func (model *OrmModel) reflectModelProperties() (map[string] interface{}, error){
	model_values := reflect.ValueOf(model.instance).Elem()
	if model_values.IsValid() {
		ins_data := make(map[string] interface{})
		for def_k, def_v := range model.rows_def {
			if !def_v.ValAutoincreament {
				field_val := model_values.FieldByName(def_v.ValName)
				if field_val.IsValid() {
					switch field_val.Type().Kind() {
					case reflect.String:
						ins_data[def_k] = field_val.String()
					case reflect.Int:
						fallthrough
					case reflect.Int64:
						ins_data[def_k] = field_val.Int()
					case reflect.Float32:
						fallthrough
					case reflect.Float64:
						ins_data[def_k] = field_val.Float()
					}
				} else {
					return nil, errors.New(fmt.Sprintf("field [%s] reflection is invalid.", def_v.ValName))
				}
			}
		}
		return ins_data, nil
	} else {
		return nil, errors.New("model reflection is invalid.")
	}
}

/**
为sql准备两个列表，一个是字段列表，一个是数据列表。
 */
func (model *OrmModel) makeSqlColumnsAndValues(properties map[string] interface{}) ([]string, []interface{}, error) {
	columns_arr 	:= make([]string, 0)
	values_arr		:= make([]interface{}, 0)
	for _tk, _tv := range properties {
		columns_arr = append(columns_arr, _tk)
		values_arr = append(values_arr, _tv)
	}
	return columns_arr, values_arr, nil
}

func (model *OrmModel)Query() sql.Rows  {
	return sql.Rows{}
}

func (model *OrmModel) insert() (int64, error) {
	reflect_properties, e := model.reflectModelProperties()
	if e == nil {
		columns, values, err := model.makeSqlColumnsAndValues(reflect_properties)
		if err == nil {
			col_sql := strings.Join(columns, ", ")
			val_sql := make([]string, len(values))
			for vi := 0; vi < len(val_sql); vi++ {
				val_sql[vi] = "?"
			}

			construct_buf := bytes.Buffer{}
			construct_buf.WriteString("INSERT INTO ")
			construct_buf.WriteString(model.__tablename__)
			construct_buf.WriteString(" (")
			construct_buf.WriteString(col_sql)
			construct_buf.WriteString(") VALUES (")
			construct_buf.WriteString(strings.Join(val_sql, ","))
			construct_buf.WriteString(")")

			return model.db_engine.Raw(construct_buf.String()).Insert(values...)
		}
	}

	return 0, errors.New("reflectModelProperties failure!")
}

func (model *OrmModel) update() (int64, error) {
	reflect_properties, e := model.reflectModelProperties()
	if e == nil {
		columns, values, err := model.makeSqlColumnsAndValues(reflect_properties)
		if err == nil {
			col_sql := make([]string, len(columns))
			for vi := 0; vi < len(col_sql); vi++ {
				col_sql[vi] = fmt.Sprintf("%s=?", columns[vi])
			}

			construct_buf := bytes.Buffer{}
			construct_buf.WriteString("UPDATE ")
			construct_buf.WriteString(model.__tablename__)
			construct_buf.WriteString(" SET ")
			construct_buf.WriteString(strings.Join(col_sql, ","))
			construct_buf.WriteString(" WHERE id=" + strconv.FormatInt(model.getPkValue(), 10))

			return model.db_engine.Raw(construct_buf.String()).Exec(values...)
		}
	}
	return 0, nil
}

func (model *OrmModel)Delete() int64{
	return 0
}

func (model *OrmModel)Save() (int64, error) {
	model.parseModelInfor()
	if model.UpdateRecord {
		return model.update()
	} else {
		return model.insert()
	}

	return 0, errors.New("Unkown modify operation!")
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

					new_data_label_field := extractDataset.FieldByName("UpdateRecord")
					if new_data_label_field.IsValid() && new_data_label_field.CanSet() && new_data_label_field.Type().Kind() == reflect.Bool {
						new_data_label_field.SetBool(true)
					}

					*ret = append(*ret, extractInstance.Interface())
				} else {
					log.Println(scan_err)
				}


			}

		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}


	return true
}
