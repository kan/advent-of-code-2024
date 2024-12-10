--- Part Two ---
Your analysis only confirmed what everyone feared: the two lists of location IDs are indeed very different.

Or are they?

The Historians can't agree on which group made the mistakes or how to read most of the Chief's handwriting, but in the commotion you notice an interesting detail: a lot of location IDs appear in both lists! Maybe the other numbers aren't location IDs at all but rather misinterpreted handwriting.

This time, you'll need to figure out exactly how often each number from the left list appears in the right list. Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.

Here are the same example lists again:

3   4
4   3
2   5
1   3
3   9
3   3
For these example lists, here is the process of finding the similarity score:

The first number in the left list is 3. It appears in the right list three times, so the similarity score increases by 3 * 3 = 9.
The second number in the left list is 4. It appears in the right list once, so the similarity score increases by 4 * 1 = 4.
The third number in the left list is 2. It does not appear in the right list, so the similarity score does not increase (2 * 0 = 0).
The fourth number, 1, also does not appear in the right list.
The fifth number, 3, appears in the right list three times; the similarity score increases by 9.
The last number, 3, appears in the right list three times; the similarity score again increases by 9.
So, for these example lists, the similarity score at the end of this process is 31 (9 + 4 + 0 + 0 + 9 + 9).

Once again consider your left and right lists. What is their similarity score?

---------------------------------------------
--- パート2 ---
あなたの分析は、誰もが恐れていたことを裏付けるものでした。2つの位置IDリストは、確かに非常に異なっています。
それとも、そうではないのでしょうか？
歴史家たちは、どちらのグループがミスを犯したのか、あるいは首長の筆跡のほとんどをどのように解読するのかについて意見が一致していませんが、この騒動の中で、興味深い詳細に気づきました。多くの位置IDが両方のリストに表示されているのです！おそらく、他の数字は位置IDではなく、誤解釈された筆跡である可能性があります。
今回は、左のリストの各数字が右のリストにどの程度頻繁に現れるかを正確に把握する必要があります。左のリストの各数字に右のリストにその数字が現れる回数を掛けたものを合計し、類似スコアの合計を算出します。
以下に、同じ例のリストを再度示します。
3 4
4 3
2 5
1 3
3 9
3 3
これらの例のリストにおける類似スコアの算出方法は以下の通りです。
左のリストの最初の数値は3です。この数値は右のリストに3回出現しているため、類似スコアは3 * 3 = 9増加します。左のリストの2番目の数値は4です。
この数値は右のリストに1回出現しているため、類似スコアは
4 * 1 = 
4増加します。
左のリストの3番目の数値は2です。
左のリストの3番目の数字は2です。右のリストには現れないため、類似スコアは増加しません（2 * 0 = 0）。
4番目の数字1も右のリストには現れません。
5番目の数字3は右のリストに3回現れるため、類似スコアは9増加します。
最後の数字、3は、正しいリストに3回表示されており、類似スコアが9増加します。
したがって、この例のリストでは、このプロセスの終了時の類似スコアは31（9 + 4 + 0 + 0 + 9 + 9）となります。もう一度、左右のリストを考えてみてください。類似スコアは何点ですか？