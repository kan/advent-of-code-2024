--- Day 12: Garden Groups ---
Why not search for the Chief Historian near the gardener and his massive farm? There's plenty of food, so The Historians grab something to eat while they search.

You're about to settle near a complex arrangement of garden plots when some Elves ask if you can lend a hand. They'd like to set up fences around each region of garden plots, but they can't figure out how much fence they need to order or how much it will cost. They hand you a map (your puzzle input) of the garden plots.

Each garden plot grows only a single type of plant and is indicated by a single letter on your map. When multiple garden plots are growing the same type of plant and are touching (horizontally or vertically), they form a region. For example:

AAAA
BBCD
BBCC
EEEC
This 4x4 arrangement includes garden plots growing five different types of plants (labeled A, B, C, D, and E), each grouped into their own region.

In order to accurately calculate the cost of the fence around a single region, you need to know that region's area and perimeter.

The area of a region is simply the number of garden plots the region contains. The above map's type A, B, and C plants are each in a region of area 4. The type E plants are in a region of area 3; the type D plants are in a region of area 1.

Each garden plot is a square and so has four sides. The perimeter of a region is the number of sides of garden plots in the region that do not touch another garden plot in the same region. The type A and C plants are each in a region with perimeter 10. The type B and E plants are each in a region with perimeter 8. The lone D plot forms its own region with perimeter 4.

Visually indicating the sides of plots in each region that contribute to the perimeter using - and |, the above map's regions' perimeters are measured as follows:

+-+-+-+-+
|A A A A|
+-+-+-+-+     +-+
              |D|
+-+-+   +-+   +-+
|B B|   |C|
+   +   + +-+
|B B|   |C C|
+-+-+   +-+ +
          |C|
+-+-+-+   +-+
|E E E|
+-+-+-+
Plants of the same type can appear in multiple separate regions, and regions can even appear within other regions. For example:

OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
The above map contains five regions, one containing all of the O garden plots, and the other four each containing a single X plot.

The four X regions each have area 1 and perimeter 4. The region containing 21 type O plants is more complicated; in addition to its outer edge contributing a perimeter of 20, its boundary with each X region contributes an additional 4 to its perimeter, for a total perimeter of 36.

Due to "modern" business practices, the price of fence required for a region is found by multiplying that region's area by its perimeter. The total price of fencing all regions on a map is found by adding together the price of fence for every region on the map.

In the first example, region A has price 4 * 10 = 40, region B has price 4 * 8 = 32, region C has price 4 * 10 = 40, region D has price 1 * 4 = 4, and region E has price 3 * 8 = 24. So, the total price for the first example is 140.

In the second example, the region with all of the O plants has price 21 * 36 = 756, and each of the four smaller X regions has price 1 * 4 = 4, for a total price of 772 (756 + 4 + 4 + 4 + 4).

Here's a larger example:

RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
It contains:

A region of R plants with price 12 * 18 = 216.
A region of I plants with price 4 * 8 = 32.
A region of C plants with price 14 * 28 = 392.
A region of F plants with price 10 * 18 = 180.
A region of V plants with price 13 * 20 = 260.
A region of J plants with price 11 * 20 = 220.
A region of C plants with price 1 * 4 = 4.
A region of E plants with price 13 * 18 = 234.
A region of I plants with price 14 * 22 = 308.
A region of M plants with price 5 * 12 = 60.
A region of S plants with price 3 * 8 = 24.
So, it has a total price of 1930.

What is the total price of fencing all regions on your map?

--- Part Two ---
Fortunately, the Elves are trying to order so much fence that they qualify for a bulk discount!

Under the bulk discount, instead of using the perimeter to calculate the price, you need to use the number of sides each region has. Each straight section of fence counts as a side, regardless of how long it is.

Consider this example again:

AAAA
BBCD
BBCC
EEEC
The region containing type A plants has 4 sides, as does each of the regions containing plants of type B, D, and E. However, the more complex region containing the plants of type C has 8 sides!

Using the new method of calculating the per-region price by multiplying the region's area by its number of sides, regions A through E have prices 16, 16, 32, 4, and 12, respectively, for a total price of 80.

The second example above (full of type X and O plants) would have a total price of 436.

Here's a map that includes an E-shaped region full of type E plants:

EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
The E-shaped region has an area of 17 and 12 sides for a price of 204. Including the two regions full of type X plants, this map has a total price of 236.

This map has a total price of 368:

AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
It includes two regions full of type B plants (each with 4 sides) and a single region full of type A plants (with 4 sides on the outside and 8 more sides on the inside, a total of 12 sides). Be especially careful when counting the fence around regions like the one full of type A plants; in particular, each section of fence has an in-side and an out-side, so the fence does not connect across the middle of the region (where the two B regions touch diagonally). (The Elves would have used the Möbius Fencing Company instead, but their contract terms were too one-sided.)

The larger example from before now has the following updated prices:

A region of R plants with price 12 * 10 = 120.
A region of I plants with price 4 * 4 = 16.
A region of C plants with price 14 * 22 = 308.
A region of F plants with price 10 * 12 = 120.
A region of V plants with price 13 * 10 = 130.
A region of J plants with price 11 * 12 = 132.
A region of C plants with price 1 * 4 = 4.
A region of E plants with price 13 * 8 = 104.
A region of I plants with price 14 * 16 = 224.
A region of M plants with price 5 * 6 = 30.
A region of S plants with price 3 * 6 = 18.
Adding these together produces its new total price of 1206.

What is the new total price of fencing all regions on your map?

---------------------------------------------

--- 12日目：ガーデングループ ---
庭師と彼の広大な農場の近くでチーフ・ヒストリアンを探してみてはいかがでしょうか？ 食べ物はたくさんあるので、ヒストリアンたちは食べ物を手に入れながら探索します。
複雑に区画された菜園の近くに落ち着こうとしたとき、エルフの何人かが手伝ってくれないかと頼んできます。彼らは菜園の各区画の周囲にフェンスを設置したいのですが、必要なフェンスの長さや費用がわからないのです。彼らはあなたに菜園の地図（パズルの入力情報）を渡します。
各菜園区画では単一の種類の植物のみが栽培されており、地図上では1文字で示されています。複数の菜園区画が同じ種類の植物を栽培しており、それらの区画が水平または垂直に接している場合、それらは地域を形成します。例えば：
AAAA
BBCD
BBCC
EEEC
この4x4の配置には、5種類の異なる植物を栽培している菜園区画（A、B、C、D、E）が含まれており、それぞれが独自の地域にグループ化されています。
1つの区域の周囲に設置するフェンスの費用を正確に算出するには、その区域の面積と周囲を知る必要があります。
区域の面積は、その区域に含まれるガーデン区画の数に単純に等しくなります。上記の地図の
A
、
B
、
C
の植物はそれぞれ面積が
4
の区域に植えられています。
E
の植物は面積が
3
の区域に植えられており、
D
の植物は面積が
1
の区域に植えられています。
各庭は正方形であり、4つの辺がある。領域の周囲とは、同じ領域内の他の庭と接していない庭の辺の数を指す。
A型
および
C型
植物は、それぞれ周囲が
10
の領域にある。
B型
および
E型
植物は、それぞれ周囲が
8
の領域にある。唯一の
D型
植物は、周囲が
4
の独自の領域を形成している。
各領域内のプロットの辺のうち、周囲に寄与する辺を視覚的に示すために、-および|を使用すると、上記の地図の領域の周囲は以下のように測定されます。
+-+-+-+-+
|A A A A|
+-+-+-+-+ +-+
              |D|
+-+-+ +-+ +-+
|B B|   |C|
+   +   + +-+
|B B|  |C C|
+-+-+  +-+ +
          |C|
+-+-+-+  +-+
|E E E|
+-+-+-+
同じタイプの植物は複数の別々の領域に存在することができ、領域は他の領域内に存在することさえあります。例えば：
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
上記の地図には5つの地域があり、そのうちの1つにはOの庭がすべて含まれ、残りの4つにはそれぞれXの庭が1つずつ含まれています。
4つの
X
の地域はそれぞれ面積が
1
、周囲が
4
です。
O
の植物が
21
種類ある地域はさらに複雑で、外周が
20
であることに加え、
X
の地域との境界線がさらに周囲に
4
加わるため、周囲は合計で
36
となります。
「近代的な」商慣習により、ある地域に必要なフェンスの価格は、その地域の面積と周囲の長さを掛け合わせることで求められます。 地図上のすべての地域にフェンスを設置する場合の合計価格は、地図上のすべての地域のフェンスの価格を合計することで求められます。
最初の例では、地域Aの価格は4×10=40、地域Bの価格は4×8=32、地域Cの価格は4×10=40、地域Dの価格は1×4=4、地域Eの価格は3×8=24です。したがって、最初の例の合計価格は140となります。
2番目の例では、Oの植物がすべてある地域は価格が21 * 36 = 756、4つのより小さいXの地域はそれぞれ価格が1 * 4 = 4で、合計価格は772（756 + 4 + 4 + 4 + 4）となります。より大きな例を以下に示します。
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
内容：
Rの植物の地域で、価格は12 * 18 = 216。Iの植物の地域で、価格は4 * 8 = 32。
Cの植物の地域で、価格は14 * 28 = 392。
Fの植物の地域で、価格は10 * 18 = 180。
Vの植物の地域で、価格は13 * 20 = 260。
Jの植物の地域で、価格は11 * 20 = 220。
C
の植物の地域で、価格は
1 * 4 = 4
。
Eの植物の地域で、価格は13 * 18 = 234。
I
の植物の地域で、価格は
14 * 22 = 308
。
M本の植物の地域で、価格は5 * 12 = 60。
S本の植物の地域で価格は3 * 8 = 24。
したがって、合計価格は1930です。
あなたの地図上のすべての地域を囲う場合の合計価格はいくらですか？

--- パート2 ---
幸いにも、エルフたちはフェンスを大量に注文しようとしているため、大量購入割引の対象となります！
大量購入割引では、価格を計算する際に周囲の長さではなく、各区域の辺の数を使用する必要があります。フェンスの直線部分の長さは関係なく、1辺として数えます。
もう一度、この例を考えてみましょう。
AAAA
BBCD
BBCC
EEEC
タイプAの植物を含む地域は、4つの辺があります。タイプB、D、Eの植物を含む地域もそれぞれ同じです。しかし、タイプCの植物を含むより複雑な地域は、8つの辺があります！
新しい方法では、地域の面積に辺の数を掛けることで地域ごとの価格を計算します。地域AからEまでの価格はそれぞれ16、16、32、4、12となり、合計価格は80となります。
上記の2番目の例（タイプXとOの植物でいっぱい）の合計価格は436となります。
以下は、タイプEの植物でいっぱいのE字型の区域を含むマップです。
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
E字型の区域は、面積が17、12面で、価格は204です。タイプXの植物でいっぱいの2つの区域を含めると、このマップの合計価格は236です。このマップの合計価格は368
です。
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
このマップには、タイプ
B
の植物でいっぱいの2つの地域（それぞれに
4つ
の面）と、タイプ
A
の植物でいっぱいの1つの地域（外側に
4つ
の面と内側に
8つ
の面があり、合計で
12つ
の面）が含まれています。特に、Aタイプの植物でいっぱいの領域のような、領域を囲むフェンスを数える際には注意が必要です。特に、フェンスの各セクションには内側と外側があるため、フェンスは領域の中央（2つのBタイプの植物の領域が斜めに接する場所）でつながっていません。（エルフたちは代わりにメビウス・フェンシング社を利用するでしょうが、契約条件が一方的に過ぎました。）先ほどの大きな例では、以下の価格が更新されました。
Rの植物の地域で、価格は12 * 10 = 120。
Iの植物の地域で、価格は4 * 4 = 16。Cの植物の地域で、価格は14 * 22 = 308。
Fの植物の地域で、価格は10 * 12 = 120。
Vの植物の地域で、価格は13 * 10 = 130。
V株の価格が13 * 10 = 130の地域
J株の価格が11 * 12 = 132の地域
C株の価格が1 * 4 = 4の地域
E株の価格が13 * 8 = 104の地域
I株の価格が14 * 16 = 224の地域
M本の植物の地域で、価格は5 * 6 = 30。
S本の植物の地域で価格は3 * 6 = 18。
これらを合計すると、新しい合計価格は1206となります。
あなたの地図上のすべての地域にフェンスを設置した場合の新しい合計価格はいくらになるでしょうか？