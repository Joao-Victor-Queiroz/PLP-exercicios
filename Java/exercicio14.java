class Configuracao {
    private static Configuracao instance;

    private Configuracao() {}

    public static Configuracao getInstance() {
        if (instance == null) {
            instance = new Configuracao();
        }
        return instance;
    }
}

public class Main {
    public static void main(String[] args) {
        Configuracao config1 = Configuracao.getInstance();
        Configuracao config2 = Configuracao.getInstance();
        System.out.println(config1 == config2);  // true
    }
}
