class ContaBancaria {
    private double saldo;
    private String titular;

    public ContaBancaria(String titular) {
        this.titular = titular;
        this.saldo = 0;
    }

    public void depositar(double valor) {
        this.saldo += valor;
    }

    public void sacar(double valor) throws Exception {
        if (valor > saldo) {
            throw new Exception("Saldo insuficiente");
        }
        this.saldo -= valor;
    }

    public double obterSaldo() {
        return saldo;
    }
}

public class Main {
    public static void main(String[] args) {
        try {
            ContaBancaria conta = new ContaBancaria("João");
            conta.depositar(100);
            System.out.println(conta.obterSaldo());
            conta.sacar(50);
            System.out.println(conta.obterSaldo());
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}
