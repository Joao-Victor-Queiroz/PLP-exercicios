import java.util.ArrayList;
import java.util.List;

class Professor {
    String nome;

    Professor(String nome) {
        this.nome = nome;
    }
}

class Escola {
    List<Professor> professores = new ArrayList<>();

    void adicionarProfessor(Professor professor) {
        professores.add(professor);
    }
}

public class Main {
    public static void main(String[] args) {
        Escola escola = new Escola();
        escola.adicionarProfessor(new Professor("Maria"));
        escola.adicionarProfessor(new Professor("José"));

        for (Professor professor : escola.professores) {
            System.out.println(professor.nome);
        }
    }
}
