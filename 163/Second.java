/*#股神#
@shaalx

Time Limit: 1000/1000 MS (Java/Others) Memory Limit: 32768/32768 K (Java/Others) Problem Description: 有股神吗？ 有，小红就是！ 经过严密的计算，小红买了一支股票，他知道从他买股票的那天开始，股票会有以下变化：第一天不变，以后涨一天，跌一天，涨两天，跌一天，涨三天，跌一天...依此类推。 为方便计算，假设每次涨和跌皆为1，股票初始单价也为1，请计算买股票的第n天每股股票值多少钱？ 输入 输入包括多组数据； 每行输入一个n，1<=n<=10^9 。 输出 请输出他每股股票多少钱，对于每组数据，输出一行。 样例输入 1 2 3 4 5 样例输出 1 2 1 2 3*/
import java.util.Scanner;

public class Main {
	public void count(long n) {
		long k = (int)Math.sqrt(2*n);
		if (k*(k+1)>2*n) {
			k -=1;
		}
		double sum = 0;
		sum = 1 + (k-2)*(k-2+1)/2;
		sum += n - (k*(k+1)/2);
		System.out.printf("%.0f\n", sum);
		
	}
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while(sc.hasNext()){
	        long n = sc.nextInt();
	        Main m = new Main();
	        m.count(n);
        }
    }
}
