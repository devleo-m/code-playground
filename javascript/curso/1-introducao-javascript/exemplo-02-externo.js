// exemplo-02-externo.js
// Este arquivo contém o JavaScript para exemplo-02-externo.html

console.log("✅ Arquivo JavaScript externo carregado!");

// Aguarda o HTML carregar completamente
document.addEventListener('DOMContentLoaded', function() {
    console.log("📄 HTML carregado, JavaScript pronto para interagir!");
    
    // Seleciona elementos
    const botao1 = document.getElementById('botao1');
    const botao2 = document.getElementById('botao2');
    const botao3 = document.getElementById('botao3');
    const resultado = document.getElementById('resultado');
    
    // Evento do Botão 1
    botao1.addEventListener('click', function() {
        console.log("🔘 Botão 1 clicado!");
        resultado.innerHTML = '<p style="color: green;">✅ Botão clicado com sucesso!</p>';
        resultado.style.backgroundColor = '#c8e6c9';
    });
    
    // Evento do Botão 2
    botao2.addEventListener('click', function() {
        console.log("🎨 Botão 2 clicado - alterando cor!");
        const cores = ['#ffebee', '#e3f2fd', '#f3e5f5', '#e8f5e9', '#fff3e0'];
        const corAleatoria = cores[Math.floor(Math.random() * cores.length)];
        document.body.style.backgroundColor = corAleatoria;
        resultado.innerHTML = '<p>🎨 Cor de fundo alterada!</p>';
        resultado.style.backgroundColor = corAleatoria;
    });
    
    // Evento do Botão 3
    botao3.addEventListener('click', function() {
        console.log("📅 Botão 3 clicado - mostrando data!");
        const agora = new Date();
        const dataFormatada = agora.toLocaleString('pt-BR', {
            day: '2-digit',
            month: '2-digit',
            year: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
        });
        resultado.innerHTML = `<p><strong>📅 Data e Hora Atual:</strong><br>${dataFormatada}</p>`;
        resultado.style.backgroundColor = '#e1f5fe';
    });
    
    // Mensagem inicial
    resultado.innerHTML = '<p>👆 Clique nos botões acima para interagir!</p>';
    
    console.log("🎯 Todos os event listeners configurados!");
});

