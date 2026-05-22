package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// App 構造体。ここに定義したPublicメソッドがフロントエンド（JS）から直接呼べるようになります
type App struct {
	ctx context.Context
}

// NewApp はApp構造体のインスタンスを生成
func NewApp() *App {
	return &App{}
}

// startup はアプリ起動時にWailsから呼ばれ、コンテキストを保持します
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// EvaluateResult はフロントエンドに返す計算結果の構造体
type EvaluateResult struct {
	Result string `json:"result"`
	IsErr  bool   `json:"isErr"`
}

// EvaluateFormula はフロントエンドから呼び出される数式評価メソッド
// 例: f(x) = ax^2 + bx + c の簡易計算ロジック
func (a *App) EvaluateFormula(aStr, bStr, cStr, xStr string) EvaluateResult {
	fa, errA := strconv.ParseFloat(strings.TrimSpace(aStr), 64)
	fb, errB := strconv.ParseFloat(strings.TrimSpace(bStr), 64)
	fc, errC := strconv.ParseFloat(strings.TrimSpace(cStr), 64)
	fx, errX := strconv.ParseFloat(strings.TrimSpace(xStr), 64)

	if errA != nil || errB != nil || errC != nil || errX != nil {
		return EvaluateResult{Result: "数値を正しく入力してください", IsErr: true}
	}

	// 構造計算: f(x) = ax^2 + bx + c
	term1 := fa * math.Pow(fx, 2)
	term2 := fb * fx
	total := term1 + term2 + fc

	return EvaluateResult{
		Result: fmt.Sprintf("%.4f", total),
		IsErr:  false,
	}
}