# 概要

数据工人，数据生成器。通过配置文件，可以批量生成一系列随机数据。

## 应用场景

- 大数据测试
- Demo 演示
- 开发过程中的单元测试

## 特性

- [x] 配置文件，定义输出的字段和一些参数。
- [x] 加载外部字典，获取更多数据。


## 如何使用

### 创建一个 json 格式的配置文件，保存在 config 目录。例如：

```json
{
  "dict_path": "dict",
  "build": {
    "count": 5,
    "format": "json",
    "name": "demo",
    "path": "out"
  },
  "fields": [
    {
      "name": "id",
      "type": "uuid"
    },
    {
      "name": "name",
      "type": "string",
      "dict": "usernames"
    },
    {
      "name": "creat_dt",
      "type": "datetime"
    }
  ]
}
```
- dict_path 指定一个存放字典文件的目录
- build 编译命令
- fields 设置需要输出哪些字段（名称、类型、字典等）

### 通过 -c 参数，指定运行的配置文件。

```shell
data-worker build -c ./config/demo.json
```


## 计划

- 输出 cvs 文件格式
- 邮箱域名字典，支持邮件地址的生成
