--- Day 10: Hoof It ---
You all arrive at a Lava Production Facility on a floating island in the sky. As the others begin to search the massive industrial complex, you feel a small nose boop your leg and look down to discover a reindeer wearing a hard hat.

The reindeer is holding a book titled "Lava Island Hiking Guide". However, when you open the book, you discover that most of it seems to have been scorched by lava! As you're about to ask how you can help, the reindeer brings you a blank topographic map of the surrounding area (your puzzle input) and looks up at you excitedly.

Perhaps you can help fill in the missing hiking trails?

The topographic map indicates the height at each position using a scale from 0 (lowest) to 9 (highest). For example:

0123
1234
8765
9876
Based on un-scorched scraps of the book, you determine that a good hiking trail is as long as possible and has an even, gradual, uphill slope. For all practical purposes, this means that a hiking trail is any path that starts at height 0, ends at height 9, and always increases by a height of exactly 1 at each step. Hiking trails never include diagonal steps - only up, down, left, or right (from the perspective of the map).

You look up from the map and notice that the reindeer has helpfully begun to construct a small pile of pencils, markers, rulers, compasses, stickers, and other equipment you might need to update the map with hiking trails.

A trailhead is any position that starts one or more hiking trails - here, these positions will always have height 0. Assembling more fragments of pages, you establish that a trailhead's score is the number of 9-height positions reachable from that trailhead via a hiking trail. In the above example, the single trailhead in the top left corner has a score of 1 because it can reach a single 9 (the one in the bottom left).

This trailhead has a score of 2:

...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9
(The positions marked . are impassable tiles to simplify these examples; they do not appear on your actual topographic map.)

This trailhead has a score of 4 because every 9 is reachable via a hiking trail except the one immediately to the left of the trailhead:

..90..9
...1.98
...2..7
6543456
765.987
876....
987....
This topographic map contains two trailheads; the trailhead at the top has a score of 1, while the trailhead at the bottom has a score of 2:

10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01
Here's a larger example:

89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
This larger example has 9 trailheads. Considering the trailheads in reading order, they have scores of 5, 6, 5, 3, 1, 3, 5, 3, and 5. Adding these scores together, the sum of the scores of all trailheads is 36.

The reindeer gleefully carries over a protractor and adds it to the pile. What is the sum of the scores of all trailheads on your topographic map?

--- Part Two ---
The reindeer spends a few minutes reviewing your hiking trail map before realizing something, disappearing for a few minutes, and finally returning with yet another slightly-charred piece of paper.

The paper describes a second way to measure a trailhead called its rating. A trailhead's rating is the number of distinct hiking trails which begin at that trailhead. For example:

.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....
The above map has a single trailhead; its rating is 3 because there are exactly three distinct hiking trails which begin at that position:

.....0.   .....0.   .....0.
..4321.   .....1.   .....1.
..5....   .....2.   .....2.
..6....   ..6543.   .....3.
..7....   ..7....   .....4.
..8....   ..8....   ..8765.
..9....   ..9....   ..9....
Here is a map containing a single trailhead with rating 13:

..90..9
...1.98
...2..7
6543456
765.987
876....
987....
This map contains a single trailhead with rating 227 (because there are 121 distinct hiking trails that lead to the 9 on the right edge and 106 that lead to the 9 on the bottom edge):

012345
123456
234567
345678
4.6789
56789.
Here's the larger example from before:

89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
Considering its trailheads in reading order, they have ratings of 20, 24, 10, 4, 1, 4, 5, 8, and 5. The sum of all trailhead ratings in this larger example topographic map is 81.

You're not sure how, but the reindeer seems to have crafted some tiny flags out of toothpicks and bits of paper and is using them to mark trailheads on your topographic map. What is the sum of the ratings of all trailheads?

------------------------------------------------------
--- 10日目：徒歩で ---
皆さんは、空に浮かぶ浮島の溶岩生産施設に到着しました。他の皆が巨大な工業団地を探索し始めると、あなたは足に小さな鼻づきを感じ、下を見るとヘルメットをかぶったトナカイがいます。
トナカイは「溶岩島ハイキングガイド」というタイトルの本を持っています。しかし、その本を開いてみると、ほとんどの部分が溶岩で焦げた跡があるようです！どうしたら手助けができるか尋ねようとしたとき、トナカイは周辺の地形図（あなたのパズル入力）をあなたに渡し、興奮した様子であなたを見上げます。
おそらく、あなたはハイキングコースの欠けている部分を埋める手助けができるでしょう？
地形図には、各位置の高さが、0（最低）から9（最高）までの目盛りで示されています。例えば：
0123
1234
8765
9876
焦げていない本の切れ端を基に、あなたは良いハイキングコースはできるだけ長く、なだらかで緩やかな上り坂であると判断します。実質的には、これはハイキングコースとは、高さ0から始まり、高さ9で終わる、各ステップで常に正確に1の高さずつ増加するあらゆる道であることを意味します。ハイキングコースには斜めのステップは含まれません。上、下、左、右（地図の視点）のみです。
地図を見上げると、トナカイが親切にもハイキングコース付きの地図を更新するために必要な鉛筆、マーカー、定規、コンパス、ステッカー、その他の備品を小さな山にしてくれているのに気づきます。
トレイルヘッドとは、1つ以上のハイキングコースの出発点となる位置を指します。ここでは、これらの位置は常に高さ0となります。ページの断片をさらに集めると、トレイルヘッドのスコアは、そのトレイルヘッドからハイキングコースを経由して到達できる高さ9の位置の数であることがわかります。上記の例では、左上隅にある1つのトレイルヘッドのスコアは1です。なぜなら、1つの9（左下隅にあるもの）に到達できるからです。
この登山口のスコアは2です：...
0......
1......
2...
6543456
7....7
8....8
9....9
（.でマークされた位置は、例を単純化するために通行できないタイルです。実際の地形図には表示されません。）
この登山口のスコアは4です。なぜなら、9の数字はすべて、登山口のすぐ左にあるものを除いて、ハイキングコースで到達できるからです。...
90..9...
1.98...
2..7
6543456
765.987
876....
987....
この地形図には2つの登山口が記載されています。上部の登山口のスコアは1、下部の登山口のスコアは2です。
10..9..
2...8..
3...7..
4567654...
8..3...
9..2.....
01
より大きな例を以下に示します。
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
この大きな例では、9つの登山口があります。 登山口を昇順に考えると、それらのスコアは5、6、5、3、1、3、5、3、そして5です。 これらのスコアを合計すると、すべての登山口のスコアの合計は36です。
トナカイは嬉しそうに分度器を持ち上げて、それを山に加える。あなたの地形図上のすべての登山口のスコアの合計は？

--- パート2 ---
トナカイは、ハイキングコースの地図を数分間眺めてから、何かを悟り、数分間姿を消し、最後にまた別の少し焦げた紙を持って戻ってきました。
その紙には、トレイルヘッドを測定する2つ目の方法が記されています。トレイルヘッドの評価は、そのトレイルヘッドから始まる異なるハイキングコースの数です。例えば、...
0....
4321....
5..2.
..6543...
7..4...
8765...
9....
上記の地図にはトレイルヘッドが1つだけあり、その評価は3です。なぜなら、その位置から始まる明確なハイキングコースがちょうど3つあるからです。.....
0.  .....0.  .....0...
4321.  .....1.  .....1...
5....  .....2.  .....2.
..6....   ..6543.   .....3...
7....   ..7....   .....4...
8....   ..8....   ..8765...
9....   ..9....   ..9....
評価が13の登山口が1つある地図は次のとおりです：..
90..9...
1.98...
2..7
6543456
765.987
876....
987....
この地図には、評価 227 の登山口が1つ含まれています（121の異なるハイキングコースが、右端の9につながり、106のコースが下端の9につながっているため）：
012345
123456
234567
345678
4.6789
56789.
以下は、以前のより大きな例です。
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
トレイルヘッドを読み順に考慮すると、それらの評価は20、24、10、4、1、4、5、8、5です。この大きな例の地形図におけるすべてのトレイルヘッドの評価の合計は81です。
その方法はわからないが、トナカイが爪楊枝と紙の切れ端で小さな旗を作り、それを地形図上の登山口に印として立てているようだ。すべての登山口の評価の合計はいくつになるだろうか？