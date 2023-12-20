docker for mac がない場合はインストールが必要

コンテナ起動

```
$ docker-compose up -d
```

コンテナが立ち上がっていることを確認

```
$ docker-compose ps
```

他のコンテナのポートに接続できてるか確認（確認につかったスクリプトとりあえずおいとく）

```
$ docker-compose exec sandbox go run protcol_check.go
```

ローカルの ./sandbox 配下をマウントしているので、docker-compose.ymlのsandboxのイメージを他の言語にかえれば上記のようにつかえる（はず）
