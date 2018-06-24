# book-writing-an-interpreter-in-go

O'Reilly Japanの『Go言語によるインタプリタ』の学習レポジトリ


## 1章 字句解析

### 1.1 字句解析

```
ソースコード -(字句解析)-> トークン列 -(構文解析器)-> 抽象構文木
```


### 1.2 トークンを定義する

トークンと識別する要素の定義。

* 特殊トークン(不定、ファイル終端)
* 識別子
* リテラル
* 演算子
* デリミタ
* 括弧
* キーワード

など


### 1.3 字句解析(レキサー)

ソースコード文字列を読み込んでトークンに分割する。

この段階では

* ソースコードは基本的にASCII
* リテラルと整数以外は一文字

といった制約付き


### 1.4 トークン集合の拡充と字句解析器の拡張

以下の種類のトークンを追加

* 1文字トークン: "!", "-", "/", "*", ">", "<"
* 2文字トークン: "==", "!="
* キーワードトークン: "true", "false", "if", "else", "return"


## 2章 構文解析

### 2.1 構文解析器(パーサー)

この章で進めて行くこと

* 前章で作成した字句解析器の出力を入力にする
* Monkeyプログラミング言語のインタプリタが要求する使用を満たして行くような独自のASTを定義
* トークンを再帰的に構文解析しながらASTのインスタンスを構築


### 2.2 パーサージェネレータじゃないの？

大体のモノの本だとこの工程はパーサージェネレーターを使ってブラックボックス的に進むことが多い。

無論、一度手書きでパーサーを作ったことがある人やプロダクションでの利用を想定している場合は効率・安全性のために利用する方が良いが、学習目的のためなので、手書きでやる。


### 2.3 Monkey言語のための構文解析器を書く

* トップダウン構文解析
    * 再帰下降構文解析
        * トップダウン演算子優先順位(Pratt構文解析器)
    * アーリー法
    * 予測的構文解析
* ボトムアップ構文解析

まずは文(ステートメント)の構文解析からはじめる。


### 2.4 構文解析器の第一歩：let文

let文の例

```
let x = 5;
let foobar = add(5.5)
let barfoo = 5 * 5 / 10 + 18 - add(5, 5) + multiply(124);
let anotherName = barfoo
```

let文はいずれも

```
let <identifier> = <expression>;
```

という形式になっている。

ASTを実装するために3つのインターフェースを実装する

* Node : ASTのNodeを表す。今回のASTは端点をLeafのような別名を持たない
    * `TokenLiteral() string` : ASTのNodeが関連づけられているトークンのリテラル値を返す　※デバッグとテストに使う
* Statement
    * `Node` : 一部ノードの属性
    * `statementNode()` : ダミーメソッド。コンパイラ支援用
* Expression
    * `Node` : 一部ノードの属性
    * `expressionNode()` : ダミーメソッド。コンパイラ支援用

まず、ASTのルートノードを表す `Program` を以下のように定義する

```.go
type Program struct {
    Statements []Statement
}

// TokenLiteral() を実装したので Program は Node と同様に扱える
func (p *Program) TokenLiteral() string {
    // 登録されている最初の Statement の TokenLiteral() を返す
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}
```

次に `LetStatement` と `Identifier` を作成する。

* `LetStatement` は token.LET トークン Name(変数名) と Value(代入値)を持つ
* `Identifier` は token.IDENT と Value(実値)を持つ

これらを用いると、例えば `let x = 5;` をAST表現すると

```
*ast.Program.Statements
   -> *ast.LetStatement
        Name  -> *ast.Identifier
        Value -> *ast.Expression
```

となる。

トークンをASTに変換する `Parser` を実装する。

Parserは

* 字句解析器インスタンスへのポインタ
* 現在見ているトークン
* １つ先のトークン

の三つのフィールドを持っている。

Newでトークンを2つ読み込み、curTokenとpeekTokenにセットする。

ParseProgramでパースする。


## 2.8 構文解析木の拡張

* `TODO` : `testInfixExpression` ヘルパ関数を使って `parser_test.go` をリファクタリングする。


### 2.8.1 真偽値リテラル

`TestParsingPrefixExpressions` の改修で一部本で記述が省略されているか、見落としかが、あった。

`prefixTests` が整数のみから `interface` に変わったのに伴って `testIntegerLiteral` のところを `testLiteralExpression` に変更する必要がある。
