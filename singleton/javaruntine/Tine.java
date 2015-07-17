package javaruntine;

import javaruntine.Runtine;

public class Tine extends Thread{
	public void run(){
		for (int i = 0; i < 10; i++) {
			try{
				Thread.sleep(500);				
			}catch(Exception e){}
			System.out.print(i+" : ");
			try {
				Runtine runtine = Runtine.getRuntine2();
				System.out.println(runtine.Times);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}
	}
	public static void main(String[]args) throws InterruptedException {
		Tine tine = new Tine();
		Thread th = new Thread(tine);
		Thread th2 = new Thread(tine);
		th.start();
		Thread.sleep(1000);
		th2.start();
		Thread.sleep(10000);
	}
}
