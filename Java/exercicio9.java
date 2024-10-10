interface Imprimivel {
    void imprimir();
}

class Relatorio implements Imprimivel {
    public void imprimir() {
        System.out.println("Imprimindo relatório");
    }
}

class Contrato implements Imprimivel {
    public void imprimir() {
        System.out.println("Imprimindo contrato");
    }
}

public class Main {
    public static void main(String[] args) {
        Imprimivel[] itens = {new Relatorio(), new Contrato()};
        for (Imprimivel item : itens) {
            item.imprimir();
        }
    }
}
