<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Exemplos de áudio e vídeo em HTML">
    <title>Áudio e Vídeo em HTML</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            line-height: 1.6;
        }
        .exemplo {
            margin: 30px 0;
            padding: 20px;
            border: 1px solid #ddd;
            border-radius: 5px;
        }
        h2 {
            color: #333;
            border-bottom: 2px solid #4CAF50;
            padding-bottom: 10px;
        }
        audio, video {
            width: 100%;
            max-width: 800px;
            margin: 20px 0;
        }
        .info {
            background: #f0f0f0;
            padding: 15px;
            border-radius: 5px;
            margin: 10px 0;
        }
    </style>
</head>
<body>
    <header>
        <h1>Exemplos de Áudio e Vídeo em HTML</h1>
        <p>Demonstração de como incorporar conteúdo de áudio e vídeo em páginas web</p>
    </header>

    <main>
        <!-- Exemplo 1: Áudio Básico -->
        <section class="exemplo">
            <h2>1. Áudio Básico com Controles</h2>
            <p>Player de áudio simples com controles:</p>
            <audio controls>
                <source src="audio/exemplo.mp3" type="audio/mpeg">
                <source src="audio/exemplo.ogg" type="audio/ogg">
                <p>
                    Seu navegador não suporta o elemento de áudio. 
                    <a href="audio/exemplo.mp3">Baixe o arquivo de áudio</a>.
                </p>
            </audio>
            <div class="info">
                <strong>Nota:</strong> Este exemplo requer arquivos de áudio reais. 
                Em produção, substitua pelos seus arquivos.
            </div>
        </section>

        <!-- Exemplo 2: Áudio com Múltiplas Fontes -->
        <section class="exemplo">
            <h2>2. Áudio com Múltiplos Formatos</h2>
            <p>Fornecer múltiplos formatos garante compatibilidade:</p>
            <audio controls preload="metadata">
                <source src="audio/podcast.mp3" type="audio/mpeg">
                <source src="audio/podcast.ogg" type="audio/ogg">
                <source src="audio/podcast.wav" type="audio/wav">
                <p>
                    Seu navegador não suporta áudio HTML5.
                    <a href="audio/podcast.mp3">Baixar MP3</a> ou
                    <a href="transcricao.txt">Ler transcrição</a>.
                </p>
            </audio>
        </section>

        <!-- Exemplo 3: Áudio com Preload Metadata -->
        <section class="exemplo">
            <h2>3. Áudio com Preload Metadata</h2>
            <p>Carrega apenas metadados (duração, etc.) sem baixar o arquivo completo:</p>
            <audio controls preload="metadata">
                <source src="audio/musica.mp3" type="audio/mpeg">
                <p>Áudio não disponível</p>
            </audio>
        </section>

        <!-- Exemplo 4: Áudio com Loop -->
        <section class="exemplo">
            <h2>4. Áudio com Loop (Repetição)</h2>
            <p>Áudio que repete automaticamente:</p>
            <audio controls loop>
                <source src="audio/ambient.mp3" type="audio/mpeg">
                <p>Áudio não disponível</p>
            </audio>
        </section>

        <!-- Exemplo 5: Vídeo Básico -->
        <section class="exemplo">
            <h2>5. Vídeo Básico com Controles</h2>
            <p>Player de vídeo simples com controles:</p>
            <video controls width="800" height="450">
                <source src="video/exemplo.mp4" type="video/mp4">
                <source src="video/exemplo.webm" type="video/webm">
                <p>
                    Seu navegador não suporta o elemento de vídeo.
                    <a href="video/exemplo.mp4">Baixe o vídeo</a>.
                </p>
            </video>
            <div class="info">
                <strong>Nota:</strong> Este exemplo requer arquivos de vídeo reais. 
                Em produção, substitua pelos seus arquivos.
            </div>
        </section>

        <!-- Exemplo 6: Vídeo com Poster -->
        <section class="exemplo">
            <h2>6. Vídeo com Imagem de Poster</h2>
            <p>Imagem exibida antes do vídeo começar a reproduzir:</p>
            <video 
                controls 
                width="800" 
                height="450"
                poster="https://via.placeholder.com/800x450?text=Thumbnail+do+Vídeo"
                preload="metadata"
            >
                <source src="video/tutorial.mp4" type="video/mp4">
                <source src="video/tutorial.webm" type="video/webm">
                <p>Vídeo não disponível</p>
            </video>
        </section>

        <!-- Exemplo 7: Vídeo com Múltiplas Fontes -->
        <section class="exemplo">
            <h2>7. Vídeo com Múltiplos Formatos</h2>
            <p>Fornecer múltiplos formatos para compatibilidade:</p>
            <video controls width="800" height="450" preload="metadata">
                <source src="video/filme.mp4" type="video/mp4">
                <source src="video/filme.webm" type="video/webm">
                <source src="video/filme.ogv" type="video/ogg">
                <p>
                    Seu navegador não suporta vídeo HTML5.
                    <a href="video/filme.mp4">Baixar MP4</a> ou
                    <a href="transcricao-video.txt">Ler transcrição</a>.
                </p>
            </video>
        </section>

        <!-- Exemplo 8: Vídeo com Legendas -->
        <section class="exemplo">
            <h2>8. Vídeo com Legendas (Subtitles)</h2>
            <p>Vídeo com suporte a legendas usando track element:</p>
            <video controls width="800" height="450">
                <source src="video/legendas.mp4" type="video/mp4">
                <track 
                    kind="subtitles" 
                    src="legendas/pt.vtt" 
                    srclang="pt" 
                    label="Português"
                    default
                >
                <track 
                    kind="subtitles" 
                    src="legendas/en.vtt" 
                    srclang="en" 
                    label="English"
                >
                <p>Vídeo não disponível</p>
            </video>
            <div class="info">
                <strong>Nota:</strong> Arquivos .vtt (WebVTT) são necessários para legendas.
                Este é um exemplo da estrutura - você precisaria criar os arquivos .vtt.
            </div>
        </section>

        <!-- Exemplo 9: Vídeo com Loop e Muted -->
        <section class="exemplo">
            <h2>9. Vídeo com Loop e Muted (Para Background)</h2>
            <p>Vídeo que repete automaticamente sem som (útil para backgrounds):</p>
            <video 
                controls 
                loop 
                muted 
                autoplay
                width="800" 
                height="450"
            >
                <source src="video/background.mp4" type="video/mp4">
                <p>Vídeo não disponível</p>
            </video>
            <div class="info">
                <strong>Nota:</strong> Muitos navegadores bloqueiam autoplay com áudio. 
                Use muted com autoplay se necessário.
            </div>
        </section>

        <!-- Exemplo 10: Vídeo Responsivo -->
        <section class="exemplo">
            <h2>10. Vídeo Responsivo</h2>
            <p>Vídeo que se adapta ao tamanho da tela:</p>
            <video 
                controls 
                width="800" 
                height="450"
                style="max-width: 100%; height: auto;"
            >
                <source src="video/responsivo.mp4" type="video/mp4">
                <p>Vídeo não disponível</p>
            </video>
        </section>

        <!-- Exemplo 11: Áudio com Transcrição -->
        <section class="exemplo">
            <h2>11. Áudio com Link para Transcrição</h2>
            <p>Fornecer transcrição para acessibilidade:</p>
            <audio controls>
                <source src="audio/podcast.mp3" type="audio/mpeg">
                <p>Áudio não disponível</p>
            </audio>
            <p>
                <a href="transcricao-podcast.txt" target="_blank">
                    📄 Ler transcrição do áudio
                </a>
            </p>
        </section>

        <!-- Exemplo 12: Vídeo com Descrição de Acessibilidade -->
        <section class="exemplo">
            <h2>12. Vídeo com Descrição para Acessibilidade</h2>
            <p>Vídeo com track de descrição para leitores de tela:</p>
            <video controls width="800" height="450">
                <source src="video/acessivel.mp4" type="video/mp4">
                <track 
                    kind="descriptions" 
                    src="descricoes/pt.vtt" 
                    srclang="pt" 
                    label="Descrição em Português"
                >
                <p>Vídeo não disponível</p>
            </video>
        </section>
    </main>

    <footer>
        <p>&copy; 2024 - Exemplos de Áudio e Vídeo em HTML</p>
        <p><small>
            <strong>Nota:</strong> Estes exemplos mostram a estrutura. 
            Em produção, você precisará de arquivos de mídia reais e otimizados.
        </small></p>
    </footer>
</body>
</html>

