// 创建管理界面数据库
// use administrator
// 创建用户信息文档
db.createCollection("sys_users", {
    capped: false,
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["username", "password", "login_time"],
            properties: {
                username: {
                    bsonType: "string",
                    minLength: 6,
                    maxLength: 20,
                    description: "用户名不能为空"
                },
                password: {
                    bsonType: "string",
                    minLength: 0,
                    description: "用户密码不能为空"
                },
                login_time: {
                    bsonType: "timestamp",
                    description: "登录时间为必须字段"
                }
            }
        }
    },
    validationLevel: "strict",
    validationAction: "error",
    storageEngine: {
        wiredTiger: {
            configString: "block_compressor=snappy"
        }
    },
    collation: {locale: "en", strength: 2}
})
// 创建索引
db.sys_user.createIndex({username: 1}, {unique: true})
// 插入默认用户信息
db.sys_user.insertOne({
    uuid: "1",
    username: "go-spider@golang.com",
    password: "$2a$10$Mw7g19d5d7k0hhv1lbUU0.czsAG/Iz4PNglzMeNryuqrS1RmYIxx.",
    login_time: "1719782220"
})