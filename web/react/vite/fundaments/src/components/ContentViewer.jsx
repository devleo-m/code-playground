import './ContentViewer.css'

function ContentViewer({ topic }) {
  if (!topic) {
    return (
      <div className="content-viewer empty">
        <p>Selecione um tópico para começar a estudar! 📚</p>
      </div>
    )
  }

  return (
    <div className="content-viewer">
      <div className="content-header">
        <h2>{topic.title}</h2>
      </div>
      <div 
        className="content-body"
        dangerouslySetInnerHTML={{ __html: topic.content }}
      />
    </div>
  )
}

export default ContentViewer

