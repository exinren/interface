package entity

import "time"

type Student struct {
	Id string	`json:"id" table:"id" type:"BIGINT UNSIGNED" isNullable:"false" len:"20" unique:"true" comment:"学生编号"`
	Name string  `json:"name" table:"name" type:"VARCHAR" isNullable:"false" len:"120" unique:"false" comment:"学生名称"`
	Age int `json:"age" table:"age" type:"INT UNSIGNED" isNullable:"false" len:"10" unique:"false" comment:"学生年龄"`
	Sex int `json:"sex" table:"sex" type:"TINY UNSIGNED" isNullable:"false" len:"3" unique:"false" comment:"学生性别:1@男，2@女"`
	IsDel int `json:"isDel" table:"is_del" type:"TINY UNSIGNED" isNullable:"false" len:"3" unique:"false" comment:"逻辑删除：1@未删除，2@删除"`
	CreateDate       time.Time `json:"createDate" table:"create_date" type:"DATETIME" isNullable:"false" len:"19" unique:"false" comment:"创建时间"`
	ModifyDate       time.Time `json:"modifyDate" table:"modify_date" type:"DATETIME" isNullable:"false" len:"19" unique:"false" comment:"最后更新"`
}

func (s *Student) TableName() string {
	return "t_student"
}

