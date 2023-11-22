# [Template] Fiber web framework made by Golang Like Express
Expressに感銘を受け、ExpressライクなGolang製WebフレームワークFiberのアーキテクチャのテンプレート

---

## 1. 動かすために必要なこと

### 1.1 .envファイルの作成
```shell
cp .env.sample .env
```
envファイルの中身を特定のものに書き換えてください

---

## 2. 起動・停止方法
### 2.1 起動方法
imageの作成
```shell
docker compose build
```

imageからコンテナの起動
```shell
docker compose up -d
```

---

### 2.2 停止方法
```shell
docker compose down
```

---

## 3. 必要最低限のエンドポイント

### <span style="color:yellow">POST</span> サインアップ処理
```text
http://localhost:8080/api/v1/signup
```

#### Request Body
```json
{
  "email": "sample@sample.com",
  "password": "sample"
}
```

---

### <span style="color:yellow">POST</span> ログイン処理
```text
http://localhost:8080/api/v1/login
```

#### Request Body
```json
{
    "email": "sample@sample.com",
    "password": "sample"
}
```

---

### <span style="color:green">GET</span> ユーザーの情報を全て取得する
```text
http://localhost:8080/api/v1/user
```

#### Request Headers
| Header Keys    | Header Values        |
|----------------|----------------------|
| Authorization  | Bearer `<jwt token>` |


---

### <span style="color:yellow">POST</span> ヘルスチェック用エンドポイント
```text
http://localhost:8080/test
```

#### Request Body
```json
{
  "userName": "Test User"
}
```

---

### <span style="color:green">GET</span> パスパラメータの取得
```text
http://localhost:8080/path/3
```

---

### <span style="color:green">GET</span> クエリパラメータの取得
```text
http://localhost:8080/query?id=56
```

---
