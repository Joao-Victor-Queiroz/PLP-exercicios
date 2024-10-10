abstract class Funcionario {
    abstract double calcularSalario();
}

class FuncionarioHorista extends Funcionario {
    int horasTrabalhadas;
    double valorPorHora;

    FuncionarioHorista(int horasTrabalhadas, double valorPorHora) {
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorPorHora = valorPorHora;
    }

    double calcularSalario() {
        return horasTrabalhadas * valorPorHora;
    }
}

class FuncionarioAssalariado extends Funcionario {
    double salario;

    FuncionarioAssalariado(double salario) {
        this.salario = salario;
    }

    double calcularSalario() {
        return salario;
    }
}

public class Main {
    public static void main(String[] args) {
        Funcionario funcionario1 = new FuncionarioHorista(160, 25);
        Funcionario funcionario2 = new FuncionarioAssalariado(3000);

        System.out.println(funcionario1.calcularSalario());
        System.out.println(funcionario2.calcularSalario());
    }
}
