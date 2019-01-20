package dao

import (
	"database/sql" // 使用 mysql 包中的初始化方法 init()，但不导出包内的标识符
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// DAO 结构体（类）
type DAO struct {
	// 连接对象
	dsn string
	// 操作数据库的抽象层对象
	db *sql.DB

	// SQL的子句
	// "table":"student"
	// stmts map[string]string

	// 字段列表子句
	field string
	// from 子句
	from string
	as   string
	// join
	joinType    []string // joinType[0] left
	joinTable   []string // joinTable[0] class
	joinTableAs []string // joinTableAs[0] c
	joinOn      []string // joinOn[0] t.class_id=c.class_id
	// where子句独立存储
	whereCond string
	whereArgs []interface{}
	// groupBy
	groupBy string
	// having
	havingCond string
	havingArgs []interface{}
	// order
	orderBy string
	// limit子句
	limit  string
	offset string
}

/**
 * 构造函数
 * @param {[type]} DSN string) (*DAO, error [description]
 */
func New(DSN string) (*DAO, error) {
	// 打开连接池
	db, err := sql.Open("mysql", DSN)

	if err != nil {
		return nil, err
	}

	// 实例化DAO对象
	dao := &DAO{
		dsn: DSN,
		db:  db,
	}
	// 初始化 属性
	dao.Reset()

	return dao, nil
}

// 重置属性
func (this *DAO) Reset() *DAO {
	this.field = "*"
	this.from = ""
	this.as = ""
	this.joinType = []string{}
	this.joinTable = []string{}
	this.joinTableAs = []string{}
	this.joinOn = []string{}
	this.whereCond = ""
	this.whereArgs = []interface{}{}
	this.groupBy = ""
	this.havingCond = ""
	this.havingArgs = []interface{}{}
	this.orderBy = ""
	this.limit = ""
	this.offset = ""
	return this
}

/**
 * 查询字段
 */
func (this *DAO) Field(f string) *DAO {
	this.field = f
	return this
}

/**
 * 设置表名
 * @param  {[type]} this *DAO)         Table(tableName string) (*DAO [description]
 * @return {[type]}      [description]
 */
func (this *DAO) Table(tableName string) *DAO {
	// this.stmts["table"] = tableName
	this.from = tableName
	return this
}

/**
 * 设置别名
 */
func (this *DAO) As(tableAs string) *DAO {
	this.as = tableAs
	return this
}

/**
 * join连接查询
 */
func (this *DAO) Join(tp, table, as, on string) *DAO {
	// 设置join属性
	this.joinType = append(this.joinType, tp)
	this.joinTable = append(this.joinTable, table)
	this.joinTableAs = append(this.joinTableAs, as)
	this.joinOn = append(this.joinOn, on)

	return this
}
func (this *DAO) LeftJoin(table, as, on string) *DAO {
	return this.Join("LEFT", table, as, on)
}
func (this *DAO) RightJoin(table, as, on string) *DAO {
	return this.Join("RIGHT", table, as, on)
}
func (this *DAO) InnerJoin(table, as, on string) *DAO {
	return this.Join("INNER", table, as, on)
}

/**
 * 设置where条件
 */
func (this *DAO) Where(where string, args ...interface{}) *DAO {
	this.whereCond = where
	this.whereArgs = args
	return this
}

/**
 * 分组
 */
func (this *DAO) GroupBy(gb string) *DAO {
	this.groupBy = gb
	return this
}
func (this *DAO) Having(cond string, args ...interface{}) *DAO {
	this.havingCond = cond
	this.havingArgs = args
	return this
}

/**
 * 排序
 */
func (this *DAO) OrderBy(ob string) *DAO {
	this.orderBy = ob
	return this
}

/**
 * 设置limit
 */
func (this *DAO) Limit(size string) *DAO {
	this.limit = size
	return this
}

/**
 * 设置offset
 */
func (this *DAO) Offset(offset string) *DAO {
	this.offset = offset
	return this
}

/**
 * [func description]
 *
 * @param  {[type]} this *DAO)         Insert() (int64, error [description]
 * @return int64	最新生成的自动增长ID
 * @return error
 * @example dao.Insert({"name": "hank", "age": 42})
 */
func (this *DAO) Insert(fields map[string]interface{}) (int64, error) {
	// INSERT INTO 表名 (字段列表) VALUES (值占位符列表)

	// 一：拼凑以上结构的SQL
	// 字段列表（fields下标），占位符列表，值变量列表（fields值）
	fieldsNum := len(fields)
	cols := make([]string, fieldsNum)
	phs := make([]string, fieldsNum)
	vals := make([]interface{}, fieldsNum)
	i := 0
	for col, val := range fields {
		cols[i] = "`" + col + "`"
		phs[i] = "?"
		vals[i] = val
		i++
	}
	// 拼凑成一个字符（join）
	colsStr := strings.Join(cols, ", ")
	phsStr := strings.Join(phs, ", ")
	// 拼凑SQL，格式化字符串
	query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", this.from, colsStr, phsStr)

	// 二，执行
	result, err := this.db.Exec(query, vals...) // append(sli1, sli2...)
	if nil != err {
		return 0, err
	}
	// 后面都是成功
	id, err := result.LastInsertId()
	if nil != err {
		// 没有自动增长的ID，但是执行SQL也没有错误
		return 0, nil
	}

	this.Reset()
	return id, nil
}

/**
 * 删除
 * @param  {[type]} this *DAO)         Delete(tableName string, where string, args ...interface{}) (int64, error [description]
 * @return int64	受影响的记录数
 * @return error
 * @example dao.Delete("gender=?", "other")
 */
// func (this *DAO) Delete(where string, args ...interface{}) (int64, error) {
func (this *DAO) Delete() (int64, error) {
	// DELETE FROM 表名 WHERE 条件

	// 一：删除不能没有条件
	if "" == this.whereCond {
		return 0, errors.New("需要指定条件，若需要匹配全部数据请使用 1 作为条件")
	}

	// 二，拼凑SQL
	query := fmt.Sprintf("DELETE FROM `%s` WHERE %s", this.from, this.whereCond)

	// 三，执行
	result, err := this.db.Exec(query, this.whereArgs...) // append(sli1, sli2...)
	if nil != err {
		return 0, err
	}
	// 后面都是成功
	affectedNum, err := result.RowsAffected()
	if nil != err {
		return 0, nil
	}

	this.Reset()
	return affectedNum, nil
}

/**
 * 更新
 * @param  {[type]} this *DAO)         Update(fields map[string]interface{}, where string, args ...interface{}) (int64, error [description]
 * @return int64 受影响的记录数
 */
// func (this *DAO) Update(fields map[string]interface{}, where string, args ...interface{}) (int64, error) {
func (this *DAO) Update(fields map[string]interface{}) (int64, error) {
	// UPDATE 表名 SET 键值列表 WHERE 条件
	// set field1=?, field2=? where gender=? where gender=?

	// 一：更新不能没有条件
	if "" == this.whereCond {
		return 0, errors.New("需要指定条件，若需要匹配全部数据请使用 1 作为条件")
	}

	// 二，拼凑SQL
	// 字段列表（fields下标），占位符列表，值变量列表（fields值）
	fieldsNum := len(fields)
	sets := make([]string, fieldsNum)
	vals := make([]interface{}, fieldsNum)
	i := 0
	for col, val := range fields {
		sets[i] = "`" + col + "`=?"
		vals[i] = val
		i++
	}
	// 拼凑成一个字符（join）
	setsStr := strings.Join(sets, ", ")
	query := fmt.Sprintf("UPDATE `%s` SET %s WHERE %s", this.from, setsStr, this.whereCond)

	// 三，执行，需要的参数有 字段参数和条件参数
	result, err := this.db.Exec(query, append(vals, this.whereArgs...)...) // append(sli1, sli2...)
	if nil != err {
		return 0, err
	}
	// 后面都是成功
	affectedNum, err := result.RowsAffected()
	if nil != err {
		return 0, nil
	}

	this.Reset()
	return affectedNum, nil
}

/**
 * 查询多条
 * @param  {[type]} this *DAO)         FetchRows() ([]map[string]interface{}, error [description]
 * @return {[type]}      [description]
 */
func (this *DAO) FetchRows() ([]map[string]string, error) {
	// SELECT 字段列表 FROM 表名 WHERE 条件 GROUP 分组 HAVING 组后过滤 ORDER BY 排序 LIMIT 结果限定

	// 一，拼凑SQL
	query := fmt.Sprintf("SELECT %s FROM `%s`", this.field, this.from)

	// 别名部分
	if "" != this.as {
		query += fmt.Sprintf(" AS `%s`", this.as)
	}

	// join 部分
	for i, _ := range this.joinType {
		// join table [as a] [on conditiong]
		query += fmt.Sprintf(" %s JOIN %s", this.joinType[i], this.joinTable[i])
		if "" != this.joinTableAs[i] {
			query += fmt.Sprintf(" AS %s", this.joinTableAs[i])
		}
		if "" != this.joinOn[i] {
			query += fmt.Sprintf(" ON %s", this.joinOn[i])
		}
	}

	//  where部分
	if "" != this.whereCond {
		query += fmt.Sprintf(" WHERE %s", this.whereCond)
	}

	// group部分
	if "" != this.groupBy {
		query += fmt.Sprintf(" GROUP BY %s", this.groupBy)
	}
	// having
	if "" != this.havingCond {
		query += fmt.Sprintf(" HAVING %s", this.havingCond)
	}

	// order by
	if "" != this.orderBy {
		query += fmt.Sprintf(" ORDER BY %s", this.orderBy)
	}

	// limit部分
	if "" != this.limit {
		// 需要 limit 部分
		query += fmt.Sprintf(" LIMIT %s", this.limit)

		// offset
		if "" != this.offset {
			// 需要 limit 部分
			query += fmt.Sprintf(" OFFSET %s", this.offset)
		}
	}
	// log.Println(query)
	// query 拼凑完毕，不需要子句参数

	// 二，执行
	rows, err := this.db.Query(query, append(this.whereArgs, this.havingArgs...)...)
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	// 三，查询select存在哪些字段
	cols, err := rows.Columns() // 字段名称
	if nil != err {
		return nil, err
	}

	// 设计可以用于Scan操作的数据结构，设计一个引用的切片，通过 切片... 语法传递进去（不确定有多少字段）
	fields := make([]interface{}, len(cols)) // [&value[0], &value[1], &value[2]]
	values := make([][]byte, len(cols))      // 存储数据，仅获取字符串（字节切片），全部的数据都是字节切片。（rows.Scan()在处理[]byte时没问题）
	for i, _ := range fields {
		fields[i] = &values[i]
	}
	rowsResult := []map[string]string{} // 存储结果

	// 遍历row集合操作
	for rows.Next() {
		rows.Scan(fields...) // &value[0], &value[1], &value[2]

		valsMap := map[string]string{} // 存储每一个行

		// 根据每个字段的类型，形成其合理的结构
		for i, val := range values {
			field := cols[i]
			valsMap[field] = string(val) // 形成 记录map [字段]=字段值 val
		}
		rowsResult = append(rowsResult, valsMap)
	}
	this.Reset() // 重置选项
	return rowsResult, nil
}

/**
 * 查询单条
 * @param  {[type]} this *DAO)         FetchRow(fields map[string]interface{}) (, error [description]
 * @return map[string]interface{}      字段与值 映射 [name: Hank, gender: male]
 * @example dao.Table("stu1").Where("gender=?", "male").FetchRow() 得到：[name: Hank, gender: male]
 */
func (this *DAO) FetchRow() (map[string]string, error) {
	// 一 保证SQL是单条查询， limit 1
	this.Limit("1")

	// 二 执行查询多条，获取第一个元素即可
	rows, err := this.FetchRows()
	if nil != err {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	} else {
		return nil, nil
	}
}

/**
 * 查询第一列
 * @param field string 需要的字段名， id, name
 */
func (this *DAO) FetchColumn(field string) ([]string, error) {
	// 执行查询多行的SQL
	rows, err := this.FetchRows()
	if nil != err {
		return nil, err
	}

	sets := make([]string, len(rows))

	// 遍历全部行，得到第一列的集合即可
	for i, row := range rows {
		sets[i] = row[field]
	}

	return sets, nil
}

/**
 * 查询标量值，得到的是第一行，第一列的值
 */
func (this *DAO) FetchValue() (string, error) {
	// 执行行查询
	row, err := this.FetchRow()
	if nil != err {
		return "", err
	}

	// 返回第一个字段的值
	for _, v := range row {
		return v, nil
	}
	return "", nil
}
