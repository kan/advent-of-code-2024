--- Day 9: Disk Fragmenter ---
Another push of the button leaves you in the familiar hallways of some friendly amphipods! Good thing you each somehow got your own personal mini submarine. The Historians jet away in search of the Chief, mostly by driving directly into walls.

While The Historians quickly figure out how to pilot these things, you notice an amphipod in the corner struggling with his computer. He's trying to make more contiguous free space by compacting all of the files, but his program isn't working; you offer to help.

He shows you the disk map (your puzzle input) he's already generated. For example:

2333133121414131402
The disk map uses a dense format to represent the layout of files and free space on the disk. The digits alternate between indicating the length of a file and the length of free space.

So, a disk map like 12345 would represent a one-block file, two blocks of free space, a three-block file, four blocks of free space, and then a five-block file. A disk map like 90909 would represent three nine-block files in a row (with no free space between them).

Each file on disk also has an ID number based on the order of the files as they appear before they are rearranged, starting with ID 0. So, the disk map 12345 has three files: a one-block file with ID 0, a three-block file with ID 1, and a five-block file with ID 2. Using one character for each block where digits are the file ID and . is free space, the disk map 12345 represents these individual blocks:

0..111....22222
The first example above, 2333133121414131402, represents these individual blocks:

00...111...2...333.44.5555.6666.777.888899
The amphipod would like to move file blocks one at a time from the end of the disk to the leftmost free space block (until there are no gaps remaining between file blocks). For the disk map 12345, the process looks like this:

0..111....22222
02.111....2222.
022111....222..
0221112...22...
02211122..2....
022111222......
The first example requires a few more steps:

00...111...2...333.44.5555.6666.777.888899
009..111...2...333.44.5555.6666.777.88889.
0099.111...2...333.44.5555.6666.777.8888..
00998111...2...333.44.5555.6666.777.888...
009981118..2...333.44.5555.6666.777.88....
0099811188.2...333.44.5555.6666.777.8.....
009981118882...333.44.5555.6666.777.......
0099811188827..333.44.5555.6666.77........
00998111888277.333.44.5555.6666.7.........
009981118882777333.44.5555.6666...........
009981118882777333644.5555.666............
00998111888277733364465555.66.............
0099811188827773336446555566..............
The final step of this file-compacting process is to update the filesystem checksum. To calculate the checksum, add up the result of multiplying each of these blocks' position with the file ID number it contains. The leftmost block is in position 0. If a block contains free space, skip it instead.

Continuing the first example, the first few blocks' position multiplied by its file ID number are 0 * 0 = 0, 1 * 0 = 0, 2 * 9 = 18, 3 * 9 = 27, 4 * 8 = 32, and so on. In this example, the checksum is the sum of these, 1928.

Compact the amphipod's hard drive using the process he requested. What is the resulting filesystem checksum? (Be careful copy/pasting the input for this puzzle; it is a single, very long line.)

--- Part Two ---
Upon completion, two things immediately become clear. First, the disk definitely has a lot more contiguous free space, just like the amphipod hoped. Second, the computer is running much more slowly! Maybe introducing all of that file system fragmentation was a bad idea?

The eager amphipod already has a new plan: rather than move individual blocks, he'd like to try compacting the files on his disk by moving whole files instead.

This time, attempt to move whole files to the leftmost span of free space blocks that could fit the file. Attempt to move each file exactly once in order of decreasing file ID number starting with the file with the highest file ID number. If there is no span of free space to the left of a file that is large enough to fit the file, the file does not move.

The first example from above now proceeds differently:

00...111...2...333.44.5555.6666.777.888899
0099.111...2...333.44.5555.6666.777.8888..
0099.1117772...333.44.5555.6666.....8888..
0099.111777244.333....5555.6666.....8888..
00992111777.44.333....5555.6666.....8888..
The process of updating the filesystem checksum is the same; now, this example's checksum would be 2858.

Start over, now compacting the amphipod's hard drive using this new method instead. What is the resulting filesystem checksum?
------------------------

--- 9日目：ディスクの断片化 ---
もう一度ボタンを押すと、見慣れたアンフィポッドの廊下に戻ります。 各自が自分の小型潜水艦を手に入れられてよかった。 歴史家たちはチーフを探して、主に壁にぶつかりながら飛び回っています。
ヒストリアンズはすぐに操縦方法を理解するが、あなたは隅で自分のコンピュータと格闘しているアンフィポッドに気づく。彼はすべてのファイルを圧縮して連続した空き領域を増やそうとしているが、プログラムがうまく動いていない。あなたは手伝うことを申し出る。
彼はすでに作成したディスクマップ（パズルの入力）を見せる。例えば：
2333133121414131402
ディスクマップは、ディスク上のファイルと空き領域のレイアウトを表現するために、緻密なフォーマットを使用しています。数字は、ファイルの長さと空き領域の長さを交互に示しています。
したがって、12345のようなディスクマップは、1ブロックのファイル、2ブロックの空き領域、3ブロックのファイル、4ブロックの空き領域、そして5ブロックのファイルを表します。90909のようなディスクマップは、9ブロックのファイルが3つ連続している（間に空き領域なし）ことを表します。
ディスク上の各ファイルには、再編成される前のファイルの順序に基づくID番号が割り当てられ、ID番号は
0
から始まります。したがって、ディスクマップ
12345
は、ID番号
0
の1ブロックのファイル、ID番号
1
の3ブロックのファイル、ID番号
2
の5ブロックのファイルの 各ブロックに1文字を使用し、数字をファイルID、.を空き領域として、ディスクマップ12345はこれらの個々のブロックを表します。
0..111....22222
上記の最初の例、2333133121414131402は、これらの個々のブロックを表します。
00...111...2...333.44.5555.6666.777.888899
このアンフィソドンは、ファイルブロックを1つずつディスクの末尾から左端の空きスペースブロックに移動させたいと考えています（ファイルブロック間に隙間がなくなるまで）。ディスクマップ12345の場合、処理は次のようになります。
0..111....22222
02.111....2222.
022111....222..
0221112...22...
02211122..2....
022111222......
最初の例では、さらにいくつかの手順が必要です。
00...111...2...333.44.5555.6666.777.888899
009..111...2...333.44.5555.6666.777.88889.
0099.111...2...333.44.5555.6666.777.8888..
00998111...2...333.44.5555.6666.777.888...
009981118..2...333.44.5555.6666.777.88....
0099811188.2...333.44.5555.6666.777.8.....
009981118882...333.44.5555.6666.777.......
0099811188827..333.44.5555.6666.77........
00998111888277.333.44.5555.6666.7.........
009981118882777333.44.5555.6666...........
009981118882777333644.5555.666............
00998111888277733364465555.66.............
0099811188827773336446555566..............
このファイル圧縮プロセスの最終ステップは、ファイルシステムのチェックサムを更新することです。チェックサムを計算するには、各ブロックの位置に含まれるファイルID番号を乗算した結果を合計します。左端のブロックは位置0です。ブロックに空き領域がある場合は、代わりにスキップします。
最初の例を続けると、最初の数ブロックの位置にファイルID番号を掛けた値は、0 * 0 = 0、1 * 0 = 0、2 * 9 = 18、3 * 9 = 27、4 * 8 = 32、となります。この例では、チェックサムはこれらの合計、1928です。
彼が依頼したプロセスを使用して、端脚類のハードドライブを圧縮します。結果として得られるファイルシステムのチェックサムは何ですか？（このパズルの入力のコピー＆ペーストには注意してください。1つの非常に長い行です。）

--- パート2 ---
完了すると、2つのことがすぐに明らかになります。まず、ディスクには確かに、両生類が期待したとおりに、連続した空き領域が大幅に増えています。次に、コンピュータの動作が大幅に遅くなっています！ ファイルシステムの断片化をすべて導入することは、悪い考えだったのでしょうか？
熱心な両生類はすでに新しい計画を立てています。個々のブロックを移動するのではなく、ファイル全体を移動してディスク上のファイルを圧縮してみることにします。
今回は、ファイルに適合する空きスペースブロックの左端のブロックに、ファイル全体を移動してみます。各ファイルは、ファイルID番号が小さい順に、最も高いファイルID番号を持つファイルから順に、それぞれ1回だけ移動してみます。ファイルの左側に、そのファイルに適合するのに十分な空きスペースブロックがない場合は、ファイルは移動しません。
上記の最初の例は、次のように異なる順序で処理されます。
00...111...2...333.44.5555.6666.777.888899
0099.111...2...333.44.5555.6666.777.8888..
0099.1117772...333.44.5555.6666.....8888..
0099.111777244.333....5555.6666.....8888..
00992111777.44.333....5555.6666.....8888..
ファイルシステムのチェックサムを更新するプロセスは同じです。この例のチェックサムは2858になります。
最初からやり直して、今度は代わりにこの新しい方法を使用してamphipodのハードドライブを圧縮します。結果として得られるファイルシステムのチェックサムは何ですか？