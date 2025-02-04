package gorm_api

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"testing"
)

type Student struct {
	Sid   uint `gorm:"primaryKey"`
	Name  string
	Login string
	Age   int
	Gpa   float32
}

func (Student) TableName() string {
	return "student"
}

type Course struct {
	Cid  string `gorm:"primaryKey"`
	Name string
}

func (Course) TableName() string {
	return "course"
}

type CourseDetail struct {
	Cid  string `gorm:"primaryKey"`
	Name string
	soft_delete.DeletedAt
}

func (CourseDetail) TableName() string { return "course_detail" }

type Enrolled struct {
	Sid   uint   `gorm:"primaryKey"`
	Cid   string `gorm:"primaryKey"`
	Grade byte
}

func (Enrolled) TableName() string {
	return "enrolled"
}

func TestCreateStuCourse(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&Student{}, &Course{}, &Enrolled{}, &CourseDetail{}))
	/*Err(t, _txTenant().Create(&[]Student{
		{53666, "RZA", "rza@cs", 55, 4.0},
		{53688, "Taylor", "swift@cs", 27, 3.9},
		{53655, "Tupac", "shakur@cs", 25, 3.5},
	}).Error)
	Err(t, _txTenant().Create(&[]Course{
		{"15-445", "Database Systems"},
		{"15-721", "Advanced Database Systems"},
		{"15-826", "Data Mining"},
		{"15-799", "Special Topics in Databases"},
	}).Error)
	Err(t, _txTenant().Create(&[]Enrolled{
		{53666, "15-445", 'C'},
		{53688, "15-721", 'A'},
		{53688, "15-826", 'B'},
		{53655, "15-445", 'B'},
		{53666, "15-721", 'C'},
	}).Error)
	Err(t, _txTenant().Create(&[]CourseDetail{
		{Cid: "6.828", Name: "Distributed System"},
		{Cid: "6.823", Name: "Computer System Architecture"},
		{Cid: "6.827", Name: "Multithreaded Parallelism: Languages and Compilers"},
	}).Error)*/
}

func TestCountStu(t *testing.T) {
	var cnt int64
	Err(t, _txTenant().Raw("SELECT COUNT(login) AS cnt FROM student WHERE login LIKE '%@cs'").Find(&cnt).Error)
	t.Log(cnt)
	Err(t, _txTenant().Raw("SELECT COUNT(*) AS cnt FROM student WHERE login LIKE '%@cs'").Find(&cnt).Error)
	t.Log(cnt)
	Err(t, _txTenant().Raw("SELECT COUNT(1) AS cnt FROM student WHERE login LIKE '%@cs'").Find(&cnt).Error)
	t.Log(cnt)
	Err(t, _txTenant().Raw("SELECT COUNT(1+1+1) AS cnt FROM student WHERE login LIKE '%@cs'").Find(&cnt).Error)
	t.Log(cnt)
}

func TestAvgGpa(t *testing.T) {
	var res struct {
		AvgGpa   float32
		CountSid int
	}
	Err(t, _txTenant().Raw(`
SELECT AVG(gpa) AS avg_gpa, COUNT(sid) AS count_sid
FROM student
WHERE login LIKE '%@cs'`).Find(&res).Error)
	t.Logf("%#v", res)
}

func TestGpaJoin(t *testing.T) {
	var resQ []struct {
		SName string
		SGpa  float32
		ECid  string
	}
	Err(t, _txTenant().Raw(`
SELECT s.name AS s_name, s.gpa AS s_gpa, e.cid AS e_cid
FROM enrolled AS e
JOIN student AS s
ON e.sid = s.sid`).Find(&resQ).Error)
	t.Log(Enc(resQ))

}

func TestGpaErr(t *testing.T) {
	var resG []struct {
		AvgGpa float32
		Cid    string
	}

	// wrong group by
	Err(t, _txTenant().Raw(`
SELECT AVG(s.gpa) AS avg_gpa, e.cid AS cid
FROM enrolled AS e
JOIN student AS s
ON e.sid = s.sid`).Find(&resG).Error)
	t.Log(Enc(resG))

	// correct group by
	Err(t, _txTenant().Raw(`
SELECT AVG(s.gpa) AS avg_gpa, e.cid AS cid
FROM enrolled AS e
JOIN student AS s
ON e.sid = s.sid
GROUP BY e.cid`).Find(&resG).Error)
	t.Log(Enc(resG))
}

func TestGpaHaving(t *testing.T) {
	var resG []struct {
		AvgGpa float32
		Cid    string
	}

	Err(t, _txTenant().Raw(`
SELECT AVG(s.gpa) AS avg_gpa, e.cid AS cid
FROM enrolled AS e
JOIN student AS s
ON e.sid = s.sid
GROUP BY e.cid
HAVING avg_gpa > 3.9
`).Find(&resG).Error)
	t.Log(Enc(resG))
}

func TestGpaNameErr(t *testing.T) {
	var resG2 []struct {
		AvgGpa float32
		Cid    string
		SName  string
	}

	// wrong group by
	Err(t, _txTenant().Raw(`
SELECT AVG(s.gpa) AS avg_gpa, e.cid AS cid, s.name as s_name
FROM enrolled AS e
JOIN student AS s
ON e.sid = s.sid
GROUP BY e.cid`).Find(&resG2).Error)
	t.Log(Enc(resG2))

	// correct group by
	Err(t, _txTenant().Raw(`
SELECT AVG(s.gpa) AS avg_gpa, e.cid AS cid, s.name as s_name
FROM enrolled AS e
JOIN student AS s
ON e.sid = s.sid
GROUP BY e.cid, s.name`).Find(&resG2).Error)
	t.Log(Enc(resG2))
}

// Output Redirection
func TestSelectInsert(t *testing.T) {
	var cid []string
	Err(t, _txTenant().Raw(`
INSERT INTO course (cid, name)
SELECT cid,name FROM course_detail WHERE deleted_at = ?
RETURNING cid
`, soft_delete.FlagActived).Scan(&cid).Error)
	t.Log(cid)
}

func TestSelectCreateTable(t *testing.T) {
	var cid []string
	Err(t, _txTenant().Raw(`
CREATE TEMP TABLE IF NOT EXISTS tmp_course
AS 
SELECT cid,name FROM course_detail WHERE deleted_at = ?
`, soft_delete.FlagActived).Scan(&cid).Error)
	t.Log(cid)
}

func TestCreateModelReturningID(t *testing.T) {
	Err(t, _txTenant().Create(&CourseDetail{
		Cid:  "6.033",
		Name: "Computer System Engineering",
	}).Error)
	Err(t, _txTenant().Model(&CourseDetail{}).Create(map[string]any{
		"Cid":  "6.092",
		"Name": "Introduction to Programming in Java",
	}).Error)
}

func TestQueryWindows(t *testing.T) {
	var gradeRank []struct {
		Enrolled
		RowLine   int
		GradeRank int
	}
	Err(t, _txTenant().Raw(`
SELECT *, 
ROW_NUMBER() OVER (PARTITION BY cid) row_line,
RANK() OVER (PARTITION BY cid ORDER BY grade DESC) grade_rank
FROM enrolled;
`).Scan(&gradeRank).Error)
	t.Log(Enc(gradeRank))
}

type Dict struct {
	gorm.Model
	Type int
	PID  uint
	Key  string
	Val  string
}

func TestCreateDict(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&Dict{}))
	Err(t, _txTenant().Create(&[]Dict{
		{Type: 1, PID: 1, Key: "Lang", Val: "C++, Go, Java"},
		{Type: 1, PID: 1, Key: "Lang", Val: "C++"},
		{Type: 1, PID: 1, Key: "Lang", Val: "Go"},
	}).Error)
}

func TestSelfJoin(t *testing.T) {
	var DictAndPDict []struct {
		Dict  `gorm:"embedded;embeddedPrefix:dict_"`
		PDict Dict `gorm:"embedded;embeddedPrefix:p_dict_"`
	}
	Err(t, _txTenant().Raw(`
SELECT
dict.id AS dict_id,
dict.created_at AS dict_created_at,
dict.type AS dict_type,
dict.key AS dict_key,
dict.val AS dict_val,
p_dict.id AS p_dict_id,
p_dict.key AS p_dict_key,
p_dict.val AS p_dict_val
FROM 
	(SELECT * FROM tbl_dict WHERE deleted_at IS NULL AND type = ?) AS dict
LEFT JOIN 
	(SELECT * FROM tbl_dict WHERE deleted_at IS NULL AND type = ?) AS p_dict
ON dict.p_id = p_dict.id
`, 1, 1).Scan(&DictAndPDict).Error)
	t.Log(Enc(DictAndPDict))
}

func TestWithCTE(t *testing.T) {
	var res struct {
		Cnt int
		Now string
	}
	Err(t, _txTenant().Raw(`
WITH
    cte1 AS (SELECT CAST(1+1 AS INTEGER) AS cnt),
    cte2 AS (SELECT CAST(DATETIME() AS DATETIME) AS now)
SELECT cnt, now FROM cte1, cte2
`).Scan(&res).Error)
	t.Log(Enc(res))
}
