/*
##
@shaalx

Time Limit: 1000/1000 MS (Java/Others) Memory Limit: 32768/32768 K (Java/Others) Problem Description: 小易经常沉迷于网络游戏。有一次，他在玩一个打怪升级的游戏，他的角色的初始能力值为a。在接下来的一段时间内，他将会依次遇见n个怪物，每个怪物的防御力为b1,b2,b3,…bn。如果遇到的怪物防御力bi小于等于小易的当前能力值c，那么他就能轻松打败怪物，并且使得自己的能力值增加bi；如果bi大于c，那他也能打败怪物，但他的能力值只能增加bi与c的最大公约数。那么问题来了，在一系列的锻炼后，小易的最终能力值为多少？ 输入 对于每组数据，第一行是两个整数n(1<=n<=100000)表示怪物的数量和a表示小易的初始能力值，第二行n个整数，b1,b2..bn.(1<=bi<=n)表示每个怪物的防御力 数据保证—— 50%的n<=100, 80%的n<=1000, 90%的n<=10000, 100%的n<=100000. 输出 对于每组数据，输出一行。每行仅包含一个整数，表示小易的最终能力值。 样例输入 3 50 50 105 200 5 20 30 20 15 40 100 样例输出 110 205
*/
import java.util.Scanner;

public class Main {
	public void Fire(Scanner sc,long n, long a) {
		float sum = a;
		long p = 0;
		for (long i = 0; i < n; i++) {
			p = sc.nextLong();
			System.out.println(p);
			if (p<=sum) {
				sum += p;
			}else{
				sum += Main.gcd(sum,p);
			}
		}
		System.out.printf("%.0f\n", sum);
	}
	public static double gcd(double a, double b) {
		if(a<b){
            double temp;
            temp=a;
            a=b;
            b=temp;	
            /*a=a^b;
			 b=a^b;
			 a=a^b;*/
	    }
	    if(0==b){
	        return a;
	    }
	    return gcd(a-b,b);
	}
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while(sc.hasNext()){
	        long n = sc.nextLong();
	        long a = sc.nextLong();
	        Main m = new Main();
	        m.Fire(sc, n, a);
        }
    }
}
