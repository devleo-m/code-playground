class Carro(){
  String _modelo;
  int ano;   //atributos
  String placa;

  Carro(this.modelo, this.ano, this.placa); //construtor

  String get modelo{
    return this._modelo;
  }

  void set modelo(String modelo){
    this._modelo = modelo;
  }

  
}