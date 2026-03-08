# not-vibe-coding

生成AIにコードを書かせずに、基礎的な部分や標準パッケージの使い方をまとめるリポジトリ
GOでの実装で不明点が出た際に参照するマニュアルになるのが理想

## 方針
- 生成AIにコードを書かせない（自分で実装して理解する）
- 標準パッケージを使いこなす
- 技術的な疑問点は生成AIに質問してOK

## 技術スタック
- Go
- PostgreSQL
- 標準パッケージ: `net/http`, `database/sql`, `encoding/json`

## アプリケーション起動

```bash
go run cmd/api/main.go
```

サーバーは `http://localhost:8080` で起動

## API仕様

```
GET    /health         ヘルスチェック
GET    /todos          Todo一覧取得
POST   /todos          Todo作成
GET    /todos/{id}     Todo取得
PATCH  /todos/{id}     Todo更新（部分更新）
DELETE /todos/{id}     Todo削除
```

## おぼえたこと

### HTTPメソッド
- GET: データ取得
- POST: 新規作成
- PATCH: 部分更新（一部のフィールドのみ）
- PUT: 完全な置き換え（全フィールド必須）
- DELETE: 削除

今回はタイトルだけ更新するのでPATCHをつかってく

### HTTPステータスコード
- 200 OK: 成功（データを返す）
- 201 Created: 作成成功
- 204 No Content: 成功（データを返さない、削除時）
- 400 Bad Request: リクエストが不正
- 404 Not Found: リソースが存在しない
- 500 Internal Server Error: サーバーエラー

削除成功時は204 No Contentだけ返せばOK

### エラーハンドリング
- `errors.Is(err, sql.ErrNoRows)` でエラーの種類を判定
- `sql.ErrNoRows` はクエリ結果が0件の場合のエラー
- validator層は`error`型を返すだけ、handler層で`WriteError()`でHTTPレスポンスに変換

### 標準パッケージ
- `strings.TrimPrefix()` 文字列の接頭辞を削除
- `strconv.Atoi()` 文字列を整数に変換
- `json.NewEncoder(w).Encode()` 構造体をJSONに変換してレスポンスに書き込む
- `json.NewDecoder(r.Body).Decode()` リクエストボディのJSONを構造体にデコード

### データベース操作
- `db.Query()` 複数行取得、`defer rows.Close()`で必ずクローズ
- `db.QueryRow()` 1行取得
- `db.Exec()` INSERT/UPDATE/DELETE実行
- `result.RowsAffected()` 影響を受けた行数を取得

### その他メモ
- ハンドラ: HTTPリクエストを処理する関数
- 高階関数: 関数を返す関数（クロージャを含む）
- `:=` は短縮変数宣言（少なくとも1つ新しい変数が必要）
- `=` は既存変数への代入
- `err`変数の再利用は問題なし（各ifブロックでreturnしてるから）

