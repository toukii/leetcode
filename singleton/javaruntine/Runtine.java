package javaruntine;

import java.text.SimpleDateFormat;
import java.util.Date;

public class Runtine {
	public static String Times;
	private static Runtine runtine;
	private Runtine (){}
	public static synchronized Runtine getRuntine() throws InterruptedException {
		if (null==runtine) {
			Thread.sleep(4000);
			SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd hh:mm:ss");
			Times = sdf.format(new Date());
			runtine = new Runtine();
		}
		return runtine;
	}
	public static Runtine getRuntine2() throws InterruptedException{
		if (null==runtine) {
			Thread.sleep(4000);
			synchronized (Runtine.class) {
				if (null==runtine) {
					SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd hh:mm:ss");
					Times = sdf.format(new Date());
					runtine = new Runtine();
				}
			}
		}
		return runtine;
	}
}
