import { useState, useEffect } from 'react'

export function Exemplo3Cronometro() {
  const [segundos, setSegundos] = useState(0)
  const [ativo, setAtivo] = useState(false)

  useEffect(() => {
    let interval = null

    if (ativo) {
      interval = setInterval(() => {
        setSegundos(s => s + 1)
      }, 1000)
    }

    // CLEANUP: Muito importante!
    // Se não limpar, o timer continua rodando mesmo se pausar ou sair da tela.
    return () => {
      if (interval) clearInterval(interval)
    }
  }, [ativo]) // Só recria o efeito se 'ativo' mudar

  return (
    <div style={{ padding: '20px', border: '1px solid #ccc', borderRadius: '8px', textAlign: 'center' }}>
      <h2 style={{ fontSize: '3rem', margin: '10px 0' }}>{segundos}s</h2>
      <div style={{ display: 'flex', gap: '10px', justifyContent: 'center' }}>
        <button onClick={() => setAtivo(!ativo)}>
          {ativo ? 'Pausar' : 'Iniciar'}
        </button>
        <button onClick={() => { setAtivo(false); setSegundos(0) }}>
          Zerar
        </button>
      </div>
    </div>
  )
}

