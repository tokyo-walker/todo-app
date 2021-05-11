# diary
## 環境構築
### コンテナを作成する

```shell script
$ docker-compose build
```

### コンテナー立ち上げ

```shell script
$ docker-compose up
```

### mysqlの中に入る

```shell script
$ docker ps 

CONTAINER ID        IMAGE                       
bc8beb62daf4        mysql:5.7.22   

$ docker exec -it bc8 bash
$ mysql -u mysql -p -P 3307
Enter password: mysql

mysql> use todo;
mysql> show tables;

```

## front end 

### Create app by Vite

```shell script
npm init @vitejs/app
npm init @vitejs/app todo-project --template vue
cd todo-project
npm install
```
### Run go app

```go
cd /todo-app/backend/app-api-todo
go build main.go
./main
```

### Run front app

```shell script
cd frontend/todo-project
npm install
npm run dev # you can access http://localhost:3000
```


参考資料
```

https://qiita.com/gold-kou/items/45a95d61d253184b0f33#select
https://qiita.com/katekichi/items/d94e078b376151858ca4
https://micnncim.com/posts/ja/go-naming-convention
https://golang.org/ref/spec#Struct_types
https://zenn.dev/keitakn/articles/go-naming-rules
https://qiita.com/high5/items/4e2580241039c950e1c4
https://qiita.com/ayasuda/items/53933c83d0fb7152c7e9
https://micnncim.com/posts/ja/go-naming-convention
https://www.kwbtblog.com/entry/2020/04/07/055735
https://zenn.dev/konnyaku256/scraps/6e8f78642cde2c
https://qiita.com/zurazurataicho/items/4a95e0daf0d960cfc2f7
https://qiita.com/munieru_jp/items/f305931aca92ef796b3b

vite
https://vitejs.dev/guide/#scaffolding-your-first-vite-project

```
