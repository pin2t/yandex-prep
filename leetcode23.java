import java.util.*;

import static java.lang.System.out;

public class leetcode23 {
     public static void main(String[] args) {
         var list = new Solution().mergeKLists(new ListNode[]{ toList(List.of(1, 4, 5)), toList(List.of(1, 3, 4)), toList(List.of(2, 6)) });
         var t = list;
         while (t != null) {
             out.print(" " + t.val);
             t = t.next;
         }
         out.println();
     }

     static ListNode toList(List<Integer> list) {
         if (list.size() == 0) return null;
         ListNode head = new ListNode(list.get(0), null);
         ListNode tail = head;
         for (int i = 1; i < list.size(); i++) {
             tail.next = new ListNode(list.get(i));
             tail = tail.next;
         }
         return head;
     }

     public static class ListNode {
         int val;
         ListNode next;
         ListNode() {}
         ListNode(int val) { this.val = val; }
         ListNode(int val, ListNode next) { this.val = val; this.next = next; }
     }

    static class Solution {
        public ListNode mergeKLists(ListNode[] lists) {
            if (lists.length == 0) return null;
            ListNode tail = null;
            ListNode head = null;
            var nodes = new ListNode[lists.length];
            for (int i = 0; i < lists.length; i++) nodes[i] = lists[i];
            while (true) {
                var i = 0;
                var min = Integer.MAX_VALUE;
                for (int j = 0; j < nodes.length; j++) {
                    if (nodes[j] != null && nodes[j].val < min) {
                        i = j;
                        min = nodes[j].val;
                    }
                }
                if (nodes[i] != null) {
                    var node = new ListNode(nodes[i].val);
                    if (tail != null)
                        tail.next = node;
                    tail = node;
                    if (head == null) head = node;
                    nodes[i] = nodes[i].next;
                } else break;
            }
            return head;
        }
    }
}
