--- Day 7: Bridge Repair ---
The Historians take you to a familiar rope bridge over a river in the middle of a jungle. The Chief isn't on this side of the bridge, though; maybe he's on the other side?

When you go to cross the bridge, you notice a group of engineers trying to repair it. (Apparently, it breaks pretty frequently.) You won't be able to cross until it's fixed.

You ask how long it'll take; the engineers tell you that it only needs final calibrations, but some young elephants were playing nearby and stole all the operators from their calibration equations! They could finish the calibrations if only someone could determine which test values could possibly be produced by placing any combination of operators into their calibration equations (your puzzle input).

For example:

190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
Each line represents a single equation. The test value appears before the colon on each line; it is your job to determine whether the remaining numbers can be combined with operators to produce the test value.

Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers in the equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of operators: add (+) and multiply (*).

Only three of the above equations can be made true by inserting operators:

190: 10 19 has only one position that accepts an operator: between 10 and 19. Choosing + would give 29, but choosing * would give the test value (10 * 19 = 190).
3267: 81 40 27 has two positions for operators. Of the four possible configurations of the operators, two cause the right side to match the test value: 81 + 40 * 27 and 81 * 40 + 27 both equal 3267 (when evaluated left-to-right)!
292: 11 6 16 20 can be solved in exactly one way: 11 + 6 * 16 + 20.
The engineers just need the total calibration result, which is the sum of the test values from just the equations that could possibly be true. In the above example, the sum of the test values for the three equations listed above is 3749.

Determine which equations could possibly be true. What is their total calibration result?

--- Part Two ---
The engineers seem concerned; the total calibration result you gave them is nowhere close to being within safety tolerances. Just then, you spot your mistake: some well-hidden elephants are holding a third type of operator.

The concatenation operator (||) combines the digits from its left and right inputs into a single number. For example, 12 || 345 would become 12345. All operators are still evaluated left-to-right.

Now, apart from the three equations that could be made true using only addition and multiplication, the above example has three more equations that can be made true by inserting operators:

156: 15 6 can be made true through a single concatenation: 15 || 6 = 156.
7290: 6 8 6 15 can be made true using 6 * 8 || 6 * 15.
192: 17 8 14 can be made true using 17 || 8 + 14.
Adding up all six test values (the three that could be made before using only + and * plus the new three that can now be made by also using ||) produces the new total calibration result of 11387.

Using your new knowledge of elephant hiding spots, determine which equations could possibly be true. What is their total calibration result?
--------------------------

--- 7日目：橋の修理 ---
歴史学者があなたをジャングルの中の川にかかる見慣れたロープ橋に案内します。 族長は橋のこちら側にいるわけではありません。もしかしたら向こう側でしょうか？
橋を渡ろうとすると、橋を修理しようとしている技術者のグループに気づきます。（どうやら、かなり頻繁に壊れるようです。）修理が完了するまでは渡ることができません。
どのくらいかかるか尋ねると、エンジニアたちは最終的な調整が必要だと言うが、近くで遊んでいた若い象たちが、調整式からすべてのオペレーターを盗んでしまったのだ！ 誰かが、どのテスト値が、どの組み合わせのオペレーターを調整式に入れることで生み出される可能性があるかを判断できれば、調整は完了する。（パズルの入力）
例えば：
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
各行は1つの方程式を表します。テスト値は各行のコロンの前に表示されます。残りの数字を演算子と組み合わせてテスト値を生成できるかどうかを判断するのはあなたの仕事です。
演算子は優先順位のルールに従ってではなく常に左から右へ評価されます。さらに、方程式内の数字は並べ替えることができません。ジャングルを覗くと、2種類の演算子（加算（+）と乗算（*））を持つ象が見えます。
上記の式のうち、演算子を挿入することで正しい式にできるのは3つだけです。
190: 10 19
 では演算子を挿入できる位置は1箇所のみで、
10
と
19
の間です。 
+
 を選択すると
29
 になりますが、
*
 を選択するとテスト値（
10 * 19 = 190
）になります。
3267: 81 40 27 には演算子を配置できる箇所が2箇所あります。 演算子の4つの可能な配置のうち、右辺がテスト値と一致するのは2つです。81 + 40 * 27と81 * 40 + 27はどちらも3267（左から右に評価した場合）に等しくなります！
292: 11 6 16 20 は、11 + 6 * 16 + 20 という唯一の方法で解くことができます。
エンジニアが知りたいのは、合計の較正結果、つまり、正しい可能性のある方程式のテスト値の合計です。上記の例では、上記の3つの方程式のテスト値の合計は3749です。
どの方程式が正しい可能性があるかを判断します。それらの合計キャリブレーション結果は何でしょうか？

--- パート2 ---
エンジニアたちは心配しているようです。あなたが彼らに提示した校正結果の合計は、安全許容範囲内に収まるには程遠いからです。その時、あなたは自分の間違いに気づきました。いくつかのよく隠された象が3番目の演算子を保持しているのです。
結合演算子（||）は、左側と右側の入力の数字を1つの数字に結合します。例えば、12 || 345は12345になります。すべての演算子は依然として左から右の順に評価されます。
さて、加算と乗算のみで正しい式となる3つの式を除いて、上記の例には演算子を挿入することで正しい式となる3つの式が残っています。
156: 15 6 は、単一の結合によって真とすることができます。15 || 6 = 156。
7290: 6 8 6 15 は、6 * 8 || 6 * 15 によって真とすることができます。
192: 17 8 14
 は、
17 || 8 + 14
 によって真とすることができます。
6つのテスト値すべてを合計すると（+と*のみを使用して導き出せる3つの値に加え、||も使用して導き出せる新しい3つの値）、新しい合計較正結果として11387が得られます。
ゾウの隠れ場所に関する新しい知識を活かして、どの方程式が正しい可能性があるかを判断してください。それらの合計較正結果は何ですか？