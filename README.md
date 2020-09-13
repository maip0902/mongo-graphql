# ローカル環境構築
1. clone

```
git clone github.com/[username]/mongo-graphql
```

2. コンテナをビルド

```
make build
```

3. コンテナを立ち上げる

```
make up
```

4. ブラウザでアクセス
http://localhost:8080

# createクエリのサンプル

```
mutation {
  createUser(input: {email: "test@test.com"}) {
    email
  }
}
```

# アプリケーションコンテナに入りたい時

```
make app
```

# 起動してるコンテナ一覧

```
make ps
```
