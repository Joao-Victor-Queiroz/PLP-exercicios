class Motor {
    int potencia;

    Motor(int potencia) {
        this.potencia = potencia;
    }
}

class Carro {
    String marca;
    String modelo;
    Motor motor;

    Carro(String marca, String modelo, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.motor = motor;
    }
}

public class Main {
    public static void main(String[] args) {
        Motor motor1 = new Motor(150);
        Carro carro = new Carro("Toyota", "Corolla", motor1);
        System.out.println(carro.marca + " " + carro.modelo + " com motor de " + carro.motor.potencia + " HP");
    }
}
