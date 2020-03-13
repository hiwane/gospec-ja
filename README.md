Go プログラミング言語仕様
==========================

- version 2019/07/31
- https://golang.org/ref/spec のなんちゃって日本語訳
- @ はどうすんべ
- 文字 = letter or character: どう区別する?

# イントロダクション

本文書は Go プログラミング言語のリファレンスマニュアルです．
他の情報源としては，[golang.org](https://golang.org/) を参照されたい．

Go はシステムプログラミングを念頭において設計された，汎用言語です．
強い型付け，ガベージコレクションを持ち，並行プログラミング@ (concurrent programming) を明示的にサポートしています．
プログラムは，
性質が依存関係の効率的な管理を実現する *パッケージ*　で構成される．

文法はコンパクト，かつ，規則的@ (regular) であり，
統合開発環境などの自動ツールによる簡易な解析が可能になります．

# 表記

構文は，Extended Backus-Naur Form（EBNF）によって示します．

```
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Production は， Term と以下の高い演算子によって構成される式です．
上から，優先順位が低いものから高いものになっています．

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

小文字の production 名は，字句トークンを識別するために使用されます．
非末端は CamelCase です．
字句トークンはダブルクオーテーション "" またはバッククォート `` で囲まれます．

形式 `a … b` は，`a` から `b` の文字列の代替に使用します．
3点リーダー `…` は @@@．
文字 `…` (3文字・・・ とは対照的に) は Go言語のトークンではありません．

# ソースコード表現


ソースコードは UTF-8 でエンコードされた Unicode テキストです．
テキストは正規化（canonicalized）されていない．
そのため，単一の・・・@@@

各符号位置@ (code point) は区別され，例えば，大文字と小文字の英字は異なる文字として扱われます．

実装上の制限：
他のツールとの互換性のために，
コンパイラはソースコード上で NUL 文字 (`U+0000`) を許可しない場合があります．

実装上の制限：
他のツールとの互換性のために，
コンパイラは，UTF-8でエンコードされたバイトオーダーマーク (BOM) (`U+FEFF`)
がソースコード上の最初の Unicoe 符号位置であるときに
無視する場合があります．

## 文字 (characters)

以下の term は，特定の Unicode 文字クラス (character class) を示すために使用されます．

```
newline        = /* Unicode 符号位置 U+000A */ .
unicode_char   = /* 任意の newline を除いた Unicode 符号位置 */ .
unicode_letter = /* "文字 / Letter" に分類される符号位置 */ .
unicode_digit  = /* "数, 10進数 / Number, decimal digit" に分類される符号位置 */ .
```

[The Unicode Standard 8.0](https://www.unicode.org/versions/Unicode8.0.0/) では，
4.5節 "General Category" 節は文字カテゴリの集合を定義しています．
Go は，Lu, Ll, Lt, Lm, Lo  @@@


訳注：

```
Lu = Letter, uppercase
Ll = Letter, lowercase
Lt = Letter, titlecase
Lm = Letter, modifier
Lo = Letter, other
Nd = Number, decimal digit
```

## 英字と数字 (letters and digits)

訳注：letter を character を区別するため，letter は英字と訳す

@@@@


# 字句要素

## コメント

コメントはプログラムのドキュメントとして機能します．
次の２つの形式があります．

1. 行コメント： 文字列 `//` から，行末まで
2. 一般的なコメント： `/*` で始まり，最初の後続の `*/` まで．（訳注：入れ子不可 `/* ABC /* DEF */ GHI */` では GHI はコメントではない）

コメントは，ルーン (@ rune) や文字列リテラル@ 内，または，コメント内では開始できません．
改行を含まない「一般的なコメント」はスペースのように機能します．
他のコメントは改行のように機能します．

## トークン

トークンは Go言語の語彙を形成します．
識別子 (identifier)，
キーワード (keyword)，
演算子 (operator) と句読点 (punctuation)，
リテラル (literal) の 4 つのクラスがあります．
空白 (`U+0020`)，
水平タブ (`U+0009`)，
キャリッジ リターン (`U+000D`)，
改行文字 (ラインフィード，LF)(`U+000A`) から形成される
ホワイトスペースは，
単一に結合するであろうトークンを分離する場合を除いて無視される．
また，改行文字とファイルの末尾はセミコロンを挿入するトリガーになる場合があります．
入力をトークンに分割する間，
次のトークンは有効なトークンを形成する最長の文字シーケンスです．

訳注：`abcde` は `abc`と`de`と分かれることはなく，最長の `abcde`である

## セミコロン


正式な文法では，セミコロン ";" を多くの production の終端として使用する．
Goプログラムでは，次の 2 つの規則を利用して，多くの場合セミコロンを省略できる．

1. 入力をトークンに分割するとき，
行の最後のトークンが以下のとき，その後ろに
セミコロンが自動的に挿入される．
  - 識別子
  - 整数，浮動小数点数，虚数，ルーン，文字列リテラル
  - キーワード break, continue, fallthrough, return
  - 演算子や句読点 ++, --, ), ], }
2. 複雑な文が 1行を専有できるようにするには，閉じカッコ ")", "}" の前では省略できる
[To allow complex statements to occupy a single line, a semicolon may be omitted before a closing ")" or "}".]


慣用的な使用を反映するために，
本ドキュメントのコード例では，これらの規則によりセミコロンを省略する．

## 識別子

識別子 (identify) は，変数 (variable) や型 (type) などのプログラムエンティティ (entity) に名付ける．
識別子は，1 つ以上の英字 (letter) と数字 (digit) の列です．
識別子は，英字から始まらなければなりません．

```
identifier = letter { letter | unicode_digit } .
```

```
a
_x9
ThisVariableIsExported
αβ
```


いくつかの識別子は事前宣言%%されています．

## キーワード

以下のキーワードは予約されていて，識別子と使用できません．

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## 演算子と句読点

以下の文字列は識別子 (代入演算子 (assignment operators) を含む) と句読点 (punctuation) です．

```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

## 整数リテラル

整数リテラル (integer literal) は，
整数定数を表現する数字の列です．
オプションの接頭辞は，
非10進数を表現する．
`0b` と `0B` は2進数，
`0o` と `0O` は8進数，
`0x` と `0X` は16進数を表現する．
単独の `0` は10進数のゼロとみなされる．
16進数では，英字 `a-f` と `A-F` がそれぞれ `10-15` の値を表す．

読みやすさのため，
アンダースコア `_` が接頭辞の後ろ，または，続く数字との間に使用される場合がある．
このアンダースコアは，リテラルの値を変更しない．

```
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
```

```
42
4_2
0600
0_600
0o600
0O600       // 2文字目は大文字 `O` である
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // 整数リテラルではなく，識別子
42_         // 不当: _ は連続する数字を区切る必要がある
4__2        // 不当: _ は一度にひとつのみ
0_xBadFace  // 不当: _ は連続する数字を区切る必要がある
```


