package fourth;

import java.util.ArrayList;
import java.util.List;

public class Main {
	int start,end;
	List<TreeNode> queue;
	public Main() {
		start = 0;
		end = 0;
		queue = new ArrayList<TreeNode>();
	}
	//层次遍历
	public void Level(TreeNode root) {
		if (root==null) {
			return;
		}
		queue.add(end++, root);
		TreeNode head = null;
		while(start<end){
			head = queue.get(start++);
			System.out.println(head.v);
			if (head.left!=null) {
				queue.add(end++,head.left);
			}
			if (head.right!=null) {
				queue.add(end++,head.right);
			}
		}
	}
	// 先序遍历
	private void Travel(TreeNode root) {
		if (root == null) {
			return;
		}
		queue.add(end++,root);
		TreeNode top = null;
		while(end>0){
			top = queue.get(--end);
			System.out.println(top.v);
			if (top.right!=null) {
				queue.add(end++,top.right);
			}
			if (top.left!=null) {
				queue.add(end++,top.left);
			}
		}
	}
	public static void main(String[] args) {
		Main main = new Main();
		TreeNode t4 = new TreeNode(4, null, null);
		TreeNode t5 = new TreeNode(5, null, null);
		TreeNode t6 = new TreeNode(6, null, null);
		TreeNode t2 = new TreeNode(2, t4, t5);
		TreeNode t3 = new TreeNode(3, null, t6);
		TreeNode root = new TreeNode(1, t2, t3);
		main.Travel(root);
	}
}

class TreeNode{
	int v;
	TreeNode left,right;
	public TreeNode(int vv,TreeNode l, TreeNode r) {
		v = vv;
		left = l;
		right = r;
	}
}