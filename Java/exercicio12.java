class Produto {
    double preco;

    Produto(double preco) {
        this.preco = preco;
    }

    Produto somar(Produto outro) {
        return new Produto(this.preco + outro.preco);
    }
}

public class Main {
    public static void main(String[] args) {
        Produto produto1 = new Produto(50);
        Produto produto2 = new Produto(30);
        Produto produto3 = produto1.somar(produto2);
        System.out.println(produto3.preco);
    }
}
