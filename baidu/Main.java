package baidu;

import java.util.HashMap;
import java.util.Scanner;

public class Main {
	public String parseUtf8(String in) {
		String utf8 = "";
		try{
			String iso = new String(in.getBytes("utf-8"), "iso-8859-1");
			utf8 = new String(iso.getBytes("iso-8859-1"), "utf-8"); 
			return utf8;
		}catch (Exception e) {
		}
		return "";
	}
	public int Check(String utf8_a, String utf8_b) {
		HashMap<String, Integer> hm = new HashMap<String,Integer>();
		char c;
		Integer it;
		for (int i = 0; i < utf8_a.length(); i++) {
			c = utf8_a.charAt(i);
			String s = String.valueOf(c);
			it = hm.get(s);
			if (null==it) {
				hm.put(s, 1);
			}else{
				hm.put(s, it+1);
			}
		}
		for (int i = 0; i < utf8_b.length(); i++) {
			c = utf8_b.charAt(i);
			String s = String.valueOf(c);
			it = hm.get(s);
			if (null==it||it.intValue()<1) {
				return 0;
			}else{
				hm.put(s, it-1);
			}
		}
		return 1;
	}
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Main m  = new Main();
		Scanner sc = new Scanner(System.in);
		String a = sc.next();
		String b = sc.next();
		String au = m.parseUtf8(a);
		String bu = m.parseUtf8(b);
		int ret = m.Check(au, bu);
		System.out.println(ret);
		
	}

}
