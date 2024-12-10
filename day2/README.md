--- Day 2: Red-Nosed Reports ---
Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?

------------------------------
--- 2日目：赤鼻のレポート ---
幸いにも、歴史家たちが最初に調査したい場所は、チーフ・ヒストリアンのオフィスからそれほど離れていない。
赤鼻のトナカイ核融合/核分裂プラントにはチーフ・ヒストリアンの姿は見当たらないようだが、エンジニアたちはあなたを見かけるなり駆け寄ってきた。どうやら彼らは、ルドルフがたった1個の電子から分子合成によって救われた時のことを今でも話題にしているようです。
そして、あなたがすでにここにいるので、赤鼻の原子炉から得られたいくつかの異常なデータの分析を手伝ってくれると非常に助かると付け加えました。あなたは、歴史家たちがあなたを待っているかどうかを確認しようと振り返りましたが、彼らはすでにグループに分かれており、現在、施設の隅々を捜索しているようです。あなたは、異常なデータの分析を手伝うことを申し出ました。
異常データ（あなたのパズル入力）は、多くの報告書から構成されており、1行に1つの報告書が記載されています。各報告書は、スペースで区切られたレベルと呼ばれる数値のリストです。例えば、以下のようになります。
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
この例のデータには、それぞれ5つのレベルを含む6つのレポートが含まれています。
エンジニアは、どのレポートが安全であるかを把握しようとしています。Red-Nosedの原子炉安全システムは、徐々に増加しているか、徐々に減少しているレベルのみを許容できます。そのため、レポートが安全であると見なされるには、以下の2つの条件が両方とも満たされている必要があります。
レベルがすべて増加しているか、すべて減少している。
隣接する任意の2つのレベルは、少なくとも1つ、最大3つ異なる。
上記の例では、以下のルールを確認することで、レポートが安全または安全でないかを判断できます。
7 6 4 2 1：レベルがすべて1または2ずつ減少しているため、安全。
1 2 7 8 9：2 7 が5増加しているため、安全でない。
9 7 6 2 1：6 2 が4減少しているため、安全でない。
1 3 2 4 5: 1 3 は増加しているが、3 2 は減少しているため、安全ではない。
8 6 4 4 1: 4 4 は増加でも減少でもないため、安全ではない。
1 3 6 7 9: レベルはすべて1、2、または3で増加しているため、安全である。
したがって、この例では2つのレポートが安全である。
エンジニアからの異常なデータを分析します。何件のレポートが安全でしょうか？
