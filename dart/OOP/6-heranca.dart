class Animal{
  String comer(){
    return "Animal esta comendo!";
  }
}

class Cachorro extends Animal{
  String latir(){
    return "Au Au";
  }
}

main(){
  Cachorro dog = Cachorro();
  print(dog.latir());
  print(dog.comer());
}