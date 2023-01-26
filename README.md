# 概要
Ginを使用した高速Blog

# 前提
- goがインストールされている必要があります。
- MySQLが必要です。
- 環境変数ファイルの作成が必要です。

# 環境変数
```.env
GIN_MODE=debug # releaseにする場合は'release'
SECRET=secret # 任意のシークレット（Sessionで使用）

DB_NAME=db # MySQLのDatabase名
DB_USER=user # MySQLのユーザー名
DB_PASS=password # MySQLのパスワード
DB_PORT=3306 # MySQLのポート（3306が多い）
DB_HOST=localhost # MySQLのホスト（ローカルの場合はlocalhost）

```

# 使用法
1. 上記の環境変数のファイルを'.env'として作成し、rootディレクトリに配置
2. 以下コマンドを実行

```
go run .
```