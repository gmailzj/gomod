package main

/**
// //插入数据
	// stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	// checkErr(err)

	// res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)
	// id := int64(2)
	// fmt.Println(id)
	// //更新数据
	// stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err := stmt.Exec("astaxieupdate", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	//查询数据
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created time.Time
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Print(gin.H{
	// 		"uid":        uid,
	// 		"username":   username,
	// 		"department": department,
	// 		"created":    created,
	// 	})
	// }

	//删除数据
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)
	// defer stmt.Close() //关闭之

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// Get user value
	// 带参数的路由
	// 获取路由匹配的参数
	router.GET("/user/:name", func(c *gin.Context) {
		// 下面两种方式都可以
		user := c.Params.ByName("name")
		user2 := c.Param("name")
		dbMap["foo"] = "aaa" + user2
		value, ok := dbMap[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// 获取querystring参数
	router.GET("/querystring", func(c *gin.Context) {

		// 获取参数?name=abc,如果没有取默认值
		name := c.DefaultQuery("name", "Guest") //可设置默认值
		// 是 c.Request.URL.Query().Get("lastname") 的简写
		fmt.Printf("%T", name)

		lastname := c.Query("lastname")
		// ct := c.Header("Content-Type","text/html; charset=utf-8")

		// 获取请求头
		headVersion := c.GetHeader("version")
		fmt.Printf("%T", headVersion)
		// 设置响应头
		c.Header("lastname", lastname)
		c.String(http.StatusOK, name)
		// fmt.Println("Hello %s", name)
	})
	// POST 传递参数
	router.POST("/post", func(c *gin.Context) {
		// 获取url里面的参数
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		message := c.PostForm("message")
		name := c.PostForm("name")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// 路由群组
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/user", func(c *gin.Context) {
			//查询数据
			rows, err := db.Query("SELECT * FROM userinfo")
			checkErr(err)
			data := gin.H{
				"errCode": 0,
			}

			list := []gin.H{}
			for rows.Next() {
				var uid int
				var username string
				var department string
				var created time.Time
				err = rows.Scan(&uid, &username, &department, &created)
				checkErr(err)
				list = append(list, gin.H{
					"uid":        uid,
					"username":   username,
					"department": department,
					"created":    created,
				})
			}
			data["list"] = list

			c.JSON(http.StatusOK, data)
		})
		v1.GET("/list", func(c *gin.Context) {
			c.String(http.StatusOK, "list-v1")
		})

	}
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	router.GET("/json", func(c *gin.Context) {
		// data := []int{1, 2, 3}
		// c.JSON(http.StatusOK, gin.H{"errCode": 0, "msg": "abc", "data": data})

		var msg struct {
			// 右边的tag用来映射返回结果中的key
			Name    string `json:"user" xml:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/uuid", func(c *gin.Context) {
		// get a UUID instance
		uuidMy := guuid.New()
		str := uuidMy.String()

		_ = uuid.NewV4()
		demo.Get()
		c.String(http.StatusOK, str)
	})

	router.GET("/xml", func(c *gin.Context) {
		data := []int{1, 2, 3}
		c.XML(http.StatusOK, gin.H{"errCode": 0, "msg": "abc", "data": data})
	})
 */
/*
   curl -X "POST" "http://127.0.0.1:8080/admin" \
		   -H 'Content-Type: application/json; charset=utf-8' \
		   -u 'foo:bar' \
		   -d $'{
		   "value": "1"
   }'
*/