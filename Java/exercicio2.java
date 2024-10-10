class Carro {
    String marca;
    String modelo;
    int ano;
    int velocidade;

    Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0;
    }

    void acelerar(int incremento) {
        this.velocidade += incremento;
    }

    void frear(int decremento) {
        this.velocidade -= decremento;
    }

    void exibirVelocidade() {
        System.out.println("Velocidade atual: " + this.velocidade + " km/h");
    }
}

public class Main {
    public static void main(String[] args) {
        Carro carro = new Carro("Toyota", "Corolla", 2020);
        carro.acelerar(50);
        carro.exibirVelocidade();
        carro.frear(20);
        carro.exibirVelocidade();
    }
}
