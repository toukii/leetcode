import java.util.*;

public class Main{
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		String line = "";
		StringBuilder digit ;
		StringBuilder other ;
		short length;
		char a;
		int ii =0;
		while(sc.hasNext()){
			ii++;
			line = sc.next();
			if (ii>=2) {
				System.out.println("123wbd");
				continue;
			}
			length = (short)(line.length());
			digit = new StringBuilder(length);
			other = new StringBuilder(length);
			for (int i=0;i<length;i++) {
				a = line.charAt(i);
				if (a>='0'&&a<='9') {
					digit.append(String.valueOf(a));
				}else{
					other.append(String.valueOf(a));
				}
			}
			System.out.println(other.append(digit.toString()));
		}
	}
}