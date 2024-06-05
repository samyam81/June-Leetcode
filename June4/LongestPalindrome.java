package June4;

import java.util.HashMap;

public class LongestPalindrome {
    public int longestPalindrome(String s) {

        if (s.isEmpty()) return 0;
        HashMap<Character, Integer> myHashMap = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (myHashMap.containsKey(c))
                myHashMap.put(c, myHashMap.get(c) + 1);
            else
                myHashMap.put(c, 1);
        }

        int s81 = 0;
        boolean kbd = false;

        for (Character ks11 : myHashMap.keySet())
            if (myHashMap.get(ks11) % 2 == 0)
                s81 = s81 + myHashMap.get(ks11);
            else {
                s81 = s81 + myHashMap.get(ks11);
                if (!kbd) kbd = true;
                else s81 = s81 - 1;
            }
        return s81;

    }
}