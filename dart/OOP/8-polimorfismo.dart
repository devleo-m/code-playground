//Polimorfismo: É a capacidade de um objeto ser tratado como seu tipo próprio ou como o 
//tipo de sua classe pai. Isso permite que métodos em classes diferentes se comportem de 
//maneira diferente, mesmo que tenham o mesmo nome.

class Animal {
  void fazerSom() {
    print("Animal fazendo som.");
  }
}

class Gato extends Animal {
  @override
  void fazerSom() {
    print("Miau!");
  }
}

class Cachorro extends Animal {
  @override
  void fazerSom() {
    print("Au Au!");
  }
}

class Cachorro2 extends Animal {
  @override
  void fazerSom() {
    print("Au Au! 2");
  }
}

void main() {
  List<Animal> animais = [Gato(), Cachorro(), Cachorro2()];
  for (Animal animal in animais) {
    animal.fazerSom(); // Polimorfismo: chamada do método específico de cada subclasse
  }
}
