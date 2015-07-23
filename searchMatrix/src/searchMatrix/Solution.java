package searchMatrix;

public class Solution {
	public boolean searchMatrix(int[][] matrix, int target) {
        int m = matrix.length;
        int n = matrix[0].length;
        int i = 0;
        int j = 0;
        for(;i<m||j<n;){
        	System.out.print(String.format("%d,%d  ",i,j));
        	if (matrix[i][j]<target) {
				i = i<m?i+1:i+0;
				j = j<n?j+1:j+0;
			}else{
				System.out.println();
				for (int j2 = j; j2 >=0; j2--) {
					System.out.print(String.format("%d,%d  ",i,j2));
					if (matrix[i][j2]==target) {
						return true;
					}
				}
				System.out.println();
				for (int i2 = i; i2 >=0; i2--) {
					System.out.print(String.format("%d,%d  ",i2,j));
					if (matrix[i2][j]==target) {
						return true;
					}
				}
				return false;
			}
        }
		return false;
    }
}
