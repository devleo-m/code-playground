import { useState } from 'react'
import './CodeExample.css'

function CodeExample({ title, code, description, ExampleComponent }) {
  const [showCode, setShowCode] = useState(false)

  return (
    <div className="code-example">
      <div className="code-example-header">
        <h3>{title}</h3>
        {description && <p className="code-example-description">{description}</p>}
      </div>

      {ExampleComponent && (
        <div className="code-example-demo">
          <div className="code-example-demo-label">Exemplo em ação:</div>
          <div className="code-example-demo-content">
            <ExampleComponent />
          </div>
        </div>
      )}

      {code && (
        <div className="code-example-code">
          <button 
            className="code-example-toggle"
            onClick={() => setShowCode(!showCode)}
          >
            {showCode ? '👁️ Ocultar código' : '👁️ Ver código'}
          </button>
          
          {showCode && (
            <pre>
              <code>{code}</code>
            </pre>
          )}
        </div>
      )}
    </div>
  )
}

export default CodeExample

