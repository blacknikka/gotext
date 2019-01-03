# 概要
Go学習用のリポジトリです。

# 環境
コマンドラインでの実行を想定しています。  
macOS上でのみ確認を行なっております。

# 機能
コマンドラインでの実行時に渡された引数からファイルを開き、  
そのファイルの内容を読み込んで中身を解析するようなアプリです。
読み込む内容はTSVを想定しております。

# 実行

- 実行（ビルド後実行）

```
$ make exec
```

- ビルド

```
$ make build
```

- コマンドライン上からの実行  
下記のように２つのテキストファイルを渡してください。

```
$ ./textAnalyzer ./data/input.txt ./data/master.txt
```
