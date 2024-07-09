'use client';
import React, { useState } from 'react';

interface EditableTextBlockProps {
  defaultText: string;
  maxLength?: number;
}

const EditableTextBlock: React.FC<EditableTextBlockProps> = ({ defaultText, maxLength = 750 }) => {
  const [text, setText] = useState(defaultText);
  const [isEditing, setIsEditing] = useState(false);

  const handleInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    const inputValue = event.target.value;
    if (inputValue.length <= maxLength) {
      setText(inputValue);
    }
  };

  const handleEditClick = () => {
    setIsEditing(true);
  };

  const handleSaveClick = () => {
    setIsEditing(false);
    // Логика для сохранения текста
    console.log('Сохранение текста:', text);
  };

  const handleCancelClick = () => {
    setText(defaultText);
    setIsEditing(false);
  };

  const renderTextContent = (textContent: string) => {
    // Регулярное выражение для обнаружения ссылок в тексте
    const urlRegex = /(https?:\/\/[^\s]+)/g;
    // Замена ссылок и символов новой строки на соответствующие HTML элементы
    const formattedText = textContent
      .replace(urlRegex, '<a href="$&" target="_blank" style="color: #0070f3; text-decoration: underline;" rel="noopener noreferrer">$&</a>')
      .replace(/\n/g, '<br />');

    return (
      <p 
        dangerouslySetInnerHTML={{ __html: formattedText }}
      />
    );
  };

  return (
    <div className="editable-text-block">
      {isEditing ? (
        <div>
          <textarea 
            value={text}
            onChange={handleInputChange}
            rows={6}
            maxLength={maxLength}
            style={{ resize: 'none', width: '100%' }}
          />
          <br />
          <span style={{ fontSize: '14px', color: '#555'}}>Осталось символов: {maxLength - text.length}</span>
          <br />
          <button onClick={handleSaveClick} style={{ fontSize: '14px', color: '#555'}}>Сохранить</button>
          <button onClick={handleCancelClick} style={{ fontSize: '14px', color: '#555', marginLeft: '20px'}}>Отмена</button>
        </div>
      ) : (
        <div>
          {renderTextContent(text)}
          <button onClick={handleEditClick} style={{ fontSize: '14px', color: '#555'}}>Редактировать</button>
        </div>
      )}
    </div>
  );
};

export default EditableTextBlock;
