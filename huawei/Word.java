import java.util.Scanner;

public class Word{
	boolean visited[][];
	String map[][];
	int m,n;
	String target;
	public void init(){
		visited = new boolean[m][n];
		map = new String[m][n];
	}
	public boolean findFrom(int i,int j,int cur){
		System.out.println(i+"-"+j);
		if (cur>=target.length()) {
			return true;
		}
		if (visited[i][j]) {
			return false;
		}
		visited[i][j] = true;
		boolean found = false;
		for(int ii=i-1;ii<=i+1;ii++){
			for (int jj=j-1;jj<=j+1;jj++){
				if (valid(ii,jj)&&!visited[ii][jj]&&map[ii][jj].equals(target.charAt(cur))) {
				 	found = findFrom(ii,jj,cur+1);
				 	if (found) {
				 		return true;
				 	}
				}
			}
		}
		visited[i][j] = false;
		return false;
	}
	public boolean from(int i,int j){
		return findFrom(i,j,0);
	}
	boolean valid(int i,int j){
		if (i>=0&&i<m&&j>=0&&j<n) {
			return true;
		}
		return false;
	}
	public void shwo(){
		System.out.println(target);
		System.out.println(m+"\t"+n);
		for(int i=0;i<m;i++){
			for (int j=0;j<n;j++){
				System.out.print(map[i][j]+" ");
			}
			System.out.println();
		}
	}
	public static void main(String[] args) {
		Word w = new Word();
		Scanner sc = new Scanner(System.in);
		boolean found = false;
		while(sc.hasNext()){
			w.m = sc.nextInt();
			w.n = sc.nextInt();
			found = false;
			w.target  = sc.next();
			w.init();
			for(int i=0;i<w.m;i++){
				for (int j=0;j<w.n;j++){
					w.map[i][j]=sc.next();
				}
			}
			w.shwo();
			for(int i=0;i<w.m;i++){
				for (int j=0;j<w.n;j++){
					if (w.map[i][j].equals(w.target.charAt(0))) {
						found = w.from(i,j);
						if (found) {
							break;
						}
					}
				}
				if (found) {
					break;
				}
			}
			if (found) {
				System.out.println(true);
			}else{
				System.out.println(false);
			}
		}
	}
}