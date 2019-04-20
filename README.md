# wschat

サーバー起動
```
$ direnv allow
$ cd src/wschat
$ go run main.go
```

確認
```
$ npm install -g wscat
$ wscat -c localhost:12345 -o localhost:12345
```

プレビュー
![wschat](https://user-images.githubusercontent.com/37661826/56454808-3ad5ab00-6391-11e9-974a-b25fb303cc3b.png)
