--- Day 13: Claw Contraption ---
Next up: the lobby of a resort on a tropical island. The Historians take a moment to admire the hexagonal floor tiles before spreading out.

Fortunately, it looks like the resort has a new arcade! Maybe you can win some prizes from the claw machines?

The claw machines here are a little unusual. Instead of a joystick or directional buttons to control the claw, these machines have two buttons labeled A and B. Worse, you can't just put in a token and play; it costs 3 tokens to push the A button and 1 token to push the B button.

With a little experimentation, you figure out that each machine's buttons are configured to move the claw a specific amount to the right (along the X axis) and a specific amount forward (along the Y axis) each time that button is pressed.

Each machine contains one prize; to win the prize, the claw must be positioned exactly above the prize on both the X and Y axes.

You wonder: what is the smallest number of tokens you would have to spend to win as many prizes as possible? You assemble a list of every machine's button behavior and prize location (your puzzle input). For example:

Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
This list describes the button configuration and prize location of four different claw machines.

For now, consider just the first claw machine in the list:

Pushing the machine's A button would move the claw 94 units along the X axis and 34 units along the Y axis.
Pushing the B button would move the claw 22 units along the X axis and 67 units along the Y axis.
The prize is located at X=8400, Y=5400; this means that from the claw's initial position, it would need to move exactly 8400 units along the X axis and exactly 5400 units along the Y axis to be perfectly aligned with the prize in this machine.
The cheapest way to win the prize is by pushing the A button 80 times and the B button 40 times. This would line up the claw along the X axis (because 80*94 + 40*22 = 8400) and along the Y axis (because 80*34 + 40*67 = 5400). Doing this would cost 80*3 tokens for the A presses and 40*1 for the B presses, a total of 280 tokens.

For the second and fourth claw machines, there is no combination of A and B presses that will ever win a prize.

For the third claw machine, the cheapest way to win the prize is by pushing the A button 38 times and the B button 86 times. Doing this would cost a total of 200 tokens.

So, the most prizes you could possibly win is two; the minimum tokens you would have to spend to win all (two) prizes is 480.

You estimate that each button would need to be pressed no more than 100 times to win a prize. How else would someone be expected to play?

Figure out how to win as many prizes as possible. What is the fewest tokens you would have to spend to win all possible prizes?

--- Part Two ---
As you go to win the first prize, you discover that the claw is nowhere near where you expected it would be. Due to a unit conversion error in your measurements, the position of every prize is actually 10000000000000 higher on both the X and Y axis!

Add 10000000000000 to the X and Y position of every prize. After making this change, the example above would now look like this:

Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=10000000008400, Y=10000000005400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=10000000012748, Y=10000000012176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=10000000007870, Y=10000000006450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=10000000018641, Y=10000000010279
Now, it is only possible to win a prize on the second and fourth claw machines. Unfortunately, it will take many more than 100 presses to do so.

Using the corrected prize coordinates, figure out how to win as many prizes as possible. What is the fewest tokens you would have to spend to win all possible prizes?

----------------------------------------------
--- 13日目：クレーンゲーム機 ---
次は、熱帯の島のリゾートのロビーです。 歴史学者たちは、六角形の床タイルをしばらく眺めてから、散らばります。
幸いにも、そのリゾートには新しいアーケードゲームセンターがあるようです！クレーンゲーム機で景品が当たるかもしれませんよ？
ここのクレーンゲーム機は少し変わっています。クレーンを制御するのにジョイスティックや方向ボタンではなく、このゲーム機には「
A」
と「
B」
と書かれた2つのボタンがあります。さらに悪いことに、コインを入れるだけでプレイできるわけではなく、
A
ボタンを押すには3枚のコインが必要で、
B
ボタンを押すには1枚のコインが必要なのです。
少し試してみると、それぞれの機械のボタンは、ボタンが押されるたびに、爪を特定の量だけ右方向（X軸に沿って）に、また特定の量だけ前方方向（Y軸に沿って）に動かすように設定されていることが分かります。
それぞれの機械には1つの賞品が入っており、賞品を獲得するには、爪を
X
軸と
Y
軸の両方で賞品の正確な上に配置する必要があります。
あなたは考えます。できるだけ多くの景品を獲得するために必要な最小限のトークン数はいくらか？あなたは、すべてのマシンのボタン動作と景品の場所のリストを作成します（これがパズルの入力となります）。例えば、
ボタンA：X+94、Y+34
ボタンB：X+22、Y+67
賞品：X=8400、Y=5400

ボタンA：X+26、Y+66
ボタンB：X+67、Y+21
賞品：X=12748、Y=12176

ボタンA：X+17、Y+86
ボタンB：X+84、Y+37
賞品：X=7870, Y=6450

ボタンA：X+69, Y+23
ボタンB：X+27, Y+71
賞品：X=18641, Y=10279
このリストは、4つの異なるクレーンゲーム機のボタン構成と賞品位置を示しています。
ここでは、リストの最初のクレーンゲーム機だけを考えてみましょう。
この機械のAボタンを押すと、爪はX軸に沿って94単位、Y軸に沿って34単位移動します。B
ボタンを押すと、爪は
X
軸に沿って
22
単位、
Y
軸に沿って
67
単位移動します。
景品は X=8400、Y=5400の位置にあります。つまり、この機械で景品を正確に獲得するには、爪の初期位置から、X軸に沿って正確に8400単位、Y軸に沿って正確に5400単位移動させる必要があります。
景品を獲得する最も安価な方法は、Aボタンを80回、Bボタンを40回押すことです。
これにより、X軸に沿って爪が並びます（80×94 + 40×22 = 8400）また、Y軸に沿って並びます（80×34 + 40×67 = 5400）。これを実行すると、Aボタンを押すのに80×3トークン、Bボタンを押すのに40×1トークンかかり、合計で280トークンになります。
2番目と4番目のクレーンゲーム機では、AとBのプレスの組み合わせで賞品を獲得できるものはありません。
3番目のクレーンゲーム機では、最も安く賞品を獲得する方法は、Aボタンを38回、Bボタンを86回押すことです。これを行うと、合計で200枚のトークンが必要になります。
つまり、獲得できる賞品の数は最大でも2つであり、2つの賞品を獲得するために最低限必要なトークン数は480です。
各ボタンを100回以内に押さなければ賞品を獲得できないと推定します。他にどのような遊び方があるでしょうか？
できるだけ多くの賞品を獲得する方法を考えましょう。すべての賞品を獲得するために必要な最低限のトークン数はいくつでしょうか？

--- パート2 ---
一等賞を獲得しようとすると、爪の位置が予想していた場所とはまったく違う場所にあることがわかります。 寸法の単位変換エラーにより、すべての賞の位置は、実際には10000000000000X軸とY軸の両方で
10000000000000をXとYの位置に追加します。この変更を加えると、上記の例は次のようになります。
ボタンA：X+94、Y+34
ボタンB：X+22、Y+67
賞品：X=10000000008400、Y=10000000005400

ボタンA：X+26、Y+66
ボタンB：X+67、Y+21
賞：X=10000000012748、Y=10000000012176

ボタンA：X+17、Y+86
ボタンB：X+84、Y+37
賞品：X=10000000007870、Y=10000000006450

ボタンA：X+69、Y+23
ボタンB：X+27、Y+71
賞品：X=10000000018641、Y=10000000010279
現在、2番目と4番目のクレーンゲーム機でしか賞品を獲得できません。残念ながら、そうするためには100回以上ボタンを押す必要があります。
修正された賞品の座標を使用して、できるだけ多くの賞品を獲得する方法を考えてみましょう。すべての賞品を獲得するために必要な最小のトークン数はいくつでしょうか？