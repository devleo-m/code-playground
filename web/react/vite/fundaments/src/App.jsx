import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Layout from './components/Layout'
import Home from './pages/Home'
import Aulas from './pages/Aulas'
import Aula from './pages/Aula'
import './App.css'

function App() {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/aulas" element={<Aulas />} />
          <Route path="/aula/:aulaId" element={<Aula />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  )
}

export default App
