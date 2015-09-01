class longestcarr {
	public int[] longestcarr(int[] arr){
		int maxl = 1;
		int pre = -1;
		int len = arr.length;
		int start = -1;
		for (int i=1;i<len;i++) {
			if (arr[i]-arr[i-1]==1) {
				if (i-pre>maxl) {
					maxl = i-pre;
					start = i-maxl;
				}
			}else{
				pre = i-1;
			}
		}
		int [] ret = new int[maxl];
		for (int i=0; i<maxl;i++ ) {
			// System.out.print(arr[i]+"\t");
			ret[i] = arr[start+1+i];
		}
		return  ret;
	}
	public static void main(String[] args) {
		longestcarr arr = new longestcarr();
		int[] a = {1,2,3,4,5,8,7,8,9,10,11};
		for (int i :a ) {
			System.out.print(i+", ");
		}
		System.out.println();
		int[] ret = arr.longestcarr(a);
		for (int i :ret ) {
			System.out.print(i+", ");
		}
		System.out.println();
	}

}