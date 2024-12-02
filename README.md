# sample-koyeb-todo-for-go
## 概要
このリポジトリは、golangで作ったKoyebで動くRestAPIのTODOアプリケーションサンプルコードです

## ローカル環境実行
```shell
$ make run
```

## 機能一覧
| 機能 | パス | メソッド |
| ---- | ---- | ---- |
| TODO一覧取得 | /todos | GET |
| TODO作成 | /todos | POST |
| TODO取得 | /todos/{id} | GET |
| TODO更新 | /todos/{id} | PUT |
| TODO削除 | /todos/{id} | DELETE |

## Example
### TODO一覧取得
```shell
$ curl http://localhost:8000/todos
{"todos":[]}
```
### TODO作成
```shell
$ curl -X POST -H "Content-Type: application/json" \
-d '{"name":"name1", "description":"sample"}' http://localhost:8000/todos
{"id":"3657736e-2b10-4c5e-a1e7-3b225fe5ae16","name":"name1","description":"sample","is_done":false}
```
### TODO取得
```shell
$ curl http://localhost:8000/todos/3657736e-2b10-4c5e-a1e7-3b225fe5ae16
{"id":"3657736e-2b10-4c5e-a1e7-3b225fe5ae16","name":"name1","description":"sample","is_done":false}
```
### TODO更新
```shell
$ curl -X PUT -H "Content-Type: application/json" \
-d '{"name":"name1", "description":"sample2"}' http://localhost:8000/todos/3657736e-2b10-4c5e-a1e7-3b225fe5ae16
{"id":"3657736e-2b10-4c5e-a1e7-3b225fe5ae16","name":"name1","description":"sample2","is_done":false}
```
### TODO削除
```shell
$ curl -X DELETE http://localhost:8000/todos/3657736e-2b10-4c5e-a1e7-3b225fe5ae16
```
