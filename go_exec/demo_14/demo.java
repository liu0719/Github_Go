public class demo {
	static long fib(long n) {
		return n <= 2 ? 1 : fib(n - 1) + fib(n - 2);
	}

	public static void main(String[] args) {
         long stime = System.currentTimeMillis();
		System.out.println(fib(45));
		long etime = System.currentTimeMillis();
        System.out.printf("%d", (etime - stime));
	}
}