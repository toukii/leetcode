package searchMatrix;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int mat[][] = {
		               {1,   4,  7, 11, 15},
		               {2,   5,  8, 12, 19},
		               {3,   6,  9, 16, 22},
		               {10, 13, 14, 17, 24},
		               {18, 21, 23, 26, 30}
		               }; 
		Solution s = new Solution();
		boolean ok = s.searchMatrix(mat,15);
		System.out.println(ok);
	}

}
