<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Exemplos práticos de títulos, parágrafos, quebras de linha e regras horizontais em HTML.">
    <title>Títulos e Parágrafos - Exemplos HTML</title>
</head>
<body>
    <header>
        <h1>Títulos e Parágrafos em HTML</h1>
        <p>Este documento demonstra o uso correto de títulos, parágrafos, quebras de linha e regras horizontais.</p>
    </header>
    
    <main>
        <!-- Seção sobre hierarquia de títulos -->
        <section>
            <h2>Hierarquia de Títulos (h1 a h6)</h2>
            
            <p>
                HTML fornece seis níveis de títulos, de <code>&lt;h1&gt;</code> (mais importante)
                a <code>&lt;h6&gt;</code> (menos importante). É essencial manter a hierarquia correta.
            </p>
            
            <h3>Exemplo de Hierarquia Correta</h3>
            <p>Observe como os títulos são renderizados:</p>
            
            <h1>Título Nível 1 (h1) - Use apenas um por página</h1>
            <h2>Título Nível 2 (h2) - Seções principais</h2>
            <h3>Título Nível 3 (h3) - Subseções</h3>
            <h4>Título Nível 4 (h4) - Sub-subseções</h4>
            <h5>Título Nível 5 (h5) - Raramente usado</h5>
            <h6>Título Nível 6 (h6) - Muito raramente usado</h6>
            
            <h3>Regras Importantes</h3>
            <ul>
                <li><strong>Apenas um h1 por página</strong> - é o título principal</li>
                <li><strong>Não pule níveis</strong> - não vá de h1 para h3</li>
                <li><strong>Use títulos para estrutura</strong>, não apenas para tamanho</li>
                <li><strong>Mantenha hierarquia lógica</strong> - h1 → h2 → h3</li>
            </ul>
        </section>
        
        <!-- Seção sobre parágrafos -->
        <section>
            <h2>Parágrafos (&lt;p&gt;)</h2>
            
            <p>
                A tag <code>&lt;p&gt;</code> define um parágrafo. Navegadores adicionam
                automaticamente uma linha em branco antes e depois de cada parágrafo.
            </p>
            
            <p>
                Este é o primeiro parágrafo. Ele contém várias frases que formam
                uma unidade de pensamento ou ideia relacionada.
            </p>
            
            <p>
                Este é o segundo parágrafo. Ele é separado do anterior por uma linha
                em branco automática. Cada parágrafo deve conter ideias relacionadas.
            </p>
            
            <p>
                Este é o terceiro parágrafo. Observe como os navegadores colapsam
                múltiplos espaços em branco dentro de um parágrafo. Por exemplo,
                este texto tem     muitos     espaços, mas o navegador os reduz
                a um único espaço.
            </p>
            
            <h3>Características dos Parágrafos</h3>
            <ul>
                <li>São elementos de bloco (ocupam toda a largura disponível)</li>
                <li>Navegadores colapsam múltiplos espaços em branco</li>
                <li>Use para blocos de texto relacionados</li>
                <li>Não use para espaçamento visual (use CSS)</li>
            </ul>
        </section>
        
        <!-- Seção sobre quebras de linha -->
        <section>
            <h2>Quebras de Linha (&lt;br&gt;)</h2>
            
            <p>
                A tag <code>&lt;br&gt;</code> cria uma quebra de linha dentro de um
                bloco de texto. É uma tag vazia (self-closing).
            </p>
            
            <h3>Quando Usar &lt;br&gt;</h3>
            
            <h4>1. Endereços</h4>
            <address>
                Rua das Flores, 123<br>
                Bairro Centro<br>
                São Paulo - SP<br>
                CEP: 01234-567
            </address>
            
            <h4>2. Poemas ou Versos</h4>
            <p>
                No meio do caminho tinha uma pedra<br>
                Tinha uma pedra no meio do caminho<br>
                Tinha uma pedra<br>
                No meio do caminho tinha uma pedra.
            </p>
            
            <h3>Quando NÃO Usar &lt;br&gt;</h3>
            <p>
                <strong>❌ Não use para criar espaço entre parágrafos:</strong>
            </p>
            <p style="background: #f0f0f0; padding: 10px;">
                Parágrafo 1<br><br><br>
                Parágrafo 2
            </p>
            
            <p>
                <strong>✅ Use parágrafos separados:</strong>
            </p>
            <p>Parágrafo 1</p>
            <p>Parágrafo 2</p>
            
            <p>
                <em>Lembre-se:</em> Para espaçamento visual, use CSS. Para layout,
                use CSS. O <code>&lt;br&gt;</code> deve ser usado apenas quando a
                quebra de linha é parte do conteúdo.
            </p>
        </section>
        
        <!-- Seção sobre regras horizontais -->
        <section>
            <h2>Regras Horizontais (&lt;hr&gt;)</h2>
            
            <p>
                A tag <code>&lt;hr&gt;</code> cria uma quebra temática horizontal.
                Visualmente, é exibida como uma linha horizontal.
            </p>
            
            <h3>Exemplo de Uso</h3>
            
            <p>
                Esta é a primeira seção do conteúdo. Ela fala sobre um assunto específico
                e contém informações relevantes sobre títulos e parágrafos.
            </p>
            
            <hr>
            
            <p>
                Esta é a segunda seção. O <code>&lt;hr&gt;</code> acima indica uma
                mudança temática. Agora estamos falando sobre regras horizontais.
            </p>
            
            <hr>
            
            <p>
                Esta é a terceira seção. Observe como o <code>&lt;hr&gt;</code> separa
                visualmente as seções, melhorando a legibilidade do documento.
            </p>
            
            <h3>Quando Usar &lt;hr&gt;</h3>
            <ul>
                <li>Para separar seções de conteúdo relacionado</li>
                <li>Para indicar mudança de assunto ou tema</li>
                <li>Em documentos longos para melhorar legibilidade</li>
                <li>Para criar separação visual entre partes do conteúdo</li>
            </ul>
            
            <h3>Características</h3>
            <ul>
                <li>É uma tag vazia (self-closing)</li>
                <li>Indica mudança temática (semântica em HTML5)</li>
                <li>Visualmente separa conteúdo</li>
                <li>Pode ser estilizado com CSS</li>
            </ul>
        </section>
        
        <!-- Seção sobre estruturação completa -->
        <section>
            <h2>Estruturação Completa de Conteúdo</h2>
            
            <p>
                Aqui está um exemplo de como combinar títulos, parágrafos e separadores
                para criar conteúdo bem estruturado:
            </p>
            
            <article>
                <h3>Artigo: Como Aprender HTML</h3>
                
                <p>
                    Aprender HTML é o primeiro passo para se tornar um desenvolvedor web.
                    HTML fornece a estrutura básica de qualquer página web.
                </p>
                
                <h4>Por Onde Começar</h4>
                <p>
                    Comece entendendo a estrutura básica de um documento HTML.
                    Aprenda sobre tags, elementos e atributos fundamentais.
                </p>
                
                <h4>Prática é Essencial</h4>
                <p>
                    Não basta apenas ler sobre HTML. Você precisa praticar criando
                    seus próprios arquivos e experimentando com diferentes tags.
                </p>
            </article>
            
            <hr>
            
            <article>
                <h3>Artigo: Boas Práticas de HTML</h3>
                
                <p>
                    Seguir boas práticas é essencial para criar código HTML profissional,
                    acessível e otimizado para mecanismos de busca.
                </p>
                
                <h4>Semântica</h4>
                <p>
                    Use tags semânticas apropriadas. Escolha a tag que melhor descreve
                    o conteúdo, não apenas a aparência visual.
                </p>
                
                <h4>Acessibilidade</h4>
                <p>
                    Sempre pense em acessibilidade. Use hierarquia de títulos correta,
                    forneça textos alternativos para imagens e crie links descritivos.
                </p>
            </article>
        </section>
    </main>
    
    <footer>
        <hr>
        <p>
            <small>
                Exemplo criado para demonstrar títulos, parágrafos, quebras de linha
                e regras horizontais em HTML. Sempre mantenha a hierarquia correta!
            </small>
        </p>
    </footer>
</body>
</html>

