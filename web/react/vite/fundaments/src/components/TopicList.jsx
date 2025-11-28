import './TopicList.css'

function TopicList({ topics, onSelectTopic, selectedTopicId }) {
  return (
    <div className="topic-list">
      <h2>Escolha um tópico para estudar:</h2>
      <div className="topics-grid">
        {topics.map((topic) => (
          <div
            key={topic.id}
            className={`topic-card ${selectedTopicId === topic.id ? 'active' : ''}`}
            onClick={() => onSelectTopic(topic)}
          >
            <h3>{topic.title}</h3>
            <p>{topic.description}</p>
          </div>
        ))}
      </div>
    </div>
  )
}

export default TopicList

