# diary
## 環境構築
### コンテナを作成する

```
$ docker-compose build
```

### コンテナー立ち上げ

```
$ docker-compose up -d
```

### apiの確認

```
$ docker-compose exec api go run main.go

Hello golang from docker!
```

### mysqlの中に入る

```
$ docker ps 

CONTAINER ID        IMAGE                       
bc8beb62daf4        mysql:5.7.22   

$ docker exec -it bc8 bash
$ mysql -u mysql -p -P 3307
Enter password: mysql

mysql> use diary;
mysql> show tables;

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

```# todo-app
