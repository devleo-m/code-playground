import { useState } from "react";

export default function Contador(){
    const [count, setCount] = useState(0)

    return (
        <div>
            <button onClick={() => setCount(count + 1)}>Incrementar</button>
            <button onClick={() => setCount(count - 1)}>Decrementar</button>
            <button onClick={() => setCount(0)}>Resetar</button>
            <p>Contador: {count}</p>
        </div>
    )
}