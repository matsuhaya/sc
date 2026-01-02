# sc - VS Code ショートカット管理CLI

## 概要

VS Codeのショートカットを忘れたときに思い出すためのCLIツール。
自分で登録したショートカットをシンプルなリスト形式で表示する。

## 技術スタック

- 言語: Go
- 設定ファイル形式: YAML

## コマンド

### `sc`

引数なしで実行すると、登録済みショートカットを一覧表示（`sc list`と同じ）。

### `sc new`

新しいショートカットを追加する。

```bash
sc new "Cmd+Shift+P" "コマンドパレット"
```

### `sc list`

登録済みショートカットをシンプルなリスト形式で表示する。

### `sc edit`

設定ファイル（YAML）をエディタで開く。
編集・削除はこのコマンドで直接YAMLを編集して行う。

### `sc help`

ヘルプを表示する。

### `sc version`

バージョンを表示する。

## 設定ファイル

### 場所

`~/.config/sc/shortcuts.yaml`

### 形式

```yaml
shortcuts:
  - key: "Cmd+Shift+P"
    description: "コマンドパレット"
  - key: "Cmd+P"
    description: "ファイル検索"
  - key: "Cmd+Shift+F"
    description: "全体検索"
```

## 表示形式

シンプルなリスト形式で出力する。

```
Cmd+Shift+P  コマンドパレット
Cmd+P        ファイル検索
Cmd+Shift+F  全体検索
```

## 実装方針

- 外部ライブラリは最小限に抑える
- 設定ディレクトリが存在しない場合は自動作成する
- `sc edit`はデフォルトで`$EDITOR`環境変数のエディタを使用する（未設定の場合は`vi`）