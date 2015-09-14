public class Card{
	private String input;
	public Card(String i){
		input = i;
	}
	public int check3(){
		int length = input.length();
		int leng = 0;
		int cur = -1;
		char T = 'T';
		for(int i=0;i<length-1;i+=2){
			if (input.charAt(i)!=T) {
				break;
			}
			int now = input.charAt(i+1)-48;
			System.out.println(now);
			if (now-1==cur) {
				leng++;
				cur = now;
			}else{
				leng = 1;
				cur = now;
			}
			if (leng>=3) {
				return 1;
			}
		}
		return 0;
	}
	public int check4(){
		int length = input.length();
		int leng = 0;
		String cur = "";
		for(int i=0;i<length-7;i+=2){
			if(i+8>length) break;
			String target = input.substring(i,i+8);
			StringBuilder curr = new StringBuilder(input.substring(i, i+2));
			for (int j = 0; j < 2; j++) {
				curr.append(curr.toString());
			}
			if(target.equals(curr.toString())) return 2;
		}
		return 0;
	}
	public int check7(){
		int length = input.length();
		int leng = 0;
		return 0;
	}
	public static void main(String[] args){

//		Card ca = new Card("T1T2T3");
//		int c3=ca.check3();
//		System.out.println(c3);
		
		Card ca = new Card("T1T1T1T1");
		int c4=ca.check4();
		System.out.println(c4);
	}
}