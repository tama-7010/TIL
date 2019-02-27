# 1. 字句解析
## 1.1 字句解析

ソースコード→トークン列→抽象構文木

## 1.2 トークンを定義する

最初に、字句解析器が出力するトークンを定義する。
> Type...TokenType  
> Literal...string
>> TokenType...string

## 1.3 字句解析器（レキサー/lexer）

字句解析器にソースコードを与えて(input)初期化し(New)、繰り返しNextToken()を呼ぶことで、ソースコードを読み進めていく。
> input...string  
> position...int  
> readPosition...int  
> ch...byte  

positionは、現在検査中のバイトchの位置を指し示す。readPositionは常に入力における「次の」位置を指し示す。

この字句解析器はASCII文字だけに対応しており、Unicode全体はカバーしていない。
そのためにはchをbyteからruneに変更し、次の文字を読む処理も、複数のバイトに対応できるよう変更する必要がある。
