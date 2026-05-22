// Wailsが自動生成するGoのバインディング関数をインポート
import { EvaluateFormula } from '../wailsjs/go/main/App.js';

const calcBtn = document.getElementById('calcBtn');
const resultOutput = document.getElementById('resultOutput');

calcBtn.addEventListener('click', async () => {
    const a = document.getElementById('paramA').value;
    const b = document.getElementById('paramB').value;
    const c = document.getElementById('paramC').value;
    const x = document.getElementById('paramX').value;

    try {
        // Goのバックエンドロジックを非同期で直接コール
        const res = await EvaluateFormula(a, b, c, x);
        
        if (res.isErr) {
            resultOutput.className = "result error";
            resultOutput.innerText = res.result;
        } else {
            resultOutput.className = "result success";
            resultOutput.innerText = `計算結果 f(${x}) = ${res.result}`;
        }
    } catch (err) {
        resultOutput.innerText = "通信エラーが発生しました: " + err;
    }
});