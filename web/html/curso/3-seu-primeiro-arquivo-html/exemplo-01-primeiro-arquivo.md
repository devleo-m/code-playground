<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <!-- Metadados e configurações -->
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Exemplo de primeiro arquivo HTML completo">
    <title>Meu Primeiro Arquivo HTML</title>
</head>
<body>
    <!-- Cabeçalho principal -->
    <header id="cabecalho-principal" class="destaque">
        <h1>Bem-vindo ao Meu Primeiro Arquivo HTML</h1>
        <p class="subtitulo">Este é um exemplo completo de estrutura HTML</p>
    </header>
    
    <!-- Conteúdo principal -->
    <main id="conteudo-principal">
        <section>
            <h2>Sobre Tags e Atributos</h2>
            <p>
                Este parágrafo demonstra o uso de tags e atributos.
                Veja como usamos <code>&lt;h2&gt;</code> para títulos de seção
                e <code>&lt;p&gt;</code> para parágrafos.
            </p>
            
            <!-- Lista com atributos title -->
            <ul>
                <li title="Tag para títulos principais">h1 - Título Principal</li>
                <li title="Tag para parágrafos de texto">p - Parágrafo</li>
                <li title="Tag para listas não ordenadas">ul - Lista</li>
            </ul>
        </section>
        
        <section>
            <h2>Links e Imagens</h2>
            <p>
                Aqui temos exemplos de links e imagens com seus atributos:
            </p>
            
            <!-- Link externo -->
            <p>
                Visite o 
                <a href="https://www.exemplo.com" 
                   target="_blank" 
                   title="Abre em nova aba">
                    site exemplo
                </a>
                para mais informações.
            </p>
            
            <!-- Imagem com todos os atributos -->
            <img src="https://via.placeholder.com/300x200" 
                 alt="Imagem de exemplo com 300 pixels de largura e 200 de altura"
                 width="300" 
                 height="200"
                 title="Imagem de exemplo">
        </section>
        
        <section>
            <h2>Formulário Simples</h2>
            <form id="formulario-contato">
                <label for="nome-usuario">Nome:</label>
                <input type="text" 
                       id="nome-usuario" 
                       name="nome" 
                       placeholder="Digite seu nome"
                       required>
                
                <label for="email-usuario">Email:</label>
                <input type="email" 
                       id="email-usuario" 
                       name="email" 
                       placeholder="seu@email.com"
                       required>
                
                <button type="submit">Enviar</button>
            </form>
        </section>
    </main>
    
    <!-- Rodapé -->
    <footer id="rodape" class="rodape">
        <p>Copyright &copy; 2024 - Meu Primeiro Arquivo HTML</p>
    </footer>
</body>
</html>

