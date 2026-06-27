# quiz-app

Go の学習用クイズアプリ。HTMX + Pico CSS で軽量に実装しています。

## 必要環境

- Go 1.22+（ルーティングパターン `GET /path` を使用）

## 実行

```bash
go run main.go
```

http://localhost:8080 を開く。

## ビルド

```bash
go build -o quiz-app .
./quiz-app
```

## 仕組み

- サーバーが `html/template` で HTML を直接描画（JSON API なし）
- HTMX がフォーム送信をインターセプトし、採点結果を部分差替
- Pico CSS がセマンティック HTML を自動整形（クラス名不要）
- 正解はサーバー側で判定（クライアントのソースからは見えない）
