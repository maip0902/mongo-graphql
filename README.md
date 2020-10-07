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

# gqlgenの実装手順
### クエリを追加したい
1. schema/query.graphql に追加したいクエリのスキーマ定義を書く
2. ルートディレクトリで下記コマンド実行

```
go run scripts/gqlgen.go
```
3. resolverの実装をする

### 新しくモデルを作りたい
1. schema/input.graphql に追加したいモデルのスキーマ定義を書く
2. ルートディレクトリで下記コマンド実行

```
go run scripts/gqlgen.go
```
3. resolverの実装をする

