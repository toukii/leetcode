import java.util.Scanner;
/*
查找最长数字串
Time Limit: 2000/1000 MS (Java/Others) Memory Limit: 65536/66536 K (Java/Others)
Problem Description:
在字符串中找出连续最长的数字串，若有相同长度的数字串则找出数字串中各字符相加后和更大的串。
*/
public class Main {
	public void Find(String in) {
		int length = in.length();
		char zero = '0';
		char nine = '9';
		char cur;
		int maxSum = 0;
		int maxLength = 0;
		int count = 0;
		int sum = 0;
		String ret = "";
		StringBuilder tmp = new StringBuilder(length);
		for (int i = 0; i < length; i++) {
			cur = in.charAt(i);
			if (cur>=zero && cur<=nine) {
				count ++;
				sum += (int)(cur - zero);
				tmp.append(cur);
				if (maxLength<count) {
					maxLength = count;
					maxSum = sum;
					ret = tmp.toString();
				}else if (maxLength==count && maxSum<sum) {
					maxSum = sum;
					ret = tmp.toString();
				}				
			}else{
				count = 0;
				sum = 0;
				tmp.delete(0, length);
			}
		}
		System.out.println(ret);
	}
	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		String in = "";
		Main m = new Main();
		while(scanner.hasNext()){
			in = scanner.next();
			m.Find(in);
		}
	}
}
