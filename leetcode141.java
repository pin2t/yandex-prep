public class leetcode141 {
     public static void main(String[] args) {
         
     }

     class ListNode {
         int val;
         ListNode next;
         ListNode(int x) {
             val = x;
             next = null;
         }
    }
    public class Solution {
        public boolean hasCycle(ListNode head) {
            if (head.next == null) return false;
            var n = head.next;
            while (n != null) {
                var nn = head;
                while (nn != n) {
                    if (n.next == nn) return true;
                    nn = nn.next;
                }
                n = n.next;
            }
            return false;
        }
    }
}
