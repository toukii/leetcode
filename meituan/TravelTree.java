package ctrip;

import java.util.Stack;

public class Main {
	public void Travel(Tree root) {
		Stack<Tree> stack = new Stack<Tree>();
		if (root==null) {
			return;
		}
		stack.push(root);
		Tree cur;
		while(!stack.isEmpty()){
			cur = stack.peek();
			if (cur.visited) {
				stack.pop();
				continue;
			}
			while(cur.left!=null&&!cur.left.visited){
				stack.push(cur.left);
				cur = cur.left;
			}
			Tree cur_right = cur.right;
			if (cur_right!=null&&!cur_right.visited) {
				stack.push(cur_right);
				continue;
			}
			System.out.print(cur.v+"-");
			cur.visited = true;
			stack.pop();
		}
	}
	public static void main(String[] args) {
		Tree r7 = new Tree(7, null, null);
		Tree r6 = new Tree(6, r7, null);
		Tree r5 = new Tree(5, null, null);
		Tree r4 = new Tree(4, null, r6);
		Tree r2 = new Tree(2, r4, r5);
		Tree r3 = new Tree(3, null, null);
		Tree r1 = new Tree(1, r2, r3);
		Main main = new Main();
		main.Travel(r1);
	}
}

class Tree{
	int v;
	boolean visited;
	Tree left,right;
	public Tree(int v,Tree l,Tree r) {
		visited = false;
		this.v = v;
		left = l;
		right = r;
	}
}