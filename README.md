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

```ebnf
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
newline        = /* Unicode 符号位置 (code point) U+000A */ .
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

## 浮動小数点リテラル

浮動小数点リテラルは，浮動小数点定数の10進数または　16進数表現である．

10進浮動小数リテラルは，整数部 (10進数)，小数点，小数部 (10進数)，
指数部 (`e` または `E` とオプションの符号，10進数）から成る．
整数部と小数部のどちらか一方は省略でき，
小数点と指数部のどちらか一方は省略できる．
指数値 `exp` は仮数 (整数部と小数部）を `10^{exp}` 倍する．


16進数浮動小数リテラルは，接頭 `0x` または `0X`，
整数部 (16進数)，
基数点 (radix point; 訳注 n進数小数点のこと)
小数部 (16進数)，
指数部 (`p` または `P` とオプションの符号，10進数）から成る．
整数部と小数部のどちらか一方は省略できる．
基数点は省略できるが，指数部は必要である．
(この構文は IEEE 754-2008 §5.12.3. で与えられる構文と一致)
指数値 `exp` は仮数 (整数部と小数部）を `2^{exp}` 倍する．

読みやすさのため，
アンダースコア `_` が接頭辞の後ろ，または，続く数字との間に使用される場合がある．
このアンダースコアは，リテラルの値を変更しない．


```
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
```

```
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction)

0x.p1        // 不当: 仮数部に数字がない
1p-2         // 不当: p 指数には 16進数仮数部が必要
0x1.5e-2     // 不当: 16進化数部は p 指数が必要
1_.5         // 不当: _ は連続する数字を区切る必要がある
1._5         // 不当: _ は連続する数字を区切る必要がある
1.5_e1       // 不当: _ は連続する数字を区切る必要がある
1.5e_1       // 不当: _ は連続する数字を区切る必要がある
1.5e1_       // 不当: _ は連続する数字を区切る必要がある
```

## 虚数リテラル

虚数リテラル (imarinary literal) は複素数定数の虚数部を表す．
虚数リテラルは，整数リテラルまたは浮動小数リテラルと，その後ろに続く小文字英字 `i` から成る．
虚数リテラルの値は，それぞれ，整数リテラルまたは浮動小数リテラルに虚数単位 `i` を掛けた値です．

```
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
```

後方互換のため，
虚数リテラルの整数部 (および場合によってはアンダースコア) が
10進数のみで構成される場合は，
`0` で始まっていたとしても 10進数と見なされる．

```
0i
0123i         // == 123i for backward-compatibility
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
```

## ルーンリテラル

ルーンリテラル (rune literal) は
Unicode 符号位置を特定する整数値である
ルーン定数を表現する．
ルーンリテラルは，`'x'` や `'\n'` のようにシングルクォートで囲まれた 1つ以上の文字たちで表現される．
シングルクォート内では，
改行とエスケープされていないシングルクォートを除く任意の文字を使用できる．
シングルクォートで囲まれた1文字は，その文字自身の Unicode 値を表すが，
バックスラッシュで始まる複数の文字たちの場合は，
様々な形式で値をエンコードする．


最もシンプルな形式は，シングルクォートで囲まれた単一の文字を表現します．
Go のソーステキストは UFT-8 でエンコードされた Unicode 文字たちなので，
複数の UTF-8 エンコードバイトは，単一の整数値を表現する場合がある．
例えば，リテラル `'a'` はリテラル `a`, Unicode `U+0061`，値 `0x61` を保持し，
`'ä'` は 2バイト (`0xc3 0xa4`) は
リテラル `a-ウムラウト`, `U+00E4`, 値 `0xe4` を保持する．

いくつかのバックスラッシュエスケープたちにより，
任意の値を
ASCII テキストとしてエンコードできる．
数定数として整数は 4 つの方法で表現できる:
`\x` と正確に 2 桁の 16進数;
`\u` と正確に 4 桁の 16進数;
`\U` と正確に 8 桁の 16進数;
素のバックスラッシュ `\` と正確に 3 桁の 8 進数.
いずれの場合も，リテラルの値は，対応する基数の数字で表される値である．

これらの表現はすべて整数を表すが，
異なる有効範囲をもつ．
8進エスケープは 0 から 255 までの値を表さなければならない．
16進エスケープは，構成からこの条件を満足する．
'\u' と '\U' のエスケープは Unicode 符号位置を表現するので，
一部の値，特に `0x10FFFF` より大きな値とサロゲートハーフたちは不正 (illegal) です．


バックスラッシュの後，特定の一文字のエスケープは特別な値を表します:

```
\a   U+0007 アラート または ベル
\b   U+0008 バックスペース
\f   U+000C form feed
\n   U+000A ラインフィード (line feed) または ニューライン (newline)
\r   U+000D キャリッジリターン (carriage return)
\t   U+0009 水平タブ
\v   U+000b 垂直タブ
\\   U+005c バックスラッシュ
\'   U+0027 シングルクォート (ルーンリテラル内でのみ有効なエスケープ)
\"   U+0022 ダブルクォート (文字列リテラル内でのみ有効なエスケープescape only within string literals)
```

バックスラッシュで始まる他のすべての列は，ルーンリテラルの中では不正 (illegal) です．

```
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
```

```
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // シングルクォート文字を含むルーンリテラル
'aa'         // illegal: 文字が多すぎる
'\xa'        // illegal: 16進数桁数が少ない
'\0'         // illegal: 8進数の桁数が少ない
'\uDFFF'     // illegal: サロゲートハーフ
'\U00110000' // illegal: 不当は Unicode 符号位置
```

## 文字列リテラル

文字列リテラル (string literal) は，
文字の列を連結して得られる文字列定数を表現する．
生の文字列リテラル (raw string literal) と
解釈された文字列リテラル (interpreted string literal) の
2 つの形式がある．

生の文字列リテラルは
`` `foo` `` のようにバッククォートで囲まれた文字列である．
バッククォート内では，
バッククォートを除く任意の文字を使用できる．
生の文字列リテラルの値は，
バッククォート間の解釈されない（暗黙的に UTF-8 でエンコードされた）
文字で構成される文字列である;
特に，
バックスラッシュには特別な意味はなく，
文字列は改行を含まれる場合がある．
生の文字列リテラル内のキャリッジリターン (`'\r'`) は
生の文字列値から破棄される．

解釈される文字列リテラルは，
`"bar"` のようにダブルクォートで囲まれた文字列である．
ダブルクォート内では，
改行とエスケープされていないダブルクォートを除く任意の文字を使用できる．
ダブルクォート間のテキストは
リテラルの値を形成し，
バックスラッシュエスケープは
ルーンリテラルたち (`\'` は不正，`\"` は正当という点を除く）
であると解釈され， 同様の制限がある．
3桁の 8進数 (`\nnn`) エスケープと 2桁の 16進数 (`\xnn`) エスケープは，
結果の文字列の個々のバイトたちを表す．
すべての他のエスケープは
個々の文字の
(マルチバイト）UTF-8 エンコードを表す．
したがって，
文字列リテラル内の `\377` と `\xFF` は
単一バイトの値 `0xFF = 255` を表し，
`ÿ`, `\u00FF`, `\U000000FF`, `\xc3\xbf` は
文字 `U+00FF` の UTF-8 エンコードの 2 バイト `0xc3 0xbf` を表す．

```ebnf
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```

```
`abc`                // "abc" と同じ
`\n
\n`                  // "\\n\n\\n" と同じ
"\n"
"\""                 // `"` と同じ
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // 不正: サロゲートハーフ
"\U00110000"         // 不正: 不当な Unicode 符号位置
```

以下の例は，すべて同じ文字列の表現である.

```
"日本語"                                // UTF-8 入力テキスト
`日本語`                                // 生のリテラルとしての UTF-8 入力テキスト
"\u65e5\u672c\u8a9e"                    // 明示的な Unicode 符号位置
"\U000065e5\U0000672c\U00008a9e"        // 明示的な Unicode 符号位置
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // 明示的な Unicode バイトたち
```

ソースコードが
アクセントと文字を含む結合形式など，
2 つの符号位置として文字を表す場合，
ルーンリテラル（単一の符号位置ではない）に配置すると，
結果はエラーとなり，
文字列リテラルに配置すると，
2 つの符号位置として使用される．


# 定数

ブール定数，ルーン定数，整数定数，浮動小数点整数，
複素数定数，文字列定数がある．
ルーン定数，整数定数，浮動小数点定数，複素数定数を総称して，数値定数と呼ばれる．

定数値は，ルーンリテラル，整数リテラル，浮動小数点リテラル，虚数リテラル，文字列リテラル，
定数を表す識別子，
定数式，
結果が定数となる変換，
任意の値に適用される `unsafe.Sizeof`，
いくつかの表現に適用される `cap`, `len`，
複素数定数に適用される `real`, `imag`，
数値定数に適用される `complex`
のような build-in 関数の復帰値
によって表現される．
ブール真偽値は
事前宣言された定数 `true` と `false` によって表現される．
事前宣言された識別子 `iota` は整数定数を示す．

一般に，複素数定数は定数式の形式であり，その節で説明される．

数値定数は任意精度の正確な値を表現し，
オーバーフローしない．
したがって，IEEE-754 の負のゼロ，無限大，非数値 (`NaN`) を示す定数はない．

定数は，型付きでも，型なしでもいい．
リテラル定数，`true`, `false`, `iota`, 型なし定数オペランドを含む定数式は，
型なしです．

