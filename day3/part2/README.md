--- Part Two ---
As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact. If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more accurate result.

There are two new instructions you'll need to handle:

The do() instruction enables future mul instructions.
The don't() instruction disables future mul instructions.
Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.

For example:

xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
This corrupted memory is similar to the example from before, but this time the mul(5,5) and mul(11,8) instructions are disabled because there is a don't() instruction before them. The other mul instructions function normally, including the one at the end that gets re-enabled by a do() instruction.

This time, the sum of the results is 48 (2*4 + 8*5).

Handle the new instructions; what do you get if you add up all of the results of just the enabled multiplications?

-------------------------------------
--- パート2 ---
破損したメモリをスキャンすると、条件文の一部も無傷のまま残っていることに気づきます。プログラム内の無傷の条件文の一部を処理すれば、さらに正確な結果を得られるかもしれません。
処理する必要がある新しい命令は2つあります。
do()命令は、将来のmul命令を有効にします。
don't()命令は、将来のmul命令を無効にします。
最も新しい do()または don't()命令のみが適用されます。プログラムの開始時には、mul命令は有効化されています。
例えば：
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
この破損したメモリは、以前の例と似ていますが、今回はmul(5,5)およびmul(11,8)命令が、その前にdon't()命令があるため無効化されています。他のmul命令は正常に機能し、最後にdo()命令によって有効化されるものも含みます。
今回は、結果の合計は48（2*4 + 8*5）となります。
新しい命令を処理します。有効になっている乗算の結果をすべて合計すると、何になりますか？