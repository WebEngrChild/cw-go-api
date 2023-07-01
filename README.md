# 初期設定

## STEP:1 リポジトリクローン

```shell
$ git clone git@github.com:WebEngrChild/cw-go-api.git
```

## STEP:2 copilotのインストール

```shell
$ brew install aws/tap/copilot-cli
```

## STEP:3 初期設定とデプロイ

```shell
# ディレクトリ移動
$ pwd
/Users/hogehoge/cw-go-api

# アプリケーション初期化
$ copilot app init

# サービス初期設定と初期デプロイ
$ copilot init
```