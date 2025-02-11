
# 文件服务
文件相关的接口
[TOC]

## 整体说明
1.	字符串都为utf8格式;
1.	HTTP Headers:
	1.	Content-Type设置为：application/json
1.	DataTime格式参考RFC3339标准

## 错误处理
错误的具体信息将在error字段中返回。

### 错误码示例
```json
{
    "code": "400",
    "message": "Param Error"
}
```


### 状态码列表
| 状态码 | 说明 |
|---|---|
| 200 | 返回正常 |
| 400 | 参数错误 |
| 401 | 无access<br> key或key无效 |
| 500 | 服务器内部错误 |


## 获取文件详情

### 请求路径
```http
GET v1/files
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `include_deleted` | `boolean` |  | 否 |  | 默认值为false |
| `limit` | `string` |  | 否 |  | 默认值为null |
| `cursor` | `string` |  | 否 |  | 默认值为null |
| `last_updated_after` | `string` | `Timestamp` | 否 |  | 默认值为null |
| `page_size` | `integer` | `Int32` | 否 |  |  |
| `page_token` | `string` |  | 否 |  |  |
| `skip` | `integer` | `Int32` | 否 |  |  |


### 返回值

#### 返回对象
| type | description |
|---|---|
| `Array<files.File>` |  |


#### `files.File`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `size` | `integer` | `Int64` | N |  |
| `permalink` | `string` |  | N |  |
| `dateAdded` | `string` | `Timestamp` | N |  |  |
| `lastUpdated` | `string` | `Timestamp` | N |  |  |
| `isDeleted` | `boolean` |  | N |  |


## 更新文件信息

### 请求路径
```http
PUT v1/files
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `size` | `integer` | `Int64` | N |  |
| `permalink` | `string` |  | N |  |
| `dateAdded` | `string` | `Timestamp` | N |  |  |
| `lastUpdated` | `string` | `Timestamp` | N |  |  |
| `isDeleted` | `boolean` |  | N |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `size` | `integer` | `Int64` | N |  |
| `permalink` | `string` |  | N |  |
| `dateAdded` | `string` | `Timestamp` | N |  |  |
| `lastUpdated` | `string` | `Timestamp` | N |  |  |
| `isDeleted` | `boolean` |  | N |  |


## 新增文件

### 请求路径
```http
POST v1/files
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `size` | `integer` | `Int64` | N |  |
| `permalink` | `string` |  | N |  |
| `dateAdded` | `string` | `Timestamp` | N |  |  |
| `lastUpdated` | `string` | `Timestamp` | N |  |  |
| `isDeleted` | `boolean` |  | N |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `size` | `integer` | `Int64` | N |  |
| `permalink` | `string` |  | N |  |
| `dateAdded` | `string` | `Timestamp` | N |  |  |
| `lastUpdated` | `string` | `Timestamp` | N |  |  |
| `isDeleted` | `boolean` |  | N |  |


## 获取指定的文件详情

### 请求路径
```http
GET v1/files/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `size` | `integer` | `Int64` | N |  |
| `permalink` | `string` |  | N |  |
| `dateAdded` | `string` | `Timestamp` | N |  |  |
| `lastUpdated` | `string` | `Timestamp` | N |  |  |
| `isDeleted` | `boolean` |  | N |  |


## 删除用户

### 请求路径
```http
DELETE v1/files/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
对象为空