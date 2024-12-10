--- Day 3: Mull It Over ---
"Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"

The computer appears to be trying to run a program, but its memory (your puzzle input) is corrupted. All of the instructions have been jumbled up!

It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).

Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the results of the multiplications?

--------------------------------
--- 3日目：よく考えて ---
「うちのコンピューターに問題が発生しているから、チーフ・ヒストリアンが在庫にあるかどうかはわからないよ。でも、倉庫を調べてみるのは大歓迎だよ」と、ノースポール・トボガンレンタルショップの店員は少し慌てた様子で言いました。ヒストリアンたちは倉庫に向かいました。
店員があなたに振り返って言いました。「うちのコンピューターにまた問題が発生している理由がわかるかな？
コンピュータはプログラムを実行しようとしているように見えますが、メモリ（パズルの入力）が破損しています。すべての指示がごちゃ混ぜになっています！
プログラムの目的は、いくつかの数字を掛けることだけのようです。
mul(X,Y)
のような命令で実行されます。ここで、
X
と
Y
はそれぞれ1～3桁の数字です。例えば、
mul(44,46)
は
44
を
46
で掛けて、結果を
2024
とします。同様に、
mul(123,4)
は
123
を
4
で掛けます
しかし、プログラムのメモリが破損しているため、mul命令の一部のように見える無効な文字も多数存在し、それらは無視する必要があります。mul(4*, mul(6,9!, ?(12,34), または mul (2, 4)のようなシーケンスは何も行いません。
例えば、破損したメモリの次の部分を考えてみましょう。
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
ハイライトされた4つのセクションのみが、実際のmul命令です。各命令の結果を合計すると、161（2*4 + 5*5 + 11*8 + 8*5）が得られます。
破損したメモリから、破損していないmul命令を探します。すべての乗算結果を合計するとどうなりますか？