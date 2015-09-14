import java.util.Collections;
import java.util.Comparator;
import java.util.Iterator;
import java.util.LinkedHashMap;
import java.util.Map;

public class Vote {
	private LinkedHashMap<String,Candi> lm;
	public Vote(){
		this.lm = new LinkedHashMap<String,Candi>(10,(float) 0.75,false);
	}
	class Candi implements Comparable{		
		private String name;
		private int count;
		public Candi(String name){
			this.name = name;
			this.count = 0;
		}
		public int compareTo(Object obj){
			Candi cad = (Candi)obj;
			return this.name.compareTo(cad.name);
		}
	}
	public boolean add(String name){
		boolean found = lm.containsKey(name);
		if (!found) {
			lm.put(name,new Candi(name));
			return true;
		}
		return false;
	}

	public boolean vote(String name){
		Candi cad = lm.get(name);
		if (cad==null) {
			return false;
		}
		cad.count ++;
		return true;
	}

	public void result(){
		Iterator itr = lm.entrySet().iterator();
		while(itr.hasNext()){
			Map.Entry<String, Candi> ety = (Map.Entry<String, Candi>)itr.next();
			System.out.println(ety.getKey()+" "+ety.getValue().count);
		}
		System.out.println();

		for(Map.Entry<String, Candi> cad : lm.entrySet()){
			System.out.println(cad.getKey()+" "+cad.getValue().count);
		}
	}
	public static void main(String[] args) {
		Vote vote = new Vote();
		vote.add("b");
		vote.add("a");
		vote.add("c");

		vote.vote("c");
		vote.vote("a");
		vote.vote("a");
		vote.vote("b");

		vote.result();
	}

}