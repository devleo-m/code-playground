// Variáveis para armazenar referências de tempo e estado do cronômetro
let segundos = 0;
let intervalo = null;

// Função para iniciar o cronômetro
function iniciarCronometro() {
  if (!intervalo) {
    intervalo = setInterval(() => {
      segundos++;
      console.log(`Tempo: ${segundos} segundos`);
    }, 1000);
  }
}

// Função para parar o cronômetro
function pararCronometro() {
  clearInterval(intervalo);
  intervalo = null;
}

// Função para resetar o cronômetro
function resetarCronometro() {
  pararCronometro();
  segundos = 0;
  console.log('Cronômetro resetado.');
}