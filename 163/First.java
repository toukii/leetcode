/*#黑默丁格的炮台#
@shaalx

Time Limit: 1000/1000 MS (Java/Others) Memory Limit: 32768/32768 K (Java/Others) Problem Description: 兰博教训了提莫之后，然后和提莫讨论起约德尔人，谈起约德尔人，自然少不了一个人，那就是黑默丁格——约德尔人历史上最伟大的科学家。 提莫说，黑默丁格最近在思考一个问题：黑默丁格有三个炮台，炮台能攻击到距离它R的敌人,(两点之间的距离为两点连线的距离,例如(3,0)和(0,4)之间的距离是5),如果一个炮台能攻击到敌人，那么会对敌人造成1X的伤害。黑默丁格将三个炮台放在N*M方格中的点上,并且给出敌人的坐标。 问：那么敌人受到伤害会是多大？ 输入 第一行9个整数，R，x1,y1,x2,y2,x3,y3,x0,y0。(0 <= R，x1,y1,x2,y2,x3,y3,x0,y0 <= 100) R 代表炮台攻击的最大距离，(x1,y1),(x2,y2),(x3,y3)代表三个炮台的坐标。(x0,y0)代表敌人的坐标。 输出 输出一行,这一行代表敌人承受的最大伤害,(如果每个炮台都不能攻击到敌人，输出0X)。 输出格式见样例。 样例输入 1 1 1 2 2 3 3 1 2 样例输出 2X*/
import java.util.Scanner;

public class Main {
	public void fire(int iR,int[][] input) {
		float R = iR*iR;
		float r = 0;
		int count = 0;
		int a = input[3][0];
		int b = input[3][1];
		for (int i = 0; i < 3; i++) {
			r = 0;
			r = (float)(Math.pow(input[i][0]-a,2)+Math.pow(input[i][1]-b,2));
			if (r<=R) {
				count++;
			}
		}
		System.out.printf("%dX",count);
	}
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int R = sc.nextInt();
        int[][] input = new int[4][2];
        for (int i = 0; i < 4; i++) {
			input[i][0] = sc.nextInt();
			input[i][1] = sc.nextInt();
		}
        Main m = new Main();
        m.fire(R, input);
    }
}
