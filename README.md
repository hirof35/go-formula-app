# 構造的数式シミュレーター (Formula App)

Go言語の高速・堅牢なバックエンドロジックと、Webフロントエンド技術を組み合わせた、軽量なネイティブデスクトップ数式評価アプリケーションです。
Wailsアーキテクチャを採用し、Electron代替としての極めて軽量なバイナリサイズ（十数MB〜数十MB）と高速な起動を実現しています。
<img width="983" height="735" alt="スクリーンショット 2026-05-23 051927" src="https://github.com/user-attachments/assets/ec26a81c-a8bb-469d-9f66-a9ca1eac5af5" />
## 🛠 構造・アーキテクチャ

本アプリケーションは、データモデル（Go）とプレゼンテーション層（HTML5/JavaScript）を疎結合に保ち、Wailsによる自動バインディング（RPC）を介してシームレスに同期する構造を持っています。

- **コアロジック (Go):** 計算評価、型安全なデータ処理、メモリ効率を最大化するバックエンド。
- **UIレイヤー (Vite + Vanilla JS):** OSネイティブのWebViewを利用した軽量かつ高自由度なユーザーインターフェース。

---

## 🚀 クイックスタート

### 前提条件

開発環境およびビルドには以下のツールが必要です。

1. **Go**: `1.18` 以上 (最新の安定版を推奨)
2. **Node.js**: `16` 以上 (フロントエンドのビルドに使用)
3. **Wails CLI**: 以下のコマンドでインストールしてください。
```bash
   go install [github.com/wailsapp/wails/v2/cmd/wails@latest](https://github.com/wailsapp/wails/v2/cmd/wails@latest)
Windows環境での注意: アプリのレンダリングには WebView2 Runtime が必要です（Windows 11等では標準搭載されています）。

開発モードの起動
コードの変更がバックエンド・フロントエンド双方にリアルタイムで反映（ホットリロード）される開発モードです。

Bash
# プロジェクトのルート（wails.jsonが存在する階層）で実行
wails dev
本番用バイナリのパッケージング
各OSネイティブの単一実行ファイル（Windowsなら .exe、macOSなら .app）を生成します。フロントエンド資産はバイナリ内部に embed されます。

Bash
wails build
生成された成果物は build/bin/ ディレクトリ内に出力されます。

📂 プロジェクト構成
Plaintext
formula-app/
├── main.go               # アプリケーションのエントリーポイント
├── app.go                # アプリケーションロジック（Goのコア・数式評価）
├── wails.json            # Wailsプロジェクト設定ファイル
├── build/                # ビルド生成物・アイコン等のアセット
└── frontend/             # フロントエンドのソースコード（Vite環境）
    ├── index.html        # メインHTML面
    ├── src/
    │   ├── main.js       # UIイベントハンドリングとGoバインディングの呼び出し
    │   └── style.css     # スタイリング
    └── wailsjs/          # Wailsにより自動生成されるJSブリッジ（Git管理除外）
💡 今後の拡張設計（ロードマップ）
抽象構文木（AST）パーサーの統合: 現在の固定数式シミュレーションから、ユーザーが任意の数式文字列（例: 3 + x * (sin(x) ^ 2)）を完全自由入力して動的評価できるインタープリタ構造へのアップグレード。

データビジュアライゼーション: フロントエンドに Chart.js や D3.js を導入し、Go側で計算した座標配列をリアルタイムにグラフ描画する機能の追加。
