import java.util.Scanner;
/*
判断m行里是否有n个true
Time Limit: 2000/1000 MS (Java/Others) Memory Limit: 65536/65536 K (Java/Others)
Problem Description:
把下面的java的代码
public static boolean atLeastTwoTrue(boolean a, boolean b, boolean c) {
  return a ? (b || c) : (b && c);
}
a,b会被检查一次，c至多检查一次（如果a和b都为true，则会跳过对c的检查）。把上述优化扩展为检查M个里是否有n个true
public static boolean nOutOfMTrue(int n, boolean... booleans) {
  
}
*/
public class Main {
	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		int n;
		String input = "";
		String TRUE = "true";
		n = scanner.nextInt();
		while(scanner.hasNext()){
			input = scanner.next();
			if (TRUE.equals(input)) {
				n --;
			}
			if (n<0) {
				break;
			}
		}
		if (n==0) {
			System.out.println("true");
		}else{
			System.out.println("false");
		}
	}
}
