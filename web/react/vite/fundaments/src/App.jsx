import { useState } from 'react'
import Header from './components/Header'
import TopicList from './components/TopicList'
import ContentViewer from './components/ContentViewer'
import { topics } from './data/topics'
import './App.css'

function App() {
  const [selectedTopic, setSelectedTopic] = useState(null)

  const handleSelectTopic = (topic) => {
    setSelectedTopic(topic)
  }

  return (
    <div className="app">
      <Header />
      <div className="app-container">
        <TopicList
          topics={topics}
          onSelectTopic={handleSelectTopic}
          selectedTopicId={selectedTopic?.id}
        />
        <ContentViewer topic={selectedTopic} />
      </div>
    </div>
  )
}

export default App
