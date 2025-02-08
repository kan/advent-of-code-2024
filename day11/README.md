--- Day 11: Plutonian Pebbles ---
The ancient civilization on Pluto was known for its ability to manipulate spacetime, and while The Historians explore their infinite corridors, you've noticed a strange set of physics-defying stones.

At first glance, they seem like normal stones: they're arranged in a perfectly straight line, and each stone has a number engraved on it.

The strange part is that every time you blink, the stones change.

Sometimes, the number engraved on a stone changes. Other times, a stone might split in two, causing all the other stones to shift over a bit to make room in their perfectly straight line.

As you observe them for a while, you find that the stones have a consistent behavior. Every time you blink, the stones each simultaneously change according to the first applicable rule in this list:

If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
No matter how the stones change, their order is preserved, and they stay on their perfectly straight line.

How will the stones evolve if you keep blinking at them? You take a note of the number engraved on each stone in the line (your puzzle input).

If you have an arrangement of five stones engraved with the numbers 0 1 10 99 999 and you blink once, the stones transform as follows:

The first stone, 0, becomes a stone marked 1.
The second stone, 1, is multiplied by 2024 to become 2024.
The third stone, 10, is split into a stone marked 1 followed by a stone marked 0.
The fourth stone, 99, is split into two stones marked 9.
The fifth stone, 999, is replaced by a stone marked 2021976.
So, after blinking once, your five stones would become an arrangement of seven stones engraved with the numbers 1 2024 1 0 9 9 2021976.

Here is a longer example:

Initial arrangement:
125 17

After 1 blink:
253000 1 7

After 2 blinks:
253 0 2024 14168

After 3 blinks:
512072 1 20 24 28676032

After 4 blinks:
512 72 2024 2 0 2 4 2867 6032

After 5 blinks:
1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32

After 6 blinks:
2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
In this example, after blinking six times, you would have 22 stones. After blinking 25 times, you would have 55312 stones!

Consider the arrangement of stones in front of you. How many stones will you have after blinking 25 times?

--- Part Two ---
The Historians sure are taking a long time. To be fair, the infinite corridors are very large.

How many stones would you have after blinking a total of 75 times?

-----------------------------------------------------
--- 11日目：冥王星の小石 ---
冥王星の古代文明は時空を操る能力で知られており、歴史学者たちが無限に続く回廊を探索している間、あなたは物理法則を無視した奇妙な石の集合に気づきました。
一見したところ、それらは普通の石のように見えます。それらは完璧に一直線に並べられており、それぞれの石には数字が刻まれています。
不思議なのは、まばたきをするたびに石が変化することです。
時には、石に刻まれた数字が変わります。またある時には、石が2つに割れて、他のすべての石が少しずつ移動して、完璧な直線上にスペースができることもあります。
しばらく観察していると、石には一定の行動パターンがあることが分かります。まばたきをするたびに、石はそれぞれ同時に、以下のリストの最初の適用可能なルールに従って変化します。
石に数字の0が刻まれている場合、数字の1が刻まれた石と入れ替わる。
石に偶数の桁を持つ数字が刻まれている場合、2つの石と入れ替わる。左半分の数字は新しい左側の石に刻まれ、右半分の数字は新しい右側の石に刻まれます。（新しい数字は余分な先頭のゼロを維持しません。1000は10と0の石になります。）
他の規則が適用されない場合、石は新しい石に置き換えられます。古い石の数字を2024倍した数字が新しい石に刻まれます。
石がどのように変化しても、その順番は維持され、石は完全に一直線に並んだままです。
目を閉じたり開いたりしている間に、石はどのように変化していくのでしょうか？ あなたは、並んだ石に刻まれた数字をメモします（これがパズルの入力情報です）。
5つの石に0、1、10、99、999と刻まれており、目を一度閉じると、石は次のように変化します。
最初の石、0は、1と書かれた石に変化します。
2番目の石、1は2024倍されて2024になります。
3番目の石、10は、1と書かれた石と0と書かれた石に分割されます。
4番目の石、
99
は、
9
と書かれた石2つに分割されます。
5番目の石、999は、2021976と刻印された石に置き換えられます。
つまり、1回点滅した後、5つの石は、1 2024 1 0 9 9 2021976と刻印された7つの石の配列になります。
以下に、より長い例を示します。
初期の配列：
125 17

1回点滅後：
253000 1 7

2回点滅後：
253 0 2024 14168

3回点滅後：
512072 1 20 24 28676032

4回点滅後：
512 72 2024 2 0 2 4 2867 6032

点滅5回後：
1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32

点滅6回後：
2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
この例では、6回点滅すると、22個の石を手にしていることになります。25回点滅すると、55312個の石を手にしていることになります！
目の前の石の配置を考えてみましょう。25回点滅すると、何個の石を手にしていることになるでしょうか？

--- パート2 ---
歴史家たちは確かに長い時間をかけている。公平に見て、無限に続く廊下はとても広い。
合計75回まばたきをした後、あなたはいくつ石を数えることができるだろうか？