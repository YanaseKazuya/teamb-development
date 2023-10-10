# teamb-development
このプロジェクトは、Go言語/Ginを使用したバックエンド、
TypeScript/Vue.js, Nuxt.jsを使用したフロントエンドで構成されています。
Dockerを使用して各サービスをコンテナ化しています。

## ディレクトリ構造（仮）
/teamb-development/<br>
│<br>
├── /backend/                  # Go言語のバックエンドコード<br>
│   ├── main.go               # Goのエントリーポイント<br>
│   └── Dockerfile            # バックエンド用のDockerfile<br>
│   ├── /routes/              # APIのルート定義<br>
│   │   └── routes.go         # APIのルート定義<br>
│   ├── /models/              # データベースモデル<br>
│   │   └── user.go           # ユーザーモデルなどのデータベースモデル<br>
│   ├── /controllers/         # コントローラー<br>
│   │   └── userController.go # ユーザー関連のAPI処理<br>
│   └── /config/              # 設定ファイルやDBの接続設定<br>
│   　   └── db.go             # DBの接続設定<br>
│<br>
├── /frontend/                 # フロントエンドコード<br>
│<br>
├── docker-compose.yml         # Docker Composeの設定ファイル<br>
├── LICENSE                    # MITライセンスなどのコピーライト記載ファイル<br>
└── README.md                  # プロジェクトの説明やセットアップ手順<br>

## セットアップ手順
1. リポジトリをクローンします。
git clone https://github.com/si-cmind/teamb-development.git

2. Dockerイメージをビルドします。
docker-compose build

3. Dockerコンテナを起動します。
docker-compose up

バックエンドは`http://localhost:8080`、フロントエンドは`http://localhost:****`でアクセスできます。

## 使用技術
- バックエンド: Go + Gin
- フロントエンド: TypeScript + Vue.js/Nuxt.js
- サーバー: AWS
- データベース: MySQL (現在はDockerコンテナ内で実行)
- インフラ: Docker

## ライセンス
このプロジェクトは[MITライセンス](LICENSE)の下で公開されています。
