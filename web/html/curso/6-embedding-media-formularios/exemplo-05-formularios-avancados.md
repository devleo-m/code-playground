<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Exemplos de formulários avançados em HTML">
    <title>Formulários Avançados em HTML</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
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
        form {
            margin: 20px 0;
        }
        label {
            display: block;
            margin: 10px 0 5px 0;
            font-weight: bold;
        }
        input, textarea, select {
            width: 100%;
            padding: 8px;
            margin-bottom: 15px;
            border: 1px solid #ddd;
            border-radius: 4px;
            box-sizing: border-box;
        }
        input[type="file"] {
            padding: 5px;
        }
        input[type="range"] {
            width: 100%;
        }
        input[type="color"] {
            width: 100px;
            height: 40px;
        }
        button {
            background-color: #4CAF50;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            margin-right: 10px;
        }
        button:hover {
            background-color: #45a049;
        }
        fieldset {
            border: 1px solid #ddd;
            border-radius: 4px;
            padding: 15px;
            margin: 15px 0;
        }
        legend {
            font-weight: bold;
            padding: 0 10px;
        }
        .info {
            background: #f0f0f0;
            padding: 15px;
            border-radius: 5px;
            margin: 10px 0;
        }
        .preview {
            margin-top: 10px;
            max-width: 300px;
        }
        .preview img {
            max-width: 100%;
            height: auto;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        input:valid {
            border-color: #4CAF50;
        }
        input:invalid:not(:placeholder-shown) {
            border-color: #f44336;
        }
        .erro {
            color: #f44336;
            font-size: 0.9em;
            margin-top: -10px;
            margin-bottom: 10px;
        }
    </style>
</head>
<body>
    <header>
        <h1>Exemplos de Formulários Avançados em HTML</h1>
        <p>Demonstração de recursos avançados de formulários</p>
    </header>

    <main>
        <!-- Exemplo 1: Upload de Arquivo -->
        <section class="exemplo">
            <h2>1. Upload de Arquivo</h2>
            <form action="/upload" method="post" enctype="multipart/form-data">
                <label for="arquivo">Selecione um arquivo:</label>
                <input 
                    type="file" 
                    id="arquivo" 
                    name="arquivo" 
                    accept="image/*"
                    required
                >
                <button type="submit">Enviar Arquivo</button>
            </form>
            <div class="info">
                <strong>Importante:</strong> Sempre use <code>enctype="multipart/form-data"</code> 
                em formulários com upload de arquivos.
            </div>
        </section>

        <!-- Exemplo 2: Upload Múltiplos Arquivos -->
        <section class="exemplo">
            <h2>2. Upload de Múltiplos Arquivos</h2>
            <form action="/upload" method="post" enctype="multipart/form-data">
                <label for="fotos">Selecione uma ou mais fotos:</label>
                <input 
                    type="file" 
                    id="fotos" 
                    name="fotos" 
                    accept="image/*"
                    multiple
                    required
                >
                <button type="submit">Enviar Fotos</button>
            </form>
        </section>

        <!-- Exemplo 3: Upload com Preview -->
        <section class="exemplo">
            <h2>3. Upload com Preview de Imagem</h2>
            <form action="/upload" method="post" enctype="multipart/form-data">
                <label for="foto-preview">Selecione uma foto:</label>
                <input 
                    type="file" 
                    id="foto-preview" 
                    name="foto-preview" 
                    accept="image/*"
                    onchange="previewImage(this)"
                >
                <div id="preview" class="preview"></div>
                <button type="submit">Enviar</button>
            </form>
        </section>

        <!-- Exemplo 4: Range (Controle Deslizante) -->
        <section class="exemplo">
            <h2>4. Range (Controle Deslizante)</h2>
            <form action="/processar" method="post">
                <label for="volume">Volume: <span id="volume-valor">50</span>%</label>
                <input 
                    type="range" 
                    id="volume" 
                    name="volume" 
                    min="0" 
                    max="100" 
                    value="50"
                    step="5"
                    oninput="document.getElementById('volume-valor').textContent = this.value"
                >
                <button type="submit">Enviar</button>
            </form>
        </section>

        <!-- Exemplo 5: Seletor de Cor -->
        <section class="exemplo">
            <h2>5. Seletor de Cor</h2>
            <form action="/processar" method="post">
                <label for="cor">Escolha uma cor:</label>
                <input 
                    type="color" 
                    id="cor" 
                    name="cor" 
                    value="#ff0000"
                >
                <button type="submit">Enviar</button>
            </form>
        </section>

        <!-- Exemplo 6: Select Múltipla Escolha -->
        <section class="exemplo">
            <h2>6. Select com Múltipla Escolha</h2>
            <form action="/processar" method="post">
                <label for="interesses">Interesses (segure Ctrl/Cmd para múltipla seleção):</label>
                <select id="interesses" name="interesses" multiple size="5">
                    <option value="tecnologia">Tecnologia</option>
                    <option value="esportes">Esportes</option>
                    <option value="musica">Música</option>
                    <option value="viagens">Viagens</option>
                    <option value="culinaria">Culinária</option>
                    <option value="leitura">Leitura</option>
                </select>
                <button type="submit">Enviar</button>
            </form>
        </section>

        <!-- Exemplo 7: Validação Customizada com JavaScript -->
        <section class="exemplo">
            <h2>7. Validação Customizada (Senha e Confirmar Senha)</h2>
            <form id="form-senha" action="/processar" method="post">
                <label for="senha-nova">Nova Senha:</label>
                <input 
                    type="password" 
                    id="senha-nova" 
                    name="senha-nova" 
                    required
                    minlength="8"
                    oninput="validarSenhas()"
                >
                
                <label for="confirmar-senha">Confirmar Senha:</label>
                <input 
                    type="password" 
                    id="confirmar-senha" 
                    name="confirmar-senha" 
                    required
                    oninput="validarSenhas()"
                >
                <div id="erro-senha" class="erro"></div>
                
                <button type="submit">Cadastrar</button>
            </form>
        </section>

        <!-- Exemplo 8: Autocomplete -->
        <section class="exemplo">
            <h2>8. Autocomplete do Navegador</h2>
            <form action="/processar" method="post">
                <label for="nome-autocomplete">Nome completo:</label>
                <input 
                    type="text" 
                    id="nome-autocomplete" 
                    name="nome-autocomplete" 
                    autocomplete="name"
                >
                
                <label for="email-autocomplete">Email:</label>
                <input 
                    type="email" 
                    id="email-autocomplete" 
                    name="email-autocomplete" 
                    autocomplete="email"
                >
                
                <label for="telefone-autocomplete">Telefone:</label>
                <input 
                    type="tel" 
                    id="telefone-autocomplete" 
                    name="telefone-autocomplete" 
                    autocomplete="tel"
                >
                
                <label for="endereco-autocomplete">Endereço:</label>
                <input 
                    type="text" 
                    id="endereco-autocomplete" 
                    name="endereco-autocomplete" 
                    autocomplete="street-address"
                >
                
                <button type="submit">Enviar</button>
            </form>
        </section>

        <!-- Exemplo 9: Formulário com Datalist -->
        <section class="exemplo">
            <h2>9. Datalist (Lista de Sugestões)</h2>
            <form action="/processar" method="post">
                <label for="cidade">Cidade:</label>
                <input 
                    type="text" 
                    id="cidade" 
                    name="cidade" 
                    list="cidades"
                    placeholder="Digite ou selecione uma cidade"
                >
                <datalist id="cidades">
                    <option value="São Paulo">
                    <option value="Rio de Janeiro">
                    <option value="Belo Horizonte">
                    <option value="Brasília">
                    <option value="Salvador">
                    <option value="Curitiba">
                    <option value="Porto Alegre">
                </datalist>
                <button type="submit">Enviar</button>
            </form>
        </section>

        <!-- Exemplo 10: Formulário Completo com Todas as Validações -->
        <section class="exemplo">
            <h2>10. Formulário Completo com Validações</h2>
            <form id="form-completo" action="/cadastrar" method="post">
                <fieldset>
                    <legend>Dados Pessoais</legend>
                    
                    <label for="nome-final">Nome completo: <span style="color: red;">*</span></label>
                    <input 
                        type="text" 
                        id="nome-final" 
                        name="nome-final" 
                        required
                        minlength="3"
                        maxlength="100"
                        pattern="[A-Za-zÀ-ÿ\s]+"
                        title="Apenas letras e espaços"
                        autocomplete="name"
                    >
                    
                    <label for="email-final">Email: <span style="color: red;">*</span></label>
                    <input 
                        type="email" 
                        id="email-final" 
                        name="email-final" 
                        required
                        autocomplete="email"
                    >
                    
                    <label for="senha-final">Senha: <span style="color: red;">*</span></label>
                    <input 
                        type="password" 
                        id="senha-final" 
                        name="senha-final" 
                        required
                        minlength="8"
                        pattern="(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}"
                        title="Mínimo 8 caracteres, incluindo maiúscula, minúscula e número"
                        autocomplete="new-password"
                    >
                    
                    <label for="data-nascimento">Data de nascimento: <span style="color: red;">*</span></label>
                    <input 
                        type="date" 
                        id="data-nascimento" 
                        name="data-nascimento" 
                        required
                        min="1900-01-01"
                        max="2024-12-31"
                    >
                </fieldset>
                
                <fieldset>
                    <legend>Preferências</legend>
                    
                    <label>
                        <input type="checkbox" name="newsletter" value="sim">
                        Desejo receber newsletter
                    </label>
                    
                    <label>
                        <input type="checkbox" name="termos" value="aceito" required>
                        Aceito os termos e condições <span style="color: red;">*</span>
                    </label>
                </fieldset>
                
                <button type="submit">Cadastrar</button>
                <button type="reset">Limpar</button>
            </form>
        </section>
    </main>

    <footer>
        <p>&copy; 2024 - Exemplos de Formulários Avançados em HTML</p>
        <p><small>
            <strong>Lembre-se:</strong> Validação no cliente é apenas para UX. 
            Sempre valide também no servidor para segurança!
        </small></p>
    </footer>

    <script>
        // Preview de imagem
        function previewImage(input) {
            const preview = document.getElementById('preview');
            preview.innerHTML = '';
            
            if (input.files && input.files[0]) {
                const reader = new FileReader();
                
                reader.onload = function(e) {
                    const img = document.createElement('img');
                    img.src = e.target.result;
                    preview.appendChild(img);
                };
                
                reader.readAsDataURL(input.files[0]);
            }
        }

        // Validação de senhas
        function validarSenhas() {
            const senha = document.getElementById('senha-nova').value;
            const confirmar = document.getElementById('confirmar-senha').value;
            const erro = document.getElementById('erro-senha');
            const form = document.getElementById('form-senha');
            
            if (confirmar && senha !== confirmar) {
                erro.textContent = 'As senhas não coincidem';
                document.getElementById('confirmar-senha').setCustomValidity('Senhas não coincidem');
            } else {
                erro.textContent = '';
                document.getElementById('confirmar-senha').setCustomValidity('');
            }
        }

        // Prevenir envio se senhas não coincidirem
        document.getElementById('form-senha').addEventListener('submit', function(e) {
            const senha = document.getElementById('senha-nova').value;
            const confirmar = document.getElementById('confirmar-senha').value;
            
            if (senha !== confirmar) {
                e.preventDefault();
                alert('As senhas não coincidem!');
                return false;
            }
        });
    </script>
</body>
</html>

