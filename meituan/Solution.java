/*
两个字符串，左右移动最少次数，使其相等。
左移动：最左边的字符移动到字符串的最右边。
右移动：最右边的字符移动到字符串的最左边。
*/
public class Solution {
    public int distance(String s1, String s2) {
    	int len1 = s1.length();
    	int len2 = s2.length();
    	if (len1!=len2) {
			return -1;
		}
    	int j = 0,i=0;
    	for (i = 0; i < len1; i++) {
    		if (s1.charAt(i)!=s2.charAt(j)) {
    			if (j>0) {
    				break;	
				}
			}else{
				j++;
			}
		}
    	int min = j;
    	i = 0;
    	//System.out.println(String.valueOf(s2.charAt(j)));
    	for(;j<len1;j++){
    		if(s2.charAt(j)!=s1.charAt(i)){
    			return -1;
    		}
    		i++;
    	}
    	//System.out.println(String.valueOf(s1.charAt(i)));
    	//System.out.printf("min:%d, i:%d", min,i);
    	return min<i?min:i;
    }
    
    public static void main(String[] args) {
		Solution solution = new Solution();
		int ret = solution.distance("abcddabcabc", "abcabcddabc");
		System.out.println(ret);
	}
}
