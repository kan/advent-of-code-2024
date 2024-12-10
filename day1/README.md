--- Day 1: Historian Hysteria ---
The Chief Historian is always present for the big Christmas sleigh launch, but nobody has seen him in months! Last anyone heard, he was visiting locations that are historically significant to the North Pole; a group of Senior Historians has asked you to accompany them as they check the places they think he was most likely to visit.

As each location is checked, they will mark it on their list with a star. They figure the Chief Historian must be in one of the first fifty places they'll look, so in order to save Christmas, you need to help them get fifty stars on their list before Santa takes off on December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You haven't even left yet and the group of Elvish Senior Historians has already hit a problem: their list of locations to check is currently empty. Eventually, someone decides that the best place to check first would be the Chief Historian's office.

Upon pouring into the office, everyone confirms that the Chief Historian is indeed nowhere to be found. Instead, the Elves discover an assortment of notes and lists of historically significant locations! This seems to be the planning the Chief Historian was doing before he left. Perhaps these notes can be used to determine which locations to search?

Throughout the Chief's office, the historically significant locations are listed not by name but by a unique number called the location ID. To make sure they don't miss anything, The Historians split into two groups, each searching the office and trying to create their own complete list of location IDs.

There's just one problem: by holding the two lists up side by side (your puzzle input), it quickly becomes clear that the lists aren't very similar. Maybe you can help The Historians reconcile their lists?

For example:

3   4
4   3
2   5
1   3
3   9
3   3
Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. Pair up the smallest number in the left list with the smallest number in the right list, then the second-smallest left number with the second-smallest right number, and so on.

Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances. For example, if you pair up a 3 from the left list with a 7 from the right list, the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.

In the example list above, the pairs and distances would be as follows:

The smallest number in the left list is 1, and the smallest number in the right list is 3. The distance between them is 2.
The second-smallest number in the left list is 2, and the second-smallest number in the right list is another 3. The distance between them is 1.
The third-smallest number in both lists is 3, so the distance between them is 0.
The next numbers to pair up are 3 and 4, a distance of 1.
The fifth-smallest numbers in each list are 3 and 5, a distance of 2.
Finally, the largest number in the left list is 4, while the largest number in the right list is 9; these are a distance 5 apart.
To find the total distance between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is 2 + 1 + 0 + 1 + 2 + 5, a total distance of 11!

Your actual left and right lists contain many location IDs. What is the total distance between your lists?

--------------------------------------------------

--- 1日目：歴史家ヒステリー ---
チーフ・ヒストリアンは、クリスマス・ソリの打ち上げには必ず出席しているが、ここ数か月、彼の姿を見た者はいない！ 最近聞いたところによると、彼は北極にとって歴史的に重要な場所を訪問しているらしい。シニア・ヒストリアンのグループが、彼が最も訪れた可能性が高いと考えられる場所を調査するので、同行してほしいと依頼してきた。
各場所が確認されるごとに、彼らはリストに星印を付けます。彼らは、主任歴史学者は最初の50か所のいずれかにいるに違いないと考えています。そのため、クリスマスを救うには、サンタが12月25日に飛び立つ前に、彼らのリストに50個の星を付けてもらう必要があります。
パズルを解いて星を集めましょう。アドベントカレンダーでは、毎日2つのパズルが利用可能になります。最初のパズルを解くと、2つ目のパズルが利用可能になります。各パズルを解くと1つ星が獲得できます。頑張ってください！
まだ出発もしていないのに、エルフのシニア・ヒストリアンのグループはすでに問題にぶつかっています。確認すべき場所のリストが現在空っぽだからです。最終的に、最初に確認すべき場所として、チーフ・ヒストリアンのオフィスが最適であるという結論に達しました。
オフィスに駆け込むと、全員がチーフ・ヒストリアンの姿が確かにどこにもないことを確認しました。代わりにエルフたちは、歴史的に重要な場所のメモやリストの数々を発見したのです！これはチーフ・ヒストリアンが旅立つ前に準備していたもののようです。これらのメモを使って、どの場所を捜索すべきかを決定できるかもしれません。
チーフのオフィス全体にわたって、歴史的に重要な場所は名称ではなく、場所IDと呼ばれる固有の番号でリスト化されていました。見落としがないように、歴史家たちは2つのグループに分かれ、それぞれがオフィス内を探索し、自分たちの場所IDの完全なリストを作成しようとしました。
ただ1つ問題があります。2つのリストを並べて比較すると（あなたのパズル入力）、リストがそれほど似ていないことがすぐに明らかになります。The Historiansがリストを一致させるのを手伝ってあげられますか？
例えば：
3 4
4 3
2 5
1 3
3 9
3 3
リストの差はわずかかもしれません！ それを確かめるために、数字をペアにして、その差を測ってみましょう。 左のリストの最小の数字と右のリストの最小の数字をペアにし、次に左のリストで2番目に小さい数字と右のリストで2番目に小さい数字をペアにし、以下同様にしてください。
各ペアにおいて、2つの数字がどの程度離れているかを把握します。そのためには、それらの距離をすべて合計する必要があります。例えば、左のリストの
3
と右のリストの
7
をペアにした場合、その距離は
4
です。また、
9
と
3
をペアにした場合、その距離は
6
です。
上記の例では、ペアと距離は以下のようになります。
左のリストの最小の数は1、右のリストの最小の数は3です。 それらの間の距離は2です。
左のリストで2番目に小さい数は2、右のリストで2番目に小さい数は同じく3です。 それらの間の距離は1です。
両方のリストで3番目に小さい数は3なので、その間の距離は0です。
次にペアになるのは3と4で、その距離は1です。
各リストで5番目に小さい数は3と5で、その距離は2です。最後に、左のリストで最大の数は4で、
右のリストで最大の数は9です。これらは5離れています。
左のリストと右のリストの間の合計距離を求めるには、見つけたすべてのペア間の距離を合計します。上記の例では、2 + 1 + 0 + 1 + 2 + 5で合計距離は11です！
実際の左のリストと右のリストには、多くの位置IDが含まれています。あなたのリストの合計距離は？