package autocode

var modelTemp = "package model\n\nimport (\n    . \"github.com/huwhy/commons/basemodel\"\n)\n\ntype {{camel .Table true}} struct {\n    {{range .Columns}}{{camel .ColumnName true}} {{typeName .DataType .Length}} `json:\"{{camel .ColumnName false}}\"`\n    {{end}}BaseModel  `gorm:\"embedded\"`\n}\n\nfunc (m *{{camel .Table true}}) TableName() string {\n\treturn \"{{.Table}}\"\n}\n\ntype {{camel .Table true}}Term struct {\n    *Term\n}\n"

var daoTemp = "package dao\n\nimport (\n\t\"github.com/huwhy/commons/base_dao\"\n\t. \"github.com/huwhy/commons/basemodel\"\n\t\"gorm.io/gorm\"\n\t\"{{.ModPath}}/constant\"\n\t\"{{.ModPath}}/model\"\n)\n\ntype {{camel .Table true}}Dao struct {\n\tbase_dao.BaseDao\n}\n\nfunc New{{camel .Table true}}Dao(db *gorm.DB) *{{camel .Table true}}Dao {\n\t{{camel .Table false}}Dao := &{{camel .Table true}}Dao{BaseDao: base_dao.NewBaseDao(db, constant.LOG)}\n\tif db == nil {\n\t\t{{camel .Table false}}Dao.DB = constant.DAO\n\t}\n\treturn {{camel .Table false}}Dao\n}\n\nfunc (dao *{{camel .Table true}}Dao) Add(po *model.{{camel .Table true}}) error {\n\trs := dao.DB.Create(po)\n\treturn rs.Error\n}\n\nfunc (dao *{{camel .Table true}}Dao) Delete(id int) error {\n\trs := dao.DB.Exec(\"update {{.Table}} set is_deleted=1 where id=?\", id)\n\treturn rs.Error\n}\n\nfunc (dao *{{camel .Table true}}Dao) Update(po *model.{{camel .Table true}}) error {\n\trs := dao.DB.Model(po).Updates(po)\n\treturn rs.Error\n}\n\nfunc (dao *{{camel .Table true}}Dao) Get(id int) (*model.{{camel .Table true}}, error) {\n\tvar po *model.{{camel .Table true}}\n\trs := dao.DB.Where(\"id=?\", id).First(&po)\n\tif rs.Error != nil {\n\t\tif rs.Error == gorm.ErrRecordNotFound {\n\t\t\treturn nil, nil\n\t\t} else {\n\t\t\treturn nil, rs.Error\n\t\t}\n\t}\n\treturn po, nil\n}\n\nfunc (dao *{{camel .Table true}}Dao) Query(term *model.{{camel .Table true}}Term) ([]model.{{camel .Table true}}, error) {\n\tvar sql = \"select * from {{.Table}} where is_deleted=0\"\n\tvar args = make([]interface{}, 0)\n\tvar pos = make([]model.{{camel .Table true}}, 0)\n\terr := dao.Paging(page, sql, \"\", args, &pos)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn pos, nil\n}\n"

var bizTemp = "package biz\n\nimport (\n\t. \"github.com/huwhy/commons/basemodel\"\n\t\"{{.ModPath}}/dao\"\n\t\"{{.ModPath}}/model\"\n)\n\ntype {{camel .Table true}}Biz struct {\n}\n\nfunc New{{camel .Table true}}Biz() *{{camel .Table true}}Biz {\n\treturn &{{camel .Table true}}Biz{}\n}\n\nfunc (b *{{camel .Table true}}Biz) Get(id int) (*model.{{camel .Table true}}, error) {\n\treturn dao.New{{camel .Table true}}Dao(nil).Get(id)\n}\n\nfunc (b *{{camel .Table true}}Biz) Add({{camel .Table false}} *model.{{camel .Table true}}) error {\n\treturn dao.New{{camel .Table true}}Dao(nil).Add({{camel .Table false}})\n}\n\nfunc (b *{{camel .Table true}}Biz) Delete(id int) bool {\n\treturn dao.New{{camel .Table true}}Dao(nil).DeleteInt(id, \"{{.Table}}\")\n}\n\nfunc (b *{{camel .Table true}}Biz) Update({{camel .Table false}} *model.{{camel .Table true}}) error {\n\treturn dao.New{{camel .Table true}}Dao(nil).Update({{camel .Table false}})\n}\n\nfunc (b *{{camel .Table true}}Biz) Query(term *model.{{camel .Table true}}Term) (*Page, error) {\n\treturn dao.New{{camel .Table true}}Dao(nil).Query(term)\n}\n"

var apiTemp = "package api\n\nimport (\n\t. \"github.com/huwhy/commons/basemodel\"\n\t\"github.com/kataras/iris/v12\"\n\t\"go.uber.org/zap\"\n\t\"{{.ModPath}}/biz\"\n\t\"{{.ModPath}}/constant\"\n\t\"{{.ModPath}}/model\"\n)\n\ntype {{camel .Table true}}Api struct {\n\t*BaseApi\n}\n\nfunc (c *{{camel .Table true}}Api) Save(ctx iris.Context) *Json {\n\tvar {{camel .Table false}} = new(model.{{camel .Table true}})\n\tvar err error\n\tif err = ctx.ReadJSON({{camel .Table false}}); err != nil {\n\t\treturn c.HandleErr(err)\n\t}\n\tif {{camel .Table false}}.ID > 0 {\n\t\terr = biz.New{{camel .Table true}}Biz().Update({{camel .Table false}})\n\t} else {\n\t\terr = biz.New{{camel .Table true}}Biz().Add({{camel .Table false}})\n\t}\n\tif err != nil {\n\t\treturn c.HandleErr(err)\n\t}\n\treturn JsonData({{camel .Table false}}.ID)\n}\n\nfunc (c *{{camel .Table true}}Api) Delete(ctx iris.Context) *Json {\n\tid, err := ctx.Params().GetInt(\"id\")\n\tif err != nil {\n\t\tconstant.LOG.Error(\"系统错误\", zap.Field{Key: \"error\", String: err.Error()})\n\t\treturn JsonFail(\"系统异常，请联系客服\")\n\t}\n\tres := biz.New{{camel .Table true}}Biz().Delete(id)\n\treturn JsonData(res)\n}\n\nfunc (c *{{camel .Table true}}Api) Get(ctx iris.Context) *Json {\n\tid, err := ctx.Params().GetInt(\"id\")\n\tif err != nil {\n\t\tconstant.LOG.Error(\"系统错误\", zap.Field{Key: \"error\", String: err.Error()})\n\t\treturn JsonFail(\"系统异常，请联系客服\")\n\t}\n\tif id == 0 {\n\t\treturn JsonFail(\"参数错误\")\n\t}\n\t{{camel .Table false}}, err := biz.New{{camel .Table true}}Biz().Get(id)\n\tif err != nil {\n\t\treturn JsonFail(err.Error())\n\t}\n\treturn JsonData({{camel .Table false}})\n}\n\nfunc (c *{{camel .Table true}}Api) List(ctx iris.Context) *Json {\n\tvar term = new(model.{{camel .Table true}}Term)\n\tif err := ctx.ReadJSON(term); err != nil {\n\t\treturn JsonFail(err.Error())\n\t}\n\t{{camel .Table false}}s, err := biz.New{{camel .Table true}}Biz().Query(term)\n\tif err != nil {\n\t\treturn JsonFail(err.Error())\n\t}\n\treturn JsonData({{camel .Table false}}s)\n}\n"
