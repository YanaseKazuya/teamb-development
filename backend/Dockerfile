# backend/Dockerfile

# ベースイメージとして公式のGoイメージを使用
FROM golang:1.21.1

# 作業ディレクトリを設定
WORKDIR /teamb-development

# 依存関係のファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# アプリケーションを実行
CMD ["./main"]
