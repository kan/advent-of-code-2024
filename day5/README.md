--- Day 5: Print Queue ---
Satisfied with their search on Ceres, the squadron of scholars suggests subsequently scanning the stationery stacks of sub-basement 17.

The North Pole printing department is busier than ever this close to Christmas, and while The Historians continue their search of this historically significant facility, an Elf operating a very familiar printer beckons you over.

The Elf must recognize you, because they waste no time explaining that the new sleigh launch safety manual updates won't print correctly. Failure to update the safety manuals would be dire indeed, so you offer your services.

Safety protocols clearly indicate that new pages for the safety manuals must be printed in a very specific order. The notation X|Y means that if both page number X and page number Y are to be produced as part of an update, page number X must be printed at some point before page number Y.

The Elf has for you both the page ordering rules and the pages to produce in each update (your puzzle input), but can't figure out whether each update has the pages in the right order.

For example:

47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
The first section specifies the page ordering rules, one per line. The first rule, 47|53, means that if an update includes both page number 47 and page number 53, then page number 47 must be printed at some point before page number 53. (47 doesn't necessarily need to be immediately before 53; other pages are allowed to be between them.)

The second section specifies the page numbers of each update. Because most safety manuals are different, the pages needed in the updates are different too. The first update, 75,47,61,53,29, means that the update consists of page numbers 75, 47, 61, 53, and 29.

To get the printers going as soon as possible, start by identifying which updates are already in the right order.

In the above example, the first update (75,47,61,53,29) is in the right order:

75 is correctly first because there are rules that put each other page after it: 75|47, 75|61, 75|53, and 75|29.
47 is correctly second because 75 must be before it (75|47) and every other page must be after it according to 47|61, 47|53, and 47|29.
61 is correctly in the middle because 75 and 47 are before it (75|61 and 47|61) and 53 and 29 are after it (61|53 and 61|29).
53 is correctly fourth because it is before page number 29 (53|29).
29 is the only page left and so is correctly last.
Because the first update does not include some page numbers, the ordering rules involving those missing page numbers are ignored.

The second and third updates are also in the correct order according to the rules. Like the first update, they also do not include every page number, and so only some of the ordering rules apply - within each update, the ordering rules that involve missing page numbers are not used.

The fourth update, 75,97,47,61,53, is not in the correct order: it would print 75 before 97, which violates the rule 97|75.

The fifth update, 61,13,29, is also not in the correct order, since it breaks the rule 29|13.

The last update, 97,13,75,29,47, is not in the correct order due to breaking several rules.

For some reason, the Elves also need to know the middle page number of each update being printed. Because you are currently only printing the correctly-ordered updates, you will need to find the middle page number of each correctly-ordered update. In the above example, the correctly-ordered updates are:

75,47,61,53,29
97,61,53,29,13
75,29,13
These have middle page numbers of 61, 53, and 29 respectively. Adding these page numbers together gives 143.

Of course, you'll need to be careful: the actual list of page ordering rules is bigger and more complicated than the above example.

Determine which updates are already in the correct order. What do you get if you add up the middle page number from those correctly-ordered updates?

---------------------------------------------------------------------------------------
--- 5日目：印刷キュー ---
セレスでの検索に満足した学者の一団は、その後、地下17階の文房具の棚をスキャンすることを提案した。
クリスマスが近づき、北極の印刷部門はこれまで以上に忙しくなっていた。歴史学者たちがこの歴史的に重要な施設の調査を続ける中、見慣れたプリンターを操作するエルフが手招きする。
エルフたちは時間を無駄にしないので、新しいそりの打ち上げ安全マニュアルの更新が正しく印刷されないことを説明します。安全マニュアルの更新を怠ると大変なことになるので、あなたは自分のサービスを申し出ます。
安全手順書には、安全マニュアルの新しいページは非常に特定の順序で印刷しなければならないことが明確に示されています。
X|Y
という表記は、ページ番号
X
とページ番号
Y
の両方を更新の一部として作成する場合、ページ番号
X
はページ番号
Y
の前に印刷しなければならないことを意味します。
Elfは、ページの順序付けルールと各更新で作成するページ（パズルの入力）の両方を提供しますが、各更新でページが正しい順序になっているかどうかは判断できません。
例えば：
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
最初のセクションでは、ページの順序付けルールが1行につき1つずつ指定されています。最初のルールである47|53は、更新にページ番号47とページ番号53の両方が含まれている場合、ページ番号47はページ番号53の前に印刷されなければならないことを意味します。（47は必ずしも53の直前にある必要はなく、間に他のページがあってもかまいません。）
2番目のセクションでは、各更新のページ番号を指定します。安全マニュアルはほとんどが異なるため、更新に必要なページも異なります。最初の更新、75,47,61,53,29は、75、47、61、53、29のページ番号で構成される更新であることを意味します。
プリンターをできるだけ早く稼働させるには、どの更新がすでに正しい順序であるかを特定することから始めます。
上記の例では、最初の更新（75,47,61,53,29）が正しい順序です。
75は正しく最初のページである。なぜなら、75|47、75|61、75|53、および75|29という
ルールがあるためである。47は正しく2番目のページである。なぜなら、75は47（75|47）の前に置かれなければならず、また、47|61、47|53、および47|29に従って、他のすべてのページは47の後になければならないからである。
61 は正しく中央に位置しています。75と47がその前（75|61および47|61）にあり、53と29がその後（61|53および61|29）にあるからです。
53は正しく4番目に位置しています。29の前のページ番号（53|29）にあるからです。
29は唯一残ったページであり、正しく最後のページです。
最初の更新にはいくつかのページ番号が含まれていないため、それらの欠落したページ番号に関する並び替えルールは無視されます。
2番目と3番目の更新も、ルールに従って正しい順序になっています。最初の更新と同様に、これらの更新にもすべてのページ番号が含まれていないため、一部の並び替えルールのみが適用されます。
4番目の更新、75、97、47、61、53は、正しい順序ではありません。75が97の前に印刷されることになり、これはルール97|75に違反しています。
5番目の更新、
61、13、29
も、ルール
29|13
を破っているため、正しい順序ではありません。
最後の更新、97,13,75,29,47は、いくつかのルールに違反しているため、正しい順序ではありません。
何らかの理由で、エルフは印刷される各更新の中間のページ番号も知る必要があります。現在、正しい順序で更新を印刷しているだけなので、正しい順序で印刷される各更新の中間のページ番号を見つける必要があります。上記の例では、正しい順序で印刷される更新は次のとおりです。
75、47、61、53、29
97、61、53、29、13
75、29、13
これらのページ番号の中央値は、それぞれ61、53、
29です。これらのページ番号を合計すると、143になります。もちろん、注意が必要です。実際のページ順序のルールのリストは、上記の例よりも多く、複雑です。
すでに正しい順序になっている更新はどれかを判断します。正しい順序になっている更新の真ん中のページ番号を合計すると何になるでしょうか？