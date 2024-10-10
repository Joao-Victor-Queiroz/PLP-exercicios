import java.util.ArrayList;
import java.util.List;

class Empregado {
    String nome;
    String cargo;
    double salario;

    Empregado(String nome, String cargo, double salario) {
        this.nome = nome;
        this.cargo = cargo;
        this.salario = salario;
    }
}

class Empresa {
    List<Empregado> empregados = new ArrayList<>();

    void adicionarEmpregado(Empregado empregado) {
        empregados.add(empregado);
    }
}

public class Main {
    public static void main(String[] args) {
        Empresa empresa = new Empresa();
        empresa.adicionarEmpregado(new Empregado("João", "Gerente", 5000));
        empresa.adicionarEmpregado(new Empregado("Ana", "Desenvolvedora", 4000));

        for (Empregado empregado : empresa.empregados) {
            System.out.println(empregado.nome + ", " + empregado.cargo + ", R$" + empregado.salario);
        }
    }
}
