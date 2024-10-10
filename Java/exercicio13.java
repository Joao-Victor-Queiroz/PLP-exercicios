class Matematica {
    static int fatorial(int n) {
        if (n == 0) {
            return 1;
        }
        return n * fatorial(n - 1);
    }
}

public class Main {
    public static void main(String[] args) {
        System.out.println(Matematica.fatorial(5));
    }
}
