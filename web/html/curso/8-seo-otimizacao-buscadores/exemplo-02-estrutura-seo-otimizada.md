<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Guia Completo de HTML5 para Iniciantes | WebDev Academy</title>
    <meta name="description" content="Guia completo e detalhado sobre HTML5. Aprenda desde o básico até conceitos avançados com exemplos práticos e exercícios.">
</head>
<body>
    <header>
        <nav aria-label="Navegação principal">
            <ul>
                <li><a href="/">Início</a></li>
                <li><a href="/cursos">Cursos</a></li>
                <li><a href="/blog">Blog</a></li>
            </ul>
        </nav>
    </header>
    
    <main>
        <!-- Apenas um H1 por página -->
        <article>
            <h1>Guia Completo de HTML5 para Iniciantes</h1>
            
            <!-- Breadcrumb (migalhas de pão) para SEO -->
            <nav aria-label="Breadcrumb">
                <ol>
                    <li><a href="/">Início</a></li>
                    <li><a href="/cursos">Cursos</a></li>
                    <li><a href="/cursos/html">HTML</a></li>
                    <li aria-current="page">Guia Completo</li>
                </ol>
            </nav>
            
            <!-- Hierarquia lógica de títulos -->
            <section>
                <h2>O que é HTML?</h2>
                <p>HTML (HyperText Markup Language) é a linguagem de marcação padrão para criar páginas web.</p>
                
                <h3>História do HTML</h3>
                <p>O HTML foi criado por Tim Berners-Lee em 1991...</p>
                
                <h3>Versões do HTML</h3>
                <p>HTML passou por várias versões ao longo dos anos...</p>
            </section>
            
            <section>
                <h2>Estrutura Básica de um Documento HTML</h2>
                <p>Todo documento HTML segue uma estrutura fundamental...</p>
                
                <h3>Elemento DOCTYPE</h3>
                <p>O DOCTYPE declara a versão do HTML sendo usada...</p>
                
                <h3>Elemento HTML</h3>
                <p>O elemento &lt;html&gt; é o elemento raiz...</p>
                
                <h3>Elemento Head</h3>
                <p>O elemento &lt;head&gt; contém metadados...</p>
                
                <h3>Elemento Body</h3>
                <p>O elemento &lt;body&gt; contém o conteúdo visível...</p>
            </section>
            
            <section>
                <h2>Elementos Semânticos HTML5</h2>
                <p>HTML5 introduziu novos elementos semânticos...</p>
                
                <h3>Header e Footer</h3>
                <p>Elementos &lt;header&gt; e &lt;footer&gt; definem...</p>
                
                <h3>Nav</h3>
                <p>O elemento &lt;nav&gt; é usado para navegação...</p>
                
                <h3>Article e Section</h3>
                <p>Elementos &lt;article&gt; e &lt;section&gt; organizam...</p>
            </section>
            
            <!-- Links internos estratégicos -->
            <section>
                <h2>Próximos Passos</h2>
                <p>
                    Agora que você entende os fundamentos do HTML5, 
                    <a href="/curso/css">aprenda CSS</a> para estilizar suas páginas,
                    ou explore nosso <a href="/curso/javascript">curso de JavaScript</a> 
                    para adicionar interatividade.
                </p>
            </section>
        </article>
        
        <!-- Sidebar com conteúdo relacionado -->
        <aside>
            <h2>Conteúdo Relacionado</h2>
            <ul>
                <li><a href="/curso/css-basico">CSS Básico</a></li>
                <li><a href="/curso/javascript-intro">Introdução ao JavaScript</a></li>
                <li><a href="/curso/responsive-design">Design Responsivo</a></li>
            </ul>
        </aside>
    </main>
    
    <footer>
        <nav aria-label="Links do rodapé">
            <ul>
                <li><a href="/sobre">Sobre Nós</a></li>
                <li><a href="/contato">Contato</a></li>
                <li><a href="/privacidade">Política de Privacidade</a></li>
            </ul>
        </nav>
        <p>&copy; 2024 WebDev Academy. Todos os direitos reservados.</p>
    </footer>
</body>
</html>

