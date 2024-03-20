package classe.pratica1;

public class Comida {
    private String comida;
    private double peso;
    public Comida(String comida, double peso) {
        this.comida = comida;
        this.peso = peso;
    }
    public String getComida() {
        return comida;
    }
    public void setComida(String comida) {
        this.comida = comida;
    }
    public double getPeso() {
        return peso;
    }
    public void setPeso(double peso) {
        this.peso = peso;
    }
}
