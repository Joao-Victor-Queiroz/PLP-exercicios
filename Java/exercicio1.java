class Carro {
    String marca;
    String modelo;
    int ano;

    Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
    }
}

public class Main {
    public static void main(String[] args) {
        Carro carro1 = new Carro("Toyota", "Corolla", 2020);
        Carro carro2 = new Carro("Honda", "Civic", 2019);
        Carro carro3 = new Carro("Ford", "Focus", 2021);

        Carro[] carros = {carro1, carro2, carro3};
        for (Carro carro : carros) {
            System.out.println(carro.marca + " " + carro.modelo + " " + carro.ano);
        }
    }
}
